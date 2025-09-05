package main

import (
	"xi/app/lib"
	"xi/app/lib/cfg"
	"xi/app/route"
	"xi/app/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Init services
	service.Init()

	// Init Gin Engine
	gin.SetMode(cfg.App.Mode)
	app := gin.Default()

	// Init routes
	lib.Route.Init(app, route.Controllers)

	// Init server
	service.Server.Init(app)
}
