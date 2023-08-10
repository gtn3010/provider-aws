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

type LoggingOptionsInitParameters struct {

	// The default logging level. Valid Values: "DEBUG", "INFO", "ERROR", "WARN", "DISABLED".
	DefaultLogLevel *string `json:"defaultLogLevel,omitempty" tf:"default_log_level,omitempty"`

	// If true all logs are disabled. The default is false.
	DisableAllLogs *bool `json:"disableAllLogs,omitempty" tf:"disable_all_logs,omitempty"`
}

type LoggingOptionsObservation struct {

	// The default logging level. Valid Values: "DEBUG", "INFO", "ERROR", "WARN", "DISABLED".
	DefaultLogLevel *string `json:"defaultLogLevel,omitempty" tf:"default_log_level,omitempty"`

	// If true all logs are disabled. The default is false.
	DisableAllLogs *bool `json:"disableAllLogs,omitempty" tf:"disable_all_logs,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type LoggingOptionsParameters struct {

	// The default logging level. Valid Values: "DEBUG", "INFO", "ERROR", "WARN", "DISABLED".
	// +kubebuilder:validation:Optional
	DefaultLogLevel *string `json:"defaultLogLevel,omitempty" tf:"default_log_level,omitempty"`

	// If true all logs are disabled. The default is false.
	// +kubebuilder:validation:Optional
	DisableAllLogs *bool `json:"disableAllLogs,omitempty" tf:"disable_all_logs,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// LoggingOptionsSpec defines the desired state of LoggingOptions
type LoggingOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoggingOptionsParameters `json:"forProvider"`
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
	InitProvider LoggingOptionsInitParameters `json:"initProvider,omitempty"`
}

// LoggingOptionsStatus defines the observed state of LoggingOptions.
type LoggingOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoggingOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingOptions is the Schema for the LoggingOptionss API. Provides a resource to manage default logging options.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LoggingOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultLogLevel) || has(self.initProvider.defaultLogLevel)",message="defaultLogLevel is a required parameter"
	Spec   LoggingOptionsSpec   `json:"spec"`
	Status LoggingOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingOptionsList contains a list of LoggingOptionss
type LoggingOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingOptions `json:"items"`
}

// Repository type metadata.
var (
	LoggingOptions_Kind             = "LoggingOptions"
	LoggingOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoggingOptions_Kind}.String()
	LoggingOptions_KindAPIVersion   = LoggingOptions_Kind + "." + CRDGroupVersion.String()
	LoggingOptions_GroupVersionKind = CRDGroupVersion.WithKind(LoggingOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&LoggingOptions{}, &LoggingOptionsList{})
}