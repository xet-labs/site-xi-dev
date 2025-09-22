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
	case errors.Is(err, e.Get("DbUnavailable").Err):
		l, status, msg = log.Error(), http.StatusServiceUnavailable, "service unavailable"

	case errors.Is(err, e.Get("InvalidUID").Err), errors.Is(err, e.Get("InvalidSlug").Err), errors.Is(err, e.Get("InvalidUser").Err):
		status, msg = http.StatusBadRequest, "invalid parameters"

	case errors.Is(err, e.Get("UserNameExists").Err):
		status, msg = http.StatusConflict, "username already exists"
	case errors.Is(err, e.Get("EmailExists").Err):
		status, msg = http.StatusConflict, "email already exists"

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
