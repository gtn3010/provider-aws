// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type TagInitParameters struct {

	// Tag name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagObservation struct {

	// Transfer Family resource identifier and key, separated by a comma (,)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Tag name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Amazon Resource Name (ARN) of the Transfer Family resource to tag.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagParameters struct {

	// Tag name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the Transfer Family resource to tag.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/transfer/v1beta1.Server
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Server in transfer to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Server in transfer to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`

	// Tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// TagSpec defines the desired state of Tag
type TagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagParameters `json:"forProvider"`
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
	InitProvider TagInitParameters `json:"initProvider,omitempty"`
}

// TagStatus defines the observed state of Tag.
type TagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tag is the Schema for the Tags API. Manages an individual Transfer Family resource tag
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   TagSpec   `json:"spec"`
	Status TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagList contains a list of Tags
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// Repository type metadata.
var (
	Tag_Kind             = "Tag"
	Tag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tag_Kind}.String()
	Tag_KindAPIVersion   = Tag_Kind + "." + CRDGroupVersion.String()
	Tag_GroupVersionKind = CRDGroupVersion.WithKind(Tag_Kind)
)

func init() {
	SchemeBuilder.Register(&Tag{}, &TagList{})
}
