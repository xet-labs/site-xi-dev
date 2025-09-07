package route

import (
	"xi/internal/app/ctrl"

	"github.com/gin-gonic/gin"
)

// Controllers holds all global controller instances to register their routes.
var Controllers = []any{
	CustomRoutes,
	ctrl.Managed,
	ctrl.Res,
	ctrl.Debug,

	ctrl.Auth,
	ctrl.Blog,
}

// CustomRoutes allows adding simple/ad-hoc routes without creating a dedicated controller.
type customRoutes struct{}
var CustomRoutes = &customRoutes{}

func (u *customRoutes) RoutesPre(r *gin.Engine) {}
func (u *customRoutes) RoutesCore(r *gin.Engine) {}
func (u *customRoutes) RoutesPost(r *gin.Engine) {}
