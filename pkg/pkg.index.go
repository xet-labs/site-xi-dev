package pkg

import (
	"xi/pkg/auth"
	"xi/pkg/ctrl"
)

type(
	AuthPkg auth.AuthPkg
	CtrlPkg ctrl.CtrlPkg
)

var(
	Auth = auth.Auth
	Ctrl = ctrl.Ctrl
)