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

type SecretRoleObservation struct {

	// The mount path for the AD backend.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last time Vault rotated this service account's password.
	LastVaultRotation *string `json:"lastVaultRotation,omitempty" tf:"last_vault_rotation,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Last time Vault set this service account's password.
	PasswordLastSet *string `json:"passwordLastSet,omitempty" tf:"password_last_set,omitempty"`

	// Name of the role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The username/logon name for the service account with which this role will be associated.
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`

	// In seconds, the default password time-to-live.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type SecretRoleParameters struct {

	// The mount path for the AD backend.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The username/logon name for the service account with which this role will be associated.
	// +kubebuilder:validation:Optional
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`

	// In seconds, the default password time-to-live.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// SecretRoleSpec defines the desired state of SecretRole
type SecretRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretRoleParameters `json:"forProvider"`
}

// SecretRoleStatus defines the observed state of SecretRole.
type SecretRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRole is the Schema for the SecretRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.role)",message="role is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.serviceAccountName)",message="serviceAccountName is a required parameter"
	Spec   SecretRoleSpec   `json:"spec"`
	Status SecretRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRoleList contains a list of SecretRoles
type SecretRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretRole `json:"items"`
}

// Repository type metadata.
var (
	SecretRole_Kind             = "SecretRole"
	SecretRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretRole_Kind}.String()
	SecretRole_KindAPIVersion   = SecretRole_Kind + "." + CRDGroupVersion.String()
	SecretRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretRole{}, &SecretRoleList{})
}
