//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackend) DeepCopyInto(out *CloudSecretBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackend.
func (in *CloudSecretBackend) DeepCopy() *CloudSecretBackend {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSecretBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackendInitParameters) DeepCopyInto(out *CloudSecretBackendInitParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BasePath != nil {
		in, out := &in.BasePath, &out.BasePath
		*out = new(string)
		**out = **in
	}
	if in.DefaultLeaseTTLSeconds != nil {
		in, out := &in.DefaultLeaseTTLSeconds, &out.DefaultLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.MaxLeaseTTLSeconds != nil {
		in, out := &in.MaxLeaseTTLSeconds, &out.MaxLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackendInitParameters.
func (in *CloudSecretBackendInitParameters) DeepCopy() *CloudSecretBackendInitParameters {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackendInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackendList) DeepCopyInto(out *CloudSecretBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudSecretBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackendList.
func (in *CloudSecretBackendList) DeepCopy() *CloudSecretBackendList {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSecretBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackendObservation) DeepCopyInto(out *CloudSecretBackendObservation) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BasePath != nil {
		in, out := &in.BasePath, &out.BasePath
		*out = new(string)
		**out = **in
	}
	if in.DefaultLeaseTTLSeconds != nil {
		in, out := &in.DefaultLeaseTTLSeconds, &out.DefaultLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxLeaseTTLSeconds != nil {
		in, out := &in.MaxLeaseTTLSeconds, &out.MaxLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackendObservation.
func (in *CloudSecretBackendObservation) DeepCopy() *CloudSecretBackendObservation {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackendObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackendParameters) DeepCopyInto(out *CloudSecretBackendParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BasePath != nil {
		in, out := &in.BasePath, &out.BasePath
		*out = new(string)
		**out = **in
	}
	if in.DefaultLeaseTTLSeconds != nil {
		in, out := &in.DefaultLeaseTTLSeconds, &out.DefaultLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.MaxLeaseTTLSeconds != nil {
		in, out := &in.MaxLeaseTTLSeconds, &out.MaxLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackendParameters.
func (in *CloudSecretBackendParameters) DeepCopy() *CloudSecretBackendParameters {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackendParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackendSpec) DeepCopyInto(out *CloudSecretBackendSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackendSpec.
func (in *CloudSecretBackendSpec) DeepCopy() *CloudSecretBackendSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretBackendStatus) DeepCopyInto(out *CloudSecretBackendStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretBackendStatus.
func (in *CloudSecretBackendStatus) DeepCopy() *CloudSecretBackendStatus {
	if in == nil {
		return nil
	}
	out := new(CloudSecretBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCreds) DeepCopyInto(out *CloudSecretCreds) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCreds.
func (in *CloudSecretCreds) DeepCopy() *CloudSecretCreds {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCreds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSecretCreds) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCredsInitParameters) DeepCopyInto(out *CloudSecretCredsInitParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCredsInitParameters.
func (in *CloudSecretCredsInitParameters) DeepCopy() *CloudSecretCredsInitParameters {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCredsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCredsList) DeepCopyInto(out *CloudSecretCredsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudSecretCreds, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCredsList.
func (in *CloudSecretCredsList) DeepCopy() *CloudSecretCredsList {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCredsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSecretCredsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCredsObservation) DeepCopyInto(out *CloudSecretCredsObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.TokenID != nil {
		in, out := &in.TokenID, &out.TokenID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCredsObservation.
func (in *CloudSecretCredsObservation) DeepCopy() *CloudSecretCredsObservation {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCredsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCredsParameters) DeepCopyInto(out *CloudSecretCredsParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCredsParameters.
func (in *CloudSecretCredsParameters) DeepCopy() *CloudSecretCredsParameters {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCredsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCredsSpec) DeepCopyInto(out *CloudSecretCredsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCredsSpec.
func (in *CloudSecretCredsSpec) DeepCopy() *CloudSecretCredsSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCredsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretCredsStatus) DeepCopyInto(out *CloudSecretCredsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretCredsStatus.
func (in *CloudSecretCredsStatus) DeepCopy() *CloudSecretCredsStatus {
	if in == nil {
		return nil
	}
	out := new(CloudSecretCredsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRole) DeepCopyInto(out *CloudSecretRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRole.
func (in *CloudSecretRole) DeepCopy() *CloudSecretRole {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSecretRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRoleInitParameters) DeepCopyInto(out *CloudSecretRoleInitParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRoleInitParameters.
func (in *CloudSecretRoleInitParameters) DeepCopy() *CloudSecretRoleInitParameters {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRoleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRoleList) DeepCopyInto(out *CloudSecretRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudSecretRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRoleList.
func (in *CloudSecretRoleList) DeepCopy() *CloudSecretRoleList {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSecretRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRoleObservation) DeepCopyInto(out *CloudSecretRoleObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRoleObservation.
func (in *CloudSecretRoleObservation) DeepCopy() *CloudSecretRoleObservation {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRoleParameters) DeepCopyInto(out *CloudSecretRoleParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(float64)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRoleParameters.
func (in *CloudSecretRoleParameters) DeepCopy() *CloudSecretRoleParameters {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRoleSpec) DeepCopyInto(out *CloudSecretRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRoleSpec.
func (in *CloudSecretRoleSpec) DeepCopy() *CloudSecretRoleSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSecretRoleStatus) DeepCopyInto(out *CloudSecretRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSecretRoleStatus.
func (in *CloudSecretRoleStatus) DeepCopy() *CloudSecretRoleStatus {
	if in == nil {
		return nil
	}
	out := new(CloudSecretRoleStatus)
	in.DeepCopyInto(out)
	return out
}
