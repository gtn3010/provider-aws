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

type ProfileInitParameters struct {

	// The number of seconds the vended session credentials are valid for. Defaults to 3600.
	DurationSeconds *float64 `json:"durationSeconds,omitempty" tf:"duration_seconds,omitempty"`

	// Whether or not the Profile is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// The name of the Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies whether instance properties are required in CreateSession requests with this profile.
	RequireInstanceProperties *bool `json:"requireInstanceProperties,omitempty" tf:"require_instance_properties,omitempty"`

	// A session policy that applies to the trust boundary of the vended session credentials.
	SessionPolicy *string `json:"sessionPolicy,omitempty" tf:"session_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProfileObservation struct {

	// Amazon Resource Name (ARN) of the Profile
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The number of seconds the vended session credentials are valid for. Defaults to 3600.
	DurationSeconds *float64 `json:"durationSeconds,omitempty" tf:"duration_seconds,omitempty"`

	// Whether or not the Profile is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Profile ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// The name of the Profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies whether instance properties are required in CreateSession requests with this profile.
	RequireInstanceProperties *bool `json:"requireInstanceProperties,omitempty" tf:"require_instance_properties,omitempty"`

	// A list of IAM roles that this profile can assume
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// A session policy that applies to the trust boundary of the vended session credentials.
	SessionPolicy *string `json:"sessionPolicy,omitempty" tf:"session_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProfileParameters struct {

	// The number of seconds the vended session credentials are valid for. Defaults to 3600.
	// +kubebuilder:validation:Optional
	DurationSeconds *float64 `json:"durationSeconds,omitempty" tf:"duration_seconds,omitempty"`

	// Whether or not the Profile is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of managed policy ARNs that apply to the vended session credentials.
	// +kubebuilder:validation:Optional
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// The name of the Profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies whether instance properties are required in CreateSession requests with this profile.
	// +kubebuilder:validation:Optional
	RequireInstanceProperties *bool `json:"requireInstanceProperties,omitempty" tf:"require_instance_properties,omitempty"`

	// A list of IAM roles that this profile can assume
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// References to Role in iam to populate roleArns.
	// +kubebuilder:validation:Optional
	RoleArnsRefs []v1.Reference `json:"roleArnsRefs,omitempty" tf:"-"`

	// Selector for a list of Role in iam to populate roleArns.
	// +kubebuilder:validation:Optional
	RoleArnsSelector *v1.Selector `json:"roleArnsSelector,omitempty" tf:"-"`

	// A session policy that applies to the trust boundary of the vended session credentials.
	// +kubebuilder:validation:Optional
	SessionPolicy *string `json:"sessionPolicy,omitempty" tf:"session_policy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileParameters `json:"forProvider"`
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
	InitProvider ProfileInitParameters `json:"initProvider,omitempty"`
}

// ProfileStatus defines the observed state of Profile.
type ProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Profile is the Schema for the Profiles API. Provides a Roles Anywhere Profile resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ProfileSpec   `json:"spec"`
	Status ProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileList contains a list of Profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Repository type metadata.
var (
	Profile_Kind             = "Profile"
	Profile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Profile_Kind}.String()
	Profile_KindAPIVersion   = Profile_Kind + "." + CRDGroupVersion.String()
	Profile_GroupVersionKind = CRDGroupVersion.WithKind(Profile_Kind)
)

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}