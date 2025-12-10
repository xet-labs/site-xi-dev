package service

import (
	// "xi/internal/app/service/app"
	"xi/pkg/lib"
	"xi/pkg/service"
)

// Init Core Libs
// xi/pkg/lib.* are designed so self init on method calls but adding them here ensures they are called once
func InitPre() {
	lib.Logger.Init()
	lib.Env.Init()
	service.Config.Init()
}
func InitCore() {
	service.Store.Init()
}
func InitPost() {
}

func Init() {
	InitPre()
	InitCore()
	InitPost()
}