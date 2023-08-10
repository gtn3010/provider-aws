/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PrivateDNSNamespace.
func (mg *PrivateDNSNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPC),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCRef,
		Selector:     mg.Spec.ForProvider.VPCSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPC")
	}
	mg.Spec.ForProvider.VPC = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DNSConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DNSConfig[i3].NamespaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DNSConfig[i3].NamespaceIDRef,
			Selector:     mg.Spec.ForProvider.DNSConfig[i3].NamespaceIDSelector,
			To: reference.To{
				List:    &PrivateDNSNamespaceList{},
				Managed: &PrivateDNSNamespace{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DNSConfig[i3].NamespaceID")
		}
		mg.Spec.ForProvider.DNSConfig[i3].NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DNSConfig[i3].NamespaceIDRef = rsp.ResolvedReference

	}

	return nil
}