package router

import (
	"sync"

	"xi/internal/config/cfg"
	model_config "xi/internal/models/config"
	"xi/internal/web"
	"xi/pkg/hook"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RouterService struct {
	hookPre  *hook.EffectHook[*gin.Engine]
	hookCore *hook.EffectHook[*gin.Engine]
	hookPost *hook.EffectHook[*gin.Engine]

	r *gin.Engine

	once sync.Once
	mu   sync.RWMutex
}

var (
	Router = &RouterService{
		hookPre:  hook.NewEffectHook[*gin.Engine](),
		hookCore: hook.NewEffectHook[*gin.Engine](),
		hookPost: hook.NewEffectHook[*gin.Engine](),
	}
)

// Router capability interfaces

// Used for setup routes, health checks, middleware, etc.
type PreRouter interface { RouterPre(r *gin.Engine) }

// Used for main application routes (APIs, business logic).
type CoreRouter interface { RouterCore(r *gin.Engine) }

// Used for fallback routes, debug endpoints, catch-alls.
type PostRouter interface { RouterPost(r *gin.Engine) }

// Init

func (rh *RouterService) Init(r *gin.Engine, ctrls any) error {
	var err error

	rh.once.Do(func() {
		rh.r = r
		rh.RegisterController(ctrls)

		if _, errs := rh.hookPre.Run(r); len(errs) > 0 {
			err = errs[0]
			return
		}
		if _, errs := rh.hookCore.Run(r); len(errs) > 0 {
			err = errs[0]
			return
		}
		if _, errs := rh.hookPost.Run(r); len(errs) > 0 {
			err = errs[0]
			return
		}

		// Register templates
		tmpl, e := web.Web.NewTmpl("main", ".html", cfg.Web.TemplateDir...)
		if e != nil {
			log.Error().Caller().Err(e).
				Msg("could not create template instance")
			err = e
			return
		}

		r.SetHTMLTemplate(tmpl)
	})

	return err
}

//
// Controller registration
//

func (rh *RouterService) RegisterController(ctrls any) {
	for _, c := range normalize(ctrls) {

		// --------------------
		// Router hooks
		// --------------------

		if v, ok := c.(PreRouter); ok {
			rh.hookPre.Add(func(r *gin.Engine) error {
				v.RouterPre(r)
				return nil
			})
		}

		if v, ok := c.(CoreRouter); ok {
			rh.hookCore.Add(func(r *gin.Engine) error {
				v.RouterCore(r)
				return nil
			})
		}

		if v, ok := c.(PostRouter); ok {
			rh.hookPost.Add(func(r *gin.Engine) error {
				v.RouterPost(r)
				return nil
			})
		}

		// --------------------
		// Sitemap hooks
		// --------------------

		if v, ok := c.(PreSitemap); ok {
			Sitemap.HookPre.Add(func(ctx *gin.Context) error {
				return v.SitemapPre(ctx)
			})
		}

		if v, ok := c.(CoreSitemap); ok {
			Sitemap.HookCore.Add(func(ctx *gin.Context) ([]model_config.MetaSitemap, error) {
				return v.SitemapCore(ctx)
			})
		}

		if v, ok := c.(PostSitemap); ok {
			Sitemap.HookPost.Add(func(args SitemapPostArgs) error {
				return v.SitemapPost(args.Ctx, args.URLs)
			})
		}
	}
}

//
// Helpers
//

func normalize(ctrls any) []any {
	switch v := ctrls.(type) {
	case []any:
		return v
	case nil:
		return nil
	default:
		return []any{v}
	}
}
