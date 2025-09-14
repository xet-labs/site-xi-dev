package res

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"xi/pkg/lib"
	"xi/pkg/lib/cfg"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CssRes struct {
	Files   sync.Map // key: string, value: []string
	BaseDir string
	RdbTTL  time.Duration

	once sync.Once
	mu   sync.RWMutex
}

var Css = &CssRes{
	BaseDir: cfg.Web.CssBaseDir,
	RdbTTL:  12 * time.Hour,
}

// Css handler: serves combined+cssMin CSS (Redis cached)
func (r *CssRes) Index(c *gin.Context) {
	rdbKey := c.Request.RequestURI
	base := cfg.Web.CssBaseDir + "/" + strings.TrimSuffix(c.Param("name"), ".css")

	if lib.Web.OutCache(c, rdbKey).Css() {
		return // Send cache
	}

	// if files list for path 'base' doesnt exists in []Files then generate
	if _, ok := r.Files.Load(base); !cfg.Web.Cache.Css.FilesList || !ok {
		var (
			files []string
			err   error
		)
		files, err = lib.Util.File.GetWithExt(".css", base)
		if err != nil {
			log.Error().Caller().Err(err).Str("Dir", base).Msg("ctrl.css.index files")
			return
		}

		r.Files.Store(base, files)
	}

	// response and cache
	if v, ok := r.Files.Load(base); ok {
		lib.Web.OutCss(c, lib.Util.File.MergeByte(v.([]string)), rdbKey)
		return
	}

	c.Status(http.StatusInternalServerError)
}
