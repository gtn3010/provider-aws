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

type AdvancedEventSelectorFieldSelectorObservation struct {
}

type AdvancedEventSelectorFieldSelectorParameters struct {

	// A list of values that includes events that match the last few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	EndsWith []*string `json:"endsWith,omitempty" tf:"ends_with,omitempty"`

	// A list of values that includes events that match the exact value of the event record field specified as the value of field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.
	// +kubebuilder:validation:Optional
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// Specifies a field in an event record on which to filter events to be logged. You can specify only the following values: readOnly, eventSource, eventName, eventCategory, resources.type, resources.ARN.
	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// A list of values that excludes events that match the last few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	NotEndsWith []*string `json:"notEndsWith,omitempty" tf:"not_ends_with,omitempty"`

	// A list of values that excludes events that match the exact value of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`

	// A list of values that excludes events that match the first few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	NotStartsWith []*string `json:"notStartsWith,omitempty" tf:"not_starts_with,omitempty"`

	// A list of values that includes events that match the first few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	StartsWith []*string `json:"startsWith,omitempty" tf:"starts_with,omitempty"`
}

type EventDataStoreAdvancedEventSelectorObservation struct {
}

type EventDataStoreAdvancedEventSelectorParameters struct {

	// Specifies the selector statements in an advanced event selector. Fields documented below.
	// +kubebuilder:validation:Optional
	FieldSelector []AdvancedEventSelectorFieldSelectorParameters `json:"fieldSelector,omitempty" tf:"field_selector,omitempty"`

	// The name of the event data store.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EventDataStoreObservation struct {

	// ARN of the event data store.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name of the event data store.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventDataStoreParameters struct {

	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see Log events by using advanced event selectors in the CloudTrail User Guide.
	// +kubebuilder:validation:Optional
	AdvancedEventSelector []EventDataStoreAdvancedEventSelectorParameters `json:"advancedEventSelector,omitempty" tf:"advanced_event_selector,omitempty"`

	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: true.
	// +kubebuilder:validation:Optional
	MultiRegionEnabled *bool `json:"multiRegionEnabled,omitempty" tf:"multi_region_enabled,omitempty"`

	// The name of the event data store.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: false.
	// +kubebuilder:validation:Optional
	OrganizationEnabled *bool `json:"organizationEnabled,omitempty" tf:"organization_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: 2555.
	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: true.
	// +kubebuilder:validation:Optional
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty" tf:"termination_protection_enabled,omitempty"`
}

// EventDataStoreSpec defines the desired state of EventDataStore
type EventDataStoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventDataStoreParameters `json:"forProvider"`
}

// EventDataStoreStatus defines the observed state of EventDataStore.
type EventDataStoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventDataStoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventDataStore is the Schema for the EventDataStores API. Provides a CloudTrail Event Data Store resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventDataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventDataStoreSpec   `json:"spec"`
	Status            EventDataStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventDataStoreList contains a list of EventDataStores
type EventDataStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventDataStore `json:"items"`
}

// Repository type metadata.
var (
	EventDataStore_Kind             = "EventDataStore"
	EventDataStore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventDataStore_Kind}.String()
	EventDataStore_KindAPIVersion   = EventDataStore_Kind + "." + CRDGroupVersion.String()
	EventDataStore_GroupVersionKind = CRDGroupVersion.WithKind(EventDataStore_Kind)
)

func init() {
	SchemeBuilder.Register(&EventDataStore{}, &EventDataStoreList{})
}