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

type TransitGatewayMulticastGroupSourceInitParameters struct {

	// The IP address assigned to the transit gateway multicast group.
	GroupIPAddress *string `json:"groupIpAddress,omitempty" tf:"group_ip_address,omitempty"`
}

type TransitGatewayMulticastGroupSourceObservation struct {

	// The IP address assigned to the transit gateway multicast group.
	GroupIPAddress *string `json:"groupIpAddress,omitempty" tf:"group_ip_address,omitempty"`

	// EC2 Transit Gateway Multicast Group Member identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The group members' network interface ID to register with the transit gateway multicast group.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`
}

type TransitGatewayMulticastGroupSourceParameters struct {

	// The IP address assigned to the transit gateway multicast group.
	// +kubebuilder:validation:Optional
	GroupIPAddress *string `json:"groupIpAddress,omitempty" tf:"group_ip_address,omitempty"`

	// The group members' network interface ID to register with the transit gateway multicast group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface in ec2 to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayMulticastDomain
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`

	// Reference to a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDRef *v1.Reference `json:"transitGatewayMulticastDomainIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDSelector *v1.Selector `json:"transitGatewayMulticastDomainIdSelector,omitempty" tf:"-"`
}

// TransitGatewayMulticastGroupSourceSpec defines the desired state of TransitGatewayMulticastGroupSource
type TransitGatewayMulticastGroupSourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayMulticastGroupSourceParameters `json:"forProvider"`
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
	InitProvider TransitGatewayMulticastGroupSourceInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayMulticastGroupSourceStatus defines the observed state of TransitGatewayMulticastGroupSource.
type TransitGatewayMulticastGroupSourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayMulticastGroupSourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastGroupSource is the Schema for the TransitGatewayMulticastGroupSources API. Manages an EC2 Transit Gateway Multicast Group Source
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayMulticastGroupSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupIpAddress) || has(self.initProvider.groupIpAddress)",message="groupIpAddress is a required parameter"
	Spec   TransitGatewayMulticastGroupSourceSpec   `json:"spec"`
	Status TransitGatewayMulticastGroupSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastGroupSourceList contains a list of TransitGatewayMulticastGroupSources
type TransitGatewayMulticastGroupSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayMulticastGroupSource `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayMulticastGroupSource_Kind             = "TransitGatewayMulticastGroupSource"
	TransitGatewayMulticastGroupSource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayMulticastGroupSource_Kind}.String()
	TransitGatewayMulticastGroupSource_KindAPIVersion   = TransitGatewayMulticastGroupSource_Kind + "." + CRDGroupVersion.String()
	TransitGatewayMulticastGroupSource_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayMulticastGroupSource_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayMulticastGroupSource{}, &TransitGatewayMulticastGroupSourceList{})
}
