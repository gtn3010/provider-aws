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

type SecurityConfigurationInitParameters struct {

	// A JSON formatted Security Configuration
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`
}

type SecurityConfigurationObservation struct {

	// A JSON formatted Security Configuration
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Date the Security Configuration was created
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// The ID of the EMR Security Configuration (Same as the name)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityConfigurationParameters struct {

	// A JSON formatted Security Configuration
	// +kubebuilder:validation:Optional
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// SecurityConfigurationSpec defines the desired state of SecurityConfiguration
type SecurityConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityConfigurationParameters `json:"forProvider"`
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
	InitProvider SecurityConfigurationInitParameters `json:"initProvider,omitempty"`
}

// SecurityConfigurationStatus defines the observed state of SecurityConfiguration.
type SecurityConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfiguration is the Schema for the SecurityConfigurations API. Provides a resource to manage AWS EMR Security Configurations
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configuration) || has(self.initProvider.configuration)",message="configuration is a required parameter"
	Spec   SecurityConfigurationSpec   `json:"spec"`
	Status SecurityConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfigurationList contains a list of SecurityConfigurations
type SecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityConfiguration `json:"items"`
}

// Repository type metadata.
var (
	SecurityConfiguration_Kind             = "SecurityConfiguration"
	SecurityConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityConfiguration_Kind}.String()
	SecurityConfiguration_KindAPIVersion   = SecurityConfiguration_Kind + "." + CRDGroupVersion.String()
	SecurityConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(SecurityConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityConfiguration{}, &SecurityConfigurationList{})
}