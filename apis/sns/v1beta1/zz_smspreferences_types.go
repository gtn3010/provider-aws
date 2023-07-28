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

type SMSPreferencesInitParameters struct {

	// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
	DefaultSMSType *string `json:"defaultSmsType,omitempty" tf:"default_sms_type,omitempty"`

	// A string, such as your business brand, that is displayed as the sender on the receiving device.
	DefaultSenderID *string `json:"defaultSenderId,omitempty" tf:"default_sender_id,omitempty"`

	// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
	DeliveryStatusSuccessSamplingRate *string `json:"deliveryStatusSuccessSamplingRate,omitempty" tf:"delivery_status_success_sampling_rate,omitempty"`

	// The maximum amount in USD that you are willing to spend each month to send SMS messages.
	MonthlySpendLimit *float64 `json:"monthlySpendLimit,omitempty" tf:"monthly_spend_limit,omitempty"`

	// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
	UsageReportS3Bucket *string `json:"usageReportS3Bucket,omitempty" tf:"usage_report_s3_bucket,omitempty"`
}

type SMSPreferencesObservation struct {

	// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
	DefaultSMSType *string `json:"defaultSmsType,omitempty" tf:"default_sms_type,omitempty"`

	// A string, such as your business brand, that is displayed as the sender on the receiving device.
	DefaultSenderID *string `json:"defaultSenderId,omitempty" tf:"default_sender_id,omitempty"`

	// The ARN of the IAM role that allows Amazon SNS to write logs about SMS deliveries in CloudWatch Logs.
	DeliveryStatusIAMRoleArn *string `json:"deliveryStatusIamRoleArn,omitempty" tf:"delivery_status_iam_role_arn,omitempty"`

	// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
	DeliveryStatusSuccessSamplingRate *string `json:"deliveryStatusSuccessSamplingRate,omitempty" tf:"delivery_status_success_sampling_rate,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum amount in USD that you are willing to spend each month to send SMS messages.
	MonthlySpendLimit *float64 `json:"monthlySpendLimit,omitempty" tf:"monthly_spend_limit,omitempty"`

	// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
	UsageReportS3Bucket *string `json:"usageReportS3Bucket,omitempty" tf:"usage_report_s3_bucket,omitempty"`
}

type SMSPreferencesParameters struct {

	// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
	// +kubebuilder:validation:Optional
	DefaultSMSType *string `json:"defaultSmsType,omitempty" tf:"default_sms_type,omitempty"`

	// A string, such as your business brand, that is displayed as the sender on the receiving device.
	// +kubebuilder:validation:Optional
	DefaultSenderID *string `json:"defaultSenderId,omitempty" tf:"default_sender_id,omitempty"`

	// The ARN of the IAM role that allows Amazon SNS to write logs about SMS deliveries in CloudWatch Logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	DeliveryStatusIAMRoleArn *string `json:"deliveryStatusIamRoleArn,omitempty" tf:"delivery_status_iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate deliveryStatusIamRoleArn.
	// +kubebuilder:validation:Optional
	DeliveryStatusIAMRoleArnRef *v1.Reference `json:"deliveryStatusIamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate deliveryStatusIamRoleArn.
	// +kubebuilder:validation:Optional
	DeliveryStatusIAMRoleArnSelector *v1.Selector `json:"deliveryStatusIamRoleArnSelector,omitempty" tf:"-"`

	// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
	// +kubebuilder:validation:Optional
	DeliveryStatusSuccessSamplingRate *string `json:"deliveryStatusSuccessSamplingRate,omitempty" tf:"delivery_status_success_sampling_rate,omitempty"`

	// The maximum amount in USD that you are willing to spend each month to send SMS messages.
	// +kubebuilder:validation:Optional
	MonthlySpendLimit *float64 `json:"monthlySpendLimit,omitempty" tf:"monthly_spend_limit,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
	// +kubebuilder:validation:Optional
	UsageReportS3Bucket *string `json:"usageReportS3Bucket,omitempty" tf:"usage_report_s3_bucket,omitempty"`
}

// SMSPreferencesSpec defines the desired state of SMSPreferences
type SMSPreferencesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SMSPreferencesParameters `json:"forProvider"`
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
	InitProvider SMSPreferencesInitParameters `json:"initProvider,omitempty"`
}

// SMSPreferencesStatus defines the observed state of SMSPreferences.
type SMSPreferencesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SMSPreferencesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SMSPreferences is the Schema for the SMSPreferencess API. Provides a way to set SNS SMS preferences.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SMSPreferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SMSPreferencesSpec   `json:"spec"`
	Status            SMSPreferencesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SMSPreferencesList contains a list of SMSPreferencess
type SMSPreferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SMSPreferences `json:"items"`
}

// Repository type metadata.
var (
	SMSPreferences_Kind             = "SMSPreferences"
	SMSPreferences_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SMSPreferences_Kind}.String()
	SMSPreferences_KindAPIVersion   = SMSPreferences_Kind + "." + CRDGroupVersion.String()
	SMSPreferences_GroupVersionKind = CRDGroupVersion.WithKind(SMSPreferences_Kind)
)

func init() {
	SchemeBuilder.Register(&SMSPreferences{}, &SMSPreferencesList{})
}
