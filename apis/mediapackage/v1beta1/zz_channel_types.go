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

type ChannelInitParameters struct {

	// A unique identifier describing the channel
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// A description of the channel
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ChannelObservation struct {

	// The ARN of the channel
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A unique identifier describing the channel
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// A description of the channel
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A single item list of HLS ingest information
	HlsIngest []HlsIngestObservation `json:"hlsIngest,omitempty" tf:"hls_ingest,omitempty"`

	// The same as channel_id
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ChannelParameters struct {

	// A unique identifier describing the channel
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// A description of the channel
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HlsIngestInitParameters struct {
}

type HlsIngestObservation struct {

	// A list of the ingest endpoints
	IngestEndpoints []IngestEndpointsObservation `json:"ingestEndpoints,omitempty" tf:"ingest_endpoints,omitempty"`
}

type HlsIngestParameters struct {
}

type IngestEndpointsInitParameters struct {
}

type IngestEndpointsObservation struct {

	// The password
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// The URL
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type IngestEndpointsParameters struct {
}

// ChannelSpec defines the desired state of Channel
type ChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChannelParameters `json:"forProvider"`
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
	InitProvider ChannelInitParameters `json:"initProvider,omitempty"`
}

// ChannelStatus defines the observed state of Channel.
type ChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Channel is the Schema for the Channels API. Provides an AWS Elemental MediaPackage Channel.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Channel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.channelId) || has(self.initProvider.channelId)",message="channelId is a required parameter"
	Spec   ChannelSpec   `json:"spec"`
	Status ChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelList contains a list of Channels
type ChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Channel `json:"items"`
}

// Repository type metadata.
var (
	Channel_Kind             = "Channel"
	Channel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Channel_Kind}.String()
	Channel_KindAPIVersion   = Channel_Kind + "." + CRDGroupVersion.String()
	Channel_GroupVersionKind = CRDGroupVersion.WithKind(Channel_Kind)
)

func init() {
	SchemeBuilder.Register(&Channel{}, &ChannelList{})
}
