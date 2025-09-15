package router

import (
	model_config "xi/internal/app/model/config"
	"xi/pkg/lib/hook"

	"github.com/gin-gonic/gin"
)

type SitemapLib struct {
	Hooks *hook.Hook
}

var Sitemap = &SitemapLib{
	Hooks: &hook.Hook{},
}

type PreSitemap interface {
	SitemapPre(c *gin.Context) (any, error)
}
type CoreSitemap interface {
	SitemapCore(c *gin.Context) (any, error)
}
type PostSitemap interface {
	SitemapPost(c *gin.Context, urls []model_config.MetaSitemap) (any, error)
}
