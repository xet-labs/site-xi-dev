package auth

import (
	"xi/internal/handlers/http/router"

	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

var (
	Api = &AuthApi{}

	// compile-time assertion
	_ router.CoreRouter   = (*AuthApi)(nil)
)

func (a *AuthApi) RouterCore(r *gin.Engine) {
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/refresh", a.Refresh)
		authApi.POST("/login", a.Login)
		authApi.POST("/logout", a.Logout)
		authApi.POST("/signup", a.Signup)
		// authApi.POST("/signout", a.Signout)
	}
}
