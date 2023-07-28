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

type CloudFormationStackInitParameters struct {

	// The ARN of the application from the Serverless Application Repository.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// A list of capabilities. Valid values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_RESOURCE_POLICY, or CAPABILITY_AUTO_EXPAND
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// The name of the stack to create. The resource deployed in AWS will be prefixed with serverlessrepo-
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion *string `json:"semanticVersion,omitempty" tf:"semantic_version,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudFormationStackObservation struct {

	// The ARN of the application from the Serverless Application Repository.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// A list of capabilities. Valid values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_RESOURCE_POLICY, or CAPABILITY_AUTO_EXPAND
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// A unique identifier of the stack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the stack to create. The resource deployed in AWS will be prefixed with serverlessrepo-
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of outputs from the stack.
	Outputs map[string]*string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion *string `json:"semanticVersion,omitempty" tf:"semantic_version,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CloudFormationStackParameters struct {

	// The ARN of the application from the Serverless Application Repository.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// A list of capabilities. Valid values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_RESOURCE_POLICY, or CAPABILITY_AUTO_EXPAND
	// +kubebuilder:validation:Optional
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// The name of the stack to create. The resource deployed in AWS will be prefixed with serverlessrepo-
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The version of the application to deploy. If not supplied, deploys the latest version.
	// +kubebuilder:validation:Optional
	SemanticVersion *string `json:"semanticVersion,omitempty" tf:"semantic_version,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CloudFormationStackSpec defines the desired state of CloudFormationStack
type CloudFormationStackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudFormationStackParameters `json:"forProvider"`
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
	InitProvider CloudFormationStackInitParameters `json:"initProvider,omitempty"`
}

// CloudFormationStackStatus defines the observed state of CloudFormationStack.
type CloudFormationStackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudFormationStackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudFormationStack is the Schema for the CloudFormationStacks API. Deploys an Application CloudFormation Stack from the Serverless Application Repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudFormationStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationId) || has(self.initProvider.applicationId)",message="applicationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.capabilities) || has(self.initProvider.capabilities)",message="capabilities is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   CloudFormationStackSpec   `json:"spec"`
	Status CloudFormationStackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudFormationStackList contains a list of CloudFormationStacks
type CloudFormationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudFormationStack `json:"items"`
}

// Repository type metadata.
var (
	CloudFormationStack_Kind             = "CloudFormationStack"
	CloudFormationStack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudFormationStack_Kind}.String()
	CloudFormationStack_KindAPIVersion   = CloudFormationStack_Kind + "." + CRDGroupVersion.String()
	CloudFormationStack_GroupVersionKind = CRDGroupVersion.WithKind(CloudFormationStack_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudFormationStack{}, &CloudFormationStackList{})
}
