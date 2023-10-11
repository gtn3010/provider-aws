// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type AccessControlAllowHeadersInitParameters struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowHeadersObservation struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowHeadersParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowMethodsInitParameters struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowMethodsObservation struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowMethodsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowOriginsInitParameters struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowOriginsObservation struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowOriginsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlExposeHeadersInitParameters struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlExposeHeadersObservation struct {
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlExposeHeadersParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type ContentSecurityPolicyInitParameters struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	ContentSecurityPolicy *string `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`
}

type ContentSecurityPolicyObservation struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	ContentSecurityPolicy *string `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`
}

type ContentSecurityPolicyParameters struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	// +kubebuilder:validation:Optional
	ContentSecurityPolicy *string `json:"contentSecurityPolicy" tf:"content_security_policy,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`
}

type ContentTypeOptionsInitParameters struct {

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`
}

type ContentTypeOptionsObservation struct {

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`
}

type ContentTypeOptionsParameters struct {

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`
}

type CorsConfigInitParameters struct {

	// A Boolean value that CloudFront uses as the value for the Access-Control-Allow-Credentials HTTP response header.
	AccessControlAllowCredentials *bool `json:"accessControlAllowCredentials,omitempty" tf:"access_control_allow_credentials,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-Headers HTTP response header.
	AccessControlAllowHeaders []AccessControlAllowHeadersInitParameters `json:"accessControlAllowHeaders,omitempty" tf:"access_control_allow_headers,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-Methods HTTP response header. Valid values: GET | POST | OPTIONS | PUT | DELETE | HEAD | ALL
	AccessControlAllowMethods []AccessControlAllowMethodsInitParameters `json:"accessControlAllowMethods,omitempty" tf:"access_control_allow_methods,omitempty"`

	// Object that contains an attribute items that contains a list of origins that CloudFront can use as the value for the Access-Control-Allow-Origin HTTP response header.
	AccessControlAllowOrigins []AccessControlAllowOriginsInitParameters `json:"accessControlAllowOrigins,omitempty" tf:"access_control_allow_origins,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-Headers HTTP response header.
	AccessControlExposeHeaders []AccessControlExposeHeadersInitParameters `json:"accessControlExposeHeaders,omitempty" tf:"access_control_expose_headers,omitempty"`

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec,omitempty"`

	// A Boolean value that determines how CloudFront behaves for the HTTP response header.
	OriginOverride *bool `json:"originOverride,omitempty" tf:"origin_override,omitempty"`
}

type CorsConfigObservation struct {

	// A Boolean value that CloudFront uses as the value for the Access-Control-Allow-Credentials HTTP response header.
	AccessControlAllowCredentials *bool `json:"accessControlAllowCredentials,omitempty" tf:"access_control_allow_credentials,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-Headers HTTP response header.
	AccessControlAllowHeaders []AccessControlAllowHeadersObservation `json:"accessControlAllowHeaders,omitempty" tf:"access_control_allow_headers,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-Methods HTTP response header. Valid values: GET | POST | OPTIONS | PUT | DELETE | HEAD | ALL
	AccessControlAllowMethods []AccessControlAllowMethodsObservation `json:"accessControlAllowMethods,omitempty" tf:"access_control_allow_methods,omitempty"`

	// Object that contains an attribute items that contains a list of origins that CloudFront can use as the value for the Access-Control-Allow-Origin HTTP response header.
	AccessControlAllowOrigins []AccessControlAllowOriginsObservation `json:"accessControlAllowOrigins,omitempty" tf:"access_control_allow_origins,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-Headers HTTP response header.
	AccessControlExposeHeaders []AccessControlExposeHeadersObservation `json:"accessControlExposeHeaders,omitempty" tf:"access_control_expose_headers,omitempty"`

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec,omitempty"`

	// A Boolean value that determines how CloudFront behaves for the HTTP response header.
	OriginOverride *bool `json:"originOverride,omitempty" tf:"origin_override,omitempty"`
}

