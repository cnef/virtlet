// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package testing

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedTest) DeepCopyInto(out *EmbeddedTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Object == nil {
		out.Object = nil
	} else {
		out.Object = in.Object.DeepCopyObject()
	}
	if in.EmptyObject == nil {
		out.EmptyObject = nil
	} else {
		out.EmptyObject = in.EmptyObject.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedTest.
func (in *EmbeddedTest) DeepCopy() *EmbeddedTest {
	if in == nil {
		return nil
	}
	out := new(EmbeddedTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmbeddedTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedTestExternal) DeepCopyInto(out *EmbeddedTestExternal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Object.DeepCopyInto(&out.Object)
	in.EmptyObject.DeepCopyInto(&out.EmptyObject)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedTestExternal.
func (in *EmbeddedTestExternal) DeepCopy() *EmbeddedTestExternal {
	if in == nil {
		return nil
	}
	out := new(EmbeddedTestExternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmbeddedTestExternal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionA) DeepCopyInto(out *ExtensionA) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionA.
func (in *ExtensionA) DeepCopy() *ExtensionA {
	if in == nil {
		return nil
	}
	out := new(ExtensionA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtensionA) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionB) DeepCopyInto(out *ExtensionB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionB.
func (in *ExtensionB) DeepCopy() *ExtensionB {
	if in == nil {
		return nil
	}
	out := new(ExtensionB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExtensionB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalComplex) DeepCopyInto(out *ExternalComplex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalComplex.
func (in *ExternalComplex) DeepCopy() *ExternalComplex {
	if in == nil {
		return nil
	}
	out := new(ExternalComplex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalComplex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalExtensionType) DeepCopyInto(out *ExternalExtensionType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Extension.DeepCopyInto(&out.Extension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalExtensionType.
func (in *ExternalExtensionType) DeepCopy() *ExternalExtensionType {
	if in == nil {
		return nil
	}
	out := new(ExternalExtensionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalExtensionType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalInternalSame) DeepCopyInto(out *ExternalInternalSame) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	out.A = in.A
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalInternalSame.
func (in *ExternalInternalSame) DeepCopy() *ExternalInternalSame {
	if in == nil {
		return nil
	}
	out := new(ExternalInternalSame)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalInternalSame) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalOptionalExtensionType) DeepCopyInto(out *ExternalOptionalExtensionType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Extension.DeepCopyInto(&out.Extension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalOptionalExtensionType.
func (in *ExternalOptionalExtensionType) DeepCopy() *ExternalOptionalExtensionType {
	if in == nil {
		return nil
	}
	out := new(ExternalOptionalExtensionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalOptionalExtensionType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSimple) DeepCopyInto(out *ExternalSimple) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSimple.
func (in *ExternalSimple) DeepCopy() *ExternalSimple {
	if in == nil {
		return nil
	}
	out := new(ExternalSimple)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSimple) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalTestType1) DeepCopyInto(out *ExternalTestType1) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	if in.M != nil {
		in, out := &in.M, &out.M
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.N != nil {
		in, out := &in.N, &out.N
		*out = make(map[string]ExternalTestType2, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.O != nil {
		in, out := &in.O, &out.O
		if *in == nil {
			*out = nil
		} else {
			*out = new(ExternalTestType2)
			**out = **in
		}
	}
	if in.P != nil {
		in, out := &in.P, &out.P
		*out = make([]ExternalTestType2, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalTestType1.
func (in *ExternalTestType1) DeepCopy() *ExternalTestType1 {
	if in == nil {
		return nil
	}
	out := new(ExternalTestType1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalTestType1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalTestType2) DeepCopyInto(out *ExternalTestType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalTestType2.
func (in *ExternalTestType2) DeepCopy() *ExternalTestType2 {
	if in == nil {
		return nil
	}
	out := new(ExternalTestType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalTestType2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalComplex) DeepCopyInto(out *InternalComplex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalComplex.
func (in *InternalComplex) DeepCopy() *InternalComplex {
	if in == nil {
		return nil
	}
	out := new(InternalComplex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternalComplex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalExtensionType) DeepCopyInto(out *InternalExtensionType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Extension == nil {
		out.Extension = nil
	} else {
		out.Extension = in.Extension.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalExtensionType.
func (in *InternalExtensionType) DeepCopy() *InternalExtensionType {
	if in == nil {
		return nil
	}
	out := new(InternalExtensionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternalExtensionType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalOptionalExtensionType) DeepCopyInto(out *InternalOptionalExtensionType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Extension == nil {
		out.Extension = nil
	} else {
		out.Extension = in.Extension.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalOptionalExtensionType.
func (in *InternalOptionalExtensionType) DeepCopy() *InternalOptionalExtensionType {
	if in == nil {
		return nil
	}
	out := new(InternalOptionalExtensionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternalOptionalExtensionType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalSimple) DeepCopyInto(out *InternalSimple) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalSimple.
func (in *InternalSimple) DeepCopy() *InternalSimple {
	if in == nil {
		return nil
	}
	out := new(InternalSimple)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InternalSimple) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTest) DeepCopyInto(out *ObjectTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = (*in)[i].DeepCopyObject()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTest.
func (in *ObjectTest) DeepCopy() *ObjectTest {
	if in == nil {
		return nil
	}
	out := new(ObjectTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTestExternal) DeepCopyInto(out *ObjectTestExternal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTestExternal.
func (in *ObjectTestExternal) DeepCopy() *ObjectTestExternal {
	if in == nil {
		return nil
	}
	out := new(ObjectTestExternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectTestExternal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestType1) DeepCopyInto(out *TestType1) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	if in.M != nil {
		in, out := &in.M, &out.M
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.N != nil {
		in, out := &in.N, &out.N
		*out = make(map[string]TestType2, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.O != nil {
		in, out := &in.O, &out.O
		if *in == nil {
			*out = nil
		} else {
			*out = new(TestType2)
			**out = **in
		}
	}
	if in.P != nil {
		in, out := &in.P, &out.P
		*out = make([]TestType2, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestType1.
func (in *TestType1) DeepCopy() *TestType1 {
	if in == nil {
		return nil
	}
	out := new(TestType1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestType1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestType2) DeepCopyInto(out *TestType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestType2.
func (in *TestType2) DeepCopy() *TestType2 {
	if in == nil {
		return nil
	}
	out := new(TestType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestType2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnknownType) DeepCopyInto(out *UnknownType) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnknownType.
func (in *UnknownType) DeepCopy() *UnknownType {
	if in == nil {
		return nil
	}
	out := new(UnknownType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnknownType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnversionedType) DeepCopyInto(out *UnversionedType) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnversionedType.
func (in *UnversionedType) DeepCopy() *UnversionedType {
	if in == nil {
		return nil
	}
	out := new(UnversionedType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UnversionedType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
