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

type SerialConsoleAccessInitParameters struct {

	// Whether or not serial console access is enabled. Valid values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SerialConsoleAccessObservation struct {

	// Whether or not serial console access is enabled. Valid values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SerialConsoleAccessParameters struct {

	// Whether or not serial console access is enabled. Valid values are true or false. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// SerialConsoleAccessSpec defines the desired state of SerialConsoleAccess
type SerialConsoleAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SerialConsoleAccessParameters `json:"forProvider"`
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
	InitProvider SerialConsoleAccessInitParameters `json:"initProvider,omitempty"`
}

// SerialConsoleAccessStatus defines the observed state of SerialConsoleAccess.
type SerialConsoleAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SerialConsoleAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SerialConsoleAccess is the Schema for the SerialConsoleAccesss API. Manages whether serial console access is enabled for your AWS account in the current AWS region.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SerialConsoleAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SerialConsoleAccessSpec   `json:"spec"`
	Status            SerialConsoleAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SerialConsoleAccessList contains a list of SerialConsoleAccesss
type SerialConsoleAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SerialConsoleAccess `json:"items"`
}

// Repository type metadata.
var (
	SerialConsoleAccess_Kind             = "SerialConsoleAccess"
	SerialConsoleAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SerialConsoleAccess_Kind}.String()
	SerialConsoleAccess_KindAPIVersion   = SerialConsoleAccess_Kind + "." + CRDGroupVersion.String()
	SerialConsoleAccess_GroupVersionKind = CRDGroupVersion.WithKind(SerialConsoleAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&SerialConsoleAccess{}, &SerialConsoleAccessList{})
}
