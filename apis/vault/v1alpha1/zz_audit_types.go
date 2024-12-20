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

type AuditInitParameters struct {

	// Human-friendly description of the audit device.
	// Human-friendly description of the audit device.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Configuration options to pass to the audit device itself.
	// Configuration options to pass to the audit device itself.
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// The path to mount the audit device. This defaults to the type.
	// Path in which to enable the audit device.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Type of the audit device, such as 'file'.
	// Type of the audit device, such as 'file'.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuditObservation struct {

	// Human-friendly description of the audit device.
	// Human-friendly description of the audit device.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Configuration options to pass to the audit device itself.
	// Configuration options to pass to the audit device itself.
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// The path to mount the audit device. This defaults to the type.
	// Path in which to enable the audit device.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Type of the audit device, such as 'file'.
	// Type of the audit device, such as 'file'.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuditParameters struct {

	// Human-friendly description of the audit device.
	// Human-friendly description of the audit device.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	// +kubebuilder:validation:Optional
	Local *bool `json:"local,omitempty" tf:"local,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Configuration options to pass to the audit device itself.
	// Configuration options to pass to the audit device itself.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// The path to mount the audit device. This defaults to the type.
	// Path in which to enable the audit device.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Type of the audit device, such as 'file'.
	// Type of the audit device, such as 'file'.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AuditSpec defines the desired state of Audit
type AuditSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuditParameters `json:"forProvider"`
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
	InitProvider AuditInitParameters `json:"initProvider,omitempty"`
}

// AuditStatus defines the observed state of Audit.
type AuditStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuditObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Audit is the Schema for the Audits API. Writes audit backends for Vault
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Audit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.options) || (has(self.initProvider) && has(self.initProvider.options))",message="spec.forProvider.options is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   AuditSpec   `json:"spec"`
	Status AuditStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuditList contains a list of Audits
type AuditList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Audit `json:"items"`
}

// Repository type metadata.
var (
	Audit_Kind             = "Audit"
	Audit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Audit_Kind}.String()
	Audit_KindAPIVersion   = Audit_Kind + "." + CRDGroupVersion.String()
	Audit_GroupVersionKind = CRDGroupVersion.WithKind(Audit_Kind)
)

func init() {
	SchemeBuilder.Register(&Audit{}, &AuditList{})
}
