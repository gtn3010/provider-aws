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

type GlobalTableInitParameters struct {

	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replica []ReplicaInitParameters `json:"replica,omitempty" tf:"replica,omitempty"`
}

type GlobalTableObservation struct {

	// The ARN of the DynamoDB Global Table
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the DynamoDB Global Table
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replica []ReplicaObservation `json:"replica,omitempty" tf:"replica,omitempty"`
}

type GlobalTableParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	// +kubebuilder:validation:Optional
	Replica []ReplicaParameters `json:"replica,omitempty" tf:"replica,omitempty"`
}

type ReplicaInitParameters struct {

	// AWS region name of replica DynamoDB TableE.g., us-east-1
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type ReplicaObservation struct {

	// AWS region name of replica DynamoDB TableE.g., us-east-1
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type ReplicaParameters struct {

	// AWS region name of replica DynamoDB TableE.g., us-east-1
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

// GlobalTableSpec defines the desired state of GlobalTable
type GlobalTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalTableParameters `json:"forProvider"`
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
	InitProvider GlobalTableInitParameters `json:"initProvider,omitempty"`
}

// GlobalTableStatus defines the observed state of GlobalTable.
type GlobalTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalTable is the Schema for the GlobalTables API. Manages DynamoDB Global Tables V1 (version 2017.11.29)
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlobalTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replica) || has(self.initProvider.replica)",message="replica is a required parameter"
	Spec   GlobalTableSpec   `json:"spec"`
	Status GlobalTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalTableList contains a list of GlobalTables
type GlobalTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalTable `json:"items"`
}

// Repository type metadata.
var (
	GlobalTable_Kind             = "GlobalTable"
	GlobalTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalTable_Kind}.String()
	GlobalTable_KindAPIVersion   = GlobalTable_Kind + "." + CRDGroupVersion.String()
	GlobalTable_GroupVersionKind = CRDGroupVersion.WithKind(GlobalTable_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalTable{}, &GlobalTableList{})
}
