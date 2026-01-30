package router

import (
	model_config "xi/internal/models/config"
	"xi/pkg/hook"

	"github.com/gin-gonic/gin"
)

// Sitemap hook container
type SitemapLib struct {
	HookPre  *hook.EffectHook[*gin.Context]
	HookCore *hook.Hook[*gin.Context, []model_config.MetaSitemap]
	HookPost *hook.EffectHook[SitemapPostArgs]
}

var Sitemap = &SitemapLib{
	HookPre:  hook.NewEffectHook[*gin.Context](),
	HookCore: hook.New[*gin.Context, []model_config.MetaSitemap](),
	HookPost: hook.NewEffectHook[SitemapPostArgs](),
}

// Args passed to post-sitemap hooks
type SitemapPostArgs struct {
	Ctx  *gin.Context
	URLs []model_config.MetaSitemap
}

// Capability interfaces

type PreSitemap interface {
	SitemapPre(c *gin.Context) error
}

type CoreSitemap interface {
	SitemapCore(c *gin.Context) ([]model_config.MetaSitemap, error)
}

type PostSitemap interface {
	SitemapPost(c *gin.Context, urls []model_config.MetaSitemap) error
}
