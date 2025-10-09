package web

import (
	"bytes"
	"net/http"
	"time"
	model_config "xi/internal/app/model/config"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/util"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Render and Cache Minified HTML
func (v *WebLib) OutHtmlLyt(c *gin.Context, p *model_config.WebPage, args ...string) bool {
	rdbKey := util.ArrFallback(args, 0, c.Request.URL.Path)

	// Render html via template
	buf := bytes.Buffer{}
	if err := v.Tcli.ExecuteTemplate(&buf, util.Str.Fallback(p.Ctrl.Layout, "layout/default"), gin.H{"P": p}); err != nil {
		log.Error().Caller().Err(err).Str("page", c.Request.URL.Path).Msg("web Page, ExecTemplate")
		c.Status(http.StatusInternalServerError)
		return false
	}
	page := buf.Bytes()

	// // Minify HTML, Response and optional Cache
	// if pageMin, err := util.Minify.Html(page); err == nil {
	// 	c.Data(http.StatusOK, "text/html; charset=utf-8", pageMin)

	// 	if p.Ctrl.Cache == nil || *p.Ctrl.Cache || cfg.App.ForceCachePage {
	// 		go func(data any) { store.Rdb.Set(rdbKey, data, 10*time.Minute) }(pageMin)
	// 	}
	// 	return true
	// } else {
	// 	log.Error().Caller().Err(err).Str("page", c.Request.URL.Path).Msg("Web.OutHtmlLyt.minify")
	// }

	// Response with optional cache if rdbKey is provided in args[0]
	c.Data(http.StatusOK, "text/html; charset=utf-8", page)
	if p.Ctrl.Cache == nil || *p.Ctrl.Cache || cfg.App.ForceCachePage {
		go func(data any) { store.Rdb.Set(rdbKey, data, 10*time.Minute) }(page)
	}
	return true
}

func (v *WebLib) OutCss(c *gin.Context, css []byte, args ...string) bool {
	// Handle empty content
	if len(css) == 0 {
		c.Status(http.StatusNoContent) // 204
		return true
	}

	// Minify the CSS
	cssMin, err := util.Minify.CssHybrid(css)
	if err != nil {
		c.Data(http.StatusOK, "text/css; charset=utf-8", css)
		log.Error().Caller().Err(err).Msg("web OutCss Minify")
		return true
	}

	// Serve the response with optional cache if rdbKey is provided in args[0]
	c.Data(http.StatusOK, "text/css; charset=utf-8", cssMin)
	if len(args) > 0 && args[0] != "" {
		go func(data any) { store.Rdb.Set(args[0], data, 10*time.Minute) }(cssMin)
	}
	return true
}

func (v *WebLib) OutJson(c *gin.Context, css []byte, args ...string) bool {
	return true
}
