package server

import "xi/internal/app/server/ctrl"

type ServerApp struct{
	Ctrl *ctrl.CtrlPkg 
}

var Server = &ServerApp{
	Ctrl: &ctrl.CtrlPkg{},
}