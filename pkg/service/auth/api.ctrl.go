package auth

import "github.com/gin-gonic/gin"

type AuthApi struct {}

var Api = &AuthApi{}

func (a *AuthApi) RouterCore(r *gin.Engine) {
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/refresh", a.Refresh)
		authApi.POST("/login", a.Login)
		authApi.POST("/logout", a.Logout)
		// authApi.POST("/signup", a.Api.Signup)
		// authApi.POST("/signout", a.Api.Signout)
	}
}