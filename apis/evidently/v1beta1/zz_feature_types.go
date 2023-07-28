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

type EvaluationRulesInitParameters struct {
}

type EvaluationRulesObservation struct {

	// The name of the experiment or launch.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This value is aws.evidently.splits if this is an evaluation rule for a launch, and it is aws.evidently.onlineab if this is an evaluation rule for an experiment.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EvaluationRulesParameters struct {
}

type FeatureInitParameters struct {

	// The name of the variation to use as the default variation. The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature. This variation must also be listed in the variations structure. If you omit default_variation, the first variation listed in the variations structure is used as the default variation.
	DefaultVariation *string `json:"defaultVariation,omitempty" tf:"default_variation,omitempty"`

	// Specifies the description of the feature.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify users that should always be served a specific variation of a feature. Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	EntityOverrides map[string]*string `json:"entityOverrides,omitempty" tf:"entity_overrides,omitempty"`

	// Specify ALL_RULES to activate the traffic allocation specified by any ongoing launches or experiments. Specify DEFAULT_VARIATION to serve the default variation to all users instead.
	EvaluationStrategy *string `json:"evaluationStrategy,omitempty" tf:"evaluation_strategy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more blocks that contain the configuration of the feature's different variations. Detailed below
	Variations []VariationsInitParameters `json:"variations,omitempty" tf:"variations,omitempty"`
}

type FeatureObservation struct {

	// The ARN of the feature.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time that the feature is created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// The name of the variation to use as the default variation. The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature. This variation must also be listed in the variations structure. If you omit default_variation, the first variation listed in the variations structure is used as the default variation.
	DefaultVariation *string `json:"defaultVariation,omitempty" tf:"default_variation,omitempty"`

	// Specifies the description of the feature.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify users that should always be served a specific variation of a feature. Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	EntityOverrides map[string]*string `json:"entityOverrides,omitempty" tf:"entity_overrides,omitempty"`

	// One or more blocks that define the evaluation rules for the feature. Detailed below
	EvaluationRules []EvaluationRulesObservation `json:"evaluationRules,omitempty" tf:"evaluation_rules,omitempty"`

	// Specify ALL_RULES to activate the traffic allocation specified by any ongoing launches or experiments. Specify DEFAULT_VARIATION to serve the default variation to all users instead.
	EvaluationStrategy *string `json:"evaluationStrategy,omitempty" tf:"evaluation_strategy,omitempty"`

	// The feature name and the project name or arn separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time that the feature was most recently updated.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty" tf:"last_updated_time,omitempty"`

	// The name or ARN of the project that is to contain the new feature.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The current state of the feature. Valid values are AVAILABLE and UPDATING.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Defines the type of value used to define the different feature variations. Valid Values: STRING, LONG, DOUBLE, BOOLEAN.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`

	// One or more blocks that contain the configuration of the feature's different variations. Detailed below
	Variations []VariationsObservation `json:"variations,omitempty" tf:"variations,omitempty"`
}

type FeatureParameters struct {

	// The name of the variation to use as the default variation. The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature. This variation must also be listed in the variations structure. If you omit default_variation, the first variation listed in the variations structure is used as the default variation.
	// +kubebuilder:validation:Optional
	DefaultVariation *string `json:"defaultVariation,omitempty" tf:"default_variation,omitempty"`

	// Specifies the description of the feature.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify users that should always be served a specific variation of a feature. Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	// +kubebuilder:validation:Optional
	EntityOverrides map[string]*string `json:"entityOverrides,omitempty" tf:"entity_overrides,omitempty"`

	// Specify ALL_RULES to activate the traffic allocation specified by any ongoing launches or experiments. Specify DEFAULT_VARIATION to serve the default variation to all users instead.
	// +kubebuilder:validation:Optional
	EvaluationStrategy *string `json:"evaluationStrategy,omitempty" tf:"evaluation_strategy,omitempty"`

	// The name or ARN of the project that is to contain the new feature.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/evidently/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in evidently to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in evidently to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more blocks that contain the configuration of the feature's different variations. Detailed below
	// +kubebuilder:validation:Optional
	Variations []VariationsParameters `json:"variations,omitempty" tf:"variations,omitempty"`
}

