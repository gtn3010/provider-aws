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

type AuthInitParameters struct {

	// The type of authentication that the proxy uses for connections from the proxy to the underlying database. One of SECRETS.
	AuthScheme *string `json:"authScheme,omitempty" tf:"auth_scheme,omitempty"`

	// The type of authentication the proxy uses for connections from clients. Valid values are MYSQL_NATIVE_PASSWORD, POSTGRES_SCRAM_SHA_256, POSTGRES_MD5, and SQL_SERVER_AUTHENTICATION.
	ClientPasswordAuthType *string `json:"clientPasswordAuthType,omitempty" tf:"client_password_auth_type,omitempty"`

	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. One of DISABLED, REQUIRED.
	IAMAuth *string `json:"iamAuth,omitempty" tf:"iam_auth,omitempty"`

	// The name of the database user to which the proxy connects.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthObservation struct {

	// The type of authentication that the proxy uses for connections from the proxy to the underlying database. One of SECRETS.
	AuthScheme *string `json:"authScheme,omitempty" tf:"auth_scheme,omitempty"`

	// The type of authentication the proxy uses for connections from clients. Valid values are MYSQL_NATIVE_PASSWORD, POSTGRES_SCRAM_SHA_256, POSTGRES_MD5, and SQL_SERVER_AUTHENTICATION.
	ClientPasswordAuthType *string `json:"clientPasswordAuthType,omitempty" tf:"client_password_auth_type,omitempty"`

	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. One of DISABLED, REQUIRED.
	IAMAuth *string `json:"iamAuth,omitempty" tf:"iam_auth,omitempty"`

	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	// The name of the database user to which the proxy connects.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type AuthParameters struct {

	// The type of authentication that the proxy uses for connections from the proxy to the underlying database. One of SECRETS.
	// +kubebuilder:validation:Optional
	AuthScheme *string `json:"authScheme,omitempty" tf:"auth_scheme,omitempty"`

	// The type of authentication the proxy uses for connections from clients. Valid values are MYSQL_NATIVE_PASSWORD, POSTGRES_SCRAM_SHA_256, POSTGRES_MD5, and SQL_SERVER_AUTHENTICATION.
	// +kubebuilder:validation:Optional
	ClientPasswordAuthType *string `json:"clientPasswordAuthType,omitempty" tf:"client_password_auth_type,omitempty"`

	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. One of DISABLED, REQUIRED.
	// +kubebuilder:validation:Optional
	IAMAuth *string `json:"iamAuth,omitempty" tf:"iam_auth,omitempty"`

	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnRef *v1.Reference `json:"secretArnRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnSelector *v1.Selector `json:"secretArnSelector,omitempty" tf:"-"`

	// The name of the database user to which the proxy connects.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ProxyInitParameters struct {

	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auth []AuthInitParameters `json:"auth,omitempty" tf:"auth,omitempty"`

	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging *bool `json:"debugLogging,omitempty" tf:"debug_logging,omitempty"`

	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are MYSQL and POSTGRESQL.
	EngineFamily *string `json:"engineFamily,omitempty" tf:"engine_family,omitempty"`

	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *float64 `json:"idleClientTimeout,omitempty" tf:"idle_client_timeout,omitempty"`

	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTLS *bool `json:"requireTls,omitempty" tf:"require_tls,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	VPCSubnetIds []*string `json:"vpcSubnetIds,omitempty" tf:"vpc_subnet_ids,omitempty"`
}

type ProxyObservation struct {

	// The Amazon Resource Name (ARN) for the proxy.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auth []AuthObservation `json:"auth,omitempty" tf:"auth,omitempty"`

	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging *bool `json:"debugLogging,omitempty" tf:"debug_logging,omitempty"`

	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are MYSQL and POSTGRESQL.
	EngineFamily *string `json:"engineFamily,omitempty" tf:"engine_family,omitempty"`

	// The Amazon Resource Name (ARN) for the proxy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *float64 `json:"idleClientTimeout,omitempty" tf:"idle_client_timeout,omitempty"`

	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTLS *bool `json:"requireTls,omitempty" tf:"require_tls,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// One or more VPC security group IDs to associate with the new proxy.
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	VPCSubnetIds []*string `json:"vpcSubnetIds,omitempty" tf:"vpc_subnet_ids,omitempty"`
}

type ProxyParameters struct {

	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	// +kubebuilder:validation:Optional
	Auth []AuthParameters `json:"auth,omitempty" tf:"auth,omitempty"`

	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	// +kubebuilder:validation:Optional
	DebugLogging *bool `json:"debugLogging,omitempty" tf:"debug_logging,omitempty"`

	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are MYSQL and POSTGRESQL.
	// +kubebuilder:validation:Optional
	EngineFamily *string `json:"engineFamily,omitempty" tf:"engine_family,omitempty"`

	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	// +kubebuilder:validation:Optional
	IdleClientTimeout *float64 `json:"idleClientTimeout,omitempty" tf:"idle_client_timeout,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	// +kubebuilder:validation:Optional
	RequireTLS *bool `json:"requireTls,omitempty" tf:"require_tls,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
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

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// One or more VPC security group IDs to associate with the new proxy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	// +kubebuilder:validation:Optional
	VPCSubnetIds []*string `json:"vpcSubnetIds,omitempty" tf:"vpc_subnet_ids,omitempty"`
}

// ProxySpec defines the desired state of Proxy
type ProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyParameters `json:"forProvider"`
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
	InitProvider ProxyInitParameters `json:"initProvider,omitempty"`
}

// ProxyStatus defines the observed state of Proxy.
type ProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Proxy is the Schema for the Proxys API. Provides an RDS DB proxy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.auth) || has(self.initProvider.auth)",message="auth is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineFamily) || has(self.initProvider.engineFamily)",message="engineFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcSubnetIds) || has(self.initProvider.vpcSubnetIds)",message="vpcSubnetIds is a required parameter"
	Spec   ProxySpec   `json:"spec"`
	Status ProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyList contains a list of Proxys
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proxy `json:"items"`
}

// Repository type metadata.
var (
	Proxy_Kind             = "Proxy"
	Proxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Proxy_Kind}.String()
	Proxy_KindAPIVersion   = Proxy_Kind + "." + CRDGroupVersion.String()
	Proxy_GroupVersionKind = CRDGroupVersion.WithKind(Proxy_Kind)
)

func init() {
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}
