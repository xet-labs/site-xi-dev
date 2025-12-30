package services

import (
	"xi/internal/handlers/http/router"
	"xi/internal/infra/store"
	"xi/internal/services/auth"
	"xi/internal/services/handler"
)

type (
	AuthService   = auth.AuthService
	StoreService  = store.StoreService
	RouterService = router.RouterService
)

var (
	Auth   = auth.Auth
	Store  = store.Store
	Err    = &handler.AppErr{}
	Router = router.Router
	Debug  = &DebugService{}
)
