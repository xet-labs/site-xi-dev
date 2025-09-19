package ctrl

import (
	"sync"
	"time"

	"xi/internal/app/ctrl/blog"
	model_config "xi/internal/app/model/config"
	model_ctrlBlog "xi/internal/app/model/ctrl/blog"
	"xi/pkg/lib"
	"xi/pkg/lib/cfg"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BlogCtrl struct {
	Http *blog.BlogHttpCtrl
	Api  *blog.BlogApiCtrl

	dbCli   *gorm.DB
	rdbCli  *redis.Client
	mu   sync.RWMutex
	once sync.Once
}

var Blog = &BlogCtrl{
	Http: blog.BlogHttp,
	Api:  blog.BlogApi,

	dbCli:  store.Db.Cli,
	rdbCli: store.Rdb.Cli,
}

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
func (b *BlogCtrl) SitemapCore(c *gin.Context) (any, error) {
	rdbKey := c.Request.URL.Path + ".blog"
	urls := []model_config.MetaSitemap{}

	// Try cache
	if err := store.Rdb.GetJson(rdbKey, &urls); err == nil {
		return urls, nil
	}

	// Try DB
	var blogs []model_ctrlBlog.BlogSitemap

	b.mu.Lock()
	err := b.dbCli.
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
		urls = append(urls, model_config.MetaSitemap{
			Loc:        cfg.Org.URL + "/blog/@" + p.Username + "/" + p.Slug,
			LastMod:    lib.Util.Str.Fallback(p.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02")),
			ChangeFreq: "daily",
			Priority:   "0.5",
		})
	}

	// Cache
	go func() { store.Rdb.SetJson(rdbKey, urls, 15*time.Minute) }()
	return urls, nil
}
