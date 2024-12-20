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

type OidcKeyInitParameters struct {

	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	// Signing algorithm to use. Signing algorithm to use. Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// : Array of role client ID allowed to use this key for signing. If
	// empty, no roles are allowed. If ["*"], all roles are allowed.
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are allowed.
	// +listType=set
	AllowedClientIds []*string `json:"allowedClientIds,omitempty" tf:"allowed_client_ids,omitempty"`

	// Name of the OIDC Key to create.
	// Name of the key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// How often to generate a new signing key in number of seconds
	// How often to generate a new signing key in number of seconds
	RotationPeriod *float64 `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	// Controls how long the public portion of a signing key will be available for verification after being rotated in seconds.
	VerificationTTL *float64 `json:"verificationTtl,omitempty" tf:"verification_ttl,omitempty"`
}

type OidcKeyObservation struct {

	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	// Signing algorithm to use. Signing algorithm to use. Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// : Array of role client ID allowed to use this key for signing. If
	// empty, no roles are allowed. If ["*"], all roles are allowed.
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are allowed.
	// +listType=set
	AllowedClientIds []*string `json:"allowedClientIds,omitempty" tf:"allowed_client_ids,omitempty"`

	// The name of the created key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the OIDC Key to create.
	// Name of the key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// How often to generate a new signing key in number of seconds
	// How often to generate a new signing key in number of seconds
	RotationPeriod *float64 `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	// Controls how long the public portion of a signing key will be available for verification after being rotated in seconds.
	VerificationTTL *float64 `json:"verificationTtl,omitempty" tf:"verification_ttl,omitempty"`
}

type OidcKeyParameters struct {

	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	// Signing algorithm to use. Signing algorithm to use. Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// : Array of role client ID allowed to use this key for signing. If
	// empty, no roles are allowed. If ["*"], all roles are allowed.
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are allowed.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedClientIds []*string `json:"allowedClientIds,omitempty" tf:"allowed_client_ids,omitempty"`

	// Name of the OIDC Key to create.
	// Name of the key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// How often to generate a new signing key in number of seconds
	// How often to generate a new signing key in number of seconds
	// +kubebuilder:validation:Optional
	RotationPeriod *float64 `json:"rotationPeriod,omitempty" tf:"rotation_period,omitempty"`

	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	// Controls how long the public portion of a signing key will be available for verification after being rotated in seconds.
	// +kubebuilder:validation:Optional
	VerificationTTL *float64 `json:"verificationTtl,omitempty" tf:"verification_ttl,omitempty"`
}

// OidcKeySpec defines the desired state of OidcKey
type OidcKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OidcKeyParameters `json:"forProvider"`
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
	InitProvider OidcKeyInitParameters `json:"initProvider,omitempty"`
}

// OidcKeyStatus defines the observed state of OidcKey.
type OidcKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OidcKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OidcKey is the Schema for the OidcKeys API. Creates an Identity OIDC Named Key for Vault
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type OidcKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OidcKeySpec   `json:"spec"`
	Status OidcKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OidcKeyList contains a list of OidcKeys
type OidcKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OidcKey `json:"items"`
}

// Repository type metadata.
var (
	OidcKey_Kind             = "OidcKey"
	OidcKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OidcKey_Kind}.String()
	OidcKey_KindAPIVersion   = OidcKey_Kind + "." + CRDGroupVersion.String()
	OidcKey_GroupVersionKind = CRDGroupVersion.WithKind(OidcKey_Kind)
)

func init() {
	SchemeBuilder.Register(&OidcKey{}, &OidcKeyList{})
}
