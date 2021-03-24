// +build !ignore_autogenerated

/*
Copyright 2020 The Flux authors

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/fluxcd/pkg/apis/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommitSpec) DeepCopyInto(out *CommitSpec) {
	*out = *in
	if in.SigningKey != nil {
		in, out := &in.SigningKey, &out.SigningKey
		*out = new(SigningKey)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommitSpec.
func (in *CommitSpec) DeepCopy() *CommitSpec {
	if in == nil {
		return nil
	}
	out := new(CommitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitCheckoutSpec) DeepCopyInto(out *GitCheckoutSpec) {
	*out = *in
	out.GitRepositoryRef = in.GitRepositoryRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitCheckoutSpec.
func (in *GitCheckoutSpec) DeepCopy() *GitCheckoutSpec {
	if in == nil {
		return nil
	}
	out := new(GitCheckoutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomation) DeepCopyInto(out *ImageUpdateAutomation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomation.
func (in *ImageUpdateAutomation) DeepCopy() *ImageUpdateAutomation {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageUpdateAutomation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomationList) DeepCopyInto(out *ImageUpdateAutomationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageUpdateAutomation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomationList.
func (in *ImageUpdateAutomationList) DeepCopy() *ImageUpdateAutomationList {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageUpdateAutomationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomationSpec) DeepCopyInto(out *ImageUpdateAutomationSpec) {
	*out = *in
	out.Checkout = in.Checkout
	out.Interval = in.Interval
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		*out = new(UpdateStrategy)
		**out = **in
	}
	in.Commit.DeepCopyInto(&out.Commit)
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = new(PushSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomationSpec.
func (in *ImageUpdateAutomationSpec) DeepCopy() *ImageUpdateAutomationSpec {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageUpdateAutomationStatus) DeepCopyInto(out *ImageUpdateAutomationStatus) {
	*out = *in
	if in.LastAutomationRunTime != nil {
		in, out := &in.LastAutomationRunTime, &out.LastAutomationRunTime
		*out = (*in).DeepCopy()
	}
	if in.LastPushTime != nil {
		in, out := &in.LastPushTime, &out.LastPushTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageUpdateAutomationStatus.
func (in *ImageUpdateAutomationStatus) DeepCopy() *ImageUpdateAutomationStatus {
	if in == nil {
		return nil
	}
	out := new(ImageUpdateAutomationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSpec) DeepCopyInto(out *PushSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSpec.
func (in *PushSpec) DeepCopy() *PushSpec {
	if in == nil {
		return nil
	}
	out := new(PushSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningKey) DeepCopyInto(out *SigningKey) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningKey.
func (in *SigningKey) DeepCopy() *SigningKey {
	if in == nil {
		return nil
	}
	out := new(SigningKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateStrategy) DeepCopyInto(out *UpdateStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateStrategy.
func (in *UpdateStrategy) DeepCopy() *UpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(UpdateStrategy)
	in.DeepCopyInto(out)
	return out
}
