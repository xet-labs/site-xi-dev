package ctrl

import (
	"net/http"
	"sort"
	"xi/pkg"
	"xi/pkg/cfg"

	"github.com/gin-gonic/gin"
)

type DebugCtrl struct{}

var Debug = &DebugCtrl{}

func (d *DebugCtrl) RoutesCore(r *gin.Engine) {
	if cfg.App.Mode != "test" {
		return
	}

	r.GET("/t", d.Index(r))

	r.GET("/t/c", func(c *gin.Context) {
		c.JSON(200, cfg.All())
	})

	r.GET("/t/cr", func(c *gin.Context) {
		cfg, err := cfg.RAll()
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to load config", "details": err.Error()})
			return
		}
		c.JSON(200, cfg)
	})

	r.GET("/t/r", func(c *gin.Context) {
		routes, _ := d.routeData(r)
		c.JSON(200, routes)
	})
}

func (d *DebugCtrl) Index(r *gin.Engine) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"route": func() []string {
				routes, _ := d.routeData(r)
				return routes
			}(),
			"conf": func() any {
				cfg, err := lib.Conf.All()
				if err != nil {
					return gin.H{"error": "failed to load config", "details": err.Error()}
				}
				return cfg
			}(),
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
