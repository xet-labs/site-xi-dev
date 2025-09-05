package ctrl

import (
	"os"
	"strings"
	"time"

	"xi/internal/app/lib"
	"xi/internal/app/lib/cfg"
	model_config "xi/internal/app/model/config"

	"github.com/gin-gonic/gin"
)

type ManagedCtrl struct {
	// mu   sync.RWMutex
	// once sync.Once
}

var Managed = &ManagedCtrl{}

// All cfg.View.Pages[@].Route are added automatically for mode={"", "managed"}
func (m *ManagedCtrl) Routes(r *gin.Engine) {
	for _, p := range cfg.View.Pages {
		if p != nil && (p.Mode == "" || strings.ToLower(p.Mode) == "managed") {

			// determine the HTTP method, default to GET
			method := "GET"
			if p.Method != "" {
				method = strings.ToUpper(p.Method)
			}

			// register route based on method
			switch method {
			case "GET":
				r.GET(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "POST":
				r.POST(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "PUT":
				r.PUT(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "PATCH":
				r.PATCH(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "DELETE":
				r.DELETE(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "HEAD":
				r.HEAD(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "OPTIONS":
				r.OPTIONS(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			case "PURGE":
				// Gin doesn't have built-in PURGE, use Handle
				r.Handle("PURGE", p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			default:
				// fallback to GET if unknown
				r.GET(p.Route, func(c *gin.Context) { lib.View.Page(c, p) })
			}
		}
	}

}

// Sitemap
func (m *ManagedCtrl) Sitemap(c *gin.Context) (any, error) {
	rdbKey := c.Request.URL.Path + ".managed"
	urls := []model_config.SitemapURL{}

	// Try cache
	if err := lib.Rdb.GetJson(rdbKey, &urls); err == nil {
		return urls, nil
	}

	for _, p := range cfg.View.Pages {

		if p == nil || p.Route == "" || !(p.Mode == "" || strings.ToLower(p.Mode) == "managed") {
			continue
		}

		// default fallbacks
		lastMod := func() string {
			// Case 1: Use UpdatedAt if present
			if p.Meta.UpdatedAt != nil {
				return p.Meta.UpdatedAt.Format("2006-01-02")
			}

			// Case 2: If page is file-based, check its mod time
			if p.Render == "file" && p.ContentFile != "" {
				if fi, err := os.Stat(p.ContentFile); err == nil {
					return fi.ModTime().Format("2006-01-02")
				}
			}

			return time.Now().Format("2006-01-02")
		}()
		changeFreq := "monthly"
		priority := "0.5"

		// If meta info available, override
		urls = append(urls, model_config.SitemapURL{
			Loc:        lib.Util.Str.Fallback(p.Meta.Canonical, cfg.Org.Url+p.Route),
			LastMod:    lib.Util.Str.Fallback(p.Meta.Sitemap.LastMod, lastMod),
			ChangeFreq: lib.Util.Str.Fallback(p.Meta.Sitemap.ChangeFreq, changeFreq),
			Priority:   lib.Util.Str.Fallback(p.Meta.Sitemap.Priority, priority),
		})
	}

	// Cache
	go func() { lib.Rdb.SetJson(rdbKey, urls, 15*time.Minute) }()
	return urls, nil
}
