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

type DelegationSetInitParameters struct {

	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`
}

type DelegationSetObservation struct {

	// The Amazon Resource Name (ARN) of the Delegation Set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The delegation set ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of authoritative name servers for the hosted zone
	// (effectively a list of NS records).
	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`

	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`
}

type DelegationSetParameters struct {

	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	// +kubebuilder:validation:Optional
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DelegationSetSpec defines the desired state of DelegationSet
type DelegationSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DelegationSetParameters `json:"forProvider"`
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
	InitProvider DelegationSetInitParameters `json:"initProvider,omitempty"`
}

// DelegationSetStatus defines the observed state of DelegationSet.
type DelegationSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DelegationSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DelegationSet is the Schema for the DelegationSets API. Provides a Route53 Delegation Set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DelegationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DelegationSetSpec   `json:"spec"`
	Status            DelegationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DelegationSetList contains a list of DelegationSets
type DelegationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DelegationSet `json:"items"`
}

// Repository type metadata.
var (
	DelegationSet_Kind             = "DelegationSet"
	DelegationSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DelegationSet_Kind}.String()
	DelegationSet_KindAPIVersion   = DelegationSet_Kind + "." + CRDGroupVersion.String()
	DelegationSet_GroupVersionKind = CRDGroupVersion.WithKind(DelegationSet_Kind)
)

func init() {
	SchemeBuilder.Register(&DelegationSet{}, &DelegationSetList{})
}