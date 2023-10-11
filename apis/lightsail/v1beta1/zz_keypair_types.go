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

type KeyPairInitParameters struct {

	// The name of the Lightsail Key Pair
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type KeyPairObservation struct {

	// The ARN of the Lightsail key pair
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The MD5 public key fingerprint for the encrypted
	// private key
	EncryptedFingerprint *string `json:"encryptedFingerprint,omitempty" tf:"encrypted_fingerprint,omitempty"`

	// – the private key material, base 64 encoded and
	// encrypted with the given pgp_key. This is only populated when creating a new
	// key and pgp_key is supplied
	EncryptedPrivateKey *string `json:"encryptedPrivateKey,omitempty" tf:"encrypted_private_key,omitempty"`

	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// The name used for this key pair
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Lightsail Key Pair
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// the private key, base64 encoded. This is only populated
	// when creating a new key, and when no pgp_key is provided
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`

	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type KeyPairParameters struct {

	// The name of the Lightsail Key Pair
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// The public key material. This public key will be
	// imported into Lightsail
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// KeyPairSpec defines the desired state of KeyPair
type KeyPairSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyPairParameters `json:"forProvider"`
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
	InitProvider KeyPairInitParameters `json:"initProvider,omitempty"`
}

// KeyPairStatus defines the observed state of KeyPair.
type KeyPairStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyPairObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPair is the Schema for the KeyPairs API. Provides an Lightsail Key Pair
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyPairSpec   `json:"spec"`
	Status            KeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPairList contains a list of KeyPairs
type KeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyPair `json:"items"`
}

// Repository type metadata.
var (
	KeyPair_Kind             = "KeyPair"
	KeyPair_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyPair_Kind}.String()
	KeyPair_KindAPIVersion   = KeyPair_Kind + "." + CRDGroupVersion.String()
	KeyPair_GroupVersionKind = CRDGroupVersion.WithKind(KeyPair_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyPair{}, &KeyPairList{})
}
