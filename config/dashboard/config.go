package dashboard

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

// IDExtractor extracts ID of the resources from "status.atProvider.id"
func IDExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.id")
		if err != nil {
			return ""
		}
		return r
	}
}

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("grafana_dashboard", func(r *config.Resource) {
		r.References["folder"] = config.Reference{
			Type:      "Folder",
			Extractor: "github.com/grafana/crossplane-provider-grafana/config/dashboard.IDExtractor()",
		}
	})
}
