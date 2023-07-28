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

type UserInGroupInitParameters struct {
}

type UserInGroupObservation struct {

	// The name of the group to which the user is to be added.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The user pool ID of the user and group.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// The username of the user to be added to the group.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserInGroupParameters struct {

	// The name of the group to which the user is to be added.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// Reference to a UserGroup in cognitoidp to populate groupName.
	// +kubebuilder:validation:Optional
	GroupNameRef *v1.Reference `json:"groupNameRef,omitempty" tf:"-"`

	// Selector for a UserGroup in cognitoidp to populate groupName.
	// +kubebuilder:validation:Optional
	GroupNameSelector *v1.Selector `json:"groupNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The user pool ID of the user and group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`

	// The username of the user to be added to the group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.User
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User in cognitoidp to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User in cognitoidp to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

// UserInGroupSpec defines the desired state of UserInGroup
type UserInGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserInGroupParameters `json:"forProvider"`
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
	InitProvider UserInGroupInitParameters `json:"initProvider,omitempty"`
}

// UserInGroupStatus defines the observed state of UserInGroup.
type UserInGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserInGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserInGroup is the Schema for the UserInGroups API. Adds the specified user to the specified group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserInGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserInGroupSpec   `json:"spec"`
	Status            UserInGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserInGroupList contains a list of UserInGroups
type UserInGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserInGroup `json:"items"`
}

// Repository type metadata.
var (
	UserInGroup_Kind             = "UserInGroup"
	UserInGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserInGroup_Kind}.String()
	UserInGroup_KindAPIVersion   = UserInGroup_Kind + "." + CRDGroupVersion.String()
	UserInGroup_GroupVersionKind = CRDGroupVersion.WithKind(UserInGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&UserInGroup{}, &UserInGroupList{})
}
