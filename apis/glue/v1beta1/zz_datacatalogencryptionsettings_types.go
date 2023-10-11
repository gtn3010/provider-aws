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

type ConnectionPasswordEncryptionInitParameters struct {

	// When set to true, passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently of the catalog encryption.
	ReturnConnectionPasswordEncrypted *bool `json:"returnConnectionPasswordEncrypted,omitempty" tf:"return_connection_password_encrypted,omitempty"`
}

type ConnectionPasswordEncryptionObservation struct {

	// A KMS key ARN that is used to encrypt the connection password. If connection password protection is enabled, the caller of CreateConnection and UpdateConnection needs at least kms:Encrypt permission on the specified AWS KMS key, to encrypt passwords before storing them in the Data Catalog.
	AwsKMSKeyID *string `json:"awsKmsKeyId,omitempty" tf:"aws_kms_key_id,omitempty"`

	// When set to true, passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently of the catalog encryption.
	ReturnConnectionPasswordEncrypted *bool `json:"returnConnectionPasswordEncrypted,omitempty" tf:"return_connection_password_encrypted,omitempty"`
}

type ConnectionPasswordEncryptionParameters struct {

	// A KMS key ARN that is used to encrypt the connection password. If connection password protection is enabled, the caller of CreateConnection and UpdateConnection needs at least kms:Encrypt permission on the specified AWS KMS key, to encrypt passwords before storing them in the Data Catalog.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	AwsKMSKeyID *string `json:"awsKmsKeyId,omitempty" tf:"aws_kms_key_id,omitempty"`

	// Reference to a Key in kms to populate awsKmsKeyId.
	// +kubebuilder:validation:Optional
	AwsKMSKeyIDRef *v1.Reference `json:"awsKmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate awsKmsKeyId.
	// +kubebuilder:validation:Optional
	AwsKMSKeyIDSelector *v1.Selector `json:"awsKmsKeyIdSelector,omitempty" tf:"-"`

	// When set to true, passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently of the catalog encryption.
	// +kubebuilder:validation:Optional
	ReturnConnectionPasswordEncrypted *bool `json:"returnConnectionPasswordEncrypted" tf:"return_connection_password_encrypted,omitempty"`
}

type DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsInitParameters struct {

	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption. see Connection Password Encryption.
	ConnectionPasswordEncryption []ConnectionPasswordEncryptionInitParameters `json:"connectionPasswordEncryption,omitempty" tf:"connection_password_encryption,omitempty"`

	// Specifies the encryption-at-rest configuration for the Data Catalog. see Encryption At Rest.
	EncryptionAtRest []EncryptionAtRestInitParameters `json:"encryptionAtRest,omitempty" tf:"encryption_at_rest,omitempty"`
}

type DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsObservation struct {

	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption. see Connection Password Encryption.
	ConnectionPasswordEncryption []ConnectionPasswordEncryptionObservation `json:"connectionPasswordEncryption,omitempty" tf:"connection_password_encryption,omitempty"`

	// Specifies the encryption-at-rest configuration for the Data Catalog. see Encryption At Rest.
	EncryptionAtRest []EncryptionAtRestObservation `json:"encryptionAtRest,omitempty" tf:"encryption_at_rest,omitempty"`
}

type DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsParameters struct {

	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption. see Connection Password Encryption.
	// +kubebuilder:validation:Optional
	ConnectionPasswordEncryption []ConnectionPasswordEncryptionParameters `json:"connectionPasswordEncryption" tf:"connection_password_encryption,omitempty"`

	// Specifies the encryption-at-rest configuration for the Data Catalog. see Encryption At Rest.
	// +kubebuilder:validation:Optional
	EncryptionAtRest []EncryptionAtRestParameters `json:"encryptionAtRest" tf:"encryption_at_rest,omitempty"`
}

type DataCatalogEncryptionSettingsInitParameters struct {

	// –  The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  The security configuration to set. see Data Catalog Encryption Settings.
	DataCatalogEncryptionSettings []DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsInitParameters `json:"dataCatalogEncryptionSettings,omitempty" tf:"data_catalog_encryption_settings,omitempty"`
}

type DataCatalogEncryptionSettingsObservation struct {

	// –  The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  The security configuration to set. see Data Catalog Encryption Settings.
	DataCatalogEncryptionSettings []DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsObservation `json:"dataCatalogEncryptionSettings,omitempty" tf:"data_catalog_encryption_settings,omitempty"`

	// The ID of the Data Catalog to set the security configuration for.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataCatalogEncryptionSettingsParameters struct {

	// –  The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  The security configuration to set. see Data Catalog Encryption Settings.
	// +kubebuilder:validation:Optional
	DataCatalogEncryptionSettings []DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsParameters `json:"dataCatalogEncryptionSettings,omitempty" tf:"data_catalog_encryption_settings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type EncryptionAtRestInitParameters struct {

	// The encryption-at-rest mode for encrypting Data Catalog data. Valid values are DISABLED and SSE-KMS.
	CatalogEncryptionMode *string `json:"catalogEncryptionMode,omitempty" tf:"catalog_encryption_mode,omitempty"`
}

