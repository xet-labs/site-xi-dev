package res

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"xi/internal/handlers/http/router"
	"xi/internal/infra/store"
	model_config "xi/internal/models/config"
	model_ctrlRes "xi/internal/models/transport/http/res"
)

type SitemapRes struct {
	*router.SitemapLib

	mu   sync.RWMutex
	once sync.Once
}

var Sitemap = &SitemapRes{
	SitemapLib: router.Sitemap, // share the SAME hook instances
}

func (s *SitemapRes) Index(c *gin.Context) {
	rdbKey := c.Request.URL.Path
	var sitemapObj model_ctrlRes.Sitemap

	// Try Cache
	if err := store.Rdb.GetJson(rdbKey, &sitemapObj); err == nil {
		c.XML(http.StatusOK, sitemapObj)
		return
	}

	var urls []model_config.MetaSitemap

	// Run Pre Hooks
	if _, errs := s.HookPre.Run(c); len(errs) > 0 {
		for _, e := range errs {
			c.Error(e)
		}
	}

	// Run Core Hooks
	// Core hooks are expected to APPEND to `urls`
	results, errs := s.HookCore.Run(c)
	for _, e := range errs {
		c.Error(e)
	}
	for _, r := range results {
		urls = append(urls, r...)
	}

	// Run Post Hooks
	if _, errs := s.HookPost.Run(router.SitemapPostArgs{
		Ctx:  c,
		URLs: urls,
	}); len(errs) > 0 {
		for _, e := range errs {
			c.Error(e)
		}
	}

	// Final sitemap obj
	sitemapObj = model_ctrlRes.Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	// Response
	c.XML(http.StatusOK, sitemapObj)

	// Cache async
	go func() { store.Rdb.SetJson(rdbKey, sitemapObj, 15*time.Minute) }()
}
