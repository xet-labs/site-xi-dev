package res

import (
	"net/http"
	"sync"
	"time"
	
	"github.com/gin-gonic/gin"

	model_config "xi/internal/app/model/config"
	model_ctrlRes "xi/internal/app/model/ctrl/res"
	"xi/pkg/store"
	"xi/pkg/lib/router"
)

type SitemapRes struct {
	router.SitemapLib

	mu   sync.RWMutex
	once sync.Once
}

var Sitemap = &SitemapRes{
	SitemapLib: router.SitemapLib{
		Hooks: router.Sitemap.Hooks, // share the same hook instance from lib.route::Sitemap
	},
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
	if _, errs := s.Hooks.RunPre(c); len(errs) > 0 {
		for _, e := range errs {
			c.Error(e)
		}
	}

	// Run Core Hooks
	results, errs := s.Hooks.RunCore(c)
	for _, e := range errs {
		c.Error(e)
	}
	for _, r := range results {
		if u, ok := r.([]model_config.MetaSitemap); ok {
			urls = append(urls, u...)
		}
	}

	// Run Post Hooks
	if _, errs := s.Hooks.RunPost(c, urls); len(errs) > 0 {
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

	// Cache
	go func() { store.Rdb.SetJson(rdbKey, sitemapObj, 15*time.Minute) }()
}
