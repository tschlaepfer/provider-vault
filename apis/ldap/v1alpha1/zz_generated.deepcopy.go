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
func (in *AuthBackend) DeepCopyInto(out *AuthBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackend.
func (in *AuthBackend) DeepCopy() *AuthBackend {
	if in == nil {
		return nil
	}
	out := new(AuthBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroup) DeepCopyInto(out *AuthBackendGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroup.
func (in *AuthBackendGroup) DeepCopy() *AuthBackendGroup {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroupInitParameters) DeepCopyInto(out *AuthBackendGroupInitParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Groupname != nil {
		in, out := &in.Groupname, &out.Groupname
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroupInitParameters.
func (in *AuthBackendGroupInitParameters) DeepCopy() *AuthBackendGroupInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroupList) DeepCopyInto(out *AuthBackendGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroupList.
func (in *AuthBackendGroupList) DeepCopy() *AuthBackendGroupList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroupObservation) DeepCopyInto(out *AuthBackendGroupObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Groupname != nil {
		in, out := &in.Groupname, &out.Groupname
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
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroupObservation.
func (in *AuthBackendGroupObservation) DeepCopy() *AuthBackendGroupObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroupParameters) DeepCopyInto(out *AuthBackendGroupParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Groupname != nil {
		in, out := &in.Groupname, &out.Groupname
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroupParameters.
func (in *AuthBackendGroupParameters) DeepCopy() *AuthBackendGroupParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroupSpec) DeepCopyInto(out *AuthBackendGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroupSpec.
func (in *AuthBackendGroupSpec) DeepCopy() *AuthBackendGroupSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendGroupStatus) DeepCopyInto(out *AuthBackendGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendGroupStatus.
func (in *AuthBackendGroupStatus) DeepCopy() *AuthBackendGroupStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendInitParameters) DeepCopyInto(out *AuthBackendInitParameters) {
	*out = *in
	if in.Binddn != nil {
		in, out := &in.Binddn, &out.Binddn
		*out = new(string)
		**out = **in
	}
	if in.CaseSensitiveNames != nil {
		in, out := &in.CaseSensitiveNames, &out.CaseSensitiveNames
		*out = new(bool)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.ClientTLSCert != nil {
		in, out := &in.ClientTLSCert, &out.ClientTLSCert
		*out = new(string)
		**out = **in
	}
	if in.DenyNullBind != nil {
		in, out := &in.DenyNullBind, &out.DenyNullBind
		*out = new(bool)
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
	if in.Discoverdn != nil {
		in, out := &in.Discoverdn, &out.Discoverdn
		*out = new(bool)
		**out = **in
	}
	if in.Groupattr != nil {
		in, out := &in.Groupattr, &out.Groupattr
		*out = new(string)
		**out = **in
	}
	if in.Groupdn != nil {
		in, out := &in.Groupdn, &out.Groupdn
		*out = new(string)
		**out = **in
	}
	if in.Groupfilter != nil {
		in, out := &in.Groupfilter, &out.Groupfilter
		*out = new(string)
		**out = **in
	}
	if in.InsecureTLS != nil {
		in, out := &in.InsecureTLS, &out.InsecureTLS
		*out = new(bool)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(bool)
		**out = **in
	}
	if in.MaxPageSize != nil {
		in, out := &in.MaxPageSize, &out.MaxPageSize
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Starttls != nil {
		in, out := &in.Starttls, &out.Starttls
		*out = new(bool)
		**out = **in
	}
	if in.TLSMaxVersion != nil {
		in, out := &in.TLSMaxVersion, &out.TLSMaxVersion
		*out = new(string)
		**out = **in
	}
	if in.TLSMinVersion != nil {
		in, out := &in.TLSMinVersion, &out.TLSMinVersion
		*out = new(string)
		**out = **in
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenExplicitMaxTTL != nil {
		in, out := &in.TokenExplicitMaxTTL, &out.TokenExplicitMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenMaxTTL != nil {
		in, out := &in.TokenMaxTTL, &out.TokenMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenNoDefaultPolicy != nil {
		in, out := &in.TokenNoDefaultPolicy, &out.TokenNoDefaultPolicy
		*out = new(bool)
		**out = **in
	}
	if in.TokenNumUses != nil {
		in, out := &in.TokenNumUses, &out.TokenNumUses
		*out = new(float64)
		**out = **in
	}
	if in.TokenPeriod != nil {
		in, out := &in.TokenPeriod, &out.TokenPeriod
		*out = new(float64)
		**out = **in
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenTTL != nil {
		in, out := &in.TokenTTL, &out.TokenTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Upndomain != nil {
		in, out := &in.Upndomain, &out.Upndomain
		*out = new(string)
		**out = **in
	}
	if in.UseTokenGroups != nil {
		in, out := &in.UseTokenGroups, &out.UseTokenGroups
		*out = new(bool)
		**out = **in
	}
	if in.Userattr != nil {
		in, out := &in.Userattr, &out.Userattr
		*out = new(string)
		**out = **in
	}
	if in.Userdn != nil {
		in, out := &in.Userdn, &out.Userdn
		*out = new(string)
		**out = **in
	}
	if in.Userfilter != nil {
		in, out := &in.Userfilter, &out.Userfilter
		*out = new(string)
		**out = **in
	}
	if in.UsernameAsAlias != nil {
		in, out := &in.UsernameAsAlias, &out.UsernameAsAlias
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendInitParameters.
func (in *AuthBackendInitParameters) DeepCopy() *AuthBackendInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendList) DeepCopyInto(out *AuthBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendList.
func (in *AuthBackendList) DeepCopy() *AuthBackendList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendObservation) DeepCopyInto(out *AuthBackendObservation) {
	*out = *in
	if in.Accessor != nil {
		in, out := &in.Accessor, &out.Accessor
		*out = new(string)
		**out = **in
	}
	if in.Binddn != nil {
		in, out := &in.Binddn, &out.Binddn
		*out = new(string)
		**out = **in
	}
	if in.CaseSensitiveNames != nil {
		in, out := &in.CaseSensitiveNames, &out.CaseSensitiveNames
		*out = new(bool)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.ClientTLSCert != nil {
		in, out := &in.ClientTLSCert, &out.ClientTLSCert
		*out = new(string)
		**out = **in
	}
	if in.DenyNullBind != nil {
		in, out := &in.DenyNullBind, &out.DenyNullBind
		*out = new(bool)
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
	if in.Discoverdn != nil {
		in, out := &in.Discoverdn, &out.Discoverdn
		*out = new(bool)
		**out = **in
	}
	if in.Groupattr != nil {
		in, out := &in.Groupattr, &out.Groupattr
		*out = new(string)
		**out = **in
	}
	if in.Groupdn != nil {
		in, out := &in.Groupdn, &out.Groupdn
		*out = new(string)
		**out = **in
	}
	if in.Groupfilter != nil {
		in, out := &in.Groupfilter, &out.Groupfilter
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InsecureTLS != nil {
		in, out := &in.InsecureTLS, &out.InsecureTLS
		*out = new(bool)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(bool)
		**out = **in
	}
	if in.MaxPageSize != nil {
		in, out := &in.MaxPageSize, &out.MaxPageSize
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Starttls != nil {
		in, out := &in.Starttls, &out.Starttls
		*out = new(bool)
		**out = **in
	}
	if in.TLSMaxVersion != nil {
		in, out := &in.TLSMaxVersion, &out.TLSMaxVersion
		*out = new(string)
		**out = **in
	}
	if in.TLSMinVersion != nil {
		in, out := &in.TLSMinVersion, &out.TLSMinVersion
		*out = new(string)
		**out = **in
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenExplicitMaxTTL != nil {
		in, out := &in.TokenExplicitMaxTTL, &out.TokenExplicitMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenMaxTTL != nil {
		in, out := &in.TokenMaxTTL, &out.TokenMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenNoDefaultPolicy != nil {
		in, out := &in.TokenNoDefaultPolicy, &out.TokenNoDefaultPolicy
		*out = new(bool)
		**out = **in
	}
	if in.TokenNumUses != nil {
		in, out := &in.TokenNumUses, &out.TokenNumUses
		*out = new(float64)
		**out = **in
	}
	if in.TokenPeriod != nil {
		in, out := &in.TokenPeriod, &out.TokenPeriod
		*out = new(float64)
		**out = **in
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenTTL != nil {
		in, out := &in.TokenTTL, &out.TokenTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Upndomain != nil {
		in, out := &in.Upndomain, &out.Upndomain
		*out = new(string)
		**out = **in
	}
	if in.UseTokenGroups != nil {
		in, out := &in.UseTokenGroups, &out.UseTokenGroups
		*out = new(bool)
		**out = **in
	}
	if in.Userattr != nil {
		in, out := &in.Userattr, &out.Userattr
		*out = new(string)
		**out = **in
	}
	if in.Userdn != nil {
		in, out := &in.Userdn, &out.Userdn
		*out = new(string)
		**out = **in
	}
	if in.Userfilter != nil {
		in, out := &in.Userfilter, &out.Userfilter
		*out = new(string)
		**out = **in
	}
	if in.UsernameAsAlias != nil {
		in, out := &in.UsernameAsAlias, &out.UsernameAsAlias
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendObservation.
func (in *AuthBackendObservation) DeepCopy() *AuthBackendObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendParameters) DeepCopyInto(out *AuthBackendParameters) {
	*out = *in
	if in.Binddn != nil {
		in, out := &in.Binddn, &out.Binddn
		*out = new(string)
		**out = **in
	}
	if in.BindpassSecretRef != nil {
		in, out := &in.BindpassSecretRef, &out.BindpassSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CaseSensitiveNames != nil {
		in, out := &in.CaseSensitiveNames, &out.CaseSensitiveNames
		*out = new(bool)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.ClientTLSCert != nil {
		in, out := &in.ClientTLSCert, &out.ClientTLSCert
		*out = new(string)
		**out = **in
	}
	if in.ClientTLSKeySecretRef != nil {
		in, out := &in.ClientTLSKeySecretRef, &out.ClientTLSKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.DenyNullBind != nil {
		in, out := &in.DenyNullBind, &out.DenyNullBind
		*out = new(bool)
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
	if in.Discoverdn != nil {
		in, out := &in.Discoverdn, &out.Discoverdn
		*out = new(bool)
		**out = **in
	}
	if in.Groupattr != nil {
		in, out := &in.Groupattr, &out.Groupattr
		*out = new(string)
		**out = **in
	}
	if in.Groupdn != nil {
		in, out := &in.Groupdn, &out.Groupdn
		*out = new(string)
		**out = **in
	}
	if in.Groupfilter != nil {
		in, out := &in.Groupfilter, &out.Groupfilter
		*out = new(string)
		**out = **in
	}
	if in.InsecureTLS != nil {
		in, out := &in.InsecureTLS, &out.InsecureTLS
		*out = new(bool)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(bool)
		**out = **in
	}
	if in.MaxPageSize != nil {
		in, out := &in.MaxPageSize, &out.MaxPageSize
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Starttls != nil {
		in, out := &in.Starttls, &out.Starttls
		*out = new(bool)
		**out = **in
	}
	if in.TLSMaxVersion != nil {
		in, out := &in.TLSMaxVersion, &out.TLSMaxVersion
		*out = new(string)
		**out = **in
	}
	if in.TLSMinVersion != nil {
		in, out := &in.TLSMinVersion, &out.TLSMinVersion
		*out = new(string)
		**out = **in
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenExplicitMaxTTL != nil {
		in, out := &in.TokenExplicitMaxTTL, &out.TokenExplicitMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenMaxTTL != nil {
		in, out := &in.TokenMaxTTL, &out.TokenMaxTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenNoDefaultPolicy != nil {
		in, out := &in.TokenNoDefaultPolicy, &out.TokenNoDefaultPolicy
		*out = new(bool)
		**out = **in
	}
	if in.TokenNumUses != nil {
		in, out := &in.TokenNumUses, &out.TokenNumUses
		*out = new(float64)
		**out = **in
	}
	if in.TokenPeriod != nil {
		in, out := &in.TokenPeriod, &out.TokenPeriod
		*out = new(float64)
		**out = **in
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TokenTTL != nil {
		in, out := &in.TokenTTL, &out.TokenTTL
		*out = new(float64)
		**out = **in
	}
	if in.TokenType != nil {
		in, out := &in.TokenType, &out.TokenType
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Upndomain != nil {
		in, out := &in.Upndomain, &out.Upndomain
		*out = new(string)
		**out = **in
	}
	if in.UseTokenGroups != nil {
		in, out := &in.UseTokenGroups, &out.UseTokenGroups
		*out = new(bool)
		**out = **in
	}
	if in.Userattr != nil {
		in, out := &in.Userattr, &out.Userattr
		*out = new(string)
		**out = **in
	}
	if in.Userdn != nil {
		in, out := &in.Userdn, &out.Userdn
		*out = new(string)
		**out = **in
	}
	if in.Userfilter != nil {
		in, out := &in.Userfilter, &out.Userfilter
		*out = new(string)
		**out = **in
	}
	if in.UsernameAsAlias != nil {
		in, out := &in.UsernameAsAlias, &out.UsernameAsAlias
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendParameters.
func (in *AuthBackendParameters) DeepCopy() *AuthBackendParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendSpec) DeepCopyInto(out *AuthBackendSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendSpec.
func (in *AuthBackendSpec) DeepCopy() *AuthBackendSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendStatus) DeepCopyInto(out *AuthBackendStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendStatus.
func (in *AuthBackendStatus) DeepCopy() *AuthBackendStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUser) DeepCopyInto(out *AuthBackendUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUser.
func (in *AuthBackendUser) DeepCopy() *AuthBackendUser {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUserInitParameters) DeepCopyInto(out *AuthBackendUserInitParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUserInitParameters.
func (in *AuthBackendUserInitParameters) DeepCopy() *AuthBackendUserInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUserInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUserList) DeepCopyInto(out *AuthBackendUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUserList.
func (in *AuthBackendUserList) DeepCopy() *AuthBackendUserList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUserObservation) DeepCopyInto(out *AuthBackendUserObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUserObservation.
func (in *AuthBackendUserObservation) DeepCopy() *AuthBackendUserObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUserParameters) DeepCopyInto(out *AuthBackendUserParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUserParameters.
func (in *AuthBackendUserParameters) DeepCopy() *AuthBackendUserParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUserSpec) DeepCopyInto(out *AuthBackendUserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUserSpec.
func (in *AuthBackendUserSpec) DeepCopy() *AuthBackendUserSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendUserStatus) DeepCopyInto(out *AuthBackendUserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendUserStatus.
func (in *AuthBackendUserStatus) DeepCopy() *AuthBackendUserStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendUserStatus)
	in.DeepCopyInto(out)
	return out
}
