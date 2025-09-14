package ctrl

import (
	"os"
	"strings"
	"time"

	"xi/pkg/lib"
	"xi/pkg/lib/cfg"
	model_config "xi/internal/app/model/config"

	"github.com/gin-gonic/gin"
)

type ManagedCtrl struct {
	// mu   sync.RWMutex
	// once sync.Once
}

var Managed = &ManagedCtrl{}

// All cfg.Web.Pages[@].Route are added automatically for mode={"", "managed"}
func (m *ManagedCtrl) RoutesCore(r *gin.Engine) {
	for _, p := range cfg.Web.Pages {
		if p != nil && (p.Ctrl.Mode == "" || strings.ToLower(p.Ctrl.Mode) == "managed") {

			// determine the HTTP method, default to GET
			method := "GET"
			if p.Ctrl.Method != "" {
				method = strings.ToUpper(p.Ctrl.Method)
			}

			// register route based on method
			switch method {
			case "GET":
				r.GET(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "POST":
				r.POST(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "PUT":
				r.PUT(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "PATCH":
				r.PATCH(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "DELETE":
				r.DELETE(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "HEAD":
				r.HEAD(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "OPTIONS":
				r.OPTIONS(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			case "PURGE":
				// Gin doesn't have built-in PURGE, use Handle
				r.Handle("PURGE", p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			default:
				// fallback to GET if unknown
				r.GET(p.Route, func(c *gin.Context) { lib.Web.Page(c, p) })
			}
		}
	}

}

// Sitemap
func (m *ManagedCtrl) SitemapCore(c *gin.Context) (any, error) {
	rdbKey := c.Request.URL.Path + ".managed"
	urls := []model_config.MetaSitemap{}

	// Try cache
	if err := lib.Rdb.GetJson(rdbKey, &urls); err == nil {
		return urls, nil
	}

	for _, p := range cfg.Web.Pages {

		if p == nil || p.Route == "" || !(p.Ctrl.Mode == "" || strings.ToLower(p.Ctrl.Mode) == "managed") {
			continue
		}

		// default fallbacks
		lastMod := func() string {
			// Case 1: Use UpdatedAt if present
			if p.Meta.UpdatedAt != nil {
				return p.Meta.UpdatedAt.Format("2006-01-02")
			}

			// Case 2: If page is file-based, check its mod time
			if p.Ctrl.Render == "file" && p.Content.File != "" {
				if fi, err := os.Stat(p.Content.File); err == nil {
					return fi.ModTime().Format("2006-01-02")
				}
			}

			return time.Now().Format("2006-01-02")
		}()
		changeFreq := "monthly"
		priority := "0.5"

		// If meta info available, override
		urls = append(urls, model_config.MetaSitemap{
			Loc:        lib.Util.Str.Fallback(p.Meta.Canonical, cfg.Org.URL + p.Route),
			LastMod:    lib.Util.Str.Fallback(p.Meta.Sitemap.LastMod, lastMod),
			ChangeFreq: lib.Util.Str.Fallback(p.Meta.Sitemap.ChangeFreq, changeFreq),
			Priority:   lib.Util.Str.Fallback(p.Meta.Sitemap.Priority, priority),
		})
	}

	// Cache
	go func() { lib.Rdb.SetJson(rdbKey, urls, 15*time.Minute) }()
	return urls, nil
}
