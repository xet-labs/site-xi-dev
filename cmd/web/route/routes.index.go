package route

import (
	"xi/internal/app/ctrl"
	pkg_ctrl "xi/pkg/ctrl"

	"github.com/gin-gonic/gin"
)

// Add global controllers instance to register its routes
var Controllers = []any{
	pkg_ctrl.Auth,
	pkg_ctrl.Debug,
	pkg_ctrl.Managed,
	pkg_ctrl.Res,
	CustomRoutes,

	ctrl.Blog,
}

// 'CustomRoutes' allows adding simple/ad-hoc routes without creating a dedicated controller.
type customRoutes struct{}

var CustomRoutes = &customRoutes{}

func (u *customRoutes) RoutesPre(r *gin.Engine)  {}
func (u *customRoutes) RoutesCore(r *gin.Engine) {}
func (u *customRoutes) RoutesPost(r *gin.Engine) {}
