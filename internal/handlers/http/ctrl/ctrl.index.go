package ctrl

import (
	"xi/internal/services"
)

type (
	CtrlPkg struct {
		Blog *BlogCtrl
	}
)

var Ctrl = &CtrlPkg{
	Blog: Blog,
}

// Add controller's global instance to register its routes,
// Global instance must have methods 'RouterPre', 'RouterCore', 'RouterPost'
var Controllers = []any{
	Blog,
	Custom,

	services.Auth.Api,
	Res,
	Managed,
	Debug,
}
