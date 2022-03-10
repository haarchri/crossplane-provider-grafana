package apikey

import (
	"fmt"

	"github.com/crossplane/terrajet/pkg/config"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

// SlugExtractor extracts the slug of a cloud stack from "status.atProvider.slug"
func SlugExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider.slug")
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
			Type:              "CloudStack",
			Extractor:         "github.com/grafana/crossplane-provider-grafana/config/apikey.SlugExtractor()",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["cloud_stack_slug"].(string); ok {
				conn["url"] = []byte(fmt.Sprintf("https://%s.grafana.net", a))
			}
			if a, ok := attr["key"].(string); ok {
				conn["auth"] = []byte(a)
			}
			return conn, nil
		}
	})
}
