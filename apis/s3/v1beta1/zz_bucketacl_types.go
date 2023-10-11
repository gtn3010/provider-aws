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

type AccessControlPolicyGrantInitParameters struct {

	// Configuration block for the person being granted permissions. See below.
	Grantee []GranteeInitParameters `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type AccessControlPolicyGrantObservation struct {

	// Configuration block for the person being granted permissions. See below.
	Grantee []GranteeObservation `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type AccessControlPolicyGrantParameters struct {

	// Configuration block for the person being granted permissions. See below.
	// +kubebuilder:validation:Optional
	Grantee []GranteeParameters `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`
}

type AccessControlPolicyInitParameters struct {

	// Set of grant configuration blocks. See below.
	Grant []AccessControlPolicyGrantInitParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// Configuration block of the bucket owner's display name and ID. See below.
	Owner []OwnerInitParameters `json:"owner,omitempty" tf:"owner,omitempty"`
}

type AccessControlPolicyObservation struct {

	// Set of grant configuration blocks. See below.
	Grant []AccessControlPolicyGrantObservation `json:"grant,omitempty" tf:"grant,omitempty"`

	// Configuration block of the bucket owner's display name and ID. See below.
	Owner []OwnerObservation `json:"owner,omitempty" tf:"owner,omitempty"`
}

type AccessControlPolicyParameters struct {

	// Set of grant configuration blocks. See below.
	// +kubebuilder:validation:Optional
	Grant []AccessControlPolicyGrantParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// Configuration block of the bucket owner's display name and ID. See below.
	// +kubebuilder:validation:Optional
	Owner []OwnerParameters `json:"owner" tf:"owner,omitempty"`
}

type BucketACLInitParameters struct {

	// Canned ACL to apply to the bucket.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Configuration block that sets the ACL permissions for an object per grantee. See below.
	AccessControlPolicy []AccessControlPolicyInitParameters `json:"accessControlPolicy,omitempty" tf:"access_control_policy,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`
}

type BucketACLObservation struct {

	// Canned ACL to apply to the bucket.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Configuration block that sets the ACL permissions for an object per grantee. See below.
	AccessControlPolicy []AccessControlPolicyObservation `json:"accessControlPolicy,omitempty" tf:"access_control_policy,omitempty"`

	// Name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket, expected_bucket_owner (if configured), and acl (if configured) separated by commas (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketACLParameters struct {

	// Canned ACL to apply to the bucket.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Configuration block that sets the ACL permissions for an object per grantee. See below.
	// +kubebuilder:validation:Optional
	AccessControlPolicy []AccessControlPolicyParameters `json:"accessControlPolicy,omitempty" tf:"access_control_policy,omitempty"`

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type GranteeInitParameters struct {

	// Email address of the grantee. See Regions and Endpoints for supported AWS regions where this argument can be specified.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// ID of the owner.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of grantee. Valid values: CanonicalUser, AmazonCustomerByEmail, Group.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URI of the grantee group.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GranteeObservation struct {

	// Display name of the owner.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Email address of the grantee. See Regions and Endpoints for supported AWS regions where this argument can be specified.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// ID of the owner.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of grantee. Valid values: CanonicalUser, AmazonCustomerByEmail, Group.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URI of the grantee group.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GranteeParameters struct {

	// Email address of the grantee. See Regions and Endpoints for supported AWS regions where this argument can be specified.
	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// ID of the owner.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of grantee. Valid values: CanonicalUser, AmazonCustomerByEmail, Group.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// URI of the grantee group.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type OwnerInitParameters struct {

	// Display name of the owner.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// ID of the owner.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OwnerObservation struct {

	// Display name of the owner.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// ID of the owner.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OwnerParameters struct {

	// Display name of the owner.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// ID of the owner.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

// BucketACLSpec defines the desired state of BucketACL
type BucketACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketACLParameters `json:"forProvider"`
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
	InitProvider BucketACLInitParameters `json:"initProvider,omitempty"`
}

// BucketACLStatus defines the observed state of BucketACL.
type BucketACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketACL is the Schema for the BucketACLs API. Provides an S3 bucket ACL resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketACLSpec   `json:"spec"`
	Status            BucketACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketACLList contains a list of BucketACLs
type BucketACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketACL `json:"items"`
}

// Repository type metadata.
var (
	BucketACL_Kind             = "BucketACL"
	BucketACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketACL_Kind}.String()
	BucketACL_KindAPIVersion   = BucketACL_Kind + "." + CRDGroupVersion.String()
	BucketACL_GroupVersionKind = CRDGroupVersion.WithKind(BucketACL_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketACL{}, &BucketACLList{})
}
