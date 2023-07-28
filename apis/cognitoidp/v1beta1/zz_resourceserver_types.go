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

type ResourceServerInitParameters struct {

	// An identifier for the resource server.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// A name for the resource server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of Authorization Scope.
	Scope []ScopeInitParameters `json:"scope,omitempty" tf:"scope,omitempty"`
}

type ResourceServerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identifier for the resource server.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// A name for the resource server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of Authorization Scope.
	Scope []ScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// A list of all scopes configured for this resource server in the format identifier/scope_name.
	ScopeIdentifiers []*string `json:"scopeIdentifiers,omitempty" tf:"scope_identifiers,omitempty"`

	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type ResourceServerParameters struct {

	// An identifier for the resource server.
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// A name for the resource server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of Authorization Scope.
	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

type ScopeInitParameters struct {

	// The scope description.
	ScopeDescription *string `json:"scopeDescription,omitempty" tf:"scope_description,omitempty"`

	// The scope name.
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

type ScopeObservation struct {

	// The scope description.
	ScopeDescription *string `json:"scopeDescription,omitempty" tf:"scope_description,omitempty"`

	// The scope name.
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

type ScopeParameters struct {

	// The scope description.
	// +kubebuilder:validation:Optional
	ScopeDescription *string `json:"scopeDescription,omitempty" tf:"scope_description,omitempty"`

	// The scope name.
	// +kubebuilder:validation:Optional
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

// ResourceServerSpec defines the desired state of ResourceServer
type ResourceServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceServerParameters `json:"forProvider"`
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
	InitProvider ResourceServerInitParameters `json:"initProvider,omitempty"`
}

// ResourceServerStatus defines the observed state of ResourceServer.
type ResourceServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceServer is the Schema for the ResourceServers API. Provides a Cognito Resource Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identifier) || has(self.initProvider.identifier)",message="identifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ResourceServerSpec   `json:"spec"`
	Status ResourceServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceServerList contains a list of ResourceServers
type ResourceServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceServer `json:"items"`
}

// Repository type metadata.
var (
	ResourceServer_Kind             = "ResourceServer"
	ResourceServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceServer_Kind}.String()
	ResourceServer_KindAPIVersion   = ResourceServer_Kind + "." + CRDGroupVersion.String()
	ResourceServer_GroupVersionKind = CRDGroupVersion.WithKind(ResourceServer_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceServer{}, &ResourceServerList{})
}
