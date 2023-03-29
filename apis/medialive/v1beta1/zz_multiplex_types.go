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

type MultiplexMultiplexSettingsObservation struct {
}

type MultiplexMultiplexSettingsParameters struct {

	// Maximum video buffer delay.
	// +kubebuilder:validation:Optional
	MaximumVideoBufferDelayMilliseconds *float64 `json:"maximumVideoBufferDelayMilliseconds,omitempty" tf:"maximum_video_buffer_delay_milliseconds,omitempty"`

	// Transport stream bit rate.
	// +kubebuilder:validation:Required
	TransportStreamBitrate *float64 `json:"transportStreamBitrate" tf:"transport_stream_bitrate,omitempty"`

	// Unique ID for each multiplex.
	// +kubebuilder:validation:Required
	TransportStreamID *float64 `json:"transportStreamId" tf:"transport_stream_id,omitempty"`

	// Transport stream reserved bit rate.
	// +kubebuilder:validation:Optional
	TransportStreamReservedBitrate *float64 `json:"transportStreamReservedBitrate,omitempty" tf:"transport_stream_reserved_bitrate,omitempty"`
}

type MultiplexObservation struct {

	// ARN of the Multiplex.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type MultiplexParameters struct {

	// A list of availability zones. You must specify exactly two.
	// +kubebuilder:validation:Required
	AvailabilityZones []*string `json:"availabilityZones" tf:"availability_zones,omitempty"`

	// Multiplex settings. See Multiplex Settings for more details.
	// +kubebuilder:validation:Optional
	MultiplexSettings []MultiplexMultiplexSettingsParameters `json:"multiplexSettings,omitempty" tf:"multiplex_settings,omitempty"`

	// name of Multiplex.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether to start the Multiplex. Defaults to false.
	// +kubebuilder:validation:Optional
	StartMultiplex *bool `json:"startMultiplex,omitempty" tf:"start_multiplex,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MultiplexSpec defines the desired state of Multiplex
type MultiplexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MultiplexParameters `json:"forProvider"`
}

// MultiplexStatus defines the observed state of Multiplex.
type MultiplexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MultiplexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Multiplex is the Schema for the Multiplexs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Multiplex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MultiplexSpec   `json:"spec"`
	Status            MultiplexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MultiplexList contains a list of Multiplexs
type MultiplexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Multiplex `json:"items"`
}

// Repository type metadata.
var (
	Multiplex_Kind             = "Multiplex"
	Multiplex_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Multiplex_Kind}.String()
	Multiplex_KindAPIVersion   = Multiplex_Kind + "." + CRDGroupVersion.String()
	Multiplex_GroupVersionKind = CRDGroupVersion.WithKind(Multiplex_Kind)
)

func init() {
	SchemeBuilder.Register(&Multiplex{}, &MultiplexList{})
}