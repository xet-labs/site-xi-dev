package ctrl

import (
	"net/http"
	"sort"
	"xi/internal/app/lib"
	"xi/internal/app/lib/cfg"

	"github.com/gin-gonic/gin"
)

type DebugCtrl struct {}
var Debug = &DebugCtrl{}

func (d *DebugCtrl) Routes(r *gin.Engine) {
	if cfg.App.Mode != "test" {
		return
	}

	r.GET("/t", d.Index(r))
	r.GET("/t/c", func(c *gin.Context) {
		c.Data(200, "application/json", lib.Conf.AllJsonPretty())
	})
	r.GET("/t/cs", func(c *gin.Context) {
		c.Data(200, "application/json", lib.Conf.AllJsonStructPretty())
	})
}

func (d *DebugCtrl) Index(r *gin.Engine) gin.HandlerFunc {

	return func(c *gin.Context) {
		routes, _ := d.routeData(r)
		c.JSON(http.StatusOK, gin.H{
			"route": routes,
			// "routeDetailed": detailed,
			"conf": lib.Conf.AllMap(),
		})
	}
}

func (d *DebugCtrl) routeData(r *gin.Engine) ([]string, []string) {
	var routes []string
	var detailed []string

	for _, rt := range r.Routes() {
		method := rt.Method
		if len(method) > 3 {
			method = method[:3]
		}
		routes = append(routes, method+" "+rt.Path)
		detailed = append(detailed, method+" "+rt.Path+" | "+rt.Handler)
	}

	// Sort by route path (strip method prefix for sorting)
	sort.Slice(routes, func(i, j int) bool {
		return routes[i][4:] < routes[j][4:]
	})
	return routes, detailed
}
