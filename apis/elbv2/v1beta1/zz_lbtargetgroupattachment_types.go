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

type LBTargetGroupAttachmentInitParameters struct {

	// The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The port on which targets receive traffic.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda. If the target type is alb, specify the arn of alb.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`
}

type LBTargetGroupAttachmentObservation struct {

	// The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// A unique identifier for the attachment
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The port on which targets receive traffic.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ARN of the target group with which to register targets
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda. If the target type is alb, specify the arn of alb.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`
}

type LBTargetGroupAttachmentParameters struct {

	// The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The port on which targets receive traffic.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the target group with which to register targets
	// +crossplane:generate:reference:type=LBTargetGroup
	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// Reference to a LBTargetGroup to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnRef *v1.Reference `json:"targetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnSelector *v1.Selector `json:"targetGroupArnSelector,omitempty" tf:"-"`

	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda. If the target type is alb, specify the arn of alb.
	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`
}

// LBTargetGroupAttachmentSpec defines the desired state of LBTargetGroupAttachment
type LBTargetGroupAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBTargetGroupAttachmentParameters `json:"forProvider"`
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
	InitProvider LBTargetGroupAttachmentInitParameters `json:"initProvider,omitempty"`
}

// LBTargetGroupAttachmentStatus defines the observed state of LBTargetGroupAttachment.
type LBTargetGroupAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBTargetGroupAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBTargetGroupAttachment is the Schema for the LBTargetGroupAttachments API. Provides the ability to register instances and containers with a LB target group
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBTargetGroupAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetId) || has(self.initProvider.targetId)",message="targetId is a required parameter"
	Spec   LBTargetGroupAttachmentSpec   `json:"spec"`
	Status LBTargetGroupAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBTargetGroupAttachmentList contains a list of LBTargetGroupAttachments
type LBTargetGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBTargetGroupAttachment `json:"items"`
}

// Repository type metadata.
var (
	LBTargetGroupAttachment_Kind             = "LBTargetGroupAttachment"
	LBTargetGroupAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBTargetGroupAttachment_Kind}.String()
	LBTargetGroupAttachment_KindAPIVersion   = LBTargetGroupAttachment_Kind + "." + CRDGroupVersion.String()
	LBTargetGroupAttachment_GroupVersionKind = CRDGroupVersion.WithKind(LBTargetGroupAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&LBTargetGroupAttachment{}, &LBTargetGroupAttachmentList{})
}