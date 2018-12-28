// +build !ignore_autogenerated

/*
Copyright 2018 zhongwei.lzw@alibaba-inc.com.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronHorizontalPodAutoscaler) DeepCopyInto(out *CronHorizontalPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronHorizontalPodAutoscaler.
func (in *CronHorizontalPodAutoscaler) DeepCopy() *CronHorizontalPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(CronHorizontalPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronHorizontalPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronHorizontalPodAutoscalerList) DeepCopyInto(out *CronHorizontalPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CronHorizontalPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronHorizontalPodAutoscalerList.
func (in *CronHorizontalPodAutoscalerList) DeepCopy() *CronHorizontalPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(CronHorizontalPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronHorizontalPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronHorizontalPodAutoscalerSpec) DeepCopyInto(out *CronHorizontalPodAutoscalerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronHorizontalPodAutoscalerSpec.
func (in *CronHorizontalPodAutoscalerSpec) DeepCopy() *CronHorizontalPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(CronHorizontalPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronHorizontalPodAutoscalerStatus) DeepCopyInto(out *CronHorizontalPodAutoscalerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronHorizontalPodAutoscalerStatus.
func (in *CronHorizontalPodAutoscalerStatus) DeepCopy() *CronHorizontalPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(CronHorizontalPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}