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

type VPCIPv4CidrBlockAssociationInitParameters struct {

	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using ipv4_netmask_length.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	IPv4IpamPoolID *string `json:"ipv4IpamPoolId,omitempty" tf:"ipv4_ipam_pool_id,omitempty"`

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a ipv4_ipam_pool_id.
	IPv4NetmaskLength *float64 `json:"ipv4NetmaskLength,omitempty" tf:"ipv4_netmask_length,omitempty"`
}

type VPCIPv4CidrBlockAssociationObservation struct {

	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using ipv4_netmask_length.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The ID of the VPC CIDR association
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	IPv4IpamPoolID *string `json:"ipv4IpamPoolId,omitempty" tf:"ipv4_ipam_pool_id,omitempty"`

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a ipv4_ipam_pool_id.
	IPv4NetmaskLength *float64 `json:"ipv4NetmaskLength,omitempty" tf:"ipv4_netmask_length,omitempty"`

	// The ID of the VPC to make the association with.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCIPv4CidrBlockAssociationParameters struct {

	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using ipv4_netmask_length.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	// +kubebuilder:validation:Optional
	IPv4IpamPoolID *string `json:"ipv4IpamPoolId,omitempty" tf:"ipv4_ipam_pool_id,omitempty"`

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a ipv4_ipam_pool_id.
	// +kubebuilder:validation:Optional
	IPv4NetmaskLength *float64 `json:"ipv4NetmaskLength,omitempty" tf:"ipv4_netmask_length,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC to make the association with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCIPv4CidrBlockAssociationSpec defines the desired state of VPCIPv4CidrBlockAssociation
type VPCIPv4CidrBlockAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIPv4CidrBlockAssociationParameters `json:"forProvider"`
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
	InitProvider VPCIPv4CidrBlockAssociationInitParameters `json:"initProvider,omitempty"`
}

// VPCIPv4CidrBlockAssociationStatus defines the observed state of VPCIPv4CidrBlockAssociation.
type VPCIPv4CidrBlockAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIPv4CidrBlockAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIPv4CidrBlockAssociation is the Schema for the VPCIPv4CidrBlockAssociations API. Associate additional IPv4 CIDR blocks with a VPC
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIPv4CidrBlockAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCIPv4CidrBlockAssociationSpec   `json:"spec"`
	Status            VPCIPv4CidrBlockAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIPv4CidrBlockAssociationList contains a list of VPCIPv4CidrBlockAssociations
type VPCIPv4CidrBlockAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIPv4CidrBlockAssociation `json:"items"`
}

// Repository type metadata.
var (
	VPCIPv4CidrBlockAssociation_Kind             = "VPCIPv4CidrBlockAssociation"
	VPCIPv4CidrBlockAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIPv4CidrBlockAssociation_Kind}.String()
	VPCIPv4CidrBlockAssociation_KindAPIVersion   = VPCIPv4CidrBlockAssociation_Kind + "." + CRDGroupVersion.String()
	VPCIPv4CidrBlockAssociation_GroupVersionKind = CRDGroupVersion.WithKind(VPCIPv4CidrBlockAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIPv4CidrBlockAssociation{}, &VPCIPv4CidrBlockAssociationList{})
}
