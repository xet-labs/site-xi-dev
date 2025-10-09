package web

import (
	"html/template"
	"sync"

	"xi/pkg/lib/cfg"

	"github.com/gin-gonic/gin"
)

type WebLib struct {
	Ecli      *gin.Engine        // Gin Engine
	Tcli      *template.Template // Current Template Cli
	RawTcli   *template.Template // Clean Template Cli
	templates []string

	// once sync.Once
	mu sync.RWMutex
}

var Web = &WebLib{
	templates: cfg.Web.TemplateDir,
}
