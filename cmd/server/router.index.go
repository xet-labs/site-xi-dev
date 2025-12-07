package main

import (
	"xi/internal/app/server/ctrl"
	ctrlPkg "xi/pkg/app/server/ctrl"
	srvcPkg "xi/pkg/service"
)

// Add controller's global instance to register its routes,
// Global instance must have methods 'RouterPre', 'RouterCore', 'RouterPost'
var Controllers = []any{
	ctrl.Blog,
	ctrl.Custom,

	srvcPkg.Auth.Api,
	ctrlPkg.Res,
	ctrlPkg.Managed,
	ctrlPkg.Debug,
}
