package services

import (
	"xi/internal/services/auth"
)

type (
	AuthService = auth.AuthService
)

var (
	Auth  = auth.Auth
	Debug = &DebugService{}
)
