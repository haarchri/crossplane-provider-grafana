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
	annotation "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/annotation"
	apikey "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/apikey"
	builtinroleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/builtinroleassignment"
	cloudapikey "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/cloudapikey"
	cloudplugininstallation "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/cloudplugininstallation"
	cloudstack "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/cloudstack"
	contactpoint "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/contactpoint"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/dashboardpermission"
	datasource "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/datasource"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/datasourcepermission"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/folderpermission"
	librarypanel "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/librarypanel"
	machinelearningjob "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/machinelearningjob"
	messagetemplate "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/messagetemplate"
	mutetiming "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/mutetiming"
	notificationpolicy "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/notificationpolicy"
	oncallescalation "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncallescalation"
	oncallescalationchain "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncallescalationchain"
	oncallintegration "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncallintegration"
	oncalloncallshift "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncalloncallshift"
	oncalloutgoingwebhook "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncalloutgoingwebhook"
	oncallroute "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncallroute"
	oncallschedule "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/oncallschedule"
	organization "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/organization"
	organizationpreferences "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/organizationpreferences"
	playlist "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/playlist"
	report "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/report"
	role "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/role"
	roleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/roleassignment"
	rulegroup "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/rulegroup"
	serviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/serviceaccount"
	serviceaccounttoken "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/serviceaccounttoken"
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
		annotation.Setup,
		apikey.Setup,
		builtinroleassignment.Setup,
		cloudapikey.Setup,
		cloudplugininstallation.Setup,
		cloudstack.Setup,
		contactpoint.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		datasource.Setup,
		datasourcepermission.Setup,
		folder.Setup,
		folderpermission.Setup,
		librarypanel.Setup,
		machinelearningjob.Setup,
		messagetemplate.Setup,
		mutetiming.Setup,
		notificationpolicy.Setup,
		oncallescalation.Setup,
		oncallescalationchain.Setup,
		oncallintegration.Setup,
		oncalloncallshift.Setup,
		oncalloutgoingwebhook.Setup,
		oncallroute.Setup,
		oncallschedule.Setup,
		organization.Setup,
		organizationpreferences.Setup,
		playlist.Setup,
		report.Setup,
		role.Setup,
		roleassignment.Setup,
		rulegroup.Setup,
		serviceaccount.Setup,
		serviceaccounttoken.Setup,
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
