package service

import (
	"xi/pkg/service/auth"
	"xi/pkg/service/config"
	"xi/pkg/service/handler"
	"xi/pkg/service/store"
	"xi/pkg/service/router"
)

type (
	AuthService   = auth.AuthService
	StoreService  = store.StoreService
	ConfigService = config.ConfigService
	RouterService = router.RouterService
)

var (
	Auth   = auth.Auth
	Store  = store.Store
	Config = config.Config
	Err    = &handler.AppErr{}
	Router = router.Router
)
