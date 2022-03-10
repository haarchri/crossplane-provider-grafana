package dashboard

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("grafana_dashboard", func(r *config.Resource) {
		r.References["folder"] = config.Reference{
			Type:      "Folder",
			Extractor: "github.com/grafana/crossplane-provider-grafana/config.IDExtractor()",
		}
	})
}
