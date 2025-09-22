package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
	model_config "xi/internal/app/model/config"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/util"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (v *WebLib) PageHandler(pageName string) gin.HandlerFunc {
	return func(c *gin.Context) { v.Page(c, cfg.Web.Pages[pageName]) }
}

func (v *WebLib) Page(c *gin.Context, p *model_config.WebPage) bool {
	rdbKey := c.Request.URL.Path

	// Try cache
	if v.OutCache(c, rdbKey).Html() {
		return true
	}

	// Process Render.Content Src
	switch p.Ctrl.Render {
	case "url":
	case "md":
	case "file":
		v.mu.Lock()
		contentBytes, err := os.ReadFile(p.Content.File)
		v.mu.Unlock()
		if err != nil {
			log.Error().Caller().Err(err).Str("page", c.Request.URL.Path).Msg("web Page, Read-file")
			c.Status(http.StatusInternalServerError)
			return false
		}
		p.Rt = map[string]any{
			"Content": template.HTML(contentBytes),
		}

	case "raw":
		p.Rt = map[string]any{
			"Content": template.HTML(p.Content.Raw),
		}
	}

	// Process Layout Type
	var page []byte
	switch p.Ctrl.Layout {
	case "raw":
		switch v := p.Rt["Content"].(type) {
		case []byte:
			page = v
		case string:
			page = []byte(v)
		case template.HTML:
			page = []byte(v)
		default:
			c.Status(http.StatusInternalServerError)
			log.Warn().Caller().Str("type", fmt.Sprintf("%T", v)).Str("Page", c.Request.URL.Path).Msg("web Page, Unsupported content type in p.Rt[\"content\"]")
			return false
		}

	default:
		buf := bytes.Buffer{}
		if err := v.Tcli.ExecuteTemplate(&buf, util.Str.Fallback(p.Ctrl.Layout, "layout/default"), gin.H{"P": p}); err != nil {
			log.Error().Caller().Err(err).Str("page", c.Request.URL.Path).Msg("web Page, ExecTemplate")
			c.Status(http.StatusInternalServerError)
			return false
		}
		page = buf.Bytes()
	}

	// Minify HTML
	pageMin, err := util.Minify.Html(page)
	if err != nil {
		// Serve the response with optional cache if rdbKey is provided in args[0]
		c.Data(http.StatusOK, "text/html; charset=utf-8", page)
		log.Error().Caller().Err(err).Str("page", c.Request.URL.Path).Msg("Web.OutHtmlLyt.minify")

		if p.Ctrl.Cache == nil || *p.Ctrl.Cache || cfg.App.ForceCachePage {
			go func(data any) { store.Rdb.Set(rdbKey, data, 10*time.Minute) }(page)
		}
		return true
	}

	// Serve the response with optional cache if rdbKey is provided in args[0]
	c.Data(http.StatusOK, "text/html; charset=utf-8", pageMin)
	if p.Ctrl.Cache == nil || *p.Ctrl.Cache || cfg.App.ForceCachePage {
		go func(data any) { store.Rdb.Set(rdbKey, data, 10*time.Minute) }(pageMin)
	}
	return true
}
