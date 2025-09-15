// ctrl/auth.go
package ctrl

import (
	"xi/pkg/service/auth"

	"github.com/gin-gonic/gin"
)

type AuthCtrl struct {
	Api *auth.AuthService
}

var Auth = &AuthCtrl{
	Api: auth.Auth,
}

func (a *AuthCtrl) RoutesCore(r *gin.Engine) {
	authApi := r.Group("/api/auth")
	{
		authApi.POST("/refresh", a.Api.Log)
		authApi.POST("/login", a.Api.ShowLogin)
		authApi.POST("/logout", a.Api.Logout)
		authApi.POST("/signup", a.Api.Signup)
		authApi.POST("/signout", a.Api.Signout)
	}
}
