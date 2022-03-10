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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	config "github.com/grafana/crossplane-provider-grafana/config"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this APIKey.
func (mg *APIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudStackSlug),
		Extract:      config.SlugExtractor(),
		Reference:    mg.Spec.ForProvider.CloudStackRef,
		Selector:     mg.Spec.ForProvider.CloudStackSelector,
		To: reference.To{
			List:    &CloudStackList{},
			Managed: &CloudStack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudStackSlug")
	}
	mg.Spec.ForProvider.CloudStackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudStackRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CloudPluginInstallation.
func (mg *CloudPluginInstallation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackSlug),
		Extract:      config.SlugExtractor(),
		Reference:    mg.Spec.ForProvider.StackRef,
		Selector:     mg.Spec.ForProvider.StackSelector,
		To: reference.To{
			List:    &CloudStackList{},
			Managed: &CloudStack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackSlug")
	}
	mg.Spec.ForProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StackRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Dashboard.
func (mg *Dashboard) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Folder),
		Extract:      config.IDExtractor(),
		Reference:    mg.Spec.ForProvider.FolderRef,
		Selector:     mg.Spec.ForProvider.FolderSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Folder")
	}
	mg.Spec.ForProvider.Folder = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	return nil
}
