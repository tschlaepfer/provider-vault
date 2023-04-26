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

type AuthBackendConfigObservation struct {

	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The configured URL for the application registered in Azure Active Directory.
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type AuthBackendConfigParameters struct {

	// Unique name of the auth backend to configure.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are required.
	// +kubebuilder:validation:Optional
	ClientIDSecretRef *v1.SecretKeySelector `json:"clientIdSecretRef,omitempty" tf:"-"`

	// The client secret for credentials to query the Azure APIs
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`

	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The configured URL for the application registered in Azure Active Directory.
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// The tenant id for the Azure Active Directory organization.
	// +kubebuilder:validation:Optional
	TenantIDSecretRef v1.SecretKeySelector `json:"tenantIdSecretRef" tf:"-"`
}

// AuthBackendConfigSpec defines the desired state of AuthBackendConfig
type AuthBackendConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendConfigParameters `json:"forProvider"`
}

// AuthBackendConfigStatus defines the observed state of AuthBackendConfig.
type AuthBackendConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendConfig is the Schema for the AuthBackendConfigs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.resource)",message="resource is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.tenantIdSecretRef)",message="tenantIdSecretRef is a required parameter"
	Spec   AuthBackendConfigSpec   `json:"spec"`
	Status AuthBackendConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendConfigList contains a list of AuthBackendConfigs
type AuthBackendConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendConfig `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendConfig_Kind             = "AuthBackendConfig"
	AuthBackendConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendConfig_Kind}.String()
	AuthBackendConfig_KindAPIVersion   = AuthBackendConfig_Kind + "." + CRDGroupVersion.String()
	AuthBackendConfig_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendConfig{}, &AuthBackendConfigList{})
}