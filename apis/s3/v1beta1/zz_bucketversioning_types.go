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

type BucketVersioningInitParameters struct {

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// Configuration block for the versioning parameters. See below.
	VersioningConfiguration []VersioningConfigurationInitParameters `json:"versioningConfiguration,omitempty" tf:"versioning_configuration,omitempty"`
}

type BucketVersioningObservation struct {

	// Name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// Configuration block for the versioning parameters. See below.
	VersioningConfiguration []VersioningConfigurationObservation `json:"versioningConfiguration,omitempty" tf:"versioning_configuration,omitempty"`
}

type BucketVersioningParameters struct {

	// Name of the S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
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

	// Concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.
	// +kubebuilder:validation:Optional
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for the versioning parameters. See below.
	// +kubebuilder:validation:Optional
	VersioningConfiguration []VersioningConfigurationParameters `json:"versioningConfiguration,omitempty" tf:"versioning_configuration,omitempty"`
}

type VersioningConfigurationInitParameters struct {

	// Specifies whether MFA delete is enabled in the bucket versioning configuration. Valid values: Enabled or Disabled.
	MfaDelete *string `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`

	// Versioning state of the bucket. Valid values: Enabled, Suspended, or Disabled. Disabled should only be used when creating or importing resources that correspond to unversioned S3 buckets.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type VersioningConfigurationObservation struct {

	// Specifies whether MFA delete is enabled in the bucket versioning configuration. Valid values: Enabled or Disabled.
	MfaDelete *string `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`

	// Versioning state of the bucket. Valid values: Enabled, Suspended, or Disabled. Disabled should only be used when creating or importing resources that correspond to unversioned S3 buckets.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type VersioningConfigurationParameters struct {

	// Specifies whether MFA delete is enabled in the bucket versioning configuration. Valid values: Enabled or Disabled.
	// +kubebuilder:validation:Optional
	MfaDelete *string `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`

	// Versioning state of the bucket. Valid values: Enabled, Suspended, or Disabled. Disabled should only be used when creating or importing resources that correspond to unversioned S3 buckets.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// BucketVersioningSpec defines the desired state of BucketVersioning
type BucketVersioningSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketVersioningParameters `json:"forProvider"`
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
	InitProvider BucketVersioningInitParameters `json:"initProvider,omitempty"`
}

// BucketVersioningStatus defines the observed state of BucketVersioning.
type BucketVersioningStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketVersioningObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketVersioning is the Schema for the BucketVersionings API. Provides an S3 bucket versioning resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketVersioning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.versioningConfiguration) || has(self.initProvider.versioningConfiguration)",message="versioningConfiguration is a required parameter"
	Spec   BucketVersioningSpec   `json:"spec"`
	Status BucketVersioningStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketVersioningList contains a list of BucketVersionings
type BucketVersioningList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketVersioning `json:"items"`
}

// Repository type metadata.
var (
	BucketVersioning_Kind             = "BucketVersioning"
	BucketVersioning_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketVersioning_Kind}.String()
	BucketVersioning_KindAPIVersion   = BucketVersioning_Kind + "." + CRDGroupVersion.String()
	BucketVersioning_GroupVersionKind = CRDGroupVersion.WithKind(BucketVersioning_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketVersioning{}, &BucketVersioningList{})
}