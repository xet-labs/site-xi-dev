package res

import (
	"xi/internal/handlers/http/router"

	"github.com/gin-gonic/gin"
)

type ResCtrl struct {
	Css     *CssRes
	Sitemap *SitemapRes
}

var (
	Res = &ResCtrl{
	Css:     Css,
	Sitemap: Sitemap,
}
	// compile-time assertion
	_ router.CoreRouter   = (*ResCtrl)(nil)
)

func (rc *ResCtrl) RouterCore(r *gin.Engine) {
	// css
	r.GET("/res/css/*name", rc.Css.Index)

	// Sitemap
	r.GET("/sitemap", rc.Sitemap.Index)
	r.GET("/sitemap.xml", rc.Sitemap.Index)
}

func (rc *ResCtrl) RouterPost(r *gin.Engine) {
	// Static
	r.NoRoute(func(c *gin.Context) { c.File("public" + c.Request.URL.Path) })
}
