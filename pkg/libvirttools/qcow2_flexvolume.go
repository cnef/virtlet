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
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/glog"
	libvirtxml "github.com/libvirt/libvirt-go-xml"

	"github.com/Mirantis/virtlet/pkg/metadata/types"
	"github.com/Mirantis/virtlet/pkg/utils"
	"github.com/Mirantis/virtlet/pkg/virt"
)

const (
	defaultVolumeCapacity     = 1024
	defaultVolumeCapacityUnit = "MB"
)

var capacityUnits = []string{
	// https://libvirt.org/formatstorage.html#StorageVolFirst
	"B", "bytes", "KB", "K", "KiB", "MB", "M", "MiB", "GB", "G",
	"GiB", "TB", "T", "TiB", "PB", "P", "PiB", "EB", "E", "EiB",
}

var capacityRx = regexp.MustCompile(`^\s*(\d+)\s*(\S*)\s*$`)

type qcow2VolumeOptions struct {
	Capacity string `json:"capacity,omitempty"`
	UUID     string `json:"uuid"`
}

// qcow2Volume denotes a volume in QCOW2 format
type qcow2Volume struct {
	volumeBase
	capacity     int
	capacityUnit string
	name         string
	uuid         string
}

var _ VMVolume = &qcow2Volume{}

func newQCOW2Volume(volumeName, configPath string, config *types.VMConfig, owner volumeOwner) (VMVolume, error) {
	var err error
	var opts qcow2VolumeOptions
	if err = utils.ReadJSON(configPath, &opts); err != nil {
		return nil, fmt.Errorf("failed to parse qcow2 volume config %q: %v", configPath, err)
	}
	v := &qcow2Volume{
		volumeBase: volumeBase{config, owner},
		name:       volumeName,
		uuid:       opts.UUID,
	}

	v.capacity, v.capacityUnit, err = parseCapacityStr(opts.Capacity)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *qcow2Volume) volumeName() string {
	if len(v.config.ParsedAnnotations.Snapshot) == 0 {
		return "virtlet-" + v.config.DomainUUID + "-" + v.name
	}
	return "virtlet-" + v.config.DomainUUID + "-" + v.name +
		"." + v.config.ParsedAnnotations.Snapshot[1]
}

func (v *qcow2Volume) createQCOW2Volume(capacity uint64, capacityUnit string) (virt.StorageVolume, error) {
	storagePool, err := v.owner.StoragePool()
	if err != nil {
		return nil, err
	}
	vol, err := storagePool.LookupVolumeByName(v.volumeName())
	if err == nil {
		return vol, nil
	} else if err != virt.ErrStorageVolumeNotFound {
		return nil, err
	}
	return storagePool.CreateStorageVol(&libvirtxml.StorageVolume{
		Name:       v.volumeName(),
		Allocation: &libvirtxml.StorageVolumeSize{Value: 0},
		Capacity:   &libvirtxml.StorageVolumeSize{Unit: capacityUnit, Value: capacity},
		Target:     &libvirtxml.StorageVolumeTarget{Format: &libvirtxml.StorageVolumeTargetFormat{Type: "qcow2"}},
	})
}

// convert root volume name foramt to data volume format
func (v *qcow2Volume) snapshotBackingfile(backingFile string) string {
	backingFile = strings.Replace(backingFile, "virtlet_root_", "virtlet-", 1)
	if strings.Contains(backingFile, ".") {
		backingFile = strings.Replace(backingFile, ".", "-"+v.name+".", 1)
	} else {
		backingFile = backingFile + "-" + v.name
	}
	return backingFile
}

func (v *qcow2Volume) createSnapshotVolume(vol, backingFile string) error {
	storagePool, err := v.owner.StoragePool()
	if err != nil {
		return err
	}

	_, err = storagePool.LookupVolumeByName(vol)
	if err == nil {
		glog.V(3).Infof("Vol %s existed, don't clone", vol)
		return nil
	}

	backingVol, err := storagePool.LookupVolumeByName(backingFile)
	if err == virt.ErrStorageVolumeNotFound {
		// when add a new vol to VM, the backing file not existed
		// so, we need to create a new vol
		return nil
	} else if err != nil {
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

func (v *qcow2Volume) UUID() string {
	return v.uuid
}

func (v *qcow2Volume) Setup() (*libvirtxml.DomainDisk, *libvirtxml.DomainFilesystem, error) {

	if len(v.volumeBase.config.ParsedAnnotations.Snapshot) > 0 {
		backingFile := v.snapshotBackingfile(v.volumeBase.config.ParsedAnnotations.Snapshot[0])
		glog.V(3).Infof("Create vol %s snapshot from backingfile %s", v.volumeName(), backingFile)
		err := v.createSnapshotVolume(v.volumeName(), backingFile)
		if err != nil {
			glog.Errorf("Failed create snapshot file %s: %v", v.volumeName(), err)
		}
	}
	vol, err := v.createQCOW2Volume(uint64(v.capacity), v.capacityUnit)
	if err != nil {
		return nil, nil, fmt.Errorf("error during creation of volume '%s' with virtlet description %s: %v", v.volumeName(), v.name, err)
	}

	path, err := vol.Path()
	if err != nil {
		return nil, nil, err
	}

	// fix:
	// 1. loss data when recreated domain use exists disk
	// 2. windows guest don't suppport ext4
	//
	// err = vol.Format()
	// if err != nil {
	// 	return nil, nil, err
	// }

	return &libvirtxml.DomainDisk{
		Device: "disk",
		Source: &libvirtxml.DomainDiskSource{File: &libvirtxml.DomainDiskSourceFile{File: path}},
		Driver: &libvirtxml.DomainDiskDriver{Name: "qemu", Type: "qcow2"},
	}, nil, nil
}

func (v *qcow2Volume) Teardown() error {
	storagePool, err := v.owner.StoragePool()
	if err != nil {
		return err
	}
	// clear volume and snaphost
	volPrefix := "virtlet-" + v.config.DomainUUID + "-" + v.name
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

func parseCapacityStr(capacityStr string) (int, string, error) {
	if capacityStr == "" {
		return defaultVolumeCapacity, defaultVolumeCapacityUnit, nil
	}

	subs := capacityRx.FindStringSubmatch(capacityStr)
	if subs == nil {
		return 0, "", fmt.Errorf("invalid capacity spec: %q", capacityStr)
	}
	capacity, err := strconv.Atoi(subs[1])
	if err != nil {
		return 0, "", fmt.Errorf("invalid capacity spec: %q", capacityStr)
	}
	capacityUnit := subs[2]
	if capacityUnit == "" {
		return capacity, defaultVolumeCapacityUnit, nil
	}
	for _, item := range capacityUnits {
		if item == capacityUnit {
			return capacity, capacityUnit, nil
		}
	}
	return 0, "", fmt.Errorf("invalid capacity unit: %q", capacityUnit)
}

func init() {
	addFlexvolumeSource("qcow2", newQCOW2Volume)
}

// TODO: this file needs a test
