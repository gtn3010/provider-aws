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

type LogDestinationConfigInitParameters struct {

	// A map describing the logging destination for the chosen log_destination_type.
	LogDestination map[string]*string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// The location to send logs to. Valid values: S3, CloudWatchLogs, KinesisDataFirehose.
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The type of log to send. Valid values: ALERT or FLOW. Alert logs report traffic that matches a StatefulRule with an action setting that sends a log message. Flow logs are standard network traffic flow logs.
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`
}

type LogDestinationConfigObservation struct {

	// A map describing the logging destination for the chosen log_destination_type.
	LogDestination map[string]*string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// The location to send logs to. Valid values: S3, CloudWatchLogs, KinesisDataFirehose.
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The type of log to send. Valid values: ALERT or FLOW. Alert logs report traffic that matches a StatefulRule with an action setting that sends a log message. Flow logs are standard network traffic flow logs.
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`
}

type LogDestinationConfigParameters struct {

	// A map describing the logging destination for the chosen log_destination_type.
	// +kubebuilder:validation:Optional
	LogDestination map[string]*string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// The location to send logs to. Valid values: S3, CloudWatchLogs, KinesisDataFirehose.
	// +kubebuilder:validation:Optional
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The type of log to send. Valid values: ALERT or FLOW. Alert logs report traffic that matches a StatefulRule with an action setting that sends a log message. Flow logs are standard network traffic flow logs.
	// +kubebuilder:validation:Optional
	LogType *string `json:"logType,omitempty" tf:"log_type,omitempty"`
}

type LoggingConfigurationInitParameters struct {

	// A configuration block describing how AWS Network Firewall performs logging for a firewall. See Logging Configuration below for details.
	LoggingConfiguration []LoggingConfigurationLoggingConfigurationInitParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`
}

type LoggingConfigurationLoggingConfigurationInitParameters struct {

	// Set of configuration blocks describing the logging details for a firewall. See Log Destination Config below for details. At most, only two blocks can be specified; one for FLOW logs and one for ALERT logs.
	LogDestinationConfig []LogDestinationConfigInitParameters `json:"logDestinationConfig,omitempty" tf:"log_destination_config,omitempty"`
}

type LoggingConfigurationLoggingConfigurationObservation struct {

	// Set of configuration blocks describing the logging details for a firewall. See Log Destination Config below for details. At most, only two blocks can be specified; one for FLOW logs and one for ALERT logs.
	LogDestinationConfig []LogDestinationConfigObservation `json:"logDestinationConfig,omitempty" tf:"log_destination_config,omitempty"`
}

type LoggingConfigurationLoggingConfigurationParameters struct {

	// Set of configuration blocks describing the logging details for a firewall. See Log Destination Config below for details. At most, only two blocks can be specified; one for FLOW logs and one for ALERT logs.
	// +kubebuilder:validation:Optional
	LogDestinationConfig []LogDestinationConfigParameters `json:"logDestinationConfig,omitempty" tf:"log_destination_config,omitempty"`
}

type LoggingConfigurationObservation struct {

	// The Amazon Resource Name (ARN) of the Network Firewall firewall.
	FirewallArn *string `json:"firewallArn,omitempty" tf:"firewall_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the associated firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A configuration block describing how AWS Network Firewall performs logging for a firewall. See Logging Configuration below for details.
	LoggingConfiguration []LoggingConfigurationLoggingConfigurationObservation `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`
}

type LoggingConfigurationParameters struct {

	// The Amazon Resource Name (ARN) of the Network Firewall firewall.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkfirewall/v1beta1.Firewall
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	FirewallArn *string `json:"firewallArn,omitempty" tf:"firewall_arn,omitempty"`

	// Reference to a Firewall in networkfirewall to populate firewallArn.
	// +kubebuilder:validation:Optional
	FirewallArnRef *v1.Reference `json:"firewallArnRef,omitempty" tf:"-"`

	// Selector for a Firewall in networkfirewall to populate firewallArn.
	// +kubebuilder:validation:Optional
	FirewallArnSelector *v1.Selector `json:"firewallArnSelector,omitempty" tf:"-"`

	// A configuration block describing how AWS Network Firewall performs logging for a firewall. See Logging Configuration below for details.
	// +kubebuilder:validation:Optional
	LoggingConfiguration []LoggingConfigurationLoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LoggingConfigurationSpec defines the desired state of LoggingConfiguration
type LoggingConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoggingConfigurationParameters `json:"forProvider"`
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
	InitProvider LoggingConfigurationInitParameters `json:"initProvider,omitempty"`
}

// LoggingConfigurationStatus defines the observed state of LoggingConfiguration.
type LoggingConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoggingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingConfiguration is the Schema for the LoggingConfigurations API. Provides an AWS Network Firewall Logging Configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loggingConfiguration) || has(self.initProvider.loggingConfiguration)",message="loggingConfiguration is a required parameter"
	Spec   LoggingConfigurationSpec   `json:"spec"`
	Status LoggingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingConfigurationList contains a list of LoggingConfigurations
type LoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	LoggingConfiguration_Kind             = "LoggingConfiguration"
	LoggingConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoggingConfiguration_Kind}.String()
	LoggingConfiguration_KindAPIVersion   = LoggingConfiguration_Kind + "." + CRDGroupVersion.String()
	LoggingConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(LoggingConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&LoggingConfiguration{}, &LoggingConfigurationList{})
}
