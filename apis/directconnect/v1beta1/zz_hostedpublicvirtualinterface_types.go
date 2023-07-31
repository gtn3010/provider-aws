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

type HostedPublicVirtualInterfaceInitParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AWS account that will own the new virtual interface.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// The VLAN ID.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type HostedPublicVirtualInterfaceObservation struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AWS account that will own the new virtual interface.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// The VLAN ID.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type HostedPublicVirtualInterfaceParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	// +kubebuilder:validation:Optional
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	// +kubebuilder:validation:Optional
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	// +kubebuilder:validation:Optional
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	// +crossplane:generate:reference:type=Connection
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The name for the virtual interface.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The AWS account that will own the new virtual interface.
	// +kubebuilder:validation:Optional
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of routes to be advertised to the AWS network in this region.
	// +kubebuilder:validation:Optional
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// The VLAN ID.
	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

// HostedPublicVirtualInterfaceSpec defines the desired state of HostedPublicVirtualInterface
type HostedPublicVirtualInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedPublicVirtualInterfaceParameters `json:"forProvider"`
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
	InitProvider HostedPublicVirtualInterfaceInitParameters `json:"initProvider,omitempty"`
}

// HostedPublicVirtualInterfaceStatus defines the observed state of HostedPublicVirtualInterface.
type HostedPublicVirtualInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedPublicVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPublicVirtualInterface is the Schema for the HostedPublicVirtualInterfaces API. Provides a Direct Connect hosted public virtual interface resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressFamily) || has(self.initProvider.addressFamily)",message="addressFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bgpAsn) || has(self.initProvider.bgpAsn)",message="bgpAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ownerAccountId) || has(self.initProvider.ownerAccountId)",message="ownerAccountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeFilterPrefixes) || has(self.initProvider.routeFilterPrefixes)",message="routeFilterPrefixes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vlan) || has(self.initProvider.vlan)",message="vlan is a required parameter"
	Spec   HostedPublicVirtualInterfaceSpec   `json:"spec"`
	Status HostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPublicVirtualInterfaceList contains a list of HostedPublicVirtualInterfaces
type HostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedPublicVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	HostedPublicVirtualInterface_Kind             = "HostedPublicVirtualInterface"
	HostedPublicVirtualInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedPublicVirtualInterface_Kind}.String()
	HostedPublicVirtualInterface_KindAPIVersion   = HostedPublicVirtualInterface_Kind + "." + CRDGroupVersion.String()
	HostedPublicVirtualInterface_GroupVersionKind = CRDGroupVersion.WithKind(HostedPublicVirtualInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedPublicVirtualInterface{}, &HostedPublicVirtualInterfaceList{})
}