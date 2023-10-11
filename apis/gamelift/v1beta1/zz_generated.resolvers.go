/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Build.
func (mg *Build) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StorageLocation[i3].BucketRef,
			Selector:     mg.Spec.ForProvider.StorageLocation[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Bucket")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Key),
			Extract:      resource.ExtractParamPath("key", false),
			Reference:    mg.Spec.ForProvider.StorageLocation[i3].KeyRef,
			Selector:     mg.Spec.ForProvider.StorageLocation[i3].KeySelector,
			To: reference.To{
				List:    &v1beta1.ObjectList{},
				Managed: &v1beta1.Object{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Key")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].KeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.StorageLocation[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].RoleArn")
		}
		mg.Spec.ForProvider.StorageLocation[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Fleet.
func (mg *Fleet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BuildIDRef,
		Selector:     mg.Spec.ForProvider.BuildIDSelector,
		To: reference.To{
			List:    &BuildList{},
			Managed: &Build{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BuildID")
	}
	mg.Spec.ForProvider.BuildID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BuildIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.InstanceRoleArnRef,
		Selector:     mg.Spec.ForProvider.InstanceRoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceRoleArn")
	}
	mg.Spec.ForProvider.InstanceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GameSessionQueue.
func (mg *GameSessionQueue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NotificationTarget),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.NotificationTargetRef,
		Selector:     mg.Spec.ForProvider.NotificationTargetSelector,
		To: reference.To{
			List:    &v1beta12.TopicList{},
			Managed: &v1beta12.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NotificationTarget")
	}
	mg.Spec.ForProvider.NotificationTarget = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NotificationTargetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Script.
func (mg *Script) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StorageLocation[i3].BucketRef,
			Selector:     mg.Spec.ForProvider.StorageLocation[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Bucket")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Key),
			Extract:      resource.ExtractParamPath("key", false),
			Reference:    mg.Spec.ForProvider.StorageLocation[i3].KeyRef,
			Selector:     mg.Spec.ForProvider.StorageLocation[i3].KeySelector,
			To: reference.To{
				List:    &v1beta1.ObjectList{},
				Managed: &v1beta1.Object{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Key")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].KeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.StorageLocation[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].RoleArn")
		}
		mg.Spec.ForProvider.StorageLocation[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef = rsp.ResolvedReference

	}

	return nil
}
