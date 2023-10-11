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
	v1beta13 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/efs/v1beta1"
	v1beta14 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta15 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/signer/v1beta1"
	v1beta17 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	v1beta16 "github.com/upbound/provider-aws/apis/sqs/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Alias.
func (mg *Alias) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FunctionNameRef,
		Selector:     mg.Spec.ForProvider.FunctionNameSelector,
		To: reference.To{
			List:    &FunctionList{},
			Managed: &Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CodeSigningConfig.
func (mg *CodeSigningConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AllowedPublishers); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.AllowedPublishers[i3].SigningProfileVersionArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.ForProvider.AllowedPublishers[i3].SigningProfileVersionArnsRefs,
			Selector:      mg.Spec.ForProvider.AllowedPublishers[i3].SigningProfileVersionArnsSelector,
			To: reference.To{
				List:    &v1beta1.SigningProfileList{},
				Managed: &v1beta1.SigningProfile{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AllowedPublishers[i3].SigningProfileVersionArns")
		}
		mg.Spec.ForProvider.AllowedPublishers[i3].SigningProfileVersionArns = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.AllowedPublishers[i3].SigningProfileVersionArnsRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this EventSourceMapping.
func (mg *EventSourceMapping) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.FunctionNameRef,
		Selector:     mg.Spec.ForProvider.FunctionNameSelector,
		To: reference.To{
			List:    &FunctionList{},
			Managed: &Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Function.
func (mg *Function) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.FileSystemConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemConfig[i3].Arn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.FileSystemConfig[i3].ArnRef,
			Selector:     mg.Spec.ForProvider.FileSystemConfig[i3].ArnSelector,
			To: reference.To{
				List:    &v1beta11.AccessPointList{},
				Managed: &v1beta11.AccessPoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemConfig[i3].Arn")
		}
		mg.Spec.ForProvider.FileSystemConfig[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.FileSystemConfig[i3].ArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1beta12.KeyList{},
			Managed: &v1beta12.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.ReplacementSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ReplacementSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.ReplacementSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta13.SecurityGroupList{},
			Managed: &v1beta13.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplacementSecurityGroupIds")
	}
	mg.Spec.ForProvider.ReplacementSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ReplacementSecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &v1beta14.RoleList{},
			Managed: &v1beta14.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.S3BucketRef,
		Selector:     mg.Spec.ForProvider.S3BucketSelector,
		To: reference.To{
			List:    &v1beta15.BucketList{},
			Managed: &v1beta15.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.S3Bucket")
	}
	mg.Spec.ForProvider.S3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.S3BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIDSelector,
			To: reference.To{
				List:    &v1beta13.SecurityGroupList{},
				Managed: &v1beta13.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIds")
		}
		mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig[i3].SecurityGroupIDRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCConfig[i3].SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.VPCConfig[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta13.SubnetList{},
				Managed: &v1beta13.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].SubnetIds")
		}
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.VPCConfig[i3].SubnetIDRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this FunctionEventInvokeConfig.
func (mg *FunctionEventInvokeConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DestinationConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DestinationConfig[i3].OnFailure); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationConfig[i3].OnFailure[i4].Destination),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.DestinationConfig[i3].OnFailure[i4].DestinationRef,
				Selector:     mg.Spec.ForProvider.DestinationConfig[i3].OnFailure[i4].DestinationSelector,
				To: reference.To{
					List:    &v1beta16.QueueList{},
					Managed: &v1beta16.Queue{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DestinationConfig[i3].OnFailure[i4].Destination")
			}
			mg.Spec.ForProvider.DestinationConfig[i3].OnFailure[i4].Destination = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DestinationConfig[i3].OnFailure[i4].DestinationRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DestinationConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess[i4].Destination),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess[i4].DestinationRef,
				Selector:     mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess[i4].DestinationSelector,
				To: reference.To{
					List:    &v1beta17.TopicList{},
					Managed: &v1beta17.Topic{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess[i4].Destination")
			}
			mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess[i4].Destination = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.DestinationConfig[i3].OnSuccess[i4].DestinationRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this FunctionURL.
func (mg *FunctionURL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FunctionNameRef,
		Selector:     mg.Spec.ForProvider.FunctionNameSelector,
		To: reference.To{
			List:    &FunctionList{},
			Managed: &Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Invocation.
func (mg *Invocation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FunctionNameRef,
		Selector:     mg.Spec.ForProvider.FunctionNameSelector,
		To: reference.To{
			List:    &FunctionList{},
			Managed: &Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Permission.
func (mg *Permission) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FunctionNameRef,
		Selector:     mg.Spec.ForProvider.FunctionNameSelector,
		To: reference.To{
			List:    &FunctionList{},
			Managed: &Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionName")
	}
	mg.Spec.ForProvider.FunctionName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Qualifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.QualifierRef,
		Selector:     mg.Spec.ForProvider.QualifierSelector,
		To: reference.To{
			List:    &AliasList{},
			Managed: &Alias{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Qualifier")
	}
	mg.Spec.ForProvider.Qualifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QualifierRef = rsp.ResolvedReference

	return nil
}
