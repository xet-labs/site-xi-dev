package ctrl

type CtrlPkg struct {
	Auth    *AuthCtrl
	Debug   *DebugCtrl
	Managed *ManagedCtrl
	Res     *ResCtrl
}

var Ctrl = &CtrlPkg{
	Auth:    Auth,
	Debug:   Debug,
	Managed: Managed,
	Res:     Res,
}
