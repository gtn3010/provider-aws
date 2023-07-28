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

type EncryptionEntitiesInitParameters struct {
	Items []EncryptionEntitiesItemsInitParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type EncryptionEntitiesItemsInitParameters struct {

	// Object that contains an attribute items that contains the list of field patterns in a field-level encryption content type profile specify the fields that you want to be encrypted.
	FieldPatterns []FieldPatternsInitParameters `json:"fieldPatterns,omitempty" tf:"field_patterns,omitempty"`

	// The provider associated with the public key being used for encryption.
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`
}

type EncryptionEntitiesItemsObservation struct {

	// Object that contains an attribute items that contains the list of field patterns in a field-level encryption content type profile specify the fields that you want to be encrypted.
	FieldPatterns []FieldPatternsObservation `json:"fieldPatterns,omitempty" tf:"field_patterns,omitempty"`

	// The provider associated with the public key being used for encryption.
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// The public key associated with a set of field-level encryption patterns, to be used when encrypting the fields that match the patterns.
	PublicKeyID *string `json:"publicKeyId,omitempty" tf:"public_key_id,omitempty"`
}

type EncryptionEntitiesItemsParameters struct {

	// Object that contains an attribute items that contains the list of field patterns in a field-level encryption content type profile specify the fields that you want to be encrypted.
	// +kubebuilder:validation:Optional
	FieldPatterns []FieldPatternsParameters `json:"fieldPatterns,omitempty" tf:"field_patterns,omitempty"`

	// The provider associated with the public key being used for encryption.
	// +kubebuilder:validation:Optional
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// The public key associated with a set of field-level encryption patterns, to be used when encrypting the fields that match the patterns.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudfront/v1beta1.PublicKey
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicKeyID *string `json:"publicKeyId,omitempty" tf:"public_key_id,omitempty"`

	// Reference to a PublicKey in cloudfront to populate publicKeyId.
	// +kubebuilder:validation:Optional
	PublicKeyIDRef *v1.Reference `json:"publicKeyIdRef,omitempty" tf:"-"`

	// Selector for a PublicKey in cloudfront to populate publicKeyId.
	// +kubebuilder:validation:Optional
	PublicKeyIDSelector *v1.Selector `json:"publicKeyIdSelector,omitempty" tf:"-"`
}

type EncryptionEntitiesObservation struct {
	Items []EncryptionEntitiesItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type EncryptionEntitiesParameters struct {

	// +kubebuilder:validation:Optional
	Items []EncryptionEntitiesItemsParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type FieldLevelEncryptionProfileInitParameters struct {

	// An optional comment about the Field Level Encryption Profile.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The encryption entities config block for field-level encryption profiles that contains an attribute items which includes the encryption key and field pattern specifications.
	EncryptionEntities []EncryptionEntitiesInitParameters `json:"encryptionEntities,omitempty" tf:"encryption_entities,omitempty"`

	// The name of the Field Level Encryption Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FieldLevelEncryptionProfileObservation struct {

	// Internal value used by CloudFront to allow future updates to the Field Level Encryption Profile.
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	// An optional comment about the Field Level Encryption Profile.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The encryption entities config block for field-level encryption profiles that contains an attribute items which includes the encryption key and field pattern specifications.
	EncryptionEntities []EncryptionEntitiesObservation `json:"encryptionEntities,omitempty" tf:"encryption_entities,omitempty"`

	// The current version of the Field Level Encryption Profile. For example: E2QWRUHAPOMQZL.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The identifier for the Field Level Encryption Profile. For example: K3D5EWEUDCCXON.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Field Level Encryption Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FieldLevelEncryptionProfileParameters struct {

	// An optional comment about the Field Level Encryption Profile.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// The encryption entities config block for field-level encryption profiles that contains an attribute items which includes the encryption key and field pattern specifications.
	// +kubebuilder:validation:Optional
	EncryptionEntities []EncryptionEntitiesParameters `json:"encryptionEntities,omitempty" tf:"encryption_entities,omitempty"`

	// The name of the Field Level Encryption Profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type FieldPatternsInitParameters struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type FieldPatternsObservation struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type FieldPatternsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

// FieldLevelEncryptionProfileSpec defines the desired state of FieldLevelEncryptionProfile
type FieldLevelEncryptionProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FieldLevelEncryptionProfileParameters `json:"forProvider"`
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
	InitProvider FieldLevelEncryptionProfileInitParameters `json:"initProvider,omitempty"`
}

// FieldLevelEncryptionProfileStatus defines the observed state of FieldLevelEncryptionProfile.
type FieldLevelEncryptionProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FieldLevelEncryptionProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FieldLevelEncryptionProfile is the Schema for the FieldLevelEncryptionProfiles API. Provides a CloudFront Field-level Encryption Profile resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FieldLevelEncryptionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.encryptionEntities) || has(self.initProvider.encryptionEntities)",message="encryptionEntities is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   FieldLevelEncryptionProfileSpec   `json:"spec"`
	Status FieldLevelEncryptionProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FieldLevelEncryptionProfileList contains a list of FieldLevelEncryptionProfiles
type FieldLevelEncryptionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FieldLevelEncryptionProfile `json:"items"`
}

// Repository type metadata.
var (
	FieldLevelEncryptionProfile_Kind             = "FieldLevelEncryptionProfile"
	FieldLevelEncryptionProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FieldLevelEncryptionProfile_Kind}.String()
	FieldLevelEncryptionProfile_KindAPIVersion   = FieldLevelEncryptionProfile_Kind + "." + CRDGroupVersion.String()
	FieldLevelEncryptionProfile_GroupVersionKind = CRDGroupVersion.WithKind(FieldLevelEncryptionProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&FieldLevelEncryptionProfile{}, &FieldLevelEncryptionProfileList{})
}
