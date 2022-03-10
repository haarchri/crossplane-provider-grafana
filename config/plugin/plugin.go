package plugin

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("grafana_cloud_plugin_installation", func(r *config.Resource) {
		r.References["stack_slug"] = config.Reference{
			Type:              "CloudStack",
			Extractor:         "github.com/grafana/crossplane-provider-grafana/config.SlugExtractor()",
			RefFieldName:      "StackRef",
			SelectorFieldName: "StackSelector",
		}
	})
}
