package main

import (
	"xi/internal/service"
	srvcPkg "xi/pkg/service"
	appPkg "xi/pkg/app"
	"xi/pkg/lib/cfg"

	"github.com/gin-gonic/gin"
)

func main() {

	service.Init() // Init services

	gin.SetMode(cfg.App.Mode) // Init Gin Engine
	engine := gin.Default()

	srvcPkg.Router.Init(engine, Controllers) 	// Init routes
	appPkg.Server.Init(engine, cfg.App.Port) 	// Init server
}