type CorsConfigParameters struct {

	// A Boolean value that CloudFront uses as the value for the Access-Control-Allow-Credentials HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlAllowCredentials *bool `json:"accessControlAllowCredentials" tf:"access_control_allow_credentials,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-Headers HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlAllowHeaders []AccessControlAllowHeadersParameters `json:"accessControlAllowHeaders" tf:"access_control_allow_headers,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-Methods HTTP response header. Valid values: GET | POST | OPTIONS | PUT | DELETE | HEAD | ALL
	// +kubebuilder:validation:Optional
	AccessControlAllowMethods []AccessControlAllowMethodsParameters `json:"accessControlAllowMethods" tf:"access_control_allow_methods,omitempty"`

	// Object that contains an attribute items that contains a list of origins that CloudFront can use as the value for the Access-Control-Allow-Origin HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlAllowOrigins []AccessControlAllowOriginsParameters `json:"accessControlAllowOrigins" tf:"access_control_allow_origins,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-Headers HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlExposeHeaders []AccessControlExposeHeadersParameters `json:"accessControlExposeHeaders,omitempty" tf:"access_control_expose_headers,omitempty"`

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec,omitempty"`

	// A Boolean value that determines how CloudFront behaves for the HTTP response header.
	// +kubebuilder:validation:Optional
	OriginOverride *bool `json:"originOverride" tf:"origin_override,omitempty"`
}