type ValueInitParameters struct {

	// If this feature uses the Boolean variation type, this field contains the Boolean value of this variation.
	BoolValue *string `json:"boolValue,omitempty" tf:"bool_value,omitempty"`

	// If this feature uses the double integer variation type, this field contains the double integer value of this variation.
	DoubleValue *string `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// If this feature uses the long variation type, this field contains the long value of this variation. Minimum value of -9007199254740991. Maximum value of 9007199254740991.
	LongValue *string `json:"longValue,omitempty" tf:"long_value,omitempty"`

	// If this feature uses the string variation type, this field contains the string value of this variation. Minimum length of 0. Maximum length of 512.
	StringValue *string `json:"stringValue,omitempty" tf:"string_value,omitempty"`
}

type ValueObservation struct {

	// If this feature uses the Boolean variation type, this field contains the Boolean value of this variation.
	BoolValue *string `json:"boolValue,omitempty" tf:"bool_value,omitempty"`

	// If this feature uses the double integer variation type, this field contains the double integer value of this variation.
	DoubleValue *string `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// If this feature uses the long variation type, this field contains the long value of this variation. Minimum value of -9007199254740991. Maximum value of 9007199254740991.
	LongValue *string `json:"longValue,omitempty" tf:"long_value,omitempty"`

	// If this feature uses the string variation type, this field contains the string value of this variation. Minimum length of 0. Maximum length of 512.
	StringValue *string `json:"stringValue,omitempty" tf:"string_value,omitempty"`
}

type ValueParameters struct {

	// If this feature uses the Boolean variation type, this field contains the Boolean value of this variation.
	// +kubebuilder:validation:Optional
	BoolValue *string `json:"boolValue,omitempty" tf:"bool_value,omitempty"`

	// If this feature uses the double integer variation type, this field contains the double integer value of this variation.
	// +kubebuilder:validation:Optional
	DoubleValue *string `json:"doubleValue,omitempty" tf:"double_value,omitempty"`

	// If this feature uses the long variation type, this field contains the long value of this variation. Minimum value of -9007199254740991. Maximum value of 9007199254740991.
	// +kubebuilder:validation:Optional
	LongValue *string `json:"longValue,omitempty" tf:"long_value,omitempty"`

	// If this feature uses the string variation type, this field contains the string value of this variation. Minimum length of 0. Maximum length of 512.
	// +kubebuilder:validation:Optional
	StringValue *string `json:"stringValue,omitempty" tf:"string_value,omitempty"`
}

type VariationsInitParameters struct {

	// The name of the variation. Minimum length of 1. Maximum length of 127.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that specifies the value assigned to this variation. Detailed below
	Value []ValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type VariationsObservation struct {

	// The name of the variation. Minimum length of 1. Maximum length of 127.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that specifies the value assigned to this variation. Detailed below
	Value []ValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type VariationsParameters struct {

	// The name of the variation. Minimum length of 1. Maximum length of 127.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that specifies the value assigned to this variation. Detailed below
	// +kubebuilder:validation:Optional
	Value []ValueParameters `json:"value,omitempty" tf:"value,omitempty"`
}

// FeatureSpec defines the desired state of Feature
type FeatureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FeatureParameters `json:"forProvider"`
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
	InitProvider FeatureInitParameters `json:"initProvider,omitempty"`
}

// FeatureStatus defines the observed state of Feature.
type FeatureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FeatureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Feature is the Schema for the Features API. Provides a CloudWatch Evidently Feature resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Feature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.variations) || has(self.initProvider.variations)",message="variations is a required parameter"
	Spec   FeatureSpec   `json:"spec"`
	Status FeatureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FeatureList contains a list of Features
type FeatureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Feature `json:"items"`
}

// Repository type metadata.
var (
	Feature_Kind             = "Feature"
	Feature_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Feature_Kind}.String()
	Feature_KindAPIVersion   = Feature_Kind + "." + CRDGroupVersion.String()
	Feature_GroupVersionKind = CRDGroupVersion.WithKind(Feature_Kind)
)

func init() {
	SchemeBuilder.Register(&Feature{}, &FeatureList{})
}
