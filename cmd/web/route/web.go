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

// Use this to add any middleware-dependent or prerequisite routes.
func (u *customRoutes) RoutesPre(r *gin.Engine) {}

// Use this for simple pages or routes that do not need a full controller.
func (u *customRoutes) Routes(r *gin.Engine) {}

// Use this for routes that should run last, e.g., catch-all, 404 pages, etc.
func (u *customRoutes) RoutesPost(r *gin.Engine) {}
