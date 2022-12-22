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

type CapacitySpecificationObservation struct {
}

type CapacitySpecificationParameters struct {

	// The throughput capacity specified for read operations defined in read capacity units (RCUs).
	// +kubebuilder:validation:Optional
	ReadCapacityUnits *float64 `json:"readCapacityUnits,omitempty" tf:"read_capacity_units,omitempty"`

	// The read/write throughput capacity mode for a table. Valid values: PAY_PER_REQUEST, PROVISIONED. The default value is PAY_PER_REQUEST.
	// +kubebuilder:validation:Optional
	ThroughputMode *string `json:"throughputMode,omitempty" tf:"throughput_mode,omitempty"`

	// The throughput capacity specified for write operations defined in write capacity units (WCUs).
	// +kubebuilder:validation:Optional
	WriteCapacityUnits *float64 `json:"writeCapacityUnits,omitempty" tf:"write_capacity_units,omitempty"`
}

type ClusteringKeyObservation struct {
}

type ClusteringKeyParameters struct {

	// The name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The order modifier. Valid values: ASC, DESC.
	// +kubebuilder:validation:Required
	OrderBy *string `json:"orderBy" tf:"order_by,omitempty"`
}

type ColumnObservation struct {
}

type ColumnParameters struct {

	// The name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The encryption option specified for the table. Valid values: AWS_OWNED_KMS_KEY, CUSTOMER_MANAGED_KMS_KEY. The default value is AWS_OWNED_KMS_KEY.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type CommentObservation struct {
}

type CommentParameters struct {

	// A description of the table.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`
}

type EncryptionSpecificationObservation struct {
}

type EncryptionSpecificationParameters struct {

	// The Amazon Resource Name (ARN) of the customer managed KMS key.
	// +kubebuilder:validation:Optional
	KMSKeyIdentifier *string `json:"kmsKeyIdentifier,omitempty" tf:"kms_key_identifier,omitempty"`

	// The encryption option specified for the table. Valid values: AWS_OWNED_KMS_KEY, CUSTOMER_MANAGED_KMS_KEY. The default value is AWS_OWNED_KMS_KEY.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PartitionKeyObservation struct {
}

type PartitionKeyParameters struct {

	// The name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type PointInTimeRecoveryObservation struct {
}

type PointInTimeRecoveryParameters struct {

	// Valid values: ENABLED, DISABLED. The default value is DISABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SchemaDefinitionObservation struct {
}

type SchemaDefinitionParameters struct {

	// The columns that are part of the clustering key of the table.
	// +kubebuilder:validation:Optional
	ClusteringKey []ClusteringKeyParameters `json:"clusteringKey,omitempty" tf:"clustering_key,omitempty"`

	// The regular columns of the table.
	// +kubebuilder:validation:Required
	Column []ColumnParameters `json:"column" tf:"column,omitempty"`

	// The columns that are part of the partition key of the table .
	// +kubebuilder:validation:Required
	PartitionKey []PartitionKeyParameters `json:"partitionKey" tf:"partition_key,omitempty"`

	// The columns that have been defined as STATIC. Static columns store values that are shared by all rows in the same partition.
	// +kubebuilder:validation:Optional
	StaticColumn []StaticColumnParameters `json:"staticColumn,omitempty" tf:"static_column,omitempty"`
}

type StaticColumnObservation struct {
}

type StaticColumnParameters struct {

	// The name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type TTLObservation struct {
}

type TTLParameters struct {

	// Valid values: ENABLED, DISABLED. The default value is DISABLED.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type TableObservation struct {

	// The ARN of the table.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TableParameters struct {

	// Specifies the read/write throughput capacity mode for the table.
	// +kubebuilder:validation:Optional
	CapacitySpecification []CapacitySpecificationParameters `json:"capacitySpecification,omitempty" tf:"capacity_specification,omitempty"`

	// A description of the table.
	// +kubebuilder:validation:Optional
	Comment []CommentParameters `json:"comment,omitempty" tf:"comment,omitempty"`

	// The default Time to Live setting in seconds for the table. More information can be found in the Developer Guide.
	// +kubebuilder:validation:Optional
	DefaultTimeToLive *float64 `json:"defaultTimeToLive,omitempty" tf:"default_time_to_live,omitempty"`

	// Specifies how the encryption key for encryption at rest is managed for the table. More information can be found in the Developer Guide.
	// +kubebuilder:validation:Optional
	EncryptionSpecification []EncryptionSpecificationParameters `json:"encryptionSpecification,omitempty" tf:"encryption_specification,omitempty"`

	// The name of the keyspace that the table is going to be created in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/keyspaces/v1beta1.Keyspace
	// +kubebuilder:validation:Optional
	KeyspaceName *string `json:"keyspaceName,omitempty" tf:"keyspace_name,omitempty"`

	// Reference to a Keyspace in keyspaces to populate keyspaceName.
	// +kubebuilder:validation:Optional
	KeyspaceNameRef *v1.Reference `json:"keyspaceNameRef,omitempty" tf:"-"`

	// Selector for a Keyspace in keyspaces to populate keyspaceName.
	// +kubebuilder:validation:Optional
	KeyspaceNameSelector *v1.Selector `json:"keyspaceNameSelector,omitempty" tf:"-"`

	// Specifies if point-in-time recovery is enabled or disabled for the table. More information can be found in the Developer Guide.
	// +kubebuilder:validation:Optional
	PointInTimeRecovery []PointInTimeRecoveryParameters `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Describes the schema of the table.
	// +kubebuilder:validation:Required
	SchemaDefinition []SchemaDefinitionParameters `json:"schemaDefinition" tf:"schema_definition,omitempty"`

	// Enables Time to Live custom settings for the table. More information can be found in the Developer Guide.
	// +kubebuilder:validation:Optional
	TTL []TTLParameters `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The name of the table.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Table is the Schema for the Tables API. Provides a Keyspaces Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}