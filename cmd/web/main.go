package main

import (
	"xi/cmd/web/route"
	"xi/pkg"
	"xi/pkg/cfg"
	"xi/internal/app/service"

	"github.com/gin-gonic/gin"
)

func main() {

	service.Init() // Init services

	gin.SetMode(cfg.App.Mode) // Init Gin Engine
	app := gin.Default()

	lib.Route.Init(app, route.Controllers) // Init routes
	service.Server.Init(app)               // Init server
}
