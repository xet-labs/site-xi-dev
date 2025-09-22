package lib

import (
	"xi/pkg/lib/config"
	"xi/pkg/lib/env"
	"xi/pkg/lib/hook"
	"xi/pkg/lib/logger"
	"xi/pkg/lib/router"
	"xi/pkg/lib/util"
	"xi/pkg/lib/web"
)

// Expose structs
type (
	ConfigLib = config.ConfigLib
	EnvLib    = env.EnvLib
	Hook      = hook.Hook // Only struct exposed
	LoggerLib = logger.LoggerLib
	RouterLib = router.RouterLib
	UtilLib   = util.UtilLib
	WebLib    = web.WebLib
)

// Expose Global instance
var (
	Config = config.Config
	Env    = env.Env
	Log    = logger.Logger.Log
	Logger = logger.Logger
	Router = router.Router
	Util   = util.Util
	Web    = web.Web
)
