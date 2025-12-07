package lib

import (
	"xi/pkg/lib/env"
	"xi/pkg/lib/hook"
	"xi/pkg/lib/logger"
	"xi/pkg/lib/util"
	"xi/pkg/lib/web"
)

// Expose structs
type (
	EnvLib    = env.EnvLib
	Hook      = hook.Hook // Only struct exposed
	LoggerLib = logger.LoggerLib
	UtilLib   = util.UtilLib
	WebLib    = web.WebLib
)

// Expose Global instance
var (
	Env    = env.Env
	Log    = logger.Logger.Log
	Logger = logger.Logger
	Util   = util.Util
	Web    = web.Web
)
