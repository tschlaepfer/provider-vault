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

type RateLimitInitParameters struct {

	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval *float64 `json:"blockInterval,omitempty" tf:"block_interval,omitempty"`

	// If set to true on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to true if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default.
	Inheritable *bool `json:"inheritable,omitempty" tf:"inheritable,omitempty"`

	// The duration in seconds to enforce rate limiting for.
	// The duration in seconds to enforce rate limiting for.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Name of the rate limit quota
	// The name of the quota.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example namespace1/ adds a quota to a full namespace,
	// namespace1/auth/userpass adds a quota to userpass in namespace1.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// auth/userpass to namespace1/auth/userpass moves this quota from being a global mount quota to
	// a namespace specific mount quota. Note, namespaces are supported in Enterprise only.
	// Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The rate must be positive.
	// The maximum number of requests at any given second to be allowed by the quota rule. The rate must be positive.
	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type RateLimitObservation struct {

	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	BlockInterval *float64 `json:"blockInterval,omitempty" tf:"block_interval,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If set to true on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to true if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default.
	Inheritable *bool `json:"inheritable,omitempty" tf:"inheritable,omitempty"`

	// The duration in seconds to enforce rate limiting for.
	// The duration in seconds to enforce rate limiting for.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Name of the rate limit quota
	// The name of the quota.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example namespace1/ adds a quota to a full namespace,
	// namespace1/auth/userpass adds a quota to userpass in namespace1.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// auth/userpass to namespace1/auth/userpass moves this quota from being a global mount quota to
	// a namespace specific mount quota. Note, namespaces are supported in Enterprise only.
	// Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The rate must be positive.
	// The maximum number of requests at any given second to be allowed by the quota rule. The rate must be positive.
	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type RateLimitParameters struct {

	// If set, when a client reaches a rate limit threshold, the client will
	// be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	// +kubebuilder:validation:Optional
	BlockInterval *float64 `json:"blockInterval,omitempty" tf:"block_interval,omitempty"`

	// If set to true on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace. The inheritable parameter cannot be set to true if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default.
	// +kubebuilder:validation:Optional
	Inheritable *bool `json:"inheritable,omitempty" tf:"inheritable,omitempty"`

	// The duration in seconds to enforce rate limiting for.
	// The duration in seconds to enforce rate limiting for.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Name of the rate limit quota
	// The name of the quota.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example namespace1/ adds a quota to a full namespace,
	// namespace1/auth/userpass adds a quota to userpass in namespace1.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// auth/userpass to namespace1/auth/userpass moves this quota from being a global mount quota to
	// a namespace specific mount quota. Note, namespaces are supported in Enterprise only.
	// Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The maximum number of requests at any given second to be allowed by the quota
	// rule. The rate must be positive.
	// The maximum number of requests at any given second to be allowed by the quota rule. The rate must be positive.
	// +kubebuilder:validation:Optional
	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// RateLimitSpec defines the desired state of RateLimit
type RateLimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RateLimitParameters `json:"forProvider"`
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
	InitProvider RateLimitInitParameters `json:"initProvider,omitempty"`
}

// RateLimitStatus defines the observed state of RateLimit.
type RateLimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RateLimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RateLimit is the Schema for the RateLimits API. Manage Rate Limit Quota
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type RateLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rate) || (has(self.initProvider) && has(self.initProvider.rate))",message="spec.forProvider.rate is a required parameter"
	Spec   RateLimitSpec   `json:"spec"`
	Status RateLimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RateLimitList contains a list of RateLimits
type RateLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateLimit `json:"items"`
}

// Repository type metadata.
var (
	RateLimit_Kind             = "RateLimit"
	RateLimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RateLimit_Kind}.String()
	RateLimit_KindAPIVersion   = RateLimit_Kind + "." + CRDGroupVersion.String()
	RateLimit_GroupVersionKind = CRDGroupVersion.WithKind(RateLimit_Kind)
)

func init() {
	SchemeBuilder.Register(&RateLimit{}, &RateLimitList{})
}
