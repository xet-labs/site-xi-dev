package router

import (
	"xi/internal/app/ctrl"
	pkg_ctrl "xi/pkg/app/ctrl"
	pkg_srvc "xi/pkg/service"

	"github.com/gin-gonic/gin"
)

// Add global controllers instance to register its routes
var Controllers = []any{
	pkg_srvc.Auth.Api,
	pkg_ctrl.Debug,
	pkg_ctrl.Managed,
	pkg_ctrl.Res,
	CustomRoutes,

	ctrl.Blog,
}

// 'CustomRoutes' allows adding simple/ad-hoc routes without creating a dedicated controller
type customRoutes struct{}

var CustomRoutes = &customRoutes{}

func (u *customRoutes) RouterPre(r *gin.Engine)  {}
func (u *customRoutes) RouterCore(r *gin.Engine) {}
func (u *customRoutes) RouterPost(r *gin.Engine) {}
