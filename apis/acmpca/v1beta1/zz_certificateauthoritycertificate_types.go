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

type CertificateAuthorityCertificateInitParameters struct {
}

type CertificateAuthorityCertificateObservation struct {

	// ARN of the Certificate Authority.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CertificateAuthorityCertificateParameters struct {

	// ARN of the Certificate Authority.
	// +crossplane:generate:reference:type=CertificateAuthority
	// +kubebuilder:validation:Optional
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// Reference to a CertificateAuthority to populate certificateAuthorityArn.
	// +kubebuilder:validation:Optional
	CertificateAuthorityArnRef *v1.Reference `json:"certificateAuthorityArnRef,omitempty" tf:"-"`

	// Selector for a CertificateAuthority to populate certificateAuthorityArn.
	// +kubebuilder:validation:Optional
	CertificateAuthorityArnSelector *v1.Selector `json:"certificateAuthorityArnSelector,omitempty" tf:"-"`

	// PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	// +kubebuilder:validation:Optional
	CertificateChainSecretRef *v1.SecretKeySelector `json:"certificateChainSecretRef,omitempty" tf:"-"`

	// PEM-encoded certificate for the Certificate Authority.
	// +kubebuilder:validation:Optional
	CertificateSecretRef v1.SecretKeySelector `json:"certificateSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// CertificateAuthorityCertificateSpec defines the desired state of CertificateAuthorityCertificate
type CertificateAuthorityCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateAuthorityCertificateParameters `json:"forProvider"`
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
	InitProvider CertificateAuthorityCertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateAuthorityCertificateStatus defines the observed state of CertificateAuthorityCertificate.
type CertificateAuthorityCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateAuthorityCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthorityCertificate is the Schema for the CertificateAuthorityCertificates API. Associates a certificate with an AWS Certificate Manager Private Certificate Authority
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CertificateAuthorityCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateSecretRef)",message="spec.forProvider.certificateSecretRef is a required parameter"
	Spec   CertificateAuthorityCertificateSpec   `json:"spec"`
	Status CertificateAuthorityCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthorityCertificateList contains a list of CertificateAuthorityCertificates
type CertificateAuthorityCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateAuthorityCertificate `json:"items"`
}

// Repository type metadata.
var (
	CertificateAuthorityCertificate_Kind             = "CertificateAuthorityCertificate"
	CertificateAuthorityCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateAuthorityCertificate_Kind}.String()
	CertificateAuthorityCertificate_KindAPIVersion   = CertificateAuthorityCertificate_Kind + "." + CRDGroupVersion.String()
	CertificateAuthorityCertificate_GroupVersionKind = CRDGroupVersion.WithKind(CertificateAuthorityCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateAuthorityCertificate{}, &CertificateAuthorityCertificateList{})
}
