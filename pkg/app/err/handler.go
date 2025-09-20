package err

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (e *AppErr) Handle(c *gin.Context, err error, asJSON ...bool) bool {
	if err == nil {
		return false // nothing handled
	}

	l, status, msg := log.Warn(), http.StatusInternalServerError, "internal server error"

	switch {
	case errors.Is(err, e.DbUnavailable):
		l, status, msg = log.Error(), http.StatusServiceUnavailable, "service unavailable"

	case errors.Is(err, e.InvalidUID), errors.Is(err, e.InvalidSlug), errors.Is(err, e.InvalidUser):
		status, msg = http.StatusBadRequest, "invalid parameters"

	case errors.Is(err, e.UserExists):
		status, msg = http.StatusConflict, "user already exists"

	case errors.Is(err, e.EmailExists):
		status, msg = http.StatusConflict, "email already registered"

	case errors.Is(err, gorm.ErrRecordNotFound):
		status, msg = http.StatusNotFound, "resource not found"
	}

	if _, file, line, ok := runtime.Caller(1); ok {

		l.Err(err).
			Str("caller", fmt.Sprintf("%s:%d", file, line)).
			Str("response", msg).
			Int("status", status).
			Msg("handler")
	}

	respondAsJSON := len(asJSON) > 0 && asJSON[0]
	if respondAsJSON {
		c.JSON(status, gin.H{"error": msg})
	} else {
		c.Status(status)
	}

	return true // handled an error
}
