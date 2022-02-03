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
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/terraform"

	key "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/api/key"
	roleassignment "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/builtin/roleassignment"
	permission "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/dashboard/permission"
	sourcepermission "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/data/sourcepermission"
	permissionfolder "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/folder/permission"
	dashboard "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/dashboard"
	folder "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/folder"
	organization "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/organization"
	playlist "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/playlist"
	role "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/role"
	team "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/team"
	user "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/grafana/user"
	learningjob "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/machine/learningjob"
	providerconfig "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/providerconfig"
	monitoringcheck "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/synthetic/monitoringcheck"
	monitoringprobe "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/synthetic/monitoringprobe"
	externalgroup "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/team/externalgroup"
	preferences "github.com/crossplane-contrib/provider-jet-grafana/internal/controller/team/preferences"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		key.Setup,
		roleassignment.Setup,
		permission.Setup,
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
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
