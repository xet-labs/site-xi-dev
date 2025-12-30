// cntr/blog.go
package blog

import (
	"html/template"
	"sync"

	"github.com/gin-gonic/gin"

	"xi/internal/config/cfg"
	model_config "xi/internal/models/config"
	model_store "xi/internal/models/store"
	"xi/internal/services/handler"
	"xi/internal/web"
	"xi/pkg/util"
)

type BlogHttpCtrl struct {
	mu sync.RWMutex
}

var BlogHttp = &BlogHttpCtrl{}

// GET /blog/@<user>/<slug>
func (b *BlogHttpCtrl) Show(c *gin.Context) {
	rdbKey := c.Request.RequestURI

	if web.Web.OutCache(c, rdbKey).Html() {
		return
	}

	rawUID := c.Param("uid")
	rawID := c.Param("id")

	if err := BlogApi.Validate(rawUID, rawID); err != nil {
		handler.Err.Handle(c, err)
		return
	}

	var blog model_store.Blog
	b.mu.Lock()
	err := BlogApi.ShowCore(&blog, rawUID, rawID)
	b.mu.Unlock()
	if err != nil {
		handler.Err.Handle(c, err)
		return
	}

	// on success prepare response
	p := *cfg.Web.Pages["blogs"]
	// page meta data
	b.PrepPageMeta(c, &p.Meta, &blog)
	// page render data
	p.R = map[string]any{
		"B":       &blog,
		"Content": template.HTML(blog.Content),
	}

	web.Web.OutHtmlLyt(c, &p, rdbKey)
}

func (b *BlogHttpCtrl) PrepPageMeta(c *gin.Context, meta *model_config.WebMeta, raw *model_store.Blog) {
	meta.Type = "Article"
	meta.Title = raw.Title
	meta.URL = util.Url.Full(c)
	meta.AltJson = util.Url.Host(c) + "/api" + c.Request.RequestURI
	meta.Description = raw.Description
	meta.Img.URL = util.Url.Host(c) + raw.FeaturedImg
	meta.Tags = raw.Tags
	meta.Author.Name = raw.User.Name
	meta.Author.Img = raw.User.AvatarURL
	meta.Author.URL = util.Url.Host(c) + "/@" + raw.User.Username
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
