package route

import (
	"xi/internal/app/ctrl"
	"xi/pkg"

	"github.com/gin-gonic/gin"
)

// Add global controllers instance to register its routes
var Controllers = []any{
	pkg.Ctrl.Auth,
	pkg.Ctrl.Debug,
	pkg.Ctrl.Managed,
	pkg.Ctrl.Res,
	CustomRoutes,

	ctrl.Blog,
}

// 'CustomRoutes' allows adding simple/ad-hoc routes without creating a dedicated controller.
type customRoutes struct{}

var CustomRoutes = &customRoutes{}

func (u *customRoutes) RoutesPre(r *gin.Engine)  {}
func (u *customRoutes) RoutesCore(r *gin.Engine) {}
func (u *customRoutes) RoutesPost(r *gin.Engine) {}
