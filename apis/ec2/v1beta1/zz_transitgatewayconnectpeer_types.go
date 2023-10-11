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

type TransitGatewayConnectPeerInitParameters struct {

	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BGPAsn *string `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks []*string `json:"insideCidrBlocks,omitempty" tf:"inside_cidr_blocks,omitempty"`

	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as transit_gateway_address
	PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as peer_address. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress *string `json:"transitGatewayAddress,omitempty" tf:"transit_gateway_address,omitempty"`
}

type TransitGatewayConnectPeerObservation struct {

	// EC2 Transit Gateway Connect Peer ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	BGPAsn *string `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// EC2 Transit Gateway Connect Peer identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	InsideCidrBlocks []*string `json:"insideCidrBlocks,omitempty" tf:"inside_cidr_blocks,omitempty"`

	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as transit_gateway_address
	PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as peer_address. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	TransitGatewayAddress *string `json:"transitGatewayAddress,omitempty" tf:"transit_gateway_address,omitempty"`

	// The Transit Gateway Connect
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`
}

type TransitGatewayConnectPeerParameters struct {

	// The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
	// +kubebuilder:validation:Optional
	BGPAsn *string `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
	// +kubebuilder:validation:Optional
	InsideCidrBlocks []*string `json:"insideCidrBlocks,omitempty" tf:"inside_cidr_blocks,omitempty"`

	// The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as transit_gateway_address
	// +kubebuilder:validation:Optional
	PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as peer_address. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
	// +kubebuilder:validation:Optional
	TransitGatewayAddress *string `json:"transitGatewayAddress,omitempty" tf:"transit_gateway_address,omitempty"`

	// The Transit Gateway Connect
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayConnect
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayConnect in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayConnect in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`
}

// TransitGatewayConnectPeerSpec defines the desired state of TransitGatewayConnectPeer
type TransitGatewayConnectPeerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayConnectPeerParameters `json:"forProvider"`
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
	InitProvider TransitGatewayConnectPeerInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayConnectPeerStatus defines the observed state of TransitGatewayConnectPeer.
type TransitGatewayConnectPeerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayConnectPeerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayConnectPeer is the Schema for the TransitGatewayConnectPeers API. Manages an EC2 Transit Gateway Connect Peer
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayConnectPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.insideCidrBlocks) || (has(self.initProvider) && has(self.initProvider.insideCidrBlocks))",message="spec.forProvider.insideCidrBlocks is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.peerAddress) || (has(self.initProvider) && has(self.initProvider.peerAddress))",message="spec.forProvider.peerAddress is a required parameter"
	Spec   TransitGatewayConnectPeerSpec   `json:"spec"`
	Status TransitGatewayConnectPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayConnectPeerList contains a list of TransitGatewayConnectPeers
type TransitGatewayConnectPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayConnectPeer `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayConnectPeer_Kind             = "TransitGatewayConnectPeer"
	TransitGatewayConnectPeer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayConnectPeer_Kind}.String()
	TransitGatewayConnectPeer_KindAPIVersion   = TransitGatewayConnectPeer_Kind + "." + CRDGroupVersion.String()
	TransitGatewayConnectPeer_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayConnectPeer_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayConnectPeer{}, &TransitGatewayConnectPeerList{})
}
