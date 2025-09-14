package route

import (
	"sync"

	"xi/pkg/lib/cfg"
	"xi/pkg/lib/hook"
	"xi/pkg/lib/web"
	model_config "xi/internal/app/model/config"

	"github.com/gin-gonic/gin"
)

type RouteLib struct {
	hooks *hook.Hook
	r     *gin.Engine

	once sync.Once
	mu   sync.RWMutex
}

var (
	Route = &RouteLib{
		hooks: &hook.Hook{},
		r:     &gin.Engine{},
	}
)

// Used for setup routes, health checks, middleware, etc.
type PreRoutable interface{ RoutesPre(r *gin.Engine) }

// Used for main application routes (APIs, business logic).
type CoreRoutable interface{ RoutesCore(r *gin.Engine) }

// Used for fallback routes, debug endpoints, catch-alls.
type PostRoutable interface{ RoutesPost(r *gin.Engine) }


// Initializes all routes and templates
func (rh *RouteLib) Init(r *gin.Engine, ctrls any) {
	// Store Gin Engine
	rh.r = r

	// Register controller routes for diffent satages
	rh.RegisterController(r, ctrls)

	
	// Run Hooks
	rh.hooks.RunPre(r, rh)
	rh.hooks.RunCore(r, rh)
	rh.hooks.RunPost(r, rh)
	// Register templates
	r.SetHTMLTemplate(web.Web.NewTmpl("main", ".html", cfg.Web.TemplateDir...))
}

// RegisterController controllers to route and sitemap
func (rh *RouteLib) RegisterController(r *gin.Engine, ctrls any) {
	
	// Normalize input: single controller or slice of controllers
	var controllers []any
	switch v := ctrls.(type) {
	case []any:
		controllers = v
	default:
		controllers = append(controllers, ctrls)
	}

	for _, c := range controllers {
		// Register Pre routes if implemented
		if pre, ok := c.(PreRoutable); ok {
			rh.hooks.AddPre(func(args ...any) (any, error) {
				pre.RoutesPre(r)
				return nil, nil
			})
		}
		// Register Core routes if implemented
		if core, ok := c.(CoreRoutable); ok {
			rh.hooks.AddCore(func(args ...any) (any, error) {
				core.RoutesCore(r)
				return nil, nil
			})
		}
		// Register Post routes if implemented
		if post, ok := c.(PostRoutable); ok {
			rh.hooks.AddPost(func(args ...any) (any, error) {
				post.RoutesPost(r)
				return nil, nil
			})
		}

		// Register Pre Sitemaps if implemented
		if pre, ok := c.(PreSitemap); ok {
			Sitemap.Hooks.AddPre(func(args ...any) (any, error) {
				if len(args) > 0 {
					if ctx, ok := args[0].(*gin.Context); ok {
						return pre.SitemapPre(ctx)
					}
				}
				return nil, nil
			})
		}
		// Register Core Sitemaps if implemented
		if core, ok := c.(CoreSitemap); ok {
			Sitemap.Hooks.AddCore(func(args ...any) (any, error) {
				if len(args) > 0 {
					if ctx, ok := args[0].(*gin.Context); ok {
						return core.SitemapCore(ctx)
					}
				}
				return nil, nil
			})
		}
		// Register Post Sitemaps if implemented
		if post, ok := c.(PostSitemap); ok {
			Sitemap.Hooks.AddPost(func(args ...any) (any, error) {
				if len(args) >= 2 {
					if ctx, ok := args[0].(*gin.Context); ok {
						if urls, ok := args[1].([]model_config.MetaSitemap); ok {
							return post.SitemapPost(ctx, urls)
						}
					}
				}
				return nil, nil
			})
		}
	}
}
