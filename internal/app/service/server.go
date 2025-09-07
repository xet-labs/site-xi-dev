// services/server
package service

import (
	"xi/internal/app/lib/cfg"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ServerService struct{}

var Server = &ServerService{}

// InitServer start the web server
func (s *ServerService) Init(app *gin.Engine) error {

	log.Info().
		Str("mode", cfg.App.Mode).
		Msgf("\a\033[1;94mapp running \033[0;34m'http://localhost:%s'%s\033[0m", cfg.App.Port,
			func() string {
				if cfg.Org.URL != "" {
					return ", '" + cfg.Org.URL + "'"
				}
				return ""
			}())

	// Start Web-Server
	if err := app.Run(":" + cfg.App.Port); err != nil {
		log.Error().Caller().Err(err).Msg("server")
		return err
	}

	return nil
}
