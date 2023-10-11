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

type InvitationAccepterInitParameters struct {

	// The AWS account ID for the account that sent the invitation.
	AdministratorAccountID *string `json:"administratorAccountId,omitempty" tf:"administrator_account_id,omitempty"`
}

type InvitationAccepterObservation struct {

	// The AWS account ID for the account that sent the invitation.
	AdministratorAccountID *string `json:"administratorAccountId,omitempty" tf:"administrator_account_id,omitempty"`

	// The unique identifier (ID) of the macie invitation accepter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier for the invitation.
	InvitationID *string `json:"invitationId,omitempty" tf:"invitation_id,omitempty"`
}

type InvitationAccepterParameters struct {

	// The AWS account ID for the account that sent the invitation.
	// +kubebuilder:validation:Optional
	AdministratorAccountID *string `json:"administratorAccountId,omitempty" tf:"administrator_account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// InvitationAccepterSpec defines the desired state of InvitationAccepter
type InvitationAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InvitationAccepterParameters `json:"forProvider"`
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
	InitProvider InvitationAccepterInitParameters `json:"initProvider,omitempty"`
}

// InvitationAccepterStatus defines the observed state of InvitationAccepter.
type InvitationAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InvitationAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InvitationAccepter is the Schema for the InvitationAccepters API. Provides a resource to manage an Amazon Macie Invitation Accepter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InvitationAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.administratorAccountId) || (has(self.initProvider) && has(self.initProvider.administratorAccountId))",message="spec.forProvider.administratorAccountId is a required parameter"
	Spec   InvitationAccepterSpec   `json:"spec"`
	Status InvitationAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InvitationAccepterList contains a list of InvitationAccepters
type InvitationAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InvitationAccepter `json:"items"`
}

// Repository type metadata.
var (
	InvitationAccepter_Kind             = "InvitationAccepter"
	InvitationAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InvitationAccepter_Kind}.String()
	InvitationAccepter_KindAPIVersion   = InvitationAccepter_Kind + "." + CRDGroupVersion.String()
	InvitationAccepter_GroupVersionKind = CRDGroupVersion.WithKind(InvitationAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&InvitationAccepter{}, &InvitationAccepterList{})
}
