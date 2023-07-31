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

type ApplicationSourceInitParameters struct {

	// ARN of a AWS CloudFormation stack.
	CloudFormationStackArn *string `json:"cloudformationStackArn,omitempty" tf:"cloudformation_stack_arn,omitempty"`

	// Set of tags.
	TagFilter []TagFilterInitParameters `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type ApplicationSourceObservation struct {

	// ARN of a AWS CloudFormation stack.
	CloudFormationStackArn *string `json:"cloudformationStackArn,omitempty" tf:"cloudformation_stack_arn,omitempty"`

	// Set of tags.
	TagFilter []TagFilterObservation `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type ApplicationSourceParameters struct {

	// ARN of a AWS CloudFormation stack.
	// +kubebuilder:validation:Optional
	CloudFormationStackArn *string `json:"cloudformationStackArn,omitempty" tf:"cloudformation_stack_arn,omitempty"`

	// Set of tags.
	// +kubebuilder:validation:Optional
	TagFilter []TagFilterParameters `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type CustomizedLoadMetricSpecificationInitParameters struct {

	// Dimensions of the metric.
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be Sum.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedLoadMetricSpecificationObservation struct {

	// Dimensions of the metric.
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be Sum.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedLoadMetricSpecificationParameters struct {

	// Dimensions of the metric.
	// +kubebuilder:validation:Optional
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be Sum.
	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedScalingMetricSpecificationInitParameters struct {

	// Dimensions of the metric.
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be Sum.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedScalingMetricSpecificationObservation struct {

	// Dimensions of the metric.
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be Sum.
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CustomizedScalingMetricSpecificationParameters struct {

	// Dimensions of the metric.
	// +kubebuilder:validation:Optional
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Name of the metric.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Namespace of the metric.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be Sum.
	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Unit of the metric.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type PredefinedLoadMetricSpecificationInitParameters struct {

	// Metric type. Valid values: ALBTargetGroupRequestCount, ASGTotalCPUUtilization, ASGTotalNetworkIn, ASGTotalNetworkOut.
	PredefinedLoadMetricType *string `json:"predefinedLoadMetricType,omitempty" tf:"predefined_load_metric_type,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedLoadMetricSpecificationObservation struct {

	// Metric type. Valid values: ALBTargetGroupRequestCount, ASGTotalCPUUtilization, ASGTotalNetworkIn, ASGTotalNetworkOut.
	PredefinedLoadMetricType *string `json:"predefinedLoadMetricType,omitempty" tf:"predefined_load_metric_type,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedLoadMetricSpecificationParameters struct {

	// Metric type. Valid values: ALBTargetGroupRequestCount, ASGTotalCPUUtilization, ASGTotalNetworkIn, ASGTotalNetworkOut.
	// +kubebuilder:validation:Optional
	PredefinedLoadMetricType *string `json:"predefinedLoadMetricType,omitempty" tf:"predefined_load_metric_type,omitempty"`

	// Identifies the resource associated with the metric type.
	// +kubebuilder:validation:Optional
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedScalingMetricSpecificationInitParameters struct {

	// Metric type. Valid values: ALBRequestCountPerTarget, ASGAverageCPUUtilization, ASGAverageNetworkIn, ASGAverageNetworkOut, DynamoDBReadCapacityUtilization, DynamoDBWriteCapacityUtilization, ECSServiceAverageCPUUtilization, ECSServiceAverageMemoryUtilization, EC2SpotFleetRequestAverageCPUUtilization, EC2SpotFleetRequestAverageNetworkIn, EC2SpotFleetRequestAverageNetworkOut, RDSReaderAverageCPUUtilization, RDSReaderAverageDatabaseConnections.
	PredefinedScalingMetricType *string `json:"predefinedScalingMetricType,omitempty" tf:"predefined_scaling_metric_type,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedScalingMetricSpecificationObservation struct {

	// Metric type. Valid values: ALBRequestCountPerTarget, ASGAverageCPUUtilization, ASGAverageNetworkIn, ASGAverageNetworkOut, DynamoDBReadCapacityUtilization, DynamoDBWriteCapacityUtilization, ECSServiceAverageCPUUtilization, ECSServiceAverageMemoryUtilization, EC2SpotFleetRequestAverageCPUUtilization, EC2SpotFleetRequestAverageNetworkIn, EC2SpotFleetRequestAverageNetworkOut, RDSReaderAverageCPUUtilization, RDSReaderAverageDatabaseConnections.
	PredefinedScalingMetricType *string `json:"predefinedScalingMetricType,omitempty" tf:"predefined_scaling_metric_type,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type PredefinedScalingMetricSpecificationParameters struct {

	// Metric type. Valid values: ALBRequestCountPerTarget, ASGAverageCPUUtilization, ASGAverageNetworkIn, ASGAverageNetworkOut, DynamoDBReadCapacityUtilization, DynamoDBWriteCapacityUtilization, ECSServiceAverageCPUUtilization, ECSServiceAverageMemoryUtilization, EC2SpotFleetRequestAverageCPUUtilization, EC2SpotFleetRequestAverageNetworkIn, EC2SpotFleetRequestAverageNetworkOut, RDSReaderAverageCPUUtilization, RDSReaderAverageDatabaseConnections.
	// +kubebuilder:validation:Optional
	PredefinedScalingMetricType *string `json:"predefinedScalingMetricType,omitempty" tf:"predefined_scaling_metric_type,omitempty"`

	// Identifies the resource associated with the metric type.
	// +kubebuilder:validation:Optional
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type ScalingInstructionInitParameters struct {

	// Customized load metric to use for predictive scaling. You must specify either customized_load_metric_specification or predefined_load_metric_specification when configuring predictive scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	CustomizedLoadMetricSpecification []CustomizedLoadMetricSpecificationInitParameters `json:"customizedLoadMetricSpecification,omitempty" tf:"customized_load_metric_specification,omitempty"`

	// Boolean controlling whether dynamic scaling by AWS Auto Scaling is disabled. Defaults to false.
	DisableDynamicScaling *bool `json:"disableDynamicScaling,omitempty" tf:"disable_dynamic_scaling,omitempty"`

	// Maximum capacity of the resource. The exception to this upper limit is if you specify a non-default setting for predictive_scaling_max_capacity_behavior.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Minimum capacity of the resource.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// Predefined load metric to use for predictive scaling. You must specify either predefined_load_metric_specification or customized_load_metric_specification when configuring predictive scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	PredefinedLoadMetricSpecification []PredefinedLoadMetricSpecificationInitParameters `json:"predefinedLoadMetricSpecification,omitempty" tf:"predefined_load_metric_specification,omitempty"`

	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	// Valid values: SetForecastCapacityToMaxCapacity, SetMaxCapacityAboveForecastCapacity, SetMaxCapacityToForecastCapacity.
	PredictiveScalingMaxCapacityBehavior *string `json:"predictiveScalingMaxCapacityBehavior,omitempty" tf:"predictive_scaling_max_capacity_behavior,omitempty"`

	// Size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	PredictiveScalingMaxCapacityBuffer *float64 `json:"predictiveScalingMaxCapacityBuffer,omitempty" tf:"predictive_scaling_max_capacity_buffer,omitempty"`

	// Predictive scaling mode. Valid values: ForecastAndScale, ForecastOnly.
	PredictiveScalingMode *string `json:"predictiveScalingMode,omitempty" tf:"predictive_scaling_mode,omitempty"`

	// ID of the resource. This string consists of the resource type and unique identifier.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Scalable dimension associated with the resource. Valid values: autoscaling:autoScalingGroup:DesiredCapacity, dynamodb:index:ReadCapacityUnits, dynamodb:index:WriteCapacityUnits, dynamodb:table:ReadCapacityUnits, dynamodb:table:WriteCapacityUnits, ecs:service:DesiredCount, ec2:spot-fleet-request:TargetCapacity, rds:cluster:ReadReplicaCount.
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// Controls whether a resource's externally created scaling policies are kept or replaced. Valid values: KeepExternalPolicies, ReplaceExternalPolicies. Defaults to KeepExternalPolicies.
	ScalingPolicyUpdateBehavior *string `json:"scalingPolicyUpdateBehavior,omitempty" tf:"scaling_policy_update_behavior,omitempty"`

	// Amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	ScheduledActionBufferTime *float64 `json:"scheduledActionBufferTime,omitempty" tf:"scheduled_action_buffer_time,omitempty"`

	// Namespace of the AWS service. Valid values: autoscaling, dynamodb, ecs, ec2, rds.
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Structure that defines new target tracking configurations. Each of these structures includes a specific scaling metric and a target value for the metric, along with various parameters to use with dynamic scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	TargetTrackingConfiguration []TargetTrackingConfigurationInitParameters `json:"targetTrackingConfiguration,omitempty" tf:"target_tracking_configuration,omitempty"`
}

type ScalingInstructionObservation struct {

	// Customized load metric to use for predictive scaling. You must specify either customized_load_metric_specification or predefined_load_metric_specification when configuring predictive scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	CustomizedLoadMetricSpecification []CustomizedLoadMetricSpecificationObservation `json:"customizedLoadMetricSpecification,omitempty" tf:"customized_load_metric_specification,omitempty"`

	// Boolean controlling whether dynamic scaling by AWS Auto Scaling is disabled. Defaults to false.
	DisableDynamicScaling *bool `json:"disableDynamicScaling,omitempty" tf:"disable_dynamic_scaling,omitempty"`

	// Maximum capacity of the resource. The exception to this upper limit is if you specify a non-default setting for predictive_scaling_max_capacity_behavior.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Minimum capacity of the resource.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// Predefined load metric to use for predictive scaling. You must specify either predefined_load_metric_specification or customized_load_metric_specification when configuring predictive scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	PredefinedLoadMetricSpecification []PredefinedLoadMetricSpecificationObservation `json:"predefinedLoadMetricSpecification,omitempty" tf:"predefined_load_metric_specification,omitempty"`

	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	// Valid values: SetForecastCapacityToMaxCapacity, SetMaxCapacityAboveForecastCapacity, SetMaxCapacityToForecastCapacity.
	PredictiveScalingMaxCapacityBehavior *string `json:"predictiveScalingMaxCapacityBehavior,omitempty" tf:"predictive_scaling_max_capacity_behavior,omitempty"`

	// Size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	PredictiveScalingMaxCapacityBuffer *float64 `json:"predictiveScalingMaxCapacityBuffer,omitempty" tf:"predictive_scaling_max_capacity_buffer,omitempty"`

	// Predictive scaling mode. Valid values: ForecastAndScale, ForecastOnly.
	PredictiveScalingMode *string `json:"predictiveScalingMode,omitempty" tf:"predictive_scaling_mode,omitempty"`

	// ID of the resource. This string consists of the resource type and unique identifier.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Scalable dimension associated with the resource. Valid values: autoscaling:autoScalingGroup:DesiredCapacity, dynamodb:index:ReadCapacityUnits, dynamodb:index:WriteCapacityUnits, dynamodb:table:ReadCapacityUnits, dynamodb:table:WriteCapacityUnits, ecs:service:DesiredCount, ec2:spot-fleet-request:TargetCapacity, rds:cluster:ReadReplicaCount.
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// Controls whether a resource's externally created scaling policies are kept or replaced. Valid values: KeepExternalPolicies, ReplaceExternalPolicies. Defaults to KeepExternalPolicies.
	ScalingPolicyUpdateBehavior *string `json:"scalingPolicyUpdateBehavior,omitempty" tf:"scaling_policy_update_behavior,omitempty"`

	// Amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	ScheduledActionBufferTime *float64 `json:"scheduledActionBufferTime,omitempty" tf:"scheduled_action_buffer_time,omitempty"`

	// Namespace of the AWS service. Valid values: autoscaling, dynamodb, ecs, ec2, rds.
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Structure that defines new target tracking configurations. Each of these structures includes a specific scaling metric and a target value for the metric, along with various parameters to use with dynamic scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	TargetTrackingConfiguration []TargetTrackingConfigurationObservation `json:"targetTrackingConfiguration,omitempty" tf:"target_tracking_configuration,omitempty"`
}

type ScalingInstructionParameters struct {

	// Customized load metric to use for predictive scaling. You must specify either customized_load_metric_specification or predefined_load_metric_specification when configuring predictive scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	// +kubebuilder:validation:Optional
	CustomizedLoadMetricSpecification []CustomizedLoadMetricSpecificationParameters `json:"customizedLoadMetricSpecification,omitempty" tf:"customized_load_metric_specification,omitempty"`

	// Boolean controlling whether dynamic scaling by AWS Auto Scaling is disabled. Defaults to false.
	// +kubebuilder:validation:Optional
	DisableDynamicScaling *bool `json:"disableDynamicScaling,omitempty" tf:"disable_dynamic_scaling,omitempty"`

	// Maximum capacity of the resource. The exception to this upper limit is if you specify a non-default setting for predictive_scaling_max_capacity_behavior.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Minimum capacity of the resource.
	// +kubebuilder:validation:Optional
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// Predefined load metric to use for predictive scaling. You must specify either predefined_load_metric_specification or customized_load_metric_specification when configuring predictive scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	// +kubebuilder:validation:Optional
	PredefinedLoadMetricSpecification []PredefinedLoadMetricSpecificationParameters `json:"predefinedLoadMetricSpecification,omitempty" tf:"predefined_load_metric_specification,omitempty"`

	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	// Valid values: SetForecastCapacityToMaxCapacity, SetMaxCapacityAboveForecastCapacity, SetMaxCapacityToForecastCapacity.
	// +kubebuilder:validation:Optional
	PredictiveScalingMaxCapacityBehavior *string `json:"predictiveScalingMaxCapacityBehavior,omitempty" tf:"predictive_scaling_max_capacity_behavior,omitempty"`

	// Size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	// +kubebuilder:validation:Optional
	PredictiveScalingMaxCapacityBuffer *float64 `json:"predictiveScalingMaxCapacityBuffer,omitempty" tf:"predictive_scaling_max_capacity_buffer,omitempty"`

	// Predictive scaling mode. Valid values: ForecastAndScale, ForecastOnly.
	// +kubebuilder:validation:Optional
	PredictiveScalingMode *string `json:"predictiveScalingMode,omitempty" tf:"predictive_scaling_mode,omitempty"`

	// ID of the resource. This string consists of the resource type and unique identifier.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Scalable dimension associated with the resource. Valid values: autoscaling:autoScalingGroup:DesiredCapacity, dynamodb:index:ReadCapacityUnits, dynamodb:index:WriteCapacityUnits, dynamodb:table:ReadCapacityUnits, dynamodb:table:WriteCapacityUnits, ecs:service:DesiredCount, ec2:spot-fleet-request:TargetCapacity, rds:cluster:ReadReplicaCount.
	// +kubebuilder:validation:Optional
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// Controls whether a resource's externally created scaling policies are kept or replaced. Valid values: KeepExternalPolicies, ReplaceExternalPolicies. Defaults to KeepExternalPolicies.
	// +kubebuilder:validation:Optional
	ScalingPolicyUpdateBehavior *string `json:"scalingPolicyUpdateBehavior,omitempty" tf:"scaling_policy_update_behavior,omitempty"`

	// Amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	// +kubebuilder:validation:Optional
	ScheduledActionBufferTime *float64 `json:"scheduledActionBufferTime,omitempty" tf:"scheduled_action_buffer_time,omitempty"`

	// Namespace of the AWS service. Valid values: autoscaling, dynamodb, ecs, ec2, rds.
	// +kubebuilder:validation:Optional
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Structure that defines new target tracking configurations. Each of these structures includes a specific scaling metric and a target value for the metric, along with various parameters to use with dynamic scaling.
	// More details can be found in the AWS Auto Scaling API Reference.
	// +kubebuilder:validation:Optional
	TargetTrackingConfiguration []TargetTrackingConfigurationParameters `json:"targetTrackingConfiguration,omitempty" tf:"target_tracking_configuration,omitempty"`
}

type ScalingPlanInitParameters struct {

	// CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource []ApplicationSourceInitParameters `json:"applicationSource,omitempty" tf:"application_source,omitempty"`

	// Name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Scaling instructions. More details can be found in the AWS Auto Scaling API Reference.
	ScalingInstruction []ScalingInstructionInitParameters `json:"scalingInstruction,omitempty" tf:"scaling_instruction,omitempty"`
}

type ScalingPlanObservation struct {

	// CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource []ApplicationSourceObservation `json:"applicationSource,omitempty" tf:"application_source,omitempty"`

	// Scaling plan identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Scaling instructions. More details can be found in the AWS Auto Scaling API Reference.
	ScalingInstruction []ScalingInstructionObservation `json:"scalingInstruction,omitempty" tf:"scaling_instruction,omitempty"`

	// The version number of the scaling plan. This value is always 1.
	ScalingPlanVersion *float64 `json:"scalingPlanVersion,omitempty" tf:"scaling_plan_version,omitempty"`
}

type ScalingPlanParameters struct {

	// CloudFormation stack or set of tags. You can create one scaling plan per application source.
	// +kubebuilder:validation:Optional
	ApplicationSource []ApplicationSourceParameters `json:"applicationSource,omitempty" tf:"application_source,omitempty"`

	// Name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Scaling instructions. More details can be found in the AWS Auto Scaling API Reference.
	// +kubebuilder:validation:Optional
	ScalingInstruction []ScalingInstructionParameters `json:"scalingInstruction,omitempty" tf:"scaling_instruction,omitempty"`
}

type TagFilterInitParameters struct {

	// Tag key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TagFilterObservation struct {

	// Tag key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag values.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TagFilterParameters struct {

	// Tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag values.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TargetTrackingConfigurationInitParameters struct {

	// Customized metric. You can specify either customized_scaling_metric_specification or predefined_scaling_metric_specification.
	// More details can be found in the AWS Auto Scaling API Reference.
	CustomizedScalingMetricSpecification []CustomizedScalingMetricSpecificationInitParameters `json:"customizedScalingMetricSpecification,omitempty" tf:"customized_scaling_metric_specification,omitempty"`

	// Boolean indicating whether scale in by the target tracking scaling policy is disabled. Defaults to false.
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// Estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	// This value is used only if the resource is an Auto Scaling group.
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup,omitempty"`

	// Predefined metric. You can specify either predefined_scaling_metric_specification or customized_scaling_metric_specification.
	// More details can be found in the AWS Auto Scaling API Reference.
	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationInitParameters `json:"predefinedScalingMetricSpecification,omitempty" tf:"predefined_scaling_metric_specification,omitempty"`

	// Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleInCooldown *float64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`

	// Amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`

	// Target value for the metric.
	TargetValue *float64 `json:"targetValue,omitempty" tf:"target_value,omitempty"`
}

type TargetTrackingConfigurationObservation struct {

	// Customized metric. You can specify either customized_scaling_metric_specification or predefined_scaling_metric_specification.
	// More details can be found in the AWS Auto Scaling API Reference.
	CustomizedScalingMetricSpecification []CustomizedScalingMetricSpecificationObservation `json:"customizedScalingMetricSpecification,omitempty" tf:"customized_scaling_metric_specification,omitempty"`

	// Boolean indicating whether scale in by the target tracking scaling policy is disabled. Defaults to false.
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// Estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	// This value is used only if the resource is an Auto Scaling group.
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup,omitempty"`

	// Predefined metric. You can specify either predefined_scaling_metric_specification or customized_scaling_metric_specification.
	// More details can be found in the AWS Auto Scaling API Reference.
	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationObservation `json:"predefinedScalingMetricSpecification,omitempty" tf:"predefined_scaling_metric_specification,omitempty"`

	// Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleInCooldown *float64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`

	// Amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`

	// Target value for the metric.
	TargetValue *float64 `json:"targetValue,omitempty" tf:"target_value,omitempty"`
}

type TargetTrackingConfigurationParameters struct {

	// Customized metric. You can specify either customized_scaling_metric_specification or predefined_scaling_metric_specification.
	// More details can be found in the AWS Auto Scaling API Reference.
	// +kubebuilder:validation:Optional
	CustomizedScalingMetricSpecification []CustomizedScalingMetricSpecificationParameters `json:"customizedScalingMetricSpecification,omitempty" tf:"customized_scaling_metric_specification,omitempty"`

	// Boolean indicating whether scale in by the target tracking scaling policy is disabled. Defaults to false.
	// +kubebuilder:validation:Optional
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`

	// Estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	// This value is used only if the resource is an Auto Scaling group.
	// +kubebuilder:validation:Optional
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup,omitempty"`

	// Predefined metric. You can specify either predefined_scaling_metric_specification or customized_scaling_metric_specification.
	// More details can be found in the AWS Auto Scaling API Reference.
	// +kubebuilder:validation:Optional
	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationParameters `json:"predefinedScalingMetricSpecification,omitempty" tf:"predefined_scaling_metric_specification,omitempty"`

	// Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	// This value is not used if the scalable resource is an Auto Scaling group.
	// +kubebuilder:validation:Optional
	ScaleInCooldown *float64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`

	// Amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	// This value is not used if the scalable resource is an Auto Scaling group.
	// +kubebuilder:validation:Optional
	ScaleOutCooldown *float64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`

	// Target value for the metric.
	// +kubebuilder:validation:Optional
	TargetValue *float64 `json:"targetValue,omitempty" tf:"target_value,omitempty"`
}

// ScalingPlanSpec defines the desired state of ScalingPlan
type ScalingPlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScalingPlanParameters `json:"forProvider"`
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
	InitProvider ScalingPlanInitParameters `json:"initProvider,omitempty"`
}

// ScalingPlanStatus defines the observed state of ScalingPlan.
type ScalingPlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScalingPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingPlan is the Schema for the ScalingPlans API. Manages an AWS Auto Scaling scaling plan.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ScalingPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationSource) || has(self.initProvider.applicationSource)",message="applicationSource is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scalingInstruction) || has(self.initProvider.scalingInstruction)",message="scalingInstruction is a required parameter"
	Spec   ScalingPlanSpec   `json:"spec"`
	Status ScalingPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingPlanList contains a list of ScalingPlans
type ScalingPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalingPlan `json:"items"`
}

// Repository type metadata.
var (
	ScalingPlan_Kind             = "ScalingPlan"
	ScalingPlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScalingPlan_Kind}.String()
	ScalingPlan_KindAPIVersion   = ScalingPlan_Kind + "." + CRDGroupVersion.String()
	ScalingPlan_GroupVersionKind = CRDGroupVersion.WithKind(ScalingPlan_Kind)
)

func init() {
	SchemeBuilder.Register(&ScalingPlan{}, &ScalingPlanList{})
}