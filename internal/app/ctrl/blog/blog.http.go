// cntr/blog.go
package blog

import (
	"html/template"
	"sync"

	"github.com/gin-gonic/gin"

	model_config "xi/internal/app/model/config"
	model_store "xi/internal/app/model/store"
	"xi/pkg/app"
	"xi/pkg/lib"
	"xi/pkg/lib/cfg"
)

type BlogHttpCtrl struct {
	mu   sync.RWMutex
}

var BlogHttp = &BlogHttpCtrl{}

// GET /blog/@<user>/<slug>
func (b *BlogHttpCtrl) Show(c *gin.Context) {
	rdbKey := c.Request.RequestURI

	if lib.Web.OutCache(c, rdbKey).Html() {
		return
	}

	rawUID := c.Param("uid")
	rawID := c.Param("id")

	if err := BlogApi.Validate(rawUID, rawID); err != nil {
		app.Err.Handle(c, err)
		return
	}

	var blog model_store.Blog
	b.mu.Lock()
	err := BlogApi.ShowCore(&blog, rawUID, rawID)
	b.mu.Unlock()
	if err != nil {
		app.Err.Handle(c, err)
		return
	}

	// Success → prepare response
	p := *cfg.Web.Pages["blogs"]
	b.PrepMeta(c, &p.Meta, &blog)
	p.Rt = map[string]any{
		"B":       &blog,
		"Content": template.HTML(blog.Content),
	}

	lib.Web.OutHtmlLyt(c, &p, rdbKey)
}

func (b *BlogHttpCtrl) PrepMeta(c *gin.Context, meta *model_config.WebMeta, raw *model_store.Blog) {
	meta.Type = "Article"
	meta.Title = raw.Title
	meta.URL = lib.Util.Url.Full(c)
	meta.AltJson = lib.Util.Url.Host(c) + "/api" + c.Request.RequestURI
	meta.Description = raw.Description
	meta.Img.URL = lib.Util.Url.Host(c) + raw.FeaturedImg
	meta.Tags = raw.Tags
	meta.Author.Name = raw.User.Name
	meta.Author.Img = raw.User.AvatarURL
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