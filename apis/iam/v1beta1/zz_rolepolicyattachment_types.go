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

type RolePolicyAttachmentInitParameters struct {
}

type RolePolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the policy you want to apply
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// The name of the IAM role to which the policy should be applied
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type RolePolicyAttachmentParameters struct {

	// The ARN of the policy you want to apply
	// +crossplane:generate:reference:type=Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// Reference to a Policy to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// Selector for a Policy to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// The name of the IAM role to which the policy should be applied
	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

// RolePolicyAttachmentSpec defines the desired state of RolePolicyAttachment
type RolePolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolePolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider RolePolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// RolePolicyAttachmentStatus defines the observed state of RolePolicyAttachment.
type RolePolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolePolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachment is the Schema for the RolePolicyAttachments API. Attaches a Managed IAM Policy to an IAM role
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RolePolicyAttachmentSpec   `json:"spec"`
	Status            RolePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachmentList contains a list of RolePolicyAttachments
type RolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	RolePolicyAttachment_Kind             = "RolePolicyAttachment"
	RolePolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RolePolicyAttachment_Kind}.String()
	RolePolicyAttachment_KindAPIVersion   = RolePolicyAttachment_Kind + "." + CRDGroupVersion.String()
	RolePolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RolePolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RolePolicyAttachment{}, &RolePolicyAttachmentList{})
}
