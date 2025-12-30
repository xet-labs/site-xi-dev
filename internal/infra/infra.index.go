package infra

import "xi/internal/infra/logger"

// Expose structs
type (
	// Hook      = hook.Hook
	LoggerLib = logger.LoggerLib
)

// Expose Global instance
var (
	Log    = logger.Logger.Log
	Logger = logger.Logger
)
