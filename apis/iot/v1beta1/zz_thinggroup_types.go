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

type AttributePayloadInitParameters struct {

	// Key-value map.
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`
}

type AttributePayloadObservation struct {

	// Key-value map.
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`
}

type AttributePayloadParameters struct {

	// Key-value map.
	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`
}

type MetadataInitParameters struct {
}

type MetadataObservation struct {
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// The name of the parent Thing Group.
	ParentGroupName *string `json:"parentGroupName,omitempty" tf:"parent_group_name,omitempty"`

	RootToParentGroups []RootToParentGroupsObservation `json:"rootToParentGroups,omitempty" tf:"root_to_parent_groups,omitempty"`
}

type MetadataParameters struct {
}

type PropertiesInitParameters struct {

	// The Thing Group attributes. Defined below.
	AttributePayload []AttributePayloadInitParameters `json:"attributePayload,omitempty" tf:"attribute_payload,omitempty"`

	// A description of the Thing Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type PropertiesObservation struct {

	// The Thing Group attributes. Defined below.
	AttributePayload []AttributePayloadObservation `json:"attributePayload,omitempty" tf:"attribute_payload,omitempty"`

	// A description of the Thing Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type PropertiesParameters struct {

	// The Thing Group attributes. Defined below.
	// +kubebuilder:validation:Optional
	AttributePayload []AttributePayloadParameters `json:"attributePayload,omitempty" tf:"attribute_payload,omitempty"`

	// A description of the Thing Group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type RootToParentGroupsInitParameters struct {
}

type RootToParentGroupsObservation struct {

	// The ARN of the Thing Group.
	GroupArn *string `json:"groupArn,omitempty" tf:"group_arn,omitempty"`

	// The name of the Thing Group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`
}

type RootToParentGroupsParameters struct {
}

type ThingGroupInitParameters struct {

	// The Thing Group properties. Defined below.
	Properties []PropertiesInitParameters `json:"properties,omitempty" tf:"properties,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ThingGroupObservation struct {

	// The ARN of the Thing Group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Thing Group ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the parent Thing Group.
	ParentGroupName *string `json:"parentGroupName,omitempty" tf:"parent_group_name,omitempty"`

	// The Thing Group properties. Defined below.
	Properties []PropertiesObservation `json:"properties,omitempty" tf:"properties,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The current version of the Thing Group record in the registry.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ThingGroupParameters struct {

	// The name of the parent Thing Group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iot/v1beta1.ThingGroup
	// +kubebuilder:validation:Optional
	ParentGroupName *string `json:"parentGroupName,omitempty" tf:"parent_group_name,omitempty"`

	// Reference to a ThingGroup in iot to populate parentGroupName.
	// +kubebuilder:validation:Optional
	ParentGroupNameRef *v1.Reference `json:"parentGroupNameRef,omitempty" tf:"-"`

	// Selector for a ThingGroup in iot to populate parentGroupName.
	// +kubebuilder:validation:Optional
	ParentGroupNameSelector *v1.Selector `json:"parentGroupNameSelector,omitempty" tf:"-"`

	// The Thing Group properties. Defined below.
	// +kubebuilder:validation:Optional
	Properties []PropertiesParameters `json:"properties,omitempty" tf:"properties,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ThingGroupSpec defines the desired state of ThingGroup
type ThingGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThingGroupParameters `json:"forProvider"`
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
	InitProvider ThingGroupInitParameters `json:"initProvider,omitempty"`
}

// ThingGroupStatus defines the observed state of ThingGroup.
type ThingGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThingGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ThingGroup is the Schema for the ThingGroups API. Manages an AWS IoT Thing Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ThingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ThingGroupSpec   `json:"spec"`
	Status            ThingGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingGroupList contains a list of ThingGroups
type ThingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThingGroup `json:"items"`
}

// Repository type metadata.
var (
	ThingGroup_Kind             = "ThingGroup"
	ThingGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ThingGroup_Kind}.String()
	ThingGroup_KindAPIVersion   = ThingGroup_Kind + "." + CRDGroupVersion.String()
	ThingGroup_GroupVersionKind = CRDGroupVersion.WithKind(ThingGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ThingGroup{}, &ThingGroupList{})
}