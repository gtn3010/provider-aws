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

type ConditionInitParameters struct {

	// Key for the condition. Valid values: aws:PrincipalOrgID.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of condition. Value values: StringEquals.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionObservation struct {

	// Key for the condition. Valid values: aws:PrincipalOrgID.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of condition. Value values: StringEquals.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value for the key.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionParameters struct {

	// Key for the condition. Valid values: aws:PrincipalOrgID.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of condition. Value values: StringEquals.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value for the key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/organizations/v1beta1.Organization
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Reference to a Organization in organizations to populate value.
	// +kubebuilder:validation:Optional
	ValueRef *v1.Reference `json:"valueRef,omitempty" tf:"-"`

	// Selector for a Organization in organizations to populate value.
	// +kubebuilder:validation:Optional
	ValueSelector *v1.Selector `json:"valueSelector,omitempty" tf:"-"`
}

type PermissionInitParameters struct {

	// The action that you are enabling the other account to perform. Defaults to events:PutEvents.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify * to permit any account to put events to your default event bus, optionally limited by condition.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// An identifier string for the external account that you are granting permissions to.
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`
}

type PermissionObservation struct {

	// The action that you are enabling the other account to perform. Defaults to events:PutEvents.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the default event bus.
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// The statement ID of the EventBridge permission.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify * to permit any account to put events to your default event bus, optionally limited by condition.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// An identifier string for the external account that you are granting permissions to.
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`
}

type PermissionParameters struct {

	// The action that you are enabling the other account to perform. Defaults to events:PutEvents.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// The name of the event bus to set the permissions on.
	// If you omit this, the permissions are set on the default event bus.
	// +crossplane:generate:reference:type=Bus
	// +kubebuilder:validation:Optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Reference to a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameRef *v1.Reference `json:"eventBusNameRef,omitempty" tf:"-"`

	// Selector for a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameSelector *v1.Selector `json:"eventBusNameSelector,omitempty" tf:"-"`

	// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify * to permit any account to put events to your default event bus, optionally limited by condition.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// An identifier string for the external account that you are granting permissions to.
	// +kubebuilder:validation:Optional
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`
}

// PermissionSpec defines the desired state of Permission
type PermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionParameters `json:"forProvider"`
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
	InitProvider PermissionInitParameters `json:"initProvider,omitempty"`
}

// PermissionStatus defines the observed state of Permission.
type PermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permission is the Schema for the Permissions API. Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principal) || has(self.initProvider.principal)",message="principal is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.statementId) || has(self.initProvider.statementId)",message="statementId is a required parameter"
	Spec   PermissionSpec   `json:"spec"`
	Status PermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionList contains a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permission `json:"items"`
}

// Repository type metadata.
var (
	Permission_Kind             = "Permission"
	Permission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permission_Kind}.String()
	Permission_KindAPIVersion   = Permission_Kind + "." + CRDGroupVersion.String()
	Permission_GroupVersionKind = CRDGroupVersion.WithKind(Permission_Kind)
)

func init() {
	SchemeBuilder.Register(&Permission{}, &PermissionList{})
}