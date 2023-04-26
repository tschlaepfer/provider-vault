//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendConfig) DeepCopyInto(out *AuthBackendConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendConfig.
func (in *AuthBackendConfig) DeepCopy() *AuthBackendConfig {
	if in == nil {
		return nil
	}
	out := new(AuthBackendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendConfigList) DeepCopyInto(out *AuthBackendConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendConfigList.
func (in *AuthBackendConfigList) DeepCopy() *AuthBackendConfigList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendConfigObservation) DeepCopyInto(out *AuthBackendConfigObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
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
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendConfigObservation.
func (in *AuthBackendConfigObservation) DeepCopy() *AuthBackendConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendConfigParameters) DeepCopyInto(out *AuthBackendConfigParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ClientIDSecretRef != nil {
		in, out := &in.ClientIDSecretRef, &out.ClientIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ClientSecretSecretRef != nil {
		in, out := &in.ClientSecretSecretRef, &out.ClientSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	out.TenantIDSecretRef = in.TenantIDSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendConfigParameters.
func (in *AuthBackendConfigParameters) DeepCopy() *AuthBackendConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendConfigSpec) DeepCopyInto(out *AuthBackendConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendConfigSpec.
func (in *AuthBackendConfigSpec) DeepCopy() *AuthBackendConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendConfigStatus) DeepCopyInto(out *AuthBackendConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendConfigStatus.
func (in *AuthBackendConfigStatus) DeepCopy() *AuthBackendConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRole) DeepCopyInto(out *AuthBackendRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRole.
func (in *AuthBackendRole) DeepCopy() *AuthBackendRole {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleList) DeepCopyInto(out *AuthBackendRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthBackendRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleList.
func (in *AuthBackendRoleList) DeepCopy() *AuthBackendRoleList {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthBackendRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleObservation) DeepCopyInto(out *AuthBackendRoleObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BoundGroupIds != nil {
		in, out := &in.BoundGroupIds, &out.BoundGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundLocations != nil {
		in, out := &in.BoundLocations, &out.BoundLocations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundResourceGroups != nil {
		in, out := &in.BoundResourceGroups, &out.BoundResourceGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundScaleSets != nil {
		in, out := &in.BoundScaleSets, &out.BoundScaleSets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundServicePrincipalIds != nil {
		in, out := &in.BoundServicePrincipalIds, &out.BoundServicePrincipalIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundSubscriptionIds != nil {
		in, out := &in.BoundSubscriptionIds, &out.BoundSubscriptionIds
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
	if in.Role != nil {
		in, out := &in.Role, &out.Role
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleObservation.
func (in *AuthBackendRoleObservation) DeepCopy() *AuthBackendRoleObservation {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleParameters) DeepCopyInto(out *AuthBackendRoleParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.BoundGroupIds != nil {
		in, out := &in.BoundGroupIds, &out.BoundGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundLocations != nil {
		in, out := &in.BoundLocations, &out.BoundLocations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundResourceGroups != nil {
		in, out := &in.BoundResourceGroups, &out.BoundResourceGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundScaleSets != nil {
		in, out := &in.BoundScaleSets, &out.BoundScaleSets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundServicePrincipalIds != nil {
		in, out := &in.BoundServicePrincipalIds, &out.BoundServicePrincipalIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BoundSubscriptionIds != nil {
		in, out := &in.BoundSubscriptionIds, &out.BoundSubscriptionIds
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
	if in.Role != nil {
		in, out := &in.Role, &out.Role
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleParameters.
func (in *AuthBackendRoleParameters) DeepCopy() *AuthBackendRoleParameters {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleSpec) DeepCopyInto(out *AuthBackendRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleSpec.
func (in *AuthBackendRoleSpec) DeepCopy() *AuthBackendRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthBackendRoleStatus) DeepCopyInto(out *AuthBackendRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthBackendRoleStatus.
func (in *AuthBackendRoleStatus) DeepCopy() *AuthBackendRoleStatus {
	if in == nil {
		return nil
	}
	out := new(AuthBackendRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureGroupsObservation) DeepCopyInto(out *AzureGroupsObservation) {
	*out = *in
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureGroupsObservation.
func (in *AzureGroupsObservation) DeepCopy() *AzureGroupsObservation {
	if in == nil {
		return nil
	}
	out := new(AzureGroupsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureGroupsParameters) DeepCopyInto(out *AzureGroupsParameters) {
	*out = *in
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureGroupsParameters.
func (in *AzureGroupsParameters) DeepCopy() *AzureGroupsParameters {
	if in == nil {
		return nil
	}
	out := new(AzureGroupsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureRolesObservation) DeepCopyInto(out *AzureRolesObservation) {
	*out = *in
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureRolesObservation.
func (in *AzureRolesObservation) DeepCopy() *AzureRolesObservation {
	if in == nil {
		return nil
	}
	out := new(AzureRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureRolesParameters) DeepCopyInto(out *AzureRolesParameters) {
	*out = *in
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureRolesParameters.
func (in *AzureRolesParameters) DeepCopy() *AzureRolesParameters {
	if in == nil {
		return nil
	}
	out := new(AzureRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackend) DeepCopyInto(out *SecretBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackend.
func (in *SecretBackend) DeepCopy() *SecretBackend {
	if in == nil {
		return nil
	}
	out := new(SecretBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendList) DeepCopyInto(out *SecretBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendList.
func (in *SecretBackendList) DeepCopy() *SecretBackendList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendObservation) DeepCopyInto(out *SecretBackendObservation) {
	*out = *in
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
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
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
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.UseMicrosoftGraphAPI != nil {
		in, out := &in.UseMicrosoftGraphAPI, &out.UseMicrosoftGraphAPI
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendObservation.
func (in *SecretBackendObservation) DeepCopy() *SecretBackendObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendParameters) DeepCopyInto(out *SecretBackendParameters) {
	*out = *in
	if in.ClientIDSecretRef != nil {
		in, out := &in.ClientIDSecretRef, &out.ClientIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ClientSecretSecretRef != nil {
		in, out := &in.ClientSecretSecretRef, &out.ClientSecretSecretRef
		*out = new(v1.SecretKeySelector)
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
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(string)
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
	out.SubscriptionIDSecretRef = in.SubscriptionIDSecretRef
	out.TenantIDSecretRef = in.TenantIDSecretRef
	if in.UseMicrosoftGraphAPI != nil {
		in, out := &in.UseMicrosoftGraphAPI, &out.UseMicrosoftGraphAPI
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendParameters.
func (in *SecretBackendParameters) DeepCopy() *SecretBackendParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRole) DeepCopyInto(out *SecretBackendRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRole.
func (in *SecretBackendRole) DeepCopy() *SecretBackendRole {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleList) DeepCopyInto(out *SecretBackendRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackendRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleList.
func (in *SecretBackendRoleList) DeepCopy() *SecretBackendRoleList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleObservation) DeepCopyInto(out *SecretBackendRoleObservation) {
	*out = *in
	if in.ApplicationObjectID != nil {
		in, out := &in.ApplicationObjectID, &out.ApplicationObjectID
		*out = new(string)
		**out = **in
	}
	if in.AzureGroups != nil {
		in, out := &in.AzureGroups, &out.AzureGroups
		*out = make([]AzureGroupsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureRoles != nil {
		in, out := &in.AzureRoles, &out.AzureRoles
		*out = make([]AzureRolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleObservation.
func (in *SecretBackendRoleObservation) DeepCopy() *SecretBackendRoleObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleParameters) DeepCopyInto(out *SecretBackendRoleParameters) {
	*out = *in
	if in.ApplicationObjectID != nil {
		in, out := &in.ApplicationObjectID, &out.ApplicationObjectID
		*out = new(string)
		**out = **in
	}
	if in.AzureGroups != nil {
		in, out := &in.AzureGroups, &out.AzureGroups
		*out = make([]AzureGroupsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AzureRoles != nil {
		in, out := &in.AzureRoles, &out.AzureRoles
		*out = make([]AzureRolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.MaxTTL != nil {
		in, out := &in.MaxTTL, &out.MaxTTL
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
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleParameters.
func (in *SecretBackendRoleParameters) DeepCopy() *SecretBackendRoleParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleSpec) DeepCopyInto(out *SecretBackendRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleSpec.
func (in *SecretBackendRoleSpec) DeepCopy() *SecretBackendRoleSpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleStatus) DeepCopyInto(out *SecretBackendRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleStatus.
func (in *SecretBackendRoleStatus) DeepCopy() *SecretBackendRoleStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendSpec) DeepCopyInto(out *SecretBackendSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendSpec.
func (in *SecretBackendSpec) DeepCopy() *SecretBackendSpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendStatus) DeepCopyInto(out *SecretBackendStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendStatus.
func (in *SecretBackendStatus) DeepCopy() *SecretBackendStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendStatus)
	in.DeepCopyInto(out)
	return out
}