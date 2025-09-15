package main

import (
	"xi/cmd/web/router"
	"xi/internal/app/service"
	"xi/pkg/lib"
	"xi/pkg/lib/cfg"

	"github.com/gin-gonic/gin"
)

func main() {

	service.Init() // Init services

	gin.SetMode(cfg.App.Mode) // Init Gin Engine
	app := gin.Default()

	lib.Router.Init(app, router.Controllers) // Init routes
	service.App.Server.Init(app)             // Init server
}
