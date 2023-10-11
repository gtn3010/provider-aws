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
	v1beta11 "github.com/upbound/provider-aws/apis/ecr/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta14 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	v1beta15 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Component.
func (mg *Component) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContainerRecipe.
func (mg *ContainerRecipe) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Component); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Component[i3].ComponentArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Component[i3].ComponentArnRef,
			Selector:     mg.Spec.ForProvider.Component[i3].ComponentArnSelector,
			To: reference.To{
				List:    &ComponentList{},
				Managed: &Component{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Component[i3].ComponentArn")
		}
		mg.Spec.ForProvider.Component[i3].ComponentArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Component[i3].ComponentArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.TargetRepository); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetRepository[i3].RepositoryName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TargetRepository[i3].RepositoryNameRef,
			Selector:     mg.Spec.ForProvider.TargetRepository[i3].RepositoryNameSelector,
			To: reference.To{
				List:    &v1beta11.RepositoryList{},
				Managed: &v1beta11.Repository{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.TargetRepository[i3].RepositoryName")
		}
		mg.Spec.ForProvider.TargetRepository[i3].RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.TargetRepository[i3].RepositoryNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Image.
func (mg *Image) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DistributionConfigurationArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.DistributionConfigurationArnRef,
		Selector:     mg.Spec.ForProvider.DistributionConfigurationArnSelector,
		To: reference.To{
			List:    &DistributionConfigurationList{},
			Managed: &DistributionConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DistributionConfigurationArn")
	}
	mg.Spec.ForProvider.DistributionConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DistributionConfigurationArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageRecipeArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ImageRecipeArnRef,
		Selector:     mg.Spec.ForProvider.ImageRecipeArnSelector,
		To: reference.To{
			List:    &ImageRecipeList{},
			Managed: &ImageRecipe{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageRecipeArn")
	}
	mg.Spec.ForProvider.ImageRecipeArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageRecipeArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InfrastructureConfigurationArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.InfrastructureConfigurationArnRef,
		Selector:     mg.Spec.ForProvider.InfrastructureConfigurationArnSelector,
		To: reference.To{
			List:    &InfrastructureConfigurationList{},
			Managed: &InfrastructureConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InfrastructureConfigurationArn")
	}
	mg.Spec.ForProvider.InfrastructureConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InfrastructureConfigurationArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ImagePipeline.
func (mg *ImagePipeline) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageRecipeArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ImageRecipeArnRef,
		Selector:     mg.Spec.ForProvider.ImageRecipeArnSelector,
		To: reference.To{
			List:    &ImageRecipeList{},
			Managed: &ImageRecipe{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageRecipeArn")
	}
	mg.Spec.ForProvider.ImageRecipeArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageRecipeArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InfrastructureConfigurationArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.InfrastructureConfigurationArnRef,
		Selector:     mg.Spec.ForProvider.InfrastructureConfigurationArnSelector,
		To: reference.To{
			List:    &InfrastructureConfigurationList{},
			Managed: &InfrastructureConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InfrastructureConfigurationArn")
	}
	mg.Spec.ForProvider.InfrastructureConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InfrastructureConfigurationArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ImageRecipe.
func (mg *ImageRecipe) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Component); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Component[i3].ComponentArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Component[i3].ComponentArnRef,
			Selector:     mg.Spec.ForProvider.Component[i3].ComponentArnSelector,
			To: reference.To{
				List:    &ComponentList{},
				Managed: &Component{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Component[i3].ComponentArn")
		}
		mg.Spec.ForProvider.Component[i3].ComponentArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Component[i3].ComponentArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this InfrastructureConfiguration.
func (mg *InfrastructureConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceProfileName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceProfileNameRef,
		Selector:     mg.Spec.ForProvider.InstanceProfileNameSelector,
		To: reference.To{
			List:    &v1beta12.InstanceProfileList{},
			Managed: &v1beta12.InstanceProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceProfileName")
	}
	mg.Spec.ForProvider.InstanceProfileName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceProfileNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1beta13.KeyPairList{},
			Managed: &v1beta13.KeyPair{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Logging); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Logging[i3].S3Logs); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Logging[i3].S3Logs[i4].S3BucketName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Logging[i3].S3Logs[i4].S3BucketNameRef,
				Selector:     mg.Spec.ForProvider.Logging[i3].S3Logs[i4].S3BucketNameSelector,
				To: reference.To{
					List:    &v1beta14.BucketList{},
					Managed: &v1beta14.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Logging[i3].S3Logs[i4].S3BucketName")
			}
			mg.Spec.ForProvider.Logging[i3].S3Logs[i4].S3BucketName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Logging[i3].S3Logs[i4].S3BucketNameRef = rsp.ResolvedReference

		}
	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta13.SecurityGroupList{},
			Managed: &v1beta13.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsTopicArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SnsTopicArnRef,
		Selector:     mg.Spec.ForProvider.SnsTopicArnSelector,
		To: reference.To{
			List:    &v1beta15.TopicList{},
			Managed: &v1beta15.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnsTopicArn")
	}
	mg.Spec.ForProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnsTopicArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta13.SubnetList{},
			Managed: &v1beta13.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}
