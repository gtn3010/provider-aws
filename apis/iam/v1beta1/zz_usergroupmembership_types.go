/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserGroupMembershipInitParameters struct {
}

type UserGroupMembershipObservation struct {

	// A list of IAM Groups to add the user to
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IAM User to add to groups
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type UserGroupMembershipParameters struct {

	// References to Group to populate groups.
	// +kubebuilder:validation:Optional
	GroupRefs []v1.Reference `json:"groupRefs,omitempty" tf:"-"`

	// Selector for a list of Group to populate groups.
	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// A list of IAM Groups to add the user to
	// +crossplane:generate:reference:type=Group
	// +crossplane:generate:reference:refFieldName=GroupRefs
	// +crossplane:generate:reference:selectorFieldName=GroupSelector
	// +kubebuilder:validation:Optional
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// The name of the IAM User to add to groups
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// UserGroupMembershipSpec defines the desired state of UserGroupMembership
type UserGroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserGroupMembershipParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider UserGroupMembershipInitParameters `json:"initProvider,omitempty"`
}

// UserGroupMembershipStatus defines the observed state of UserGroupMembership.
type UserGroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserGroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupMembership is the Schema for the UserGroupMemberships API. Provides a resource for adding an IAM User to IAM Groups without conflicting with itself.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserGroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserGroupMembershipSpec   `json:"spec"`
	Status            UserGroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupMembershipList contains a list of UserGroupMemberships
type UserGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserGroupMembership `json:"items"`
}

// Repository type metadata.
var (
	UserGroupMembership_Kind             = "UserGroupMembership"
	UserGroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserGroupMembership_Kind}.String()
	UserGroupMembership_KindAPIVersion   = UserGroupMembership_Kind + "." + CRDGroupVersion.String()
	UserGroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(UserGroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&UserGroupMembership{}, &UserGroupMembershipList{})
}
