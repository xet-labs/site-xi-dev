package router

import (
	"xi/internal/app/ctrl"
	pkg_ctrl "xi/pkg/app/ctrl"
	pkg_srvc "xi/pkg/service"
)

// Add controller's global instance to register its routes,
// Global instance must have methods 'RouterPre', 'RouterCore', 'RouterPost'
var Controllers = []any{
	ctrl.Blog,
	ctrl.Custom,

	pkg_ctrl.Debug,
	pkg_ctrl.Managed,
	pkg_ctrl.Res,
	pkg_srvc.Auth.Api,
}
