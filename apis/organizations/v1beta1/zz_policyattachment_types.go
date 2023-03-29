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

type PolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyAttachmentParameters struct {

	// The unique identifier (ID) of the policy that you want to attach to the target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/organizations/v1beta1.Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a Policy in organizations to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a Policy in organizations to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// If set to true, destroy will not detach the policy and instead just remove the resource from state. This can be useful in situations where the attachment must be preserved to meet the AWS minimum requirement of 1 attached policy.
	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	// +kubebuilder:validation:Required
	TargetID *string `json:"targetId" tf:"target_id,omitempty"`
}

// PolicyAttachmentSpec defines the desired state of PolicyAttachment
type PolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyAttachmentParameters `json:"forProvider"`
}

// PolicyAttachmentStatus defines the observed state of PolicyAttachment.
type PolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyAttachment is the Schema for the PolicyAttachments API. Provides a resource to attach an AWS Organizations policy to an organization account, root, or unit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyAttachmentSpec   `json:"spec"`
	Status            PolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyAttachmentList contains a list of PolicyAttachments
type PolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	PolicyAttachment_Kind             = "PolicyAttachment"
	PolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyAttachment_Kind}.String()
	PolicyAttachment_KindAPIVersion   = PolicyAttachment_Kind + "." + CRDGroupVersion.String()
	PolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyAttachment{}, &PolicyAttachmentList{})
}