// cntr/blog.go
package blog

import (
	"html/template"
	"net/http"
	"sync"
	"xi/internal/app/lib"
	"xi/internal/app/lib/cfg"
	model_config "xi/internal/app/model/config"
	model_db "xi/internal/app/model/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BlogHttpCtrl struct {
	db    *gorm.DB
	rdb   *redis.Client

	mu   sync.RWMutex
	once sync.Once
}

// Singleton controller
var BlogHttp = &BlogHttpCtrl{
	db: lib.Db.GetCli(),
}

// GET /blog
// func (b *BlogHttpCtrl) Index(c *gin.Context) {}

func (b *BlogHttpCtrl) Show(c *gin.Context) {
	rdbKey := c.Request.RequestURI

	if lib.Web.OutCache(c, rdbKey).Html() {
		return
	} // Try cache

	// On cache miss fetch data from DB
	rawUID := c.Param("uid") // @username or UID
	rawID := c.Param("id")   // blog ID or slug

	blog := model_db.Blog{}
	if err := BlogApi.Validate(rawUID, rawID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b.mu.Lock()
	if err := BlogApi.ShowCore(&blog, rawUID, rawID); err != nil { // Fallback to DB
		status := http.StatusNotFound
		if err == ErrInvalidUID {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	b.mu.Unlock()

	p := cfg.Web.Pages["blogs"]
	b.PrepMeta(c, &p.Meta, &blog)
	p.Rt = map[string]any{
		"B":       blog,
		"Content": template.HTML(blog.Content),
	}

	lib.Web.OutHtmlLyt(c, p, rdbKey)
}

func (b *BlogHttpCtrl) PrepMeta(c *gin.Context, meta *model_config.Meta, raw *model_db.Blog) {
	meta.Type = "Article"
	meta.Title = raw.Title
	meta.URL = lib.Util.Url.Full(c)
	meta.AltJson = lib.Util.Url.Host(c) + "/api" + c.Request.RequestURI
	meta.Description = raw.Description
	meta.Img.URL = lib.Util.Url.Host(c) + raw.FeaturedImg
	meta.Tags = raw.Tags
	meta.Author.Name = raw.User.Name
	meta.Author.Img = raw.User.ProfileImg
	meta.Author.URL = lib.Util.Url.Host(c) + "/@" + raw.User.Username
	meta.CreatedAt = raw.CreatedAt
	meta.UpdatedAt = raw.UpdatedAt
	// meta.Category = raw.Tags
}

// POST api/blog/uid/id
func (b *BlogHttpCtrl) Post(c *gin.Context) {}

// PUT api/blog/uid/id
func (b *BlogHttpCtrl) Put(c *gin.Context) {}

// DELETE api/blog/uid/id
func (b *BlogHttpCtrl) Delete(c *gin.Context) {}
