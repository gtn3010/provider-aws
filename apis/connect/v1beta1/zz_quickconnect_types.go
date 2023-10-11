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

type PhoneConfigInitParameters struct {

	// Specifies the phone number in in E.164 format.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type PhoneConfigObservation struct {

	// Specifies the phone number in in E.164 format.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type PhoneConfigParameters struct {

	// Specifies the phone number in in E.164 format.
	// +kubebuilder:validation:Optional
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

type QueueConfigInitParameters struct {

	// Specifies the identifier of the contact flow.
	ContactFlowID *string `json:"contactFlowId,omitempty" tf:"contact_flow_id,omitempty"`

	// Specifies the identifier for the queue.
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`
}

type QueueConfigObservation struct {

	// Specifies the identifier of the contact flow.
	ContactFlowID *string `json:"contactFlowId,omitempty" tf:"contact_flow_id,omitempty"`

	// Specifies the identifier for the queue.
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`
}

type QueueConfigParameters struct {

	// Specifies the identifier of the contact flow.
	// +kubebuilder:validation:Optional
	ContactFlowID *string `json:"contactFlowId" tf:"contact_flow_id,omitempty"`

	// Specifies the identifier for the queue.
	// +kubebuilder:validation:Optional
	QueueID *string `json:"queueId" tf:"queue_id,omitempty"`
}

type QuickConnectConfigInitParameters struct {

	// Specifies the phone configuration of the Quick Connect. This is required only if quick_connect_type is PHONE_NUMBER. The phone_config block is documented below.
	PhoneConfig []PhoneConfigInitParameters `json:"phoneConfig,omitempty" tf:"phone_config,omitempty"`

	// Specifies the queue configuration of the Quick Connect. This is required only if quick_connect_type is QUEUE. The queue_config block is documented below.
	QueueConfig []QueueConfigInitParameters `json:"queueConfig,omitempty" tf:"queue_config,omitempty"`

	// Specifies the configuration type of the quick connect. valid values are PHONE_NUMBER, QUEUE, USER.
	QuickConnectType *string `json:"quickConnectType,omitempty" tf:"quick_connect_type,omitempty"`

	// Specifies the user configuration of the Quick Connect. This is required only if quick_connect_type is USER. The user_config block is documented below.
	UserConfig []UserConfigInitParameters `json:"userConfig,omitempty" tf:"user_config,omitempty"`
}

type QuickConnectConfigObservation struct {

	// Specifies the phone configuration of the Quick Connect. This is required only if quick_connect_type is PHONE_NUMBER. The phone_config block is documented below.
	PhoneConfig []PhoneConfigObservation `json:"phoneConfig,omitempty" tf:"phone_config,omitempty"`

	// Specifies the queue configuration of the Quick Connect. This is required only if quick_connect_type is QUEUE. The queue_config block is documented below.
	QueueConfig []QueueConfigObservation `json:"queueConfig,omitempty" tf:"queue_config,omitempty"`

	// Specifies the configuration type of the quick connect. valid values are PHONE_NUMBER, QUEUE, USER.
	QuickConnectType *string `json:"quickConnectType,omitempty" tf:"quick_connect_type,omitempty"`

	// Specifies the user configuration of the Quick Connect. This is required only if quick_connect_type is USER. The user_config block is documented below.
	UserConfig []UserConfigObservation `json:"userConfig,omitempty" tf:"user_config,omitempty"`
}

type QuickConnectConfigParameters struct {

	// Specifies the phone configuration of the Quick Connect. This is required only if quick_connect_type is PHONE_NUMBER. The phone_config block is documented below.
	// +kubebuilder:validation:Optional
	PhoneConfig []PhoneConfigParameters `json:"phoneConfig,omitempty" tf:"phone_config,omitempty"`

	// Specifies the queue configuration of the Quick Connect. This is required only if quick_connect_type is QUEUE. The queue_config block is documented below.
	// +kubebuilder:validation:Optional
	QueueConfig []QueueConfigParameters `json:"queueConfig,omitempty" tf:"queue_config,omitempty"`

	// Specifies the configuration type of the quick connect. valid values are PHONE_NUMBER, QUEUE, USER.
	// +kubebuilder:validation:Optional
	QuickConnectType *string `json:"quickConnectType" tf:"quick_connect_type,omitempty"`

	// Specifies the user configuration of the Quick Connect. This is required only if quick_connect_type is USER. The user_config block is documented below.
	// +kubebuilder:validation:Optional
	UserConfig []UserConfigParameters `json:"userConfig,omitempty" tf:"user_config,omitempty"`
}

type QuickConnectInitParameters struct {

	// Specifies the description of the Quick Connect.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the Quick Connect.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that defines the configuration information for the Quick Connect: quick_connect_type and one of phone_config, queue_config, user_config . The Quick Connect Config block is documented below.
	QuickConnectConfig []QuickConnectConfigInitParameters `json:"quickConnectConfig,omitempty" tf:"quick_connect_config,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type QuickConnectObservation struct {

	// The Amazon Resource Name (ARN) of the Quick Connect.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies the description of the Quick Connect.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Quick Connect separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the name of the Quick Connect.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that defines the configuration information for the Quick Connect: quick_connect_type and one of phone_config, queue_config, user_config . The Quick Connect Config block is documented below.
	QuickConnectConfig []QuickConnectConfigObservation `json:"quickConnectConfig,omitempty" tf:"quick_connect_config,omitempty"`

	// The identifier for the Quick Connect.
	QuickConnectID *string `json:"quickConnectId,omitempty" tf:"quick_connect_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type QuickConnectParameters struct {

	// Specifies the description of the Quick Connect.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Quick Connect.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that defines the configuration information for the Quick Connect: quick_connect_type and one of phone_config, queue_config, user_config . The Quick Connect Config block is documented below.
	// +kubebuilder:validation:Optional
	QuickConnectConfig []QuickConnectConfigParameters `json:"quickConnectConfig,omitempty" tf:"quick_connect_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type UserConfigInitParameters struct {

	// Specifies the identifier of the contact flow.
	ContactFlowID *string `json:"contactFlowId,omitempty" tf:"contact_flow_id,omitempty"`

	// Specifies the identifier for the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserConfigObservation struct {

	// Specifies the identifier of the contact flow.
	ContactFlowID *string `json:"contactFlowId,omitempty" tf:"contact_flow_id,omitempty"`

	// Specifies the identifier for the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserConfigParameters struct {

	// Specifies the identifier of the contact flow.
	// +kubebuilder:validation:Optional
	ContactFlowID *string `json:"contactFlowId" tf:"contact_flow_id,omitempty"`

	// Specifies the identifier for the user.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId" tf:"user_id,omitempty"`
}

// QuickConnectSpec defines the desired state of QuickConnect
type QuickConnectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QuickConnectParameters `json:"forProvider"`
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
	InitProvider QuickConnectInitParameters `json:"initProvider,omitempty"`
}

// QuickConnectStatus defines the observed state of QuickConnect.
type QuickConnectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QuickConnectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QuickConnect is the Schema for the QuickConnects API. Provides details about a specific Amazon Quick Connect
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type QuickConnect struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.quickConnectConfig) || (has(self.initProvider) && has(self.initProvider.quickConnectConfig))",message="spec.forProvider.quickConnectConfig is a required parameter"
	Spec   QuickConnectSpec   `json:"spec"`
	Status QuickConnectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuickConnectList contains a list of QuickConnects
type QuickConnectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuickConnect `json:"items"`
}

// Repository type metadata.
var (
	QuickConnect_Kind             = "QuickConnect"
	QuickConnect_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QuickConnect_Kind}.String()
	QuickConnect_KindAPIVersion   = QuickConnect_Kind + "." + CRDGroupVersion.String()
	QuickConnect_GroupVersionKind = CRDGroupVersion.WithKind(QuickConnect_Kind)
)

func init() {
	SchemeBuilder.Register(&QuickConnect{}, &QuickConnectList{})
}
