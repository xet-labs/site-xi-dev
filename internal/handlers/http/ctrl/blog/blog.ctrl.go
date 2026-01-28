package blog

import (
	"sync"
	"time"

	"xi/internal/config/cfg"
	"xi/internal/handlers/http/router"
	"xi/internal/infra/store"
	model_config "xi/internal/models/config"
	model_ctrlBlog "xi/internal/models/transport/http/blog"
	"xi/pkg/util"

	"github.com/gin-gonic/gin"
)

type BlogCtrl struct {
	Http *BlogHttpCtrl
	Api  *BlogApiCtrl

	mu sync.RWMutex
}

var Blogs *BlogCtrl = &BlogCtrl{
    Http: BlogHttp,
    Api:  BlogApi,
}

var (
	Blog = &BlogCtrl{
		Http: BlogHttp,
		Api:  BlogApi,
	}

	// compile-time assertion
	_ router.CoreRouter   = (*BlogCtrl)(nil)
	_ router.CoreSitemap  = (*BlogCtrl)(nil)
)

// Blog Routes
func (b *BlogCtrl) RouterCore(r *gin.Engine) {
	api := r.Group("api/blog") // route /api/blog
	{
		api.GET("", b.Api.Index)
		api.GET("/:uid/:id", b.Api.Show)
		api.POST("/:uid/:id", b.Api.Post)
		api.PUT("/:uid/:id", b.Api.Put)
		api.DELETE("/:uid/:id", b.Api.Delete)
	}

	blogs := r.Group("/blog/:uid/:id") // route /blog/*
	{
		blogs.GET("", b.Http.Show)
		blogs.POST("", b.Http.Post)
		blogs.PUT("", b.Http.Put)
		blogs.DELETE("", b.Http.Delete)
	}
}

// Blog Sitemap
func (b *BlogCtrl) SitemapCore(c *gin.Context) ([]model_config.MetaSitemap, error) {
	rdbKey := c.Request.URL.Path + ".blog"
	urls := []model_config.MetaSitemap{}

	// Try cache
	if err := store.Rdb.GetJson(rdbKey, &urls); err == nil {
		return urls, nil
	}

	// Try DB
	var blogs []model_ctrlBlog.BlogSitemap

	db := store.Db.Cli()
	if db.Error != nil {
		return nil, db.Error
	}

	b.mu.Lock()
	err := db.
		Table("blogs").
		Select("users.username, blogs.slug, blogs.updated_at").
		Joins("join users on users.id = blogs.uid").
		Where("blogs.status = ?", "published").
		Find(&blogs).Error
	if err != nil {
		b.mu.Unlock()
		return nil, err
	}
	b.mu.Unlock()

	for _, p := range blogs {
		urls = append(urls, model_config.MetaSitemap{
			Loc:        cfg.Org.URL + "/blog/@" + p.Username + "/" + p.Slug,
			LastMod:    util.Str.Fallback(p.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02")),
			ChangeFreq: "daily",
			Priority:   "0.5",
		})
	}

	// Cache
	go func() { store.Rdb.SetJson(rdbKey, urls, 15*time.Minute) }()
	return urls, nil
}
