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

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terrajet/pkg/terraform"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/grafana/crossplane-provider-grafana/apis/v1alpha1"
)

const (
	keyURL           = "url"
	keyAuth          = "auth"
	keyOrgID         = "org_id"
	keyCloudAPIKey   = "cloud_api_key"
	keySMAccessToken = "sm_access_token"

	// Grafana credentials environment variable names
	envAuth          = "GRAFANA_AUTH"
	envCloudAPIKey   = "GRAFANA_CLOUD_API_KEY"
	envSMAccessToken = "GRAFANA_SM_ACCESS_TOKEN"
)

const (
	fmtEnvVar = "%s=%s"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal grafana credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
// nolint: gocyclo
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		grafanaCreds := map[string]string{}
		// Connections export configs as direct secret attributes
		if pc.Spec.Credentials.Source == "Connection" {
			var err error
			grafanaCreds, err = ExtractFullSecret(ctx, client, pc.Spec.Credentials.ConnectionSecretRef)
			if err != nil {
				return ps, errors.Wrap(err, errExtractCredentials)
			}
		} else {
			data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractCredentials)
			}
			if err := json.Unmarshal(data, &grafanaCreds); err != nil {
				return ps, errors.Wrap(err, errUnmarshalCredentials)
			}
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{}
		if url, ok := grafanaCreds[keyURL]; ok {
			ps.Configuration[keyURL] = url
		}

		// default to  OrgID 1 if not specified
		ps.Configuration[keyOrgID] = 1
		orgid := pc.Spec.OrgID
		if orgid > 0 {
			ps.Configuration[keyOrgID] = orgid
		}

		// set environment variables for sensitive provider configuration
		if auth, ok := grafanaCreds[keyAuth]; ok {
			ps.Env = append(ps.Env, fmt.Sprintf(fmtEnvVar, envAuth, auth))
		}
		if cloudAPIKey, ok := grafanaCreds[keyCloudAPIKey]; ok {
			ps.Env = append(ps.Env, fmt.Sprintf(fmtEnvVar, envCloudAPIKey, cloudAPIKey))
		}
		if smAccessToken, ok := grafanaCreds[keySMAccessToken]; ok {
			ps.Env = append(ps.Env, fmt.Sprintf(fmtEnvVar, envSMAccessToken, smAccessToken))
		}

		return ps, nil
	}
}

// ExtractFullSecret extracts credentials from a Kubernetes secret.
func ExtractFullSecret(ctx context.Context, client client.Client, s *xpv1.SecretReference) (map[string]string, error) {
	if s == nil {
		return nil, errors.New("no secret reference provided")
	}
	secret := &corev1.Secret{}
	if err := client.Get(ctx, types.NamespacedName{Namespace: s.Namespace, Name: s.Name}, secret); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced secret")
	}
	data := map[string]string{}
	for k, v := range secret.Data {
		data[k] = string(v)
	}

	return data, nil
}
