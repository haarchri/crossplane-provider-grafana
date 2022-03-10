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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	alertnotification "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/alertnotification"
	apikey "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/apikey"
	builtinroleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/builtinroleassignment"
	cloudapikey "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/cloudapikey"
	cloudplugininstallation "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/cloudplugininstallation"
	cloudstack "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/cloudstack"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/dashboardpermission"
	datasource "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/datasource"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/datasourcepermission"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/folderpermission"
	librarypanel "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/librarypanel"
	machinelearningjob "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/machinelearningjob"
	organization "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/organization"
	playlist "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/playlist"
	report "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/report"
	role "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/role"
	syntheticmonitoringcheck "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/syntheticmonitoringcheck"
	syntheticmonitoringinstallation "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/syntheticmonitoringinstallation"
	syntheticmonitoringprobe "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/syntheticmonitoringprobe"
	team "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/team"
	teamexternalgroup "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/teamexternalgroup"
	teampreferences "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/teampreferences"
	user "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/user"
	providerconfig "github.com/grafana/crossplane-provider-grafana/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertnotification.Setup,
		apikey.Setup,
		builtinroleassignment.Setup,
		cloudapikey.Setup,
		cloudplugininstallation.Setup,
		cloudstack.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		datasource.Setup,
		datasourcepermission.Setup,
		folder.Setup,
		folderpermission.Setup,
		librarypanel.Setup,
		machinelearningjob.Setup,
		organization.Setup,
		playlist.Setup,
		report.Setup,
		role.Setup,
		syntheticmonitoringcheck.Setup,
		syntheticmonitoringinstallation.Setup,
		syntheticmonitoringprobe.Setup,
		team.Setup,
		teamexternalgroup.Setup,
		teampreferences.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
