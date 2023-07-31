/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-aws/apis/cognitoidp/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LB.
func (mg *LB) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AccessLogs); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessLogs[i3].Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AccessLogs[i3].BucketRef,
			Selector:     mg.Spec.ForProvider.AccessLogs[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AccessLogs[i3].Bucket")
		}
		mg.Spec.ForProvider.AccessLogs[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AccessLogs[i3].BucketRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SubnetMapping); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetMapping[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta11.SubnetList{},
				Managed: &v1beta11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SubnetMapping[i3].SubnetID")
		}
		mg.Spec.ForProvider.SubnetMapping[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetRefs,
		Selector:      mg.Spec.ForProvider.SubnetSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this LBListener.
func (mg *LBListener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultAction); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DefaultAction[i3].Forward); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].Arn),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].ArnRef,
					Selector:     mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].ArnSelector,
					To: reference.To{
						List:    &LBTargetGroupList{},
						Managed: &LBTargetGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].Arn")
				}
				mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.DefaultAction[i3].Forward[i4].TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultAction); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnRef,
			Selector:     mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnSelector,
			To: reference.To{
				List:    &LBTargetGroupList{},
				Managed: &LBTargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerArnRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerArnSelector,
		To: reference.To{
			List:    &LBList{},
			Managed: &LB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerArn")
	}
	mg.Spec.ForProvider.LoadBalancerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBListenerRule.
func (mg *LBListenerRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Action[i3].AuthenticateCognito); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolArnRef,
				Selector:     mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolArnSelector,
				To: reference.To{
					List:    &v1beta12.UserPoolList{},
					Managed: &v1beta12.UserPool{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolArn")
			}
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Action[i3].AuthenticateCognito); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolClientID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolClientIDRef,
				Selector:     mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolClientIDSelector,
				To: reference.To{
					List:    &v1beta12.UserPoolClientList{},
					Managed: &v1beta12.UserPoolClient{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolClientID")
			}
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolClientID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolClientIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Action[i3].AuthenticateCognito); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolDomain),
				Extract:      resource.ExtractParamPath("domain", false),
				Reference:    mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolDomainRef,
				Selector:     mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolDomainSelector,
				To: reference.To{
					List:    &v1beta12.UserPoolDomainList{},
					Managed: &v1beta12.UserPoolDomain{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolDomain")
			}
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolDomain = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito[i4].UserPoolDomainRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Action[i3].Forward); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup[i5].Arn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup[i5].ArnRef,
					Selector:     mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup[i5].ArnSelector,
					To: reference.To{
						List:    &LBTargetGroupList{},
						Managed: &LBTargetGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup[i5].Arn")
				}
				mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Action[i3].Forward[i4].TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].TargetGroupArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Action[i3].TargetGroupArnRef,
			Selector:     mg.Spec.ForProvider.Action[i3].TargetGroupArnSelector,
			To: reference.To{
				List:    &LBTargetGroupList{},
				Managed: &LBTargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.Action[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Action[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ListenerArnRef,
		Selector:     mg.Spec.ForProvider.ListenerArnSelector,
		To: reference.To{
			List:    &LBListenerList{},
			Managed: &LBListener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerArn")
	}
	mg.Spec.ForProvider.ListenerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBTargetGroup.
func (mg *LBTargetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCList{},
			Managed: &v1beta11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetGroupArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TargetGroupArnRef,
		Selector:     mg.Spec.ForProvider.TargetGroupArnSelector,
		To: reference.To{
			List:    &LBTargetGroupList{},
			Managed: &LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetGroupArn")
	}
	mg.Spec.ForProvider.TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetGroupArnRef = rsp.ResolvedReference

	return nil
}