// services/server
package service

import (
	"xi/app/lib/cfg"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)
type ServerService struct{}
var Server = &ServerService{}

// InitServer start the web server
func (s *ServerService) Init(app *gin.Engine) error {

	log.Info().Msgf("\a\033[1;94mApp running \033[0;34m'http://localhost:%s'%s\033[0m\n", cfg.App.Port,
		func() string {
			if cfg.Org.Url != "" {
				return ", '" + cfg.Org.Url + "'"
			}
			return ""
		}())

	// Start Web-Server
	if err := app.Run(":" + cfg.App.Port); err != nil {
		log.Error().Msgf("Failed to start server: %v", err)
		return err
	}

	return nil
}
