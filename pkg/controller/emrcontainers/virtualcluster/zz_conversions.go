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

package virtualcluster

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/emrcontainers"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/emrcontainers/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeVirtualClusterInput returns input for read
// operation.
func GenerateDescribeVirtualClusterInput(cr *svcapitypes.VirtualCluster) *svcsdk.DescribeVirtualClusterInput {
	res := &svcsdk.DescribeVirtualClusterInput{}

	if cr.Status.AtProvider.ID != nil {
		res.SetId(*cr.Status.AtProvider.ID)
	}

	return res
}

// GenerateVirtualCluster returns the current state in the form of *svcapitypes.VirtualCluster.
func GenerateVirtualCluster(resp *svcsdk.DescribeVirtualClusterOutput) *svcapitypes.VirtualCluster {
	cr := &svcapitypes.VirtualCluster{}

	if resp.VirtualCluster.Arn != nil {
		cr.Status.AtProvider.ARN = resp.VirtualCluster.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.VirtualCluster.ContainerProvider != nil {
		f1 := &svcapitypes.ContainerProvider{}
		if resp.VirtualCluster.ContainerProvider.Id != nil {
			f1.ID = resp.VirtualCluster.ContainerProvider.Id
		}
		if resp.VirtualCluster.ContainerProvider.Info != nil {
			f1f1 := &svcapitypes.ContainerInfo{}
			if resp.VirtualCluster.ContainerProvider.Info.EksInfo != nil {
				f1f1f0 := &svcapitypes.EKSInfo{}
				if resp.VirtualCluster.ContainerProvider.Info.EksInfo.Namespace != nil {
					f1f1f0.Namespace = resp.VirtualCluster.ContainerProvider.Info.EksInfo.Namespace
				}
				f1f1.EKSInfo = f1f1f0
			}
			f1.Info = f1f1
		}
		if resp.VirtualCluster.ContainerProvider.Type != nil {
			f1.Type = resp.VirtualCluster.ContainerProvider.Type
		}
		cr.Spec.ForProvider.ContainerProvider = f1
	} else {
		cr.Spec.ForProvider.ContainerProvider = nil
	}
	if resp.VirtualCluster.Id != nil {
		cr.Status.AtProvider.ID = resp.VirtualCluster.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.VirtualCluster.Name != nil {
		cr.Status.AtProvider.Name = resp.VirtualCluster.Name
	} else {
		cr.Status.AtProvider.Name = nil
	}
	if resp.VirtualCluster.Tags != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range resp.VirtualCluster.Tags {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		cr.Spec.ForProvider.Tags = f6
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return cr
}

// GenerateCreateVirtualClusterInput returns a create input.
func GenerateCreateVirtualClusterInput(cr *svcapitypes.VirtualCluster) *svcsdk.CreateVirtualClusterInput {
	res := &svcsdk.CreateVirtualClusterInput{}

	if cr.Spec.ForProvider.ContainerProvider != nil {
		f0 := &svcsdk.ContainerProvider{}
		if cr.Spec.ForProvider.ContainerProvider.ID != nil {
			f0.SetId(*cr.Spec.ForProvider.ContainerProvider.ID)
		}
		if cr.Spec.ForProvider.ContainerProvider.Info != nil {
			f0f1 := &svcsdk.ContainerInfo{}
			if cr.Spec.ForProvider.ContainerProvider.Info.EKSInfo != nil {
				f0f1f0 := &svcsdk.EksInfo{}
				if cr.Spec.ForProvider.ContainerProvider.Info.EKSInfo.Namespace != nil {
					f0f1f0.SetNamespace(*cr.Spec.ForProvider.ContainerProvider.Info.EKSInfo.Namespace)
				}
				f0f1.SetEksInfo(f0f1f0)
			}
			f0.SetInfo(f0f1)
		}
		if cr.Spec.ForProvider.ContainerProvider.Type != nil {
			f0.SetType(*cr.Spec.ForProvider.ContainerProvider.Type)
		}
		res.SetContainerProvider(f0)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range cr.Spec.ForProvider.Tags {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		res.SetTags(f1)
	}

	return res
}

// GenerateDeleteVirtualClusterInput returns a deletion input.
func GenerateDeleteVirtualClusterInput(cr *svcapitypes.VirtualCluster) *svcsdk.DeleteVirtualClusterInput {
	res := &svcsdk.DeleteVirtualClusterInput{}

	if cr.Status.AtProvider.ID != nil {
		res.SetId(*cr.Status.AtProvider.ID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}