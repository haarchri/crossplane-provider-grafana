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

	apikey "github.com/grafana/provider-jet-grafana/internal/controller/grafana/apikey"
	builtinroleassignment "github.com/grafana/provider-jet-grafana/internal/controller/grafana/builtinroleassignment"
	dashboard "github.com/grafana/provider-jet-grafana/internal/controller/grafana/dashboard"
	dashboardpermission "github.com/grafana/provider-jet-grafana/internal/controller/grafana/dashboardpermission"
	datasourcepermission "github.com/grafana/provider-jet-grafana/internal/controller/grafana/datasourcepermission"
	folder "github.com/grafana/provider-jet-grafana/internal/controller/grafana/folder"
	folderpermission "github.com/grafana/provider-jet-grafana/internal/controller/grafana/folderpermission"
	organization "github.com/grafana/provider-jet-grafana/internal/controller/grafana/organization"
	playlist "github.com/grafana/provider-jet-grafana/internal/controller/grafana/playlist"
	role "github.com/grafana/provider-jet-grafana/internal/controller/grafana/role"
	team "github.com/grafana/provider-jet-grafana/internal/controller/grafana/team"
	teamexternalgroup "github.com/grafana/provider-jet-grafana/internal/controller/grafana/teamexternalgroup"
	teampreferences "github.com/grafana/provider-jet-grafana/internal/controller/grafana/teampreferences"
	user "github.com/grafana/provider-jet-grafana/internal/controller/grafana/user"
	job "github.com/grafana/provider-jet-grafana/internal/controller/machinelearning/job"
	providerconfig "github.com/grafana/provider-jet-grafana/internal/controller/providerconfig"
	check "github.com/grafana/provider-jet-grafana/internal/controller/syntheticmonitoring/check"
	probe "github.com/grafana/provider-jet-grafana/internal/controller/syntheticmonitoring/probe"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		apikey.Setup,
		builtinroleassignment.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		datasourcepermission.Setup,
		folder.Setup,
		folderpermission.Setup,
		organization.Setup,
		playlist.Setup,
		role.Setup,
		team.Setup,
		teamexternalgroup.Setup,
		teampreferences.Setup,
		user.Setup,
		job.Setup,
		providerconfig.Setup,
		check.Setup,
		probe.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
