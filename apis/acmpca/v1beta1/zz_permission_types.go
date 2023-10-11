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

type PermissionInitParameters struct {

	// Actions that the specified AWS service principal can use. These include IssueCertificate, GetCertificate, and ListPermissions. Note that in order for ACM to automatically rotate certificates issued by a PCA, it must be granted permission on all 3 actions, as per the example above.
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// ID of the calling account
	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`
}

type PermissionObservation struct {

	// Actions that the specified AWS service principal can use. These include IssueCertificate, GetCertificate, and ListPermissions. Note that in order for ACM to automatically rotate certificates issued by a PCA, it must be granted permission on all 3 actions, as per the example above.
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// ARN of the CA that grants the permissions.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IAM policy that is associated with the permission.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// ID of the calling account
	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`
}

type PermissionParameters struct {

	// Actions that the specified AWS service principal can use. These include IssueCertificate, GetCertificate, and ListPermissions. Note that in order for ACM to automatically rotate certificates issued by a PCA, it must be granted permission on all 3 actions, as per the example above.
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// ARN of the CA that grants the permissions.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/acmpca/v1beta1.CertificateAuthority
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// Reference to a CertificateAuthority in acmpca to populate certificateAuthorityArn.
	// +kubebuilder:validation:Optional
	CertificateAuthorityArnRef *v1.Reference `json:"certificateAuthorityArnRef,omitempty" tf:"-"`

	// Selector for a CertificateAuthority in acmpca to populate certificateAuthorityArn.
	// +kubebuilder:validation:Optional
	CertificateAuthorityArnSelector *v1.Selector `json:"certificateAuthorityArnSelector,omitempty" tf:"-"`

	// AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the calling account
	// +kubebuilder:validation:Optional
	SourceAccount *string `json:"sourceAccount,omitempty" tf:"source_account,omitempty"`
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

// Permission is the Schema for the Permissions API. Provides a resource to manage an AWS Certificate Manager Private Certificate Authorities Permission
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actions) || (has(self.initProvider) && has(self.initProvider.actions))",message="spec.forProvider.actions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principal) || (has(self.initProvider) && has(self.initProvider.principal))",message="spec.forProvider.principal is a required parameter"
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
