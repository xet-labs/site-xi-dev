package logger

import (
	"os"
	"sync"
	"time"

	"xi/app/lib/cfg"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerLib struct {
	Log zerolog.Logger

	mu   sync.RWMutex
	once sync.Once
}

var Logger = &LoggerLib{}

// func init() { Logger.Init() }

func (l *LoggerLib) Init() { l.once.Do(l.InitCore) }

func (l *LoggerLib) InitCore() {
	// Set timestamp behavior globally
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	// Create console writer with custom formatting
	writer := zerolog.ConsoleWriter{
		Out:     os.Stderr,
		NoColor: false,
		FormatTimestamp: func(i any) string {
			switch v := i.(type) {
			case time.Time:
				return "\x1b[90m" + v.UTC().Format(time.RFC3339) + "\x1b[0m"
			case string:
				return "\x1b[90m" + v + "\x1b[0m"
			default:
				return ""
			}
		},
		// - for messages to appear normal rather than bold
		// FormatMessage: func(i any) string {
		// 	return "\x1b[0m" + i.(string) + "\x1b[0m"
		// },
	}

	// Create a new logger and store in struct
	l.Log = zerolog.New(writer).With().Timestamp().Logger()

	// Update global `log.Logger` too, if needed
	log.Logger = l.Log

	if cfg.Build.Mode == "release" {
		log.Info().
			Str("date", cfg.Build.Date).
			Str("mode", cfg.Build.Mode).
			Str("name", cfg.Build.Name).
			Str("revision", cfg.Build.Revision).
			Str("version", cfg.Build.Version).
			Msg("App build")
	} else {
		log.Info().
			Str("date", cfg.Build.Date).
			Str("name", cfg.Build.Name).
			Str("revision", cfg.Build.Revision).
			Str("version", cfg.Build.Version).
			Msg("App build")
	}
}
