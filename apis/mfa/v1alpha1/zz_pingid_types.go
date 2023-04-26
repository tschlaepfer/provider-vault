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

type PingidObservation struct {

	// Admin URL computed by Vault.
	AdminURL *string `json:"adminUrl,omitempty" tf:"admin_url,omitempty"`

	// Authenticator URL computed by Vault.
	AuthenticatorURL *string `json:"authenticatorUrl,omitempty" tf:"authenticator_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IDP URL computed by Vault.
	IdpURL *string `json:"idpUrl,omitempty" tf:"idp_url,omitempty"`

	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Name of the MFA method.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Namespace ID computed by Vault.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Org Alias computed by Vault.
	OrgAlias *string `json:"orgAlias,omitempty" tf:"org_alias,omitempty"`

	// A base64-encoded third-party settings file retrieved from PingID's configuration page.
	SettingsFileBase64 *string `json:"settingsFileBase64,omitempty" tf:"settings_file_base64,omitempty"`

	// Type of configuration computed by Vault.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// If set, enables use of PingID signature. Computed by Vault
	UseSignature *bool `json:"useSignature,omitempty" tf:"use_signature,omitempty"`

	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

type PingidParameters struct {

	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	// +kubebuilder:validation:Optional
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Name of the MFA method.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A base64-encoded third-party settings file retrieved from PingID's configuration page.
	// +kubebuilder:validation:Optional
	SettingsFileBase64 *string `json:"settingsFileBase64,omitempty" tf:"settings_file_base64,omitempty"`

	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	// +kubebuilder:validation:Optional
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

// PingidSpec defines the desired state of Pingid
type PingidSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PingidParameters `json:"forProvider"`
}

// PingidStatus defines the observed state of Pingid.
type PingidStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PingidObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pingid is the Schema for the Pingids API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Pingid struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.mountAccessor)",message="mountAccessor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.settingsFileBase64)",message="settingsFileBase64 is a required parameter"
	Spec   PingidSpec   `json:"spec"`
	Status PingidStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PingidList contains a list of Pingids
type PingidList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pingid `json:"items"`
}

// Repository type metadata.
var (
	Pingid_Kind             = "Pingid"
	Pingid_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pingid_Kind}.String()
	Pingid_KindAPIVersion   = Pingid_Kind + "." + CRDGroupVersion.String()
	Pingid_GroupVersionKind = CRDGroupVersion.WithKind(Pingid_Kind)
)

func init() {
	SchemeBuilder.Register(&Pingid{}, &PingidList{})
}