// +build !ignore_autogenerated

/*

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

// autogenerated by controller-gen object, do not modify manually

package v2alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthStrategy) DeepCopyInto(out *AuthStrategy) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthStrategy.
func (in *AuthStrategy) DeepCopy() *AuthStrategy {
	if in == nil {
		return nil
	}
	out := new(AuthStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gate) DeepCopyInto(out *Gate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gate.
func (in *Gate) DeepCopy() *Gate {
	if in == nil {
		return nil
	}
	out := new(Gate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GateList) DeepCopyInto(out *GateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GateList.
func (in *GateList) DeepCopy() *GateList {
	if in == nil {
		return nil
	}
	out := new(GateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GateSpec) DeepCopyInto(out *GateSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(Service)
		(*in).DeepCopyInto(*out)
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(AuthStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GateSpec.
func (in *GateSpec) DeepCopy() *GateSpec {
	if in == nil {
		return nil
	}
	out := new(GateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GateStatus) DeepCopyInto(out *GateStatus) {
	*out = *in
	if in.LastProcessedTime != nil {
		in, out := &in.LastProcessedTime, &out.LastProcessedTime
		*out = new(v1.Time)
		(*in).DeepCopyInto(*out)
	}
	if in.GateStatus != nil {
		in, out := &in.GateStatus, &out.GateStatus
		*out = new(GatewayResourceStatus)
		**out = **in
	}
	if in.VirtualServiceStatus != nil {
		in, out := &in.VirtualServiceStatus, &out.VirtualServiceStatus
		*out = new(GatewayResourceStatus)
		**out = **in
	}
	if in.PolicyServiceStatus != nil {
		in, out := &in.PolicyServiceStatus, &out.PolicyServiceStatus
		*out = new(GatewayResourceStatus)
		**out = **in
	}
	if in.AccessRuleStatus != nil {
		in, out := &in.AccessRuleStatus, &out.AccessRuleStatus
		*out = new(GatewayResourceStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GateStatus.
func (in *GateStatus) DeepCopy() *GateStatus {
	if in == nil {
		return nil
	}
	out := new(GateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayResourceStatus) DeepCopyInto(out *GatewayResourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayResourceStatus.
func (in *GatewayResourceStatus) DeepCopy() *GatewayResourceStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthModeConfig) DeepCopyInto(out *OauthModeConfig) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]Option, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthModeConfig.
func (in *OauthModeConfig) DeepCopy() *OauthModeConfig {
	if in == nil {
		return nil
	}
	out := new(OauthModeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Option) DeepCopyInto(out *Option) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Option.
func (in *Option) DeepCopy() *Option {
	if in == nil {
		return nil
	}
	out := new(Option)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.IsExternal != nil {
		in, out := &in.IsExternal, &out.IsExternal
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}
