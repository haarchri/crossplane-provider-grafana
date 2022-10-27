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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ContactPoint
func (mg *ContactPoint) GetTerraformResourceType() string {
	return "grafana_contact_point"
}

// GetConnectionDetailsMapping for this ContactPoint
func (tr *ContactPoint) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"alertmanager[*].basic_auth_password": "spec.forProvider.alertmanager[*].basicAuthPasswordSecretRef", "alertmanager[*].settings": "spec.forProvider.alertmanager[*].settingsSecretRef", "dingding[*].settings": "spec.forProvider.dingding[*].settingsSecretRef", "discord[*].settings": "spec.forProvider.discord[*].settingsSecretRef", "discord[*].url": "spec.forProvider.discord[*].urlSecretRef", "email[*].settings": "spec.forProvider.email[*].settingsSecretRef", "googlechat[*].settings": "spec.forProvider.googlechat[*].settingsSecretRef", "googlechat[*].url": "spec.forProvider.googlechat[*].urlSecretRef", "kafka[*].rest_proxy_url": "spec.forProvider.kafka[*].restProxyUrlSecretRef", "kafka[*].settings": "spec.forProvider.kafka[*].settingsSecretRef", "opsgenie[*].api_key": "spec.forProvider.opsgenie[*].apiKeySecretRef", "opsgenie[*].settings": "spec.forProvider.opsgenie[*].settingsSecretRef", "pagerduty[*].integration_key": "spec.forProvider.pagerduty[*].integrationKeySecretRef", "pagerduty[*].settings": "spec.forProvider.pagerduty[*].settingsSecretRef", "pushover[*].api_token": "spec.forProvider.pushover[*].apiTokenSecretRef", "pushover[*].settings": "spec.forProvider.pushover[*].settingsSecretRef", "pushover[*].user_key": "spec.forProvider.pushover[*].userKeySecretRef", "sensugo[*].api_key": "spec.forProvider.sensugo[*].apiKeySecretRef", "sensugo[*].settings": "spec.forProvider.sensugo[*].settingsSecretRef", "slack[*].settings": "spec.forProvider.slack[*].settingsSecretRef", "slack[*].token": "spec.forProvider.slack[*].tokenSecretRef", "slack[*].url": "spec.forProvider.slack[*].urlSecretRef", "teams[*].settings": "spec.forProvider.teams[*].settingsSecretRef", "teams[*].url": "spec.forProvider.teams[*].urlSecretRef", "telegram[*].settings": "spec.forProvider.telegram[*].settingsSecretRef", "telegram[*].token": "spec.forProvider.telegram[*].tokenSecretRef", "threema[*].api_secret": "spec.forProvider.threema[*].apiSecretSecretRef", "threema[*].settings": "spec.forProvider.threema[*].settingsSecretRef", "victorops[*].settings": "spec.forProvider.victorops[*].settingsSecretRef", "webhook[*].authorization_credentials": "spec.forProvider.webhook[*].authorizationCredentialsSecretRef", "webhook[*].basic_auth_password": "spec.forProvider.webhook[*].basicAuthPasswordSecretRef", "webhook[*].settings": "spec.forProvider.webhook[*].settingsSecretRef", "wecom[*].settings": "spec.forProvider.wecom[*].settingsSecretRef", "wecom[*].url": "spec.forProvider.wecom[*].urlSecretRef"}
}

// GetObservation of this ContactPoint
func (tr *ContactPoint) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ContactPoint
func (tr *ContactPoint) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ContactPoint
func (tr *ContactPoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ContactPoint
func (tr *ContactPoint) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ContactPoint
func (tr *ContactPoint) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ContactPoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ContactPoint) LateInitialize(attrs []byte) (bool, error) {
	params := &ContactPointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ContactPoint) GetTerraformSchemaVersion() int {
	return 0
}
