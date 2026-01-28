package ctrl

import (
	"xi/internal/handlers/http/ctrl/blog"
	"xi/internal/handlers/http/ctrl/res"
	"xi/internal/services"
)

type (
	CtrlPkg struct{}
)

var Ctrl = &CtrlPkg{}

// Add controller's global instance to register its routes,
// Global instance must have methods 'RouterPre', 'RouterCore', 'RouterPost'
var Controllers = []any{
	blog.Blog,
	Custom,

	services.Auth.Api,
	res.Res,
	Managed,
	Debug,
}