type CustomHeadersConfigInitParameters struct {
	Items []CustomHeadersConfigItemsInitParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type CustomHeadersConfigItemsInitParameters struct {

	// The HTTP response header name.
	Header *string `json:"header,omitempty" tf:"header,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// The value for the HTTP response header.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeadersConfigItemsObservation struct {

	// The HTTP response header name.
	Header *string `json:"header,omitempty" tf:"header,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// The value for the HTTP response header.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomHeadersConfigItemsParameters struct {

	// The HTTP response header name.
	// +kubebuilder:validation:Optional
	Header *string `json:"header" tf:"header,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`

	// The value for the HTTP response header.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type CustomHeadersConfigObservation struct {
	Items []CustomHeadersConfigItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type CustomHeadersConfigParameters struct {

	// +kubebuilder:validation:Optional
	Items []CustomHeadersConfigItemsParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type FrameOptionsInitParameters struct {

	// The value of the X-Frame-Options HTTP response header. Valid values: DENY | SAMEORIGIN
	FrameOption *string `json:"frameOption,omitempty" tf:"frame_option,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`
}

type FrameOptionsObservation struct {

	// The value of the X-Frame-Options HTTP response header. Valid values: DENY | SAMEORIGIN
	FrameOption *string `json:"frameOption,omitempty" tf:"frame_option,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`
}

type FrameOptionsParameters struct {

	// The value of the X-Frame-Options HTTP response header. Valid values: DENY | SAMEORIGIN
	// +kubebuilder:validation:Optional
	FrameOption *string `json:"frameOption" tf:"frame_option,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`
}

type ReferrerPolicyInitParameters struct {

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	ReferrerPolicy *string `json:"referrerPolicy,omitempty" tf:"referrer_policy,omitempty"`
}

type ReferrerPolicyObservation struct {

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	ReferrerPolicy *string `json:"referrerPolicy,omitempty" tf:"referrer_policy,omitempty"`
}

type ReferrerPolicyParameters struct {

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	// +kubebuilder:validation:Optional
	ReferrerPolicy *string `json:"referrerPolicy" tf:"referrer_policy,omitempty"`
}

type RemoveHeadersConfigInitParameters struct {
	Items []RemoveHeadersConfigItemsInitParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type RemoveHeadersConfigItemsInitParameters struct {

	// The HTTP response header name.
	Header *string `json:"header,omitempty" tf:"header,omitempty"`
}

type RemoveHeadersConfigItemsObservation struct {

	// The HTTP response header name.
	Header *string `json:"header,omitempty" tf:"header,omitempty"`
}

type RemoveHeadersConfigItemsParameters struct {

	// The HTTP response header name.
	// +kubebuilder:validation:Optional
	Header *string `json:"header" tf:"header,omitempty"`
}

type RemoveHeadersConfigObservation struct {
	Items []RemoveHeadersConfigItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type RemoveHeadersConfigParameters struct {

	// +kubebuilder:validation:Optional
	Items []RemoveHeadersConfigItemsParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type ResponseHeadersPolicyInitParameters struct {

	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig []CorsConfigInitParameters `json:"corsConfig,omitempty" tf:"cors_config,omitempty"`

	// Object that contains an attribute items that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig []CustomHeadersConfigInitParameters `json:"customHeadersConfig,omitempty" tf:"custom_headers_config,omitempty"`

	// The current version of the response headers policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A unique name to identify the response headers policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A configuration for a set of HTTP headers to remove from the HTTP response. Object that contains an attribute items that contains a list of headers. See Remove Header for more information.
	RemoveHeadersConfig []RemoveHeadersConfigInitParameters `json:"removeHeadersConfig,omitempty" tf:"remove_headers_config,omitempty"`

	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig []SecurityHeadersConfigInitParameters `json:"securityHeadersConfig,omitempty" tf:"security_headers_config,omitempty"`

	// A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
	ServerTimingHeadersConfig []ServerTimingHeadersConfigInitParameters `json:"serverTimingHeadersConfig,omitempty" tf:"server_timing_headers_config,omitempty"`
}

type ResponseHeadersPolicyObservation struct {

	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig []CorsConfigObservation `json:"corsConfig,omitempty" tf:"cors_config,omitempty"`

	// Object that contains an attribute items that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig []CustomHeadersConfigObservation `json:"customHeadersConfig,omitempty" tf:"custom_headers_config,omitempty"`

	// The current version of the response headers policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The identifier for the response headers policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name to identify the response headers policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A configuration for a set of HTTP headers to remove from the HTTP response. Object that contains an attribute items that contains a list of headers. See Remove Header for more information.
	RemoveHeadersConfig []RemoveHeadersConfigObservation `json:"removeHeadersConfig,omitempty" tf:"remove_headers_config,omitempty"`

	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig []SecurityHeadersConfigObservation `json:"securityHeadersConfig,omitempty" tf:"security_headers_config,omitempty"`

	// A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
	ServerTimingHeadersConfig []ServerTimingHeadersConfigObservation `json:"serverTimingHeadersConfig,omitempty" tf:"server_timing_headers_config,omitempty"`
}

type ResponseHeadersPolicyParameters struct {

	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	// +kubebuilder:validation:Optional
	CorsConfig []CorsConfigParameters `json:"corsConfig,omitempty" tf:"cors_config,omitempty"`

	// Object that contains an attribute items that contains a list of custom headers. See Custom Header for more information.
	// +kubebuilder:validation:Optional
	CustomHeadersConfig []CustomHeadersConfigParameters `json:"customHeadersConfig,omitempty" tf:"custom_headers_config,omitempty"`

	// The current version of the response headers policy.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A unique name to identify the response headers policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A configuration for a set of HTTP headers to remove from the HTTP response. Object that contains an attribute items that contains a list of headers. See Remove Header for more information.
	// +kubebuilder:validation:Optional
	RemoveHeadersConfig []RemoveHeadersConfigParameters `json:"removeHeadersConfig,omitempty" tf:"remove_headers_config,omitempty"`

	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	// +kubebuilder:validation:Optional
	SecurityHeadersConfig []SecurityHeadersConfigParameters `json:"securityHeadersConfig,omitempty" tf:"security_headers_config,omitempty"`

	// A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
	// +kubebuilder:validation:Optional
	ServerTimingHeadersConfig []ServerTimingHeadersConfigParameters `json:"serverTimingHeadersConfig,omitempty" tf:"server_timing_headers_config,omitempty"`
}

type SecurityHeadersConfigInitParameters struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	ContentSecurityPolicy []ContentSecurityPolicyInitParameters `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy,omitempty"`

	// Determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff. See Content Type Options for more information.
	ContentTypeOptions []ContentTypeOptionsInitParameters `json:"contentTypeOptions,omitempty" tf:"content_type_options,omitempty"`

	// Determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value. See Frame Options for more information.
	FrameOptions []FrameOptionsInitParameters `json:"frameOptions,omitempty" tf:"frame_options,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	ReferrerPolicy []ReferrerPolicyInitParameters `json:"referrerPolicy,omitempty" tf:"referrer_policy,omitempty"`

	// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value. See Strict Transport Security for more information.
	StrictTransportSecurity []StrictTransportSecurityInitParameters `json:"strictTransportSecurity,omitempty" tf:"strict_transport_security,omitempty"`

	// Determine whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value. See XSS Protection for more information.
	XSSProtection []XSSProtectionInitParameters `json:"xssProtection,omitempty" tf:"xss_protection,omitempty"`
}

type SecurityHeadersConfigObservation struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	ContentSecurityPolicy []ContentSecurityPolicyObservation `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy,omitempty"`

	// Determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff. See Content Type Options for more information.
	ContentTypeOptions []ContentTypeOptionsObservation `json:"contentTypeOptions,omitempty" tf:"content_type_options,omitempty"`

	// Determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value. See Frame Options for more information.
	FrameOptions []FrameOptionsObservation `json:"frameOptions,omitempty" tf:"frame_options,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	ReferrerPolicy []ReferrerPolicyObservation `json:"referrerPolicy,omitempty" tf:"referrer_policy,omitempty"`

	// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value. See Strict Transport Security for more information.
	StrictTransportSecurity []StrictTransportSecurityObservation `json:"strictTransportSecurity,omitempty" tf:"strict_transport_security,omitempty"`

	// Determine whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value. See XSS Protection for more information.
	XSSProtection []XSSProtectionObservation `json:"xssProtection,omitempty" tf:"xss_protection,omitempty"`
}

type SecurityHeadersConfigParameters struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	// +kubebuilder:validation:Optional
	ContentSecurityPolicy []ContentSecurityPolicyParameters `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy,omitempty"`

	// Determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff. See Content Type Options for more information.
	// +kubebuilder:validation:Optional
	ContentTypeOptions []ContentTypeOptionsParameters `json:"contentTypeOptions,omitempty" tf:"content_type_options,omitempty"`

	// Determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value. See Frame Options for more information.
	// +kubebuilder:validation:Optional
	FrameOptions []FrameOptionsParameters `json:"frameOptions,omitempty" tf:"frame_options,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	// +kubebuilder:validation:Optional
	ReferrerPolicy []ReferrerPolicyParameters `json:"referrerPolicy,omitempty" tf:"referrer_policy,omitempty"`

	// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value. See Strict Transport Security for more information.
	// +kubebuilder:validation:Optional
	StrictTransportSecurity []StrictTransportSecurityParameters `json:"strictTransportSecurity,omitempty" tf:"strict_transport_security,omitempty"`

	// Determine whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value. See XSS Protection for more information.
	// +kubebuilder:validation:Optional
	XSSProtection []XSSProtectionParameters `json:"xssProtection,omitempty" tf:"xss_protection,omitempty"`
}

type ServerTimingHeadersConfigInitParameters struct {

	// A Whether CloudFront adds the Server-Timing header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the Server-Timing header to. Valid range: Minimum value of 0.0. Maximum value of 100.0.
	SamplingRate *float64 `json:"samplingRate,omitempty" tf:"sampling_rate,omitempty"`
}

type ServerTimingHeadersConfigObservation struct {

	// A Whether CloudFront adds the Server-Timing header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the Server-Timing header to. Valid range: Minimum value of 0.0. Maximum value of 100.0.
	SamplingRate *float64 `json:"samplingRate,omitempty" tf:"sampling_rate,omitempty"`
}

type ServerTimingHeadersConfigParameters struct {

	// A Whether CloudFront adds the Server-Timing header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// A number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the Server-Timing header to. Valid range: Minimum value of 0.0. Maximum value of 100.0.
	// +kubebuilder:validation:Optional
	SamplingRate *float64 `json:"samplingRate" tf:"sampling_rate,omitempty"`
}

type StrictTransportSecurityInitParameters struct {

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec,omitempty"`

	// Whether CloudFront includes the includeSubDomains directive in the Strict-Transport-Security HTTP response header.
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// Whether CloudFront includes the preload directive in the Strict-Transport-Security HTTP response header.
	Preload *bool `json:"preload,omitempty" tf:"preload,omitempty"`
}

type StrictTransportSecurityObservation struct {

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec,omitempty"`

	// Whether CloudFront includes the includeSubDomains directive in the Strict-Transport-Security HTTP response header.
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// Whether CloudFront includes the preload directive in the Strict-Transport-Security HTTP response header.
	Preload *bool `json:"preload,omitempty" tf:"preload,omitempty"`
}

type StrictTransportSecurityParameters struct {

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec" tf:"access_control_max_age_sec,omitempty"`

	// Whether CloudFront includes the includeSubDomains directive in the Strict-Transport-Security HTTP response header.
	// +kubebuilder:validation:Optional
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`

	// Whether CloudFront includes the preload directive in the Strict-Transport-Security HTTP response header.
	// +kubebuilder:validation:Optional
	Preload *bool `json:"preload,omitempty" tf:"preload,omitempty"`
}

type XSSProtectionInitParameters struct {

	// Whether CloudFront includes the mode=block directive in the X-XSS-Protection header.
	ModeBlock *bool `json:"modeBlock,omitempty" tf:"mode_block,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// A Boolean value that determines the value of the X-XSS-Protection HTTP response header. When this setting is true, the value of the X-XSS-Protection header is 1. When this setting is false, the value of the X-XSS-Protection header is 0.
	Protection *bool `json:"protection,omitempty" tf:"protection,omitempty"`

	// A reporting URI, which CloudFront uses as the value of the report directive in the X-XSS-Protection header. You cannot specify a report_uri when mode_block is true.
	ReportURI *string `json:"reportUri,omitempty" tf:"report_uri,omitempty"`
}

type XSSProtectionObservation struct {

	// Whether CloudFront includes the mode=block directive in the X-XSS-Protection header.
	ModeBlock *bool `json:"modeBlock,omitempty" tf:"mode_block,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// A Boolean value that determines the value of the X-XSS-Protection HTTP response header. When this setting is true, the value of the X-XSS-Protection header is 1. When this setting is false, the value of the X-XSS-Protection header is 0.
	Protection *bool `json:"protection,omitempty" tf:"protection,omitempty"`

	// A reporting URI, which CloudFront uses as the value of the report directive in the X-XSS-Protection header. You cannot specify a report_uri when mode_block is true.
	ReportURI *string `json:"reportUri,omitempty" tf:"report_uri,omitempty"`
}

type XSSProtectionParameters struct {

	// Whether CloudFront includes the mode=block directive in the X-XSS-Protection header.
	// +kubebuilder:validation:Optional
	ModeBlock *bool `json:"modeBlock,omitempty" tf:"mode_block,omitempty"`

	// Whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override" tf:"override,omitempty"`

	// A Boolean value that determines the value of the X-XSS-Protection HTTP response header. When this setting is true, the value of the X-XSS-Protection header is 1. When this setting is false, the value of the X-XSS-Protection header is 0.
	// +kubebuilder:validation:Optional
	Protection *bool `json:"protection" tf:"protection,omitempty"`

	// A reporting URI, which CloudFront uses as the value of the report directive in the X-XSS-Protection header. You cannot specify a report_uri when mode_block is true.
	// +kubebuilder:validation:Optional
	ReportURI *string `json:"reportUri,omitempty" tf:"report_uri,omitempty"`
}

// ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy
type ResponseHeadersPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResponseHeadersPolicyParameters `json:"forProvider"`
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
	InitProvider ResponseHeadersPolicyInitParameters `json:"initProvider,omitempty"`
}

// ResponseHeadersPolicyStatus defines the observed state of ResponseHeadersPolicy.
type ResponseHeadersPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResponseHeadersPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicys API. Provides a CloudFront response headers policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResponseHeadersPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ResponseHeadersPolicySpec   `json:"spec"`
	Status ResponseHeadersPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResponseHeadersPolicyList contains a list of ResponseHeadersPolicys
type ResponseHeadersPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResponseHeadersPolicy `json:"items"`
}

// Repository type metadata.
var (
	ResponseHeadersPolicy_Kind             = "ResponseHeadersPolicy"
	ResponseHeadersPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResponseHeadersPolicy_Kind}.String()
	ResponseHeadersPolicy_KindAPIVersion   = ResponseHeadersPolicy_Kind + "." + CRDGroupVersion.String()
	ResponseHeadersPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ResponseHeadersPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ResponseHeadersPolicy{}, &ResponseHeadersPolicyList{})
}
