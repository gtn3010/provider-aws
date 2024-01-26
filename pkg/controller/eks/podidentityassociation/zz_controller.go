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

package podidentityassociation

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/eks"
	svcsdk "github.com/aws/aws-sdk-go/service/eks"
	svcsdkapi "github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/eks/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an PodIdentityAssociation resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create PodIdentityAssociation in AWS"
	errUpdate        = "cannot update PodIdentityAssociation in AWS"
	errDescribe      = "failed to describe PodIdentityAssociation"
	errDelete        = "failed to delete PodIdentityAssociation"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.PodIdentityAssociation)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.PodIdentityAssociation)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribePodIdentityAssociationInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribePodIdentityAssociationWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GeneratePodIdentityAssociation(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.PodIdentityAssociation)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreatePodIdentityAssociationInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreatePodIdentityAssociationWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.Association.AssociationArn != nil {
		cr.Status.AtProvider.AssociationARN = resp.Association.AssociationArn
	} else {
		cr.Status.AtProvider.AssociationARN = nil
	}
	if resp.Association.AssociationId != nil {
		cr.Status.AtProvider.AssociationID = resp.Association.AssociationId
	} else {
		cr.Status.AtProvider.AssociationID = nil
	}
	if resp.Association.ClusterName != nil {
		cr.Spec.ForProvider.ClusterName = resp.Association.ClusterName
	} else {
		cr.Spec.ForProvider.ClusterName = nil
	}
	if resp.Association.CreatedAt != nil {
		cr.Status.AtProvider.CreatedAt = &metav1.Time{*resp.Association.CreatedAt}
	} else {
		cr.Status.AtProvider.CreatedAt = nil
	}
	if resp.Association.ModifiedAt != nil {
		cr.Status.AtProvider.ModifiedAt = &metav1.Time{*resp.Association.ModifiedAt}
	} else {
		cr.Status.AtProvider.ModifiedAt = nil
	}
	if resp.Association.Namespace != nil {
		cr.Spec.ForProvider.Namespace = resp.Association.Namespace
	} else {
		cr.Spec.ForProvider.Namespace = nil
	}
	if resp.Association.RoleArn != nil {
		cr.Spec.ForProvider.RoleARN = resp.Association.RoleArn
	} else {
		cr.Spec.ForProvider.RoleARN = nil
	}
	if resp.Association.ServiceAccount != nil {
		cr.Spec.ForProvider.ServiceAccount = resp.Association.ServiceAccount
	} else {
		cr.Spec.ForProvider.ServiceAccount = nil
	}
	if resp.Association.Tags != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range resp.Association.Tags {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		cr.Spec.ForProvider.Tags = f8
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.PodIdentityAssociation)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdatePodIdentityAssociationInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdatePodIdentityAssociationWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.PodIdentityAssociation)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeletePodIdentityAssociationInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeletePodIdentityAssociationWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EKSAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EKSAPI
	preObserve     func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DescribePodIdentityAssociationInput) error
	postObserve    func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DescribePodIdentityAssociationOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.PodIdentityAssociationParameters, *svcsdk.DescribePodIdentityAssociationOutput) error
	isUpToDate     func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DescribePodIdentityAssociationOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.CreatePodIdentityAssociationInput) error
	postCreate     func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.CreatePodIdentityAssociationOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DeletePodIdentityAssociationInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DeletePodIdentityAssociationOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.UpdatePodIdentityAssociationInput) error
	postUpdate     func(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.UpdatePodIdentityAssociationOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DescribePodIdentityAssociationInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.PodIdentityAssociation, _ *svcsdk.DescribePodIdentityAssociationOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.PodIdentityAssociationParameters, *svcsdk.DescribePodIdentityAssociationOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DescribePodIdentityAssociationOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.CreatePodIdentityAssociationInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.PodIdentityAssociation, _ *svcsdk.CreatePodIdentityAssociationOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.DeletePodIdentityAssociationInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.PodIdentityAssociation, _ *svcsdk.DeletePodIdentityAssociationOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.PodIdentityAssociation, *svcsdk.UpdatePodIdentityAssociationInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.PodIdentityAssociation, _ *svcsdk.UpdatePodIdentityAssociationOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
