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

type EncryptionConfigObservation struct {

	// Region name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EncryptionConfigParameters struct {

	// An AWS KMS customer master key (CMK) ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The type of encryption. Set to KMS to use your own key for encryption. Set to NONE for default encryption.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// EncryptionConfigSpec defines the desired state of EncryptionConfig
type EncryptionConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EncryptionConfigParameters `json:"forProvider"`
}

// EncryptionConfigStatus defines the observed state of EncryptionConfig.
type EncryptionConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EncryptionConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EncryptionConfig is the Schema for the EncryptionConfigs API. Creates and manages an AWS XRay Encryption Config.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EncryptionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EncryptionConfigSpec   `json:"spec"`
	Status            EncryptionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EncryptionConfigList contains a list of EncryptionConfigs
type EncryptionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EncryptionConfig `json:"items"`
}

// Repository type metadata.
var (
	EncryptionConfig_Kind             = "EncryptionConfig"
	EncryptionConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EncryptionConfig_Kind}.String()
	EncryptionConfig_KindAPIVersion   = EncryptionConfig_Kind + "." + CRDGroupVersion.String()
	EncryptionConfig_GroupVersionKind = CRDGroupVersion.WithKind(EncryptionConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&EncryptionConfig{}, &EncryptionConfigList{})
}