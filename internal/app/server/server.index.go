package server


import (
	"xi/internal/app/server/ctrl"
	ctrlPkg "xi/pkg/app/server/ctrl"
	srvcPkg "xi/pkg/service"
)


type ServerApp struct{
	Ctrl *ctrl.CtrlPkg 
}

var Server = &ServerApp{
	Ctrl: &ctrl.CtrlPkg{},
}

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
