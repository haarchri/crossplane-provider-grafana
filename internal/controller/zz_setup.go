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

	notification "github.com/grafana/crossplane-provider-grafana/internal/controller/alert/notification"
	key "github.com/grafana/crossplane-provider-grafana/internal/controller/api/key"
	roleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/builtin/roleassignment"
	permission "github.com/grafana/crossplane-provider-grafana/internal/controller/dashboard/permission"
	source "github.com/grafana/crossplane-provider-grafana/internal/controller/data/source"
	sourcepermission "github.com/grafana/crossplane-provider-grafana/internal/controller/data/sourcepermission"
	permissionfolder "github.com/grafana/crossplane-provider-grafana/internal/controller/folder/permission"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/dashboard"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/folder"
	organization "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/organization"
	playlist "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/playlist"
	role "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/role"
	team "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/team"
	user "github.com/grafana/crossplane-provider-grafana/internal/controller/grafana/user"
	learningjob "github.com/grafana/crossplane-provider-grafana/internal/controller/machine/learningjob"
	providerconfig "github.com/grafana/crossplane-provider-grafana/internal/controller/providerconfig"
	monitoringcheck "github.com/grafana/crossplane-provider-grafana/internal/controller/synthetic/monitoringcheck"
	monitoringprobe "github.com/grafana/crossplane-provider-grafana/internal/controller/synthetic/monitoringprobe"
	externalgroup "github.com/grafana/crossplane-provider-grafana/internal/controller/team/externalgroup"
	preferences "github.com/grafana/crossplane-provider-grafana/internal/controller/team/preferences"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		notification.Setup,
		key.Setup,
		roleassignment.Setup,
		permission.Setup,
		source.Setup,
		sourcepermission.Setup,
		permissionfolder.Setup,
		dashboard.Setup,
		folder.Setup,
		organization.Setup,
		playlist.Setup,
		role.Setup,
		team.Setup,
		user.Setup,
		learningjob.Setup,
		providerconfig.Setup,
		monitoringcheck.Setup,
		monitoringprobe.Setup,
		externalgroup.Setup,
		preferences.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
