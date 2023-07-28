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

type ApplicationInitParameters struct {

	// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
	AutoConfigEnabled *bool `json:"autoConfigEnabled,omitempty" tf:"auto_config_enabled,omitempty"`

	// Configures all of the resources in the resource group by applying the recommended configurations.
	AutoCreate *bool `json:"autoCreate,omitempty" tf:"auto_create,omitempty"`

	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
	CweMonitorEnabled *bool `json:"cweMonitorEnabled,omitempty" tf:"cwe_monitor_enabled,omitempty"`

	// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to ACCOUNT_BASED.
	GroupingType *string `json:"groupingType,omitempty" tf:"grouping_type,omitempty"`

	// When set to true, creates opsItems for any problems detected on an application.
	OpsCenterEnabled *bool `json:"opsCenterEnabled,omitempty" tf:"ops_center_enabled,omitempty"`

	// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
	OpsItemSnsTopicArn *string `json:"opsItemSnsTopicArn,omitempty" tf:"ops_item_sns_topic_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationObservation struct {

	// ARN of the Application.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
	AutoConfigEnabled *bool `json:"autoConfigEnabled,omitempty" tf:"auto_config_enabled,omitempty"`

	// Configures all of the resources in the resource group by applying the recommended configurations.
	AutoCreate *bool `json:"autoCreate,omitempty" tf:"auto_create,omitempty"`

	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
	CweMonitorEnabled *bool `json:"cweMonitorEnabled,omitempty" tf:"cwe_monitor_enabled,omitempty"`

	// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to ACCOUNT_BASED.
	GroupingType *string `json:"groupingType,omitempty" tf:"grouping_type,omitempty"`

	// Name of the resource group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When set to true, creates opsItems for any problems detected on an application.
	OpsCenterEnabled *bool `json:"opsCenterEnabled,omitempty" tf:"ops_center_enabled,omitempty"`

	// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
	OpsItemSnsTopicArn *string `json:"opsItemSnsTopicArn,omitempty" tf:"ops_item_sns_topic_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ApplicationParameters struct {

	// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
	// +kubebuilder:validation:Optional
	AutoConfigEnabled *bool `json:"autoConfigEnabled,omitempty" tf:"auto_config_enabled,omitempty"`

	// Configures all of the resources in the resource group by applying the recommended configurations.
	// +kubebuilder:validation:Optional
	AutoCreate *bool `json:"autoCreate,omitempty" tf:"auto_create,omitempty"`

	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
	// +kubebuilder:validation:Optional
	CweMonitorEnabled *bool `json:"cweMonitorEnabled,omitempty" tf:"cwe_monitor_enabled,omitempty"`

	// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to ACCOUNT_BASED.
	// +kubebuilder:validation:Optional
	GroupingType *string `json:"groupingType,omitempty" tf:"grouping_type,omitempty"`

	// When set to true, creates opsItems for any problems detected on an application.
	// +kubebuilder:validation:Optional
	OpsCenterEnabled *bool `json:"opsCenterEnabled,omitempty" tf:"ops_center_enabled,omitempty"`

	// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
	// +kubebuilder:validation:Optional
	OpsItemSnsTopicArn *string `json:"opsItemSnsTopicArn,omitempty" tf:"ops_item_sns_topic_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Provides a CloudWatch Application Insights Application resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
