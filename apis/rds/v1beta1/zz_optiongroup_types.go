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

type OptionGroupInitParameters struct {

	// Specifies the name of the engine that this option group should be associated with.
	EngineName *string `json:"engineName,omitempty" tf:"engine_name,omitempty"`

	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion *string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`

	// A list of Options to apply.
	Option []OptionInitParameters `json:"option,omitempty" tf:"option,omitempty"`

	// The description of the option group.
	OptionGroupDescription *string `json:"optionGroupDescription,omitempty" tf:"option_group_description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OptionGroupObservation struct {

	// The ARN of the db option group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies the name of the engine that this option group should be associated with.
	EngineName *string `json:"engineName,omitempty" tf:"engine_name,omitempty"`

	// The db option group name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion *string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`

	// A list of Options to apply.
	Option []OptionObservation `json:"option,omitempty" tf:"option,omitempty"`

	// The description of the option group.
	OptionGroupDescription *string `json:"optionGroupDescription,omitempty" tf:"option_group_description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type OptionGroupParameters struct {

	// Specifies the name of the engine that this option group should be associated with.
	// +kubebuilder:validation:Optional
	EngineName *string `json:"engineName,omitempty" tf:"engine_name,omitempty"`

	// Specifies the major version of the engine that this option group should be associated with.
	// +kubebuilder:validation:Optional
	MajorEngineVersion *string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`

	// A list of Options to apply.
	// +kubebuilder:validation:Optional
	Option []OptionParameters `json:"option,omitempty" tf:"option,omitempty"`

	// The description of the option group.
	// +kubebuilder:validation:Optional
	OptionGroupDescription *string `json:"optionGroupDescription,omitempty" tf:"option_group_description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OptionInitParameters struct {

	// A list of DB Security Groups for which the option is enabled.
	DBSecurityGroupMemberships []*string `json:"dbSecurityGroupMemberships,omitempty" tf:"db_security_group_memberships,omitempty"`

	// The Name of the Option (e.g., MEMCACHED).
	OptionName *string `json:"optionName,omitempty" tf:"option_name,omitempty"`

	// A list of option settings to apply.
	OptionSettings []OptionSettingsInitParameters `json:"optionSettings,omitempty" tf:"option_settings,omitempty"`

	// The Port number when connecting to the Option (e.g., 11211).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// A list of VPC Security Groups for which the option is enabled.
	VPCSecurityGroupMemberships []*string `json:"vpcSecurityGroupMemberships,omitempty" tf:"vpc_security_group_memberships,omitempty"`

	// The version of the option (e.g., 13.1.0.0).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OptionObservation struct {

	// A list of DB Security Groups for which the option is enabled.
	DBSecurityGroupMemberships []*string `json:"dbSecurityGroupMemberships,omitempty" tf:"db_security_group_memberships,omitempty"`

	// The Name of the Option (e.g., MEMCACHED).
	OptionName *string `json:"optionName,omitempty" tf:"option_name,omitempty"`

	// A list of option settings to apply.
	OptionSettings []OptionSettingsObservation `json:"optionSettings,omitempty" tf:"option_settings,omitempty"`

	// The Port number when connecting to the Option (e.g., 11211).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// A list of VPC Security Groups for which the option is enabled.
	VPCSecurityGroupMemberships []*string `json:"vpcSecurityGroupMemberships,omitempty" tf:"vpc_security_group_memberships,omitempty"`

	// The version of the option (e.g., 13.1.0.0).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OptionParameters struct {

	// A list of DB Security Groups for which the option is enabled.
	// +kubebuilder:validation:Optional
	DBSecurityGroupMemberships []*string `json:"dbSecurityGroupMemberships,omitempty" tf:"db_security_group_memberships,omitempty"`

	// The Name of the Option (e.g., MEMCACHED).
	// +kubebuilder:validation:Optional
	OptionName *string `json:"optionName" tf:"option_name,omitempty"`

	// A list of option settings to apply.
	// +kubebuilder:validation:Optional
	OptionSettings []OptionSettingsParameters `json:"optionSettings,omitempty" tf:"option_settings,omitempty"`

	// The Port number when connecting to the Option (e.g., 11211).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// A list of VPC Security Groups for which the option is enabled.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupMemberships []*string `json:"vpcSecurityGroupMemberships,omitempty" tf:"vpc_security_group_memberships,omitempty"`

	// The version of the option (e.g., 13.1.0.0).
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OptionSettingsInitParameters struct {

	// The name of the option group. Must be lowercase, to match as it is stored in AWS.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Value of the setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OptionSettingsObservation struct {

	// The name of the option group. Must be lowercase, to match as it is stored in AWS.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Value of the setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OptionSettingsParameters struct {

	// The name of the option group. Must be lowercase, to match as it is stored in AWS.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Value of the setting.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// OptionGroupSpec defines the desired state of OptionGroup
type OptionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OptionGroupParameters `json:"forProvider"`
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
	InitProvider OptionGroupInitParameters `json:"initProvider,omitempty"`
}

// OptionGroupStatus defines the observed state of OptionGroup.
type OptionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OptionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OptionGroup is the Schema for the OptionGroups API. Provides an RDS DB option group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OptionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineName) || (has(self.initProvider) && has(self.initProvider.engineName))",message="spec.forProvider.engineName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.majorEngineVersion) || (has(self.initProvider) && has(self.initProvider.majorEngineVersion))",message="spec.forProvider.majorEngineVersion is a required parameter"
	Spec   OptionGroupSpec   `json:"spec"`
	Status OptionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OptionGroupList contains a list of OptionGroups
type OptionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OptionGroup `json:"items"`
}

// Repository type metadata.
var (
	OptionGroup_Kind             = "OptionGroup"
	OptionGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OptionGroup_Kind}.String()
	OptionGroup_KindAPIVersion   = OptionGroup_Kind + "." + CRDGroupVersion.String()
	OptionGroup_GroupVersionKind = CRDGroupVersion.WithKind(OptionGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&OptionGroup{}, &OptionGroupList{})
}
