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

type GroupPoliciesObservation struct {

	// Should the resource manage policies exclusively? Beware of race conditions when disabling exclusive management
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// ID of the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Name of the group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Policies to be tied to the group.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type GroupPoliciesParameters struct {

	// Should the resource manage policies exclusively? Beware of race conditions when disabling exclusive management
	// +kubebuilder:validation:Optional
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// ID of the group.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Policies to be tied to the group.
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// GroupPoliciesSpec defines the desired state of GroupPolicies
type GroupPoliciesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupPoliciesParameters `json:"forProvider"`
}

// GroupPoliciesStatus defines the observed state of GroupPolicies.
type GroupPoliciesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupPoliciesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicies is the Schema for the GroupPoliciess API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type GroupPolicies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.groupId)",message="groupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.policies)",message="policies is a required parameter"
	Spec   GroupPoliciesSpec   `json:"spec"`
	Status GroupPoliciesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPoliciesList contains a list of GroupPoliciess
type GroupPoliciesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupPolicies `json:"items"`
}

// Repository type metadata.
var (
	GroupPolicies_Kind             = "GroupPolicies"
	GroupPolicies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupPolicies_Kind}.String()
	GroupPolicies_KindAPIVersion   = GroupPolicies_Kind + "." + CRDGroupVersion.String()
	GroupPolicies_GroupVersionKind = CRDGroupVersion.WithKind(GroupPolicies_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupPolicies{}, &GroupPoliciesList{})
}