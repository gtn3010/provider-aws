/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DistributionParameters defines the desired state of Distribution
type DistributionParameters struct {
	// Region is which region the Distribution will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The distribution's configuration information.
	// +kubebuilder:validation:Required
	DistributionConfig           *DistributionConfig `json:"distributionConfig"`
	CustomDistributionParameters `json:",inline"`
}

// DistributionSpec defines the desired state of Distribution
type DistributionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DistributionParameters `json:"forProvider"`
}

// DistributionObservation defines the observed state of Distribution
type DistributionObservation struct {
	// The distribution's information.
	Distribution *Distribution_SDK `json:"distribution,omitempty"`
	// The current version of the distribution created.
	ETag *string `json:"eTag,omitempty"`
	// The fully qualified URI of the new distribution resource just created.
	Location *string `json:"location,omitempty"`
}

// DistributionStatus defines the observed state of Distribution.
type DistributionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DistributionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Distribution is the Schema for the Distributions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Distribution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DistributionSpec   `json:"spec"`
	Status            DistributionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributionList contains a list of Distributions
type DistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Distribution `json:"items"`
}

// Repository type metadata.
var (
	DistributionKind             = "Distribution"
	DistributionGroupKind        = schema.GroupKind{Group: Group, Kind: DistributionKind}.String()
	DistributionKindAPIVersion   = DistributionKind + "." + GroupVersion.String()
	DistributionGroupVersionKind = GroupVersion.WithKind(DistributionKind)
)

func init() {
	SchemeBuilder.Register(&Distribution{}, &DistributionList{})
}