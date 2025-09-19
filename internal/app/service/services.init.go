package service

import (
	// "xi/internal/app/service/app"
	"xi/pkg/lib"
	"xi/pkg/service"
)

// xi/pkg/lib.* are designed so self init on method calls but adding them here ensures they are called once
func Init() {
	// Init Core Libs
	lib.Logger.Init()
	lib.Env.Init()
	lib.Conf.Init()
	service.Store.Init()

	// app.Debug.MemD(60 * 5)
}
