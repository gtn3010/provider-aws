/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type ApplyMethod string

const (
	ApplyMethod_immediate      ApplyMethod = "immediate"
	ApplyMethod_pending_reboot ApplyMethod = "pending-reboot"
)

type SourceType string

const (
	SourceType_db_instance         SourceType = "db-instance"
	SourceType_db_parameter_group  SourceType = "db-parameter-group"
	SourceType_db_security_group   SourceType = "db-security-group"
	SourceType_db_snapshot         SourceType = "db-snapshot"
	SourceType_db_cluster          SourceType = "db-cluster"
	SourceType_db_cluster_snapshot SourceType = "db-cluster-snapshot"
)