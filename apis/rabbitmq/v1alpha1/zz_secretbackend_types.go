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

type SecretBackendInitParameters struct {

	// Specifies the RabbitMQ connection URI.
	// Specifies the RabbitMQ connection URI.
	ConnectionURI *string `json:"connectionUri,omitempty" tf:"connection_uri,omitempty"`

	// The default TTL for credentials
	// issued by this backend.
	// Default lease duration for secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	// Maximum possible lease duration for secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
	// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
	PasswordPolicy *string `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// Specifies the RabbitMQ management administrator password.
	// Specifies the RabbitMQ management administrator password
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The unique path this backend should be mounted at. Must
	// not begin or end with a /. Defaults to rabbitmq.
	// The path of the RabbitMQ Secret Backend where the connection should be configured
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the RabbitMQ management administrator username.
	// Specifies the RabbitMQ management administrator username
	UsernameSecretRef v1.SecretKeySelector `json:"usernameSecretRef" tf:"-"`

	// Template describing how dynamic usernames are generated.
	// Template describing how dynamic usernames are generated.
	UsernameTemplate *string `json:"usernameTemplate,omitempty" tf:"username_template,omitempty"`

	// Specifies whether to verify connection URI, username, and password.
	// Defaults to true.
	// Specifies whether to verify connection URI, username, and password.
	VerifyConnection *bool `json:"verifyConnection,omitempty" tf:"verify_connection,omitempty"`
}

type SecretBackendObservation struct {

	// Specifies the RabbitMQ connection URI.
	// Specifies the RabbitMQ connection URI.
	ConnectionURI *string `json:"connectionUri,omitempty" tf:"connection_uri,omitempty"`

	// The default TTL for credentials
	// issued by this backend.
	// Default lease duration for secrets in seconds
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	// Maximum possible lease duration for secrets in seconds
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
	// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
	PasswordPolicy *string `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// The unique path this backend should be mounted at. Must
	// not begin or end with a /. Defaults to rabbitmq.
	// The path of the RabbitMQ Secret Backend where the connection should be configured
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Template describing how dynamic usernames are generated.
	// Template describing how dynamic usernames are generated.
	UsernameTemplate *string `json:"usernameTemplate,omitempty" tf:"username_template,omitempty"`

	// Specifies whether to verify connection URI, username, and password.
	// Defaults to true.
	// Specifies whether to verify connection URI, username, and password.
	VerifyConnection *bool `json:"verifyConnection,omitempty" tf:"verify_connection,omitempty"`
}

type SecretBackendParameters struct {

	// Specifies the RabbitMQ connection URI.
	// Specifies the RabbitMQ connection URI.
	// +kubebuilder:validation:Optional
	ConnectionURI *string `json:"connectionUri,omitempty" tf:"connection_uri,omitempty"`

	// The default TTL for credentials
	// issued by this backend.
	// Default lease duration for secrets in seconds
	// +kubebuilder:validation:Optional
	DefaultLeaseTTLSeconds *float64 `json:"defaultLeaseTtlSeconds,omitempty" tf:"default_lease_ttl_seconds,omitempty"`

	// A human-friendly description for this backend.
	// Human-friendly description of the mount for the backend.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set, opts out of mount migration on path updates.
	// See here for more info on Mount Migration
	// If set, opts out of mount migration on path updates.
	// +kubebuilder:validation:Optional
	DisableRemount *bool `json:"disableRemount,omitempty" tf:"disable_remount,omitempty"`

	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	// Maximum possible lease duration for secrets in seconds
	// +kubebuilder:validation:Optional
	MaxLeaseTTLSeconds *float64 `json:"maxLeaseTtlSeconds,omitempty" tf:"max_lease_ttl_seconds,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
	// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
	// +kubebuilder:validation:Optional
	PasswordPolicy *string `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`

	// Specifies the RabbitMQ management administrator password.
	// Specifies the RabbitMQ management administrator password
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The unique path this backend should be mounted at. Must
	// not begin or end with a /. Defaults to rabbitmq.
	// The path of the RabbitMQ Secret Backend where the connection should be configured
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the RabbitMQ management administrator username.
	// Specifies the RabbitMQ management administrator username
	// +kubebuilder:validation:Optional
	UsernameSecretRef v1.SecretKeySelector `json:"usernameSecretRef" tf:"-"`

	// Template describing how dynamic usernames are generated.
	// Template describing how dynamic usernames are generated.
	// +kubebuilder:validation:Optional
	UsernameTemplate *string `json:"usernameTemplate,omitempty" tf:"username_template,omitempty"`

	// Specifies whether to verify connection URI, username, and password.
	// Defaults to true.
	// Specifies whether to verify connection URI, username, and password.
	// +kubebuilder:validation:Optional
	VerifyConnection *bool `json:"verifyConnection,omitempty" tf:"verify_connection,omitempty"`
}

// SecretBackendSpec defines the desired state of SecretBackend
type SecretBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendParameters `json:"forProvider"`
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
	InitProvider SecretBackendInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendStatus defines the observed state of SecretBackend.
type SecretBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackend is the Schema for the SecretBackends API. Creates an RabbitMQ secret backend for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionUri) || (has(self.initProvider) && has(self.initProvider.connectionUri))",message="spec.forProvider.connectionUri is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.usernameSecretRef)",message="spec.forProvider.usernameSecretRef is a required parameter"
	Spec   SecretBackendSpec   `json:"spec"`
	Status SecretBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendList contains a list of SecretBackends
type SecretBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackend `json:"items"`
}

// Repository type metadata.
var (
	SecretBackend_Kind             = "SecretBackend"
	SecretBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackend_Kind}.String()
	SecretBackend_KindAPIVersion   = SecretBackend_Kind + "." + CRDGroupVersion.String()
	SecretBackend_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackend{}, &SecretBackendList{})
}
