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

type TransitGatewayVPCAttachmentAccepterObservation struct {

	// Whether Appliance Mode support is enabled. Valid values: disable, enable.
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// Whether DNS support is enabled. Valid values: disable, enable.
	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	// EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether IPv6 support is enabled. Valid values: disable, enable.
	IPv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`

	// Identifiers of EC2 Subnets.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Identifier of EC2 VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Identifier of the AWS account that owns the EC2 VPC.
	VPCOwnerID *string `json:"vpcOwnerId,omitempty" tf:"vpc_owner_id,omitempty"`
}

type TransitGatewayVPCAttachmentAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the EC2 Transit Gateway Attachment to manage.
	// +crossplane:generate:reference:type=TransitGatewayVPCAttachment
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayVPCAttachment to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayVPCAttachment to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: true.
	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: true.
	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`
}

// TransitGatewayVPCAttachmentAccepterSpec defines the desired state of TransitGatewayVPCAttachmentAccepter
type TransitGatewayVPCAttachmentAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayVPCAttachmentAccepterParameters `json:"forProvider"`
}

// TransitGatewayVPCAttachmentAccepterStatus defines the observed state of TransitGatewayVPCAttachmentAccepter.
type TransitGatewayVPCAttachmentAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayVPCAttachmentAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachmentAccepter is the Schema for the TransitGatewayVPCAttachmentAccepters API. Manages the accepter's side of an EC2 Transit Gateway VPC Attachment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayVPCAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayVPCAttachmentAccepterSpec   `json:"spec"`
	Status            TransitGatewayVPCAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachmentAccepterList contains a list of TransitGatewayVPCAttachmentAccepters
type TransitGatewayVPCAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayVPCAttachmentAccepter `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayVPCAttachmentAccepter_Kind             = "TransitGatewayVPCAttachmentAccepter"
	TransitGatewayVPCAttachmentAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayVPCAttachmentAccepter_Kind}.String()
	TransitGatewayVPCAttachmentAccepter_KindAPIVersion   = TransitGatewayVPCAttachmentAccepter_Kind + "." + CRDGroupVersion.String()
	TransitGatewayVPCAttachmentAccepter_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayVPCAttachmentAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayVPCAttachmentAccepter{}, &TransitGatewayVPCAttachmentAccepterList{})
}