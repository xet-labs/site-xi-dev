package service

import (
	"xi/pkg/service/auth"
	"xi/pkg/service/store"
)

type (
	AuthService  = auth.AuthService
	StoreService = store.StoreService
)

var (
	Auth  = auth.Auth
	Store = store.Store
)
