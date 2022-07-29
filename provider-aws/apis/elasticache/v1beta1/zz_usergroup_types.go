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

type UserGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type UserGroupParameters struct {

	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Required
	Engine *string `json:"engine" tf:"engine,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UserIDRefs []v1.Reference `json:"userIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=User
	// +crossplane:generate:reference:refFieldName=UserIDRefs
	// +crossplane:generate:reference:selectorFieldName=UserIDSelector
	// +kubebuilder:validation:Optional
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`
}

// UserGroupSpec defines the desired state of UserGroup
type UserGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserGroupParameters `json:"forProvider"`
}

// UserGroupStatus defines the observed state of UserGroup.
type UserGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroup is the Schema for the UserGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserGroupSpec   `json:"spec"`
	Status            UserGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupList contains a list of UserGroups
type UserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserGroup `json:"items"`
}

// Repository type metadata.
var (
	UserGroup_Kind             = "UserGroup"
	UserGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserGroup_Kind}.String()
	UserGroup_KindAPIVersion   = UserGroup_Kind + "." + CRDGroupVersion.String()
	UserGroup_GroupVersionKind = CRDGroupVersion.WithKind(UserGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&UserGroup{}, &UserGroupList{})
}