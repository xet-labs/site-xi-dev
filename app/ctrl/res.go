// ctrl/res
package ctrl

import (
	"xi/app/ctrl/res"

	"github.com/gin-gonic/gin"
)

type ResCtrl struct {
	Css     *res.CssRes
	Sitemap *res.SitemapRes
}

var Res = &ResCtrl {
	Css:     res.Css,
	Sitemap: res.Sitemap,
}

func (rc *ResCtrl) Routes(r *gin.Engine) {
	// css
	r.GET("/res/css/*name", rc.Css.Index)
	
	// sitemap
	r.GET("/sitemap", rc.Sitemap.Index)
	r.GET("/sitemap.xml", rc.Sitemap.Index)

	r.NoRoute(func(c *gin.Context) { c.File("public" + c.Request.URL.Path) })
}
