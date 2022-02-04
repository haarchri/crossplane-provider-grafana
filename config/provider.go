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

package config

import (
	"strings"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	tjname "github.com/crossplane/terrajet/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourcePrefix = "grafana"
	modulePath     = "github.com/grafana/provider-jet-grafana"
)

// GetProvider returns provider configuration
func GetProvider(resourceMap map[string]*schema.Resource) *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		words := strings.Split(name, "_")[1:] // Remove "grafana_" prefix
		group := "grafana"
		switch words[0] {
		case "synthetic":
			group = "syntheticmonitoring"
			words = words[2:]
		case "machine":
			group = "machinelearning"
			words = words[2:]
		case "cloud":
			group = "cloud"
			words = words[1:]
		}

		kind := tjname.NewFromSnake(strings.Join(words, "_")).Camel
		r := &tjconfig.Resource{
			Name:              name,
			TerraformResource: terraformResource,
			ShortGroup:        group,
			Kind:              kind,
			Version:           "v1alpha1",
			ExternalName:      tjconfig.IdentifierFromProvider,
			References:        map[string]tjconfig.Reference{},
			Sensitive:         tjconfig.NopSensitive,
		}
		return r
	}

	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithSkipList([]string{
			// Ignoring some resources that have trouble with generation
			// TODO: Fix those
			"grafana_alert_notification$",
			"grafana_data_source$",
		}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
