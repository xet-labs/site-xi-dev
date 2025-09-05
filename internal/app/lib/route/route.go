package route

import (
	"fmt"
	"sync"

	"xi/internal/app/lib/cfg"
	"xi/internal/app/lib/hook"
	"xi/internal/app/lib/view"
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

// ================== Route Hook Interfaces ==================
//
// These interfaces are *optional* for controllers to implement.
// The central route registry (lib/route) will detect and call
// them during Init(), in the correct stage (Pre/Core/Post).
//
// Usage: add the method to your controller struct to hook
// routes into that stage without touching central registry.
//

// Used for setup routes, health checks, middleware, etc.
type PreRoutable interface{ RoutesPre(r *gin.Engine) }

// Used for main application routes (APIs, business logic).
type CoreRoutable interface{ Routes(r *gin.Engine) }

// Used for fallback routes, debug endpoints, catch-alls.
type PostRoutable interface{ RoutesPost(r *gin.Engine) }

// Initializes all routes and templates
func (rh *RouteLib) Init(r *gin.Engine, ctrls any) {
	// Store Gin Engine
	rh.r = r

	// Register controller routes for diffent satages
	rh.RegisterController(r, ctrls)

	// Register templates
	r.SetHTMLTemplate(view.View.NewTmpl("main", ".html", cfg.View.TemplateDirs...))

	// Run Hooks
	rh.hooks.RunPre(r, rh)
	rh.hooks.RunCore(r, rh)
	rh.hooks.RunPost(r, rh)
}

// RegisterController registers one or multiple controllers to Pre, Core, and Post hooks
// It safely checks which hook methods each controller implements before registering.
func (rh *RouteLib) RegisterController(r *gin.Engine, ctrls any) {
	controllers := make([]any, 0)

	// Normalize input: single controller or slice of controllers
	switch v := ctrls.(type) {
	case []any:
		controllers = v
	default:
		controllers = append(controllers, ctrls)
	}

	for _, c := range controllers {
		// Register Pre routes if implemented
		if pre, ok := c.(PreRoutable); ok {
			rh.hooks.AddPre("pre_"+fmt.Sprintf("%T", c), func(args ...any) (any, error) {
				pre.RoutesPre(r)
				return nil, nil
			})
		}

		// Register Core routes if implemented
		if core, ok := c.(CoreRoutable); ok {
			rh.hooks.AddCore("core_"+fmt.Sprintf("%T", c), func(args ...any) (any, error) {
				core.Routes(r)
				return nil, nil
			})
		}

		// Register Post routes if implemented
		if post, ok := c.(PostRoutable); ok {
			rh.hooks.AddPost("post_"+fmt.Sprintf("%T", c), func(args ...any) (any, error) {
				post.RoutesPost(r)
				return nil, nil
			})
		}

		// Register Pre Sitemaps if implemented
		if pre, ok := c.(PreSitemap); ok {
			Sitemap.Hooks.AddPre("pre_"+fmt.Sprintf("%T", c), func(args ...any) (any, error) {
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
			Sitemap.Hooks.AddCore("core_"+fmt.Sprintf("%T", c), func(args ...any) (any, error) {
				if len(args) > 0 {
					if ctx, ok := args[0].(*gin.Context); ok {
						return core.Sitemap(ctx)
					}
				}
				return nil, nil
			})
		}

		// Register Post Sitemaps if implemented
		if post, ok := c.(PostSitemap); ok {
			Sitemap.Hooks.AddPost("post_"+fmt.Sprintf("%T", c), func(args ...any) (any, error) {
				if len(args) >= 2 {
					if ctx, ok := args[0].(*gin.Context); ok {
						if urls, ok := args[1].([]model_config.SitemapURL); ok {
							return post.SitemapPost(ctx, urls)
						}
					}
				}
				return nil, nil
			})
		}

	}
}