type EncryptionAtRestObservation struct {

	// The encryption-at-rest mode for encrypting Data Catalog data. Valid values are DISABLED and SSE-KMS.
	CatalogEncryptionMode *string `json:"catalogEncryptionMode,omitempty" tf:"catalog_encryption_mode,omitempty"`

	// The ARN of the AWS KMS key to use for encryption at rest.
	SseAwsKMSKeyID *string `json:"sseAwsKmsKeyId,omitempty" tf:"sse_aws_kms_key_id,omitempty"`
}

type EncryptionAtRestParameters struct {

	// The encryption-at-rest mode for encrypting Data Catalog data. Valid values are DISABLED and SSE-KMS.
	// +kubebuilder:validation:Optional
	CatalogEncryptionMode *string `json:"catalogEncryptionMode" tf:"catalog_encryption_mode,omitempty"`

	// The ARN of the AWS KMS key to use for encryption at rest.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SseAwsKMSKeyID *string `json:"sseAwsKmsKeyId,omitempty" tf:"sse_aws_kms_key_id,omitempty"`

	// Reference to a Key in kms to populate sseAwsKmsKeyId.
	// +kubebuilder:validation:Optional
	SseAwsKMSKeyIDRef *v1.Reference `json:"sseAwsKmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate sseAwsKmsKeyId.
	// +kubebuilder:validation:Optional
	SseAwsKMSKeyIDSelector *v1.Selector `json:"sseAwsKmsKeyIdSelector,omitempty" tf:"-"`
}

// DataCatalogEncryptionSettingsSpec defines the desired state of DataCatalogEncryptionSettings
type DataCatalogEncryptionSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataCatalogEncryptionSettingsParameters `json:"forProvider"`
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
	InitProvider DataCatalogEncryptionSettingsInitParameters `json:"initProvider,omitempty"`
}

// DataCatalogEncryptionSettingsStatus defines the observed state of DataCatalogEncryptionSettings.
type DataCatalogEncryptionSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataCatalogEncryptionSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataCatalogEncryptionSettings is the Schema for the DataCatalogEncryptionSettingss API. Provides a Glue Data Catalog Encryption Settings resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DataCatalogEncryptionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataCatalogEncryptionSettings) || (has(self.initProvider) && has(self.initProvider.dataCatalogEncryptionSettings))",message="spec.forProvider.dataCatalogEncryptionSettings is a required parameter"
	Spec   DataCatalogEncryptionSettingsSpec   `json:"spec"`
	Status DataCatalogEncryptionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataCatalogEncryptionSettingsList contains a list of DataCatalogEncryptionSettingss
type DataCatalogEncryptionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogEncryptionSettings `json:"items"`
}

// Repository type metadata.
var (
	DataCatalogEncryptionSettings_Kind             = "DataCatalogEncryptionSettings"
	DataCatalogEncryptionSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataCatalogEncryptionSettings_Kind}.String()
	DataCatalogEncryptionSettings_KindAPIVersion   = DataCatalogEncryptionSettings_Kind + "." + CRDGroupVersion.String()
	DataCatalogEncryptionSettings_GroupVersionKind = CRDGroupVersion.WithKind(DataCatalogEncryptionSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&DataCatalogEncryptionSettings{}, &DataCatalogEncryptionSettingsList{})
}
