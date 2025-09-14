package lib

import (
	"xi/pkg/lib/conf"
	"xi/pkg/lib/db"
	"xi/pkg/lib/env"
	"xi/pkg/lib/hook"
	"xi/pkg/lib/logger"
	"xi/pkg/lib/route"
	"xi/pkg/lib/util"
	"xi/pkg/lib/web"
)

// Expose structs
type (
	ConfLib   = conf.ConfLib
	DbLib     = db.DbLib
	RdbLib    = db.RdbLib
	EnvLib    = env.EnvLib
	Hook      = hook.Hook // Only struct exposed
	LoggerLib = logger.LoggerLib
	RouteLib  = route.RouteLib
	UtilLib   = util.UtilLib
	WebLib    = web.WebLib
)

// Expose Global instance
var (
	Conf   = conf.Conf
	Db     = db.Db
	Rdb    = db.Rdb
	Env    = env.Env
	Log    = logger.Logger.Log
	Logger = logger.Logger
	Route  = route.Route
	Util   = util.Util
	Web    = web.Web
)
