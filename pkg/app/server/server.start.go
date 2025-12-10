// services/server
package server

import (
	"xi/internal/app/server"
	"xi/pkg/lib/cfg"
	srvcPkg "xi/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// InitServer start the web server
func (s *ServerApp) Start(port string) error {

	gin.SetMode(cfg.App.Mode) // Init Gin Engine
	engine := gin.Default()
	
	srvcPkg.Router.Init(engine, server.Controllers) // Init routes

	log.Info().Str("mode", cfg.App.Mode).
		Msgf("\a\033[1;94mapp running \033[0;34m'http://localhost:%s'%s\033[0m", port,
			func() string {
				if cfg.Org.URL != "" {
					return ", '" + cfg.Org.URL + "'"
				}
				return ""
			}())

	// Start Web-Server
	if err := engine.Run(":" + port); err != nil {
		log.Error().Caller().Err(err).Msg("server")
		return err
	}

	return nil
}
