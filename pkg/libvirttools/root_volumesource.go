/*
Copyright 2017 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package libvirttools

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	libvirtxml "github.com/libvirt/libvirt-go-xml"

	"github.com/Mirantis/virtlet/pkg/diskimage"
	"github.com/Mirantis/virtlet/pkg/metadata/types"
	"github.com/Mirantis/virtlet/pkg/virt"
)

// rootVolume denotes the root disk of the VM
type rootVolume struct {
	volumeBase
}

var _ VMVolume = &rootVolume{}

// GetRootVolume returns volume source for root volume clone.
func GetRootVolume(config *types.VMConfig, owner volumeOwner) ([]VMVolume, error) {
	var vol VMVolume
	rootDev := config.RootVolumeDevice()
	if rootDev != nil {
		vol = &persistentRootVolume{
			volumeBase: volumeBase{config, owner},
			dev:        *rootDev,
		}
	} else {
		vol = &rootVolume{
			volumeBase{config, owner},
		}
	}
	return []VMVolume{vol}, nil
}

func (v *rootVolume) volumeName() string {
	if len(v.config.ParsedAnnotations.Snapshot) == 0 {
		return "virtlet_root_" + v.config.DomainUUID
	}
	return "virtlet_root_" + v.config.DomainUUID +
		"." + v.config.ParsedAnnotations.Snapshot[1]
}

func (v *rootVolume) createVolume() (virt.StorageVolume, error) {
	imagePath, _, virtualSize, err := v.owner.ImageManager().GetImagePathDigestAndVirtualSize(v.config.Image)
	if err != nil {
		return nil, err
	}

	if v.config.ParsedAnnotations != nil && v.config.ParsedAnnotations.RootVolumeSize > 0 &&
		uint64(v.config.ParsedAnnotations.RootVolumeSize) > virtualSize {
		virtualSize = uint64(v.config.ParsedAnnotations.RootVolumeSize)
	}

	storagePool, err := v.owner.StoragePool()
	if err != nil {
		return nil, err
	}

	vols, err := storagePool.ListVolumes()
	if err != nil {
		glog.Errorf("listVolumes failed: %v", err)
		return nil, err
	}
	for n, v := range vols {
		glog.V(3).Infof("storage pool existed vol: %d, %v", n, v.Name())
	}

	exist, err := storagePool.LookupVolumeByName(v.volumeName())
	if err == nil {
		glog.V(3).Infof("vol %v existed, skip create", v.volumeName())
		return exist, err
	}

	glog.V(3).Infof("createVolume lookup %s err: %v", v.volumeName(), err)

	return storagePool.CreateStorageVol(&libvirtxml.StorageVolume{
		Type: "file",
		Name: v.volumeName(),
		Allocation: &libvirtxml.StorageVolumeSize{
			Unit:  "b",
			Value: 0,
		},
		Capacity: &libvirtxml.StorageVolumeSize{
			Unit:  "b",
			Value: virtualSize,
		},
		Target: &libvirtxml.StorageVolumeTarget{
			Format: &libvirtxml.StorageVolumeTargetFormat{Type: "qcow2"},
		},
		BackingStore: &libvirtxml.StorageVolumeBackingStore{
			Path:   imagePath,
			Format: &libvirtxml.StorageVolumeTargetFormat{Type: "qcow2"},
		},
	})
}

func (v *rootVolume) createSnapshotVolume(vol, backingFile string) error {
	// don't create snapshot volume when backingfile is empty(-)
	// this is a feature that to support reset the vm os
	if backingFile == "-" {
		return nil
	}
	storagePool, err := v.owner.StoragePool()
	if err != nil {
		return err
	}

	_, err = storagePool.LookupVolumeByName(vol)
	if err == nil {
		glog.V(3).Infof("Vol %s existed, don't clone", vol)
		return nil
	}

	backingVol, err := storagePool.LookupVolumeByName(filepath.Base(backingFile))
	if err != nil {
		return err
	}
	virtualSize, err := backingVol.Size()
	if err != nil {
		return err
	}

	glog.V(3).Infof("createVolume %s backing: %s", v.volumeName(), backingFile)

	_, err = storagePool.CreateStorageVol(&libvirtxml.StorageVolume{
		Type: "file",
		Name: v.volumeName(),
		Allocation: &libvirtxml.StorageVolumeSize{
			Unit:  "b",
			Value: 0,
		},
		Capacity: &libvirtxml.StorageVolumeSize{
			Unit:  "b",
			Value: virtualSize,
		},
		Target: &libvirtxml.StorageVolumeTarget{
			Format: &libvirtxml.StorageVolumeTargetFormat{Type: "qcow2"},
		},
		BackingStore: &libvirtxml.StorageVolumeBackingStore{
			Path:   backingFile,
			Format: &libvirtxml.StorageVolumeTargetFormat{Type: "qcow2"},
		},
	})

	return err
}

func (v *rootVolume) UUID() string { return "" }

func (v *rootVolume) Setup() (*libvirtxml.DomainDisk, *libvirtxml.DomainFilesystem, error) {

	if len(v.volumeBase.config.ParsedAnnotations.Snapshot) > 0 {
		backingFile := v.volumeBase.config.ParsedAnnotations.Snapshot[0]
		glog.V(3).Infof("Create vol %s snapshot from backingfile %s", v.volumeName(), backingFile)
		err := v.createSnapshotVolume(v.volumeName(), backingFile)
		if err != nil {
			glog.Errorf("Failed create snapshot file %s: %v", v.volumeName(), err)
		}
	}

	vol, err := v.createVolume()
	if err != nil {
		return nil, nil, err
	}

	volPath, err := vol.Path()
	if err != nil {
		return nil, nil, fmt.Errorf("error getting root volume path: %v", err)
	}

	if err := diskimage.Put(volPath, v.config.ParsedAnnotations.FilesForRootfs); err != nil {
		return nil, nil, fmt.Errorf("error tuning rootfs with files from configmap: %v", err)
	}

	return &libvirtxml.DomainDisk{
		Device: "disk",
		Driver: &libvirtxml.DomainDiskDriver{Name: "qemu", Type: "qcow2"},
		Source: &libvirtxml.DomainDiskSource{File: &libvirtxml.DomainDiskSourceFile{File: volPath}},
	}, nil, nil
}

func (v *rootVolume) Teardown() error {
	storagePool, err := v.owner.StoragePool()
	if err != nil {
		return err
	}
	// clear volume and snaphost
	volPrefix := "virtlet_root_" + v.config.DomainUUID
	vols, err := storagePool.ListVolumes()
	if err != nil {
		return err
	}
	for _, vol := range vols {
		if strings.HasPrefix(vol.Name(), volPrefix) {
			err = storagePool.RemoveVolumeByName(vol.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
