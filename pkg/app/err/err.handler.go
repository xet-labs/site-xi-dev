package err

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (e *AppErr) Handle(c *gin.Context, err error, asJSON ...bool) bool {
	if err == nil {
		return false // nothing handled
	}

	l, status, msg := log.Error(), http.StatusInternalServerError, "internal server error"
	_, file, line, _ := runtime.Caller(1);

	errEntry, ok := e.E[err.Error()]
	if !ok || !errors.Is(err, errEntry.Err) {
		l.Err(err).
			Str("caller", fmt.Sprintf("%s:%d", file, line)).
			Msg("err unhandled")
		return true
	}

	switch errEntry.LogLevel {
	case "debug":
		l = log.Debug()
	case "error":
		l = log.Error()
	case "fatal":
		l = log.Fatal()
	case "info":
		l = log.Info()
	case "log":
		l = log.Log()
	case "panic":
		l = log.Panic()
	case "trace":
		l = log.Trace()
	default:
		l = log.Warn() // fallback if unknown level
	}
	
	if errEntry.HttpStatus != 0 {status = errEntry.HttpStatus}

	l.Err(err).
		Str("caller", fmt.Sprintf("%s:%d", file, line)).
		Int("httpStatus", errEntry.HttpStatus).
		Str("response", errEntry.Msg).
		Msg("err handled")

	respondAsJSON := len(asJSON) > 0 && asJSON[0]
	if respondAsJSON {
		c.JSON(status, gin.H{"error": msg})
	} else {
		c.Status(status)
	}

	return true // handled an error
}
