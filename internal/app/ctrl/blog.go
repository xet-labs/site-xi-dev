package ctrl

import (
	"sync"
	"time"

	"xi/internal/app/ctrl/blog"
	"xi/internal/app/lib"
	"xi/internal/app/lib/cfg"
	model_config "xi/internal/app/model/config"
	model_ctrlBlog "xi/internal/app/model/ctrl/blog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogCtrl struct {
	Http *blog.BlogHttpCtrl
	Api  *blog.BlogApiCtrl
	db   *gorm.DB

	mu   sync.RWMutex
	once sync.Once
}

var Blog = &BlogCtrl{
	Http: blog.BlogHttp,
	Api:  blog.BlogApi,
	db:   lib.Db.GetCli(),
}

// Blog Routes
func (b *BlogCtrl) RoutesCore(r *gin.Engine) {
	api := r.Group("api/blog") // route /api/blog
	{
		api.GET("", Blog.Api.Index)
		api.GET("/:uid/:id", Blog.Api.Show)
		api.POST("/:uid/:id", Blog.Api.Post)
		api.PUT("/:uid/:id", Blog.Api.Put)
		api.DELETE("/:uid/:id", Blog.Api.Delete)
	}

	blogs := r.Group("/blog/:uid/:id") // route /blog/*
	{
		blogs.GET("", Blog.Http.Show)
		blogs.POST("", Blog.Http.Post)
		blogs.PUT("", Blog.Http.Put)
		blogs.DELETE("", Blog.Http.Delete)
	}
}

// Blog Sitemap
func (b *BlogCtrl) SitemapCore(c *gin.Context) (any, error) {
	rdbKey := c.Request.URL.Path + ".blog"
	urls := []model_config.Sitemap{}

	// Try cache
	if err := lib.Rdb.GetJson(rdbKey, &urls); err == nil {
		return urls, nil
	}

	// Try DB
	var blogs []model_ctrlBlog.BlogSitemap
	
	b.mu.Lock()
	err := b.db.
		Table("blogs").
		Select("users.username, blogs.slug, blogs.updated_at").
		Joins("join users on users.uid = blogs.uid").
		Where("blogs.status = ?", "published").
		Find(&blogs).Error
	if err != nil {
		b.mu.Unlock()
		return nil, err
	}
	b.mu.Unlock()

	for _, p := range blogs {

		// If meta info available, override
		urls = append(urls, model_config.Sitemap{
			Loc:        cfg.Org.URL + "/blog/@" + p.Username + "/" + p.Slug,
			LastMod:    lib.Util.Str.Fallback(p.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02")),
			ChangeFreq: "daily",
			Priority:   "0.5",
		})
	}

	// Cache
	go func() { lib.Rdb.SetJson(rdbKey, urls, 15*time.Minute) }()
	return urls, nil
}
