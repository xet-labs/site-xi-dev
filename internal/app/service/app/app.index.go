package app

import (
// "xi/internal/app/service/app"
)

type AppService struct {
	Server ServerApp
	Debug  DebugApp
}

var App = &AppService{
	Server: ServerApp{},
	Debug:  DebugApp{},
}

// Shortcuts
var (
	Server = &App.Server
	Debug  = &App.Debug
)
