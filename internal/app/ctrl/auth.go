// ctrl/auth.go
package ctrl

import (
	"xi/internal/app/ctrl/auth"

	"github.com/gin-gonic/gin"
)

type AuthCtrl struct {
	auth.AuthCtrl
	Api *auth.AuthApiCtrl
}

var Auth = &AuthCtrl{
	Api: auth.AuthApi,
}

func (a *AuthCtrl) Routes(r *gin.Engine) {

	authApi := r.Group("/api")
	{
		authApi.GET("/login", a.ShowLogin)
		authApi.GET("/login/:uid/:id", a.ShowLogin)
		authApi.POST("login", a.Login)

		authApi.GET("/logout", a.ShowLogout)
		authApi.GET("/logout/:uid/:id", a.ShowLogout)
		authApi.POST("/logout", a.Logout)

		authApi.GET("signup", a.ShowSignup)
		authApi.GET("/signup/:uid/:id", a.ShowSignup)
		authApi.POST("/signup", a.Signup)

		authApi.GET("/signout", a.ShowSignout)
		authApi.GET("/signout/:uid/:id", a.ShowSignout)
		authApi.POST("/signout", a.Signout)
	}

	login := r.Group("/login")
	{
		login.GET("", a.ShowLogin)
		login.GET("/:uid/:id", a.ShowLogin)
		login.POST("", a.Login)
	}
	logout := r.Group("/logout")
	{
		logout.GET("/logout", a.ShowLogout)
		logout.POST("/logout", a.Logout)
	}
	signup := r.Group("/signup")
	{
		signup.GET("", a.ShowSignup)
		signup.GET("/:uid/:id", a.ShowSignup)
		signup.POST("", a.Signup)
	}
	signout := r.Group("/logout")
	{
		signout.GET("/signout", a.ShowSignout)
		signout.POST("/signout", a.Signout)
	}
}
