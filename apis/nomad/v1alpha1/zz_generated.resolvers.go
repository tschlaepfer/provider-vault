/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SecretRole.
func (mg *SecretRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      resource.ExtractParamPath("backend", false),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &SecretBackendList{},
			Managed: &SecretBackend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Backend),
		Extract:      resource.ExtractParamPath("backend", false),
		Reference:    mg.Spec.InitProvider.BackendRef,
		Selector:     mg.Spec.InitProvider.BackendSelector,
		To: reference.To{
			List:    &SecretBackendList{},
			Managed: &SecretBackend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Backend")
	}
	mg.Spec.InitProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackendRef = rsp.ResolvedReference

	return nil
}
