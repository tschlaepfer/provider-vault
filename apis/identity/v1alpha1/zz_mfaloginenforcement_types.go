/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MfaLoginEnforcementInitParameters struct {

	// Set of auth method accessor IDs.
	// Set of auth method accessor IDs.
	// +listType=set
	AuthMethodAccessors []*string `json:"authMethodAccessors,omitempty" tf:"auth_method_accessors,omitempty"`

	// Set of auth method types.
	// Set of auth method types.
	// +listType=set
	AuthMethodTypes []*string `json:"authMethodTypes,omitempty" tf:"auth_method_types,omitempty"`

	// Set of identity entity IDs.
	// Set of identity entity IDs.
	// +listType=set
	IdentityEntityIds []*string `json:"identityEntityIds,omitempty" tf:"identity_entity_ids,omitempty"`

	// Set of identity group IDs.
	// Set of identity group IDs.
	// +listType=set
	IdentityGroupIds []*string `json:"identityGroupIds,omitempty" tf:"identity_group_ids,omitempty"`

	// Set of MFA method UUIDs.
	// Set of MFA method UUIDs.
	// +listType=set
	MfaMethodIds []*string `json:"mfaMethodIds,omitempty" tf:"mfa_method_ids,omitempty"`

	// Login enforcement name.
	// Login enforcement name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type MfaLoginEnforcementObservation struct {

	// Set of auth method accessor IDs.
	// Set of auth method accessor IDs.
	// +listType=set
	AuthMethodAccessors []*string `json:"authMethodAccessors,omitempty" tf:"auth_method_accessors,omitempty"`

	// Set of auth method types.
	// Set of auth method types.
	// +listType=set
	AuthMethodTypes []*string `json:"authMethodTypes,omitempty" tf:"auth_method_types,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of identity entity IDs.
	// Set of identity entity IDs.
	// +listType=set
	IdentityEntityIds []*string `json:"identityEntityIds,omitempty" tf:"identity_entity_ids,omitempty"`

	// Set of identity group IDs.
	// Set of identity group IDs.
	// +listType=set
	IdentityGroupIds []*string `json:"identityGroupIds,omitempty" tf:"identity_group_ids,omitempty"`

	// Set of MFA method UUIDs.
	// Set of MFA method UUIDs.
	// +listType=set
	MfaMethodIds []*string `json:"mfaMethodIds,omitempty" tf:"mfa_method_ids,omitempty"`

	// Login enforcement name.
	// Login enforcement name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Method's namespace ID.
	// Method's namespace ID.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Method's namespace path.
	// Method's namespace path.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// Resource UUID.
	// Resource UUID.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type MfaLoginEnforcementParameters struct {

	// Set of auth method accessor IDs.
	// Set of auth method accessor IDs.
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthMethodAccessors []*string `json:"authMethodAccessors,omitempty" tf:"auth_method_accessors,omitempty"`

	// Set of auth method types.
	// Set of auth method types.
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthMethodTypes []*string `json:"authMethodTypes,omitempty" tf:"auth_method_types,omitempty"`

	// Set of identity entity IDs.
	// Set of identity entity IDs.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityEntityIds []*string `json:"identityEntityIds,omitempty" tf:"identity_entity_ids,omitempty"`

	// Set of identity group IDs.
	// Set of identity group IDs.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityGroupIds []*string `json:"identityGroupIds,omitempty" tf:"identity_group_ids,omitempty"`

	// Set of MFA method UUIDs.
	// Set of MFA method UUIDs.
	// +kubebuilder:validation:Optional
	// +listType=set
	MfaMethodIds []*string `json:"mfaMethodIds,omitempty" tf:"mfa_method_ids,omitempty"`

	// Login enforcement name.
	// Login enforcement name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// MfaLoginEnforcementSpec defines the desired state of MfaLoginEnforcement
type MfaLoginEnforcementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MfaLoginEnforcementParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MfaLoginEnforcementInitParameters `json:"initProvider,omitempty"`
}

// MfaLoginEnforcementStatus defines the observed state of MfaLoginEnforcement.
type MfaLoginEnforcementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MfaLoginEnforcementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MfaLoginEnforcement is the Schema for the MfaLoginEnforcements API. Resource for configuring MFA login-enforcement
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type MfaLoginEnforcement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mfaMethodIds) || (has(self.initProvider) && has(self.initProvider.mfaMethodIds))",message="spec.forProvider.mfaMethodIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MfaLoginEnforcementSpec   `json:"spec"`
	Status MfaLoginEnforcementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MfaLoginEnforcementList contains a list of MfaLoginEnforcements
type MfaLoginEnforcementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MfaLoginEnforcement `json:"items"`
}

// Repository type metadata.
var (
	MfaLoginEnforcement_Kind             = "MfaLoginEnforcement"
	MfaLoginEnforcement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MfaLoginEnforcement_Kind}.String()
	MfaLoginEnforcement_KindAPIVersion   = MfaLoginEnforcement_Kind + "." + CRDGroupVersion.String()
	MfaLoginEnforcement_GroupVersionKind = CRDGroupVersion.WithKind(MfaLoginEnforcement_Kind)
)

func init() {
	SchemeBuilder.Register(&MfaLoginEnforcement{}, &MfaLoginEnforcementList{})
}
