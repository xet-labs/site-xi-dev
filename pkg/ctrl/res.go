// ctrl/res
package ctrl

import (
	"xi/pkg/ctrl/res"

	"github.com/gin-gonic/gin"
)

type ResCtrl struct {
	Css     *res.CssRes
	Sitemap *res.SitemapRes
}

var Res = &ResCtrl{
	Css:     res.Css,
	Sitemap: res.Sitemap,
}

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
