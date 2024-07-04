//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSnapshotHandles) DeepCopyInto(out *GroupSnapshotHandles) {
	*out = *in
	if in.VolumeSnapshotHandles != nil {
		in, out := &in.VolumeSnapshotHandles, &out.VolumeSnapshotHandles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSnapshotHandles.
func (in *GroupSnapshotHandles) DeepCopy() *GroupSnapshotHandles {
	if in == nil {
		return nil
	}
	out := new(GroupSnapshotHandles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCVolumeSnapshotPair) DeepCopyInto(out *PVCVolumeSnapshotPair) {
	*out = *in
	out.PersistentVolumeClaimRef = in.PersistentVolumeClaimRef
	out.VolumeSnapshotRef = in.VolumeSnapshotRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCVolumeSnapshotPair.
func (in *PVCVolumeSnapshotPair) DeepCopy() *PVCVolumeSnapshotPair {
	if in == nil {
		return nil
	}
	out := new(PVCVolumeSnapshotPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVVolumeSnapshotContentPair) DeepCopyInto(out *PVVolumeSnapshotContentPair) {
	*out = *in
	out.PersistentVolumeRef = in.PersistentVolumeRef
	out.VolumeSnapshotContentRef = in.VolumeSnapshotContentRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVVolumeSnapshotContentPair.
func (in *PVVolumeSnapshotContentPair) DeepCopy() *PVVolumeSnapshotContentPair {
	if in == nil {
		return nil
	}
	out := new(PVVolumeSnapshotContentPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshot) DeepCopyInto(out *VolumeGroupSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VolumeGroupSnapshotStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshot.
func (in *VolumeGroupSnapshot) DeepCopy() *VolumeGroupSnapshot {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeGroupSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotClass) DeepCopyInto(out *VolumeGroupSnapshotClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotClass.
func (in *VolumeGroupSnapshotClass) DeepCopy() *VolumeGroupSnapshotClass {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeGroupSnapshotClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotClassList) DeepCopyInto(out *VolumeGroupSnapshotClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeGroupSnapshotClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotClassList.
func (in *VolumeGroupSnapshotClassList) DeepCopy() *VolumeGroupSnapshotClassList {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeGroupSnapshotClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotContent) DeepCopyInto(out *VolumeGroupSnapshotContent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VolumeGroupSnapshotContentStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotContent.
func (in *VolumeGroupSnapshotContent) DeepCopy() *VolumeGroupSnapshotContent {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeGroupSnapshotContent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotContentList) DeepCopyInto(out *VolumeGroupSnapshotContentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeGroupSnapshotContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotContentList.
func (in *VolumeGroupSnapshotContentList) DeepCopy() *VolumeGroupSnapshotContentList {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotContentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeGroupSnapshotContentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotContentSource) DeepCopyInto(out *VolumeGroupSnapshotContentSource) {
	*out = *in
	if in.VolumeHandles != nil {
		in, out := &in.VolumeHandles, &out.VolumeHandles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GroupSnapshotHandles != nil {
		in, out := &in.GroupSnapshotHandles, &out.GroupSnapshotHandles
		*out = new(GroupSnapshotHandles)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotContentSource.
func (in *VolumeGroupSnapshotContentSource) DeepCopy() *VolumeGroupSnapshotContentSource {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotContentSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotContentSpec) DeepCopyInto(out *VolumeGroupSnapshotContentSpec) {
	*out = *in
	out.VolumeGroupSnapshotRef = in.VolumeGroupSnapshotRef
	if in.VolumeGroupSnapshotClassName != nil {
		in, out := &in.VolumeGroupSnapshotClassName, &out.VolumeGroupSnapshotClassName
		*out = new(string)
		**out = **in
	}
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotContentSpec.
func (in *VolumeGroupSnapshotContentSpec) DeepCopy() *VolumeGroupSnapshotContentSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotContentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotContentStatus) DeepCopyInto(out *VolumeGroupSnapshotContentStatus) {
	*out = *in
	if in.VolumeGroupSnapshotHandle != nil {
		in, out := &in.VolumeGroupSnapshotHandle, &out.VolumeGroupSnapshotHandle
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(int64)
		**out = **in
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(v1.VolumeSnapshotError)
		(*in).DeepCopyInto(*out)
	}
	if in.PVVolumeSnapshotContentList != nil {
		in, out := &in.PVVolumeSnapshotContentList, &out.PVVolumeSnapshotContentList
		*out = make([]PVVolumeSnapshotContentPair, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotContentStatus.
func (in *VolumeGroupSnapshotContentStatus) DeepCopy() *VolumeGroupSnapshotContentStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotContentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotList) DeepCopyInto(out *VolumeGroupSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeGroupSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotList.
func (in *VolumeGroupSnapshotList) DeepCopy() *VolumeGroupSnapshotList {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeGroupSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotSource) DeepCopyInto(out *VolumeGroupSnapshotSource) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeGroupSnapshotContentName != nil {
		in, out := &in.VolumeGroupSnapshotContentName, &out.VolumeGroupSnapshotContentName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotSource.
func (in *VolumeGroupSnapshotSource) DeepCopy() *VolumeGroupSnapshotSource {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotSpec) DeepCopyInto(out *VolumeGroupSnapshotSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.VolumeGroupSnapshotClassName != nil {
		in, out := &in.VolumeGroupSnapshotClassName, &out.VolumeGroupSnapshotClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotSpec.
func (in *VolumeGroupSnapshotSpec) DeepCopy() *VolumeGroupSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeGroupSnapshotStatus) DeepCopyInto(out *VolumeGroupSnapshotStatus) {
	*out = *in
	if in.BoundVolumeGroupSnapshotContentName != nil {
		in, out := &in.BoundVolumeGroupSnapshotContentName, &out.BoundVolumeGroupSnapshotContentName
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(v1.VolumeSnapshotError)
		(*in).DeepCopyInto(*out)
	}
	if in.PVCVolumeSnapshotRefList != nil {
		in, out := &in.PVCVolumeSnapshotRefList, &out.PVCVolumeSnapshotRefList
		*out = make([]PVCVolumeSnapshotPair, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeGroupSnapshotStatus.
func (in *VolumeGroupSnapshotStatus) DeepCopy() *VolumeGroupSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeGroupSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}
