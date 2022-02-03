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

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terrajet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-jet-grafana/apis/v1alpha1"
)

const (
	keyUrl           = "url"
	keyAuth          = "auth"
	keyOrgId         = "org_id"
	keyCloudAPIKey   = "cloud_api_key"
	keySMAccessToken = "sm_access_token"

	// Grafana credentials environment variable names
	envUrl           = "GRAFANA_URL"
	envAuth          = "GRAFANA_AUTH"
	envOrgId         = "GRAFANA_ORG_ID"
	envCloudAPIKey   = "GRAFANA_CLOUD_API_URL"
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

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		grafanaCreds := map[string]string{}
		if err := json.Unmarshal(data, &grafanaCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{}
		if url, ok := grafanaCreds[keyUrl]; ok {
			ps.Configuration[keyUrl] = url
		}
		if orgId, ok := grafanaCreds[keyOrgId]; ok {
			ps.Configuration[keyOrgId] = orgId
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
