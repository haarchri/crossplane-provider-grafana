package apikey

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

func SlugExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.slug")
		if err != nil {
			return ""
		}
		return r
	}
}

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("grafana_api_key", func(r *config.Resource) {
		r.References["cloud_stack_slug"] = config.Reference{
			Type:      "CloudStack",
			Extractor: "github.com/grafana/crossplane-provider-grafana/config/apikey.SlugExtractor()",
		}
	})
}
