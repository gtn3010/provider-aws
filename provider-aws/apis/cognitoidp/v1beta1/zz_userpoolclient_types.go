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

type AnalyticsConfigurationObservation struct {
}

type AnalyticsConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationArn *string `json:"applicationArn,omitempty" tf:"application_arn,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	UserDataShared *bool `json:"userDataShared,omitempty" tf:"user_data_shared,omitempty"`
}

type TokenValidityUnitsObservation struct {
}

type TokenValidityUnitsParameters struct {

	// +kubebuilder:validation:Optional
	AccessToken *string `json:"accessToken,omitempty" tf:"access_token,omitempty"`

	// +kubebuilder:validation:Optional
	IDToken *string `json:"idToken,omitempty" tf:"id_token,omitempty"`

	// +kubebuilder:validation:Optional
	RefreshToken *string `json:"refreshToken,omitempty" tf:"refresh_token,omitempty"`
}

type UserPoolClientObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserPoolClientParameters struct {

	// +kubebuilder:validation:Optional
	AccessTokenValidity *float64 `json:"accessTokenValidity,omitempty" tf:"access_token_validity,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedOauthFlows []*string `json:"allowedOauthFlows,omitempty" tf:"allowed_oauth_flows,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedOauthFlowsUserPoolClient *bool `json:"allowedOauthFlowsUserPoolClient,omitempty" tf:"allowed_oauth_flows_user_pool_client,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedOauthScopes []*string `json:"allowedOauthScopes,omitempty" tf:"allowed_oauth_scopes,omitempty"`

	// +kubebuilder:validation:Optional
	AnalyticsConfiguration []AnalyticsConfigurationParameters `json:"analyticsConfiguration,omitempty" tf:"analytics_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	CallbackUrls []*string `json:"callbackUrls,omitempty" tf:"callback_urls,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTokenRevocation *bool `json:"enableTokenRevocation,omitempty" tf:"enable_token_revocation,omitempty"`

	// +kubebuilder:validation:Optional
	ExplicitAuthFlows []*string `json:"explicitAuthFlows,omitempty" tf:"explicit_auth_flows,omitempty"`

	// +kubebuilder:validation:Optional
	GenerateSecret *bool `json:"generateSecret,omitempty" tf:"generate_secret,omitempty"`

	// +kubebuilder:validation:Optional
	IDTokenValidity *float64 `json:"idTokenValidity,omitempty" tf:"id_token_validity,omitempty"`

	// +kubebuilder:validation:Optional
	LogoutUrls []*string `json:"logoutUrls,omitempty" tf:"logout_urls,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors,omitempty" tf:"prevent_user_existence_errors,omitempty"`

	// +kubebuilder:validation:Optional
	ReadAttributes []*string `json:"readAttributes,omitempty" tf:"read_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	RefreshTokenValidity *float64 `json:"refreshTokenValidity,omitempty" tf:"refresh_token_validity,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SupportedIdentityProviders []*string `json:"supportedIdentityProviders,omitempty" tf:"supported_identity_providers,omitempty"`

	// +kubebuilder:validation:Optional
	TokenValidityUnits []TokenValidityUnitsParameters `json:"tokenValidityUnits,omitempty" tf:"token_validity_units,omitempty"`

	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	WriteAttributes []*string `json:"writeAttributes,omitempty" tf:"write_attributes,omitempty"`
}

// UserPoolClientSpec defines the desired state of UserPoolClient
type UserPoolClientSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolClientParameters `json:"forProvider"`
}

// UserPoolClientStatus defines the observed state of UserPoolClient.
type UserPoolClientStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolClient is the Schema for the UserPoolClients API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserPoolClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolClientSpec   `json:"spec"`
	Status            UserPoolClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolClientList contains a list of UserPoolClients
type UserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPoolClient `json:"items"`
}

// Repository type metadata.
var (
	UserPoolClient_Kind             = "UserPoolClient"
	UserPoolClient_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPoolClient_Kind}.String()
	UserPoolClient_KindAPIVersion   = UserPoolClient_Kind + "." + CRDGroupVersion.String()
	UserPoolClient_GroupVersionKind = CRDGroupVersion.WithKind(UserPoolClient_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPoolClient{}, &UserPoolClientList{})
}
