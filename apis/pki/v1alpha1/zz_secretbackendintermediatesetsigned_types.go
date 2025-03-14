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

type SecretBackendIntermediateSetSignedInitParameters struct {

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/vault/v1alpha1.Mount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	// The certificate.
	// +crossplane:generate:reference:type=SecretBackendRootSignIntermediate
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-vault/config/common.ExtractCrt()
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Reference to a SecretBackendRootSignIntermediate to populate certificate.
	// +kubebuilder:validation:Optional
	CertificateRef *v1.Reference `json:"certificateRef,omitempty" tf:"-"`

	// Selector for a SecretBackendRootSignIntermediate to populate certificate.
	// +kubebuilder:validation:Optional
	CertificateSelector *v1.Selector `json:"certificateSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SecretBackendIntermediateSetSignedObservation struct {

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	// The certificate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The imported issuers indicating which issuers were created as part of
	// this request.
	// The imported issuers.
	ImportedIssuers []*string `json:"importedIssuers,omitempty" tf:"imported_issuers,omitempty"`

	// The imported keys indicating which keys were created as part of this request.
	// The imported keys.
	ImportedKeys []*string `json:"importedKeys,omitempty" tf:"imported_keys,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SecretBackendIntermediateSetSignedParameters struct {

	// The PKI secret backend the resource belongs to.
	// The PKI secret backend the resource belongs to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-vault/apis/vault/v1alpha1.Mount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("path",false)
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Mount in vault to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// Specifies the PEM encoded certificate. May optionally append additional
	// CA certificates to populate the whole chain, which will then enable returning the full chain from
	// issue and sign operations.
	// The certificate.
	// +crossplane:generate:reference:type=SecretBackendRootSignIntermediate
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-vault/config/common.ExtractCrt()
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Reference to a SecretBackendRootSignIntermediate to populate certificate.
	// +kubebuilder:validation:Optional
	CertificateRef *v1.Reference `json:"certificateRef,omitempty" tf:"-"`

	// Selector for a SecretBackendRootSignIntermediate to populate certificate.
	// +kubebuilder:validation:Optional
	CertificateSelector *v1.Selector `json:"certificateSelector,omitempty" tf:"-"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// SecretBackendIntermediateSetSignedSpec defines the desired state of SecretBackendIntermediateSetSigned
type SecretBackendIntermediateSetSignedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendIntermediateSetSignedParameters `json:"forProvider"`
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
	InitProvider SecretBackendIntermediateSetSignedInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendIntermediateSetSignedStatus defines the observed state of SecretBackendIntermediateSetSigned.
type SecretBackendIntermediateSetSignedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendIntermediateSetSignedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackendIntermediateSetSigned is the Schema for the SecretBackendIntermediateSetSigneds API. Submit the PKI CA certificate.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendIntermediateSetSigned struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretBackendIntermediateSetSignedSpec   `json:"spec"`
	Status            SecretBackendIntermediateSetSignedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendIntermediateSetSignedList contains a list of SecretBackendIntermediateSetSigneds
type SecretBackendIntermediateSetSignedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendIntermediateSetSigned `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendIntermediateSetSigned_Kind             = "SecretBackendIntermediateSetSigned"
	SecretBackendIntermediateSetSigned_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendIntermediateSetSigned_Kind}.String()
	SecretBackendIntermediateSetSigned_KindAPIVersion   = SecretBackendIntermediateSetSigned_Kind + "." + CRDGroupVersion.String()
	SecretBackendIntermediateSetSigned_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendIntermediateSetSigned_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendIntermediateSetSigned{}, &SecretBackendIntermediateSetSignedList{})
}
