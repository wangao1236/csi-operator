// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSI) DeepCopyInto(out *CSI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSI.
func (in *CSI) DeepCopy() *CSI {
	if in == nil {
		return nil
	}
	out := new(CSI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIComponent) DeepCopyInto(out *CSIComponent) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIComponent.
func (in *CSIComponent) DeepCopy() *CSIComponent {
	if in == nil {
		return nil
	}
	out := new(CSIComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSICondition) DeepCopyInto(out *CSICondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSICondition.
func (in *CSICondition) DeepCopy() *CSICondition {
	if in == nil {
		return nil
	}
	out := new(CSICondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIController) DeepCopyInto(out *CSIController) {
	*out = *in
	if in.Provisioner != nil {
		in, out := &in.Provisioner, &out.Provisioner
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Attacher != nil {
		in, out := &in.Attacher, &out.Attacher
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Resizer != nil {
		in, out := &in.Resizer, &out.Resizer
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshotter != nil {
		in, out := &in.Snapshotter, &out.Snapshotter
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterRegistrar != nil {
		in, out := &in.ClusterRegistrar, &out.ClusterRegistrar
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIController.
func (in *CSIController) DeepCopy() *CSIController {
	if in == nil {
		return nil
	}
	out := new(CSIController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIDriverTemplate) DeepCopyInto(out *CSIDriverTemplate) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIDriverTemplate.
func (in *CSIDriverTemplate) DeepCopy() *CSIDriverTemplate {
	if in == nil {
		return nil
	}
	out := new(CSIDriverTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIList) DeepCopyInto(out *CSIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CSI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIList.
func (in *CSIList) DeepCopy() *CSIList {
	if in == nil {
		return nil
	}
	out := new(CSIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSINode) DeepCopyInto(out *CSINode) {
	*out = *in
	if in.NodeRegistrar != nil {
		in, out := &in.NodeRegistrar, &out.NodeRegistrar
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(CSIComponent)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSINode.
func (in *CSINode) DeepCopy() *CSINode {
	if in == nil {
		return nil
	}
	out := new(CSINode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSISpec) DeepCopyInto(out *CSISpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DriverTemplate != nil {
		in, out := &in.DriverTemplate, &out.DriverTemplate
		*out = new(CSIDriverTemplate)
		(*in).DeepCopyInto(*out)
	}
	in.Node.DeepCopyInto(&out.Node)
	in.Controller.DeepCopyInto(&out.Controller)
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]corev1.Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]storagev1.StorageClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSISpec.
func (in *CSISpec) DeepCopy() *CSISpec {
	if in == nil {
		return nil
	}
	out := new(CSISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIStatus) DeepCopyInto(out *CSIStatus) {
	*out = *in
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]Generation, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CSICondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIStatus.
func (in *CSIStatus) DeepCopy() *CSIStatus {
	if in == nil {
		return nil
	}
	out := new(CSIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Generation) DeepCopyInto(out *Generation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Generation.
func (in *Generation) DeepCopy() *Generation {
	if in == nil {
		return nil
	}
	out := new(Generation)
	in.DeepCopyInto(out)
	return out
}
