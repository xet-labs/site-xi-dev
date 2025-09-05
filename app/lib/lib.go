package lib

import (
	"xi/app/lib/conf"
	"xi/app/lib/db"
	"xi/app/lib/env"
	"xi/app/lib/hook"
	"xi/app/lib/logger"
	"xi/app/lib/route"
	"xi/app/lib/util"
	"xi/app/lib/view"
)

type (
	ConfLib   = conf.ConfLib
	DbLib     = db.DbLib
	RdbLib    = db.RdbLib
	EnvLib    = env.EnvLib
	Hook      = hook.Hook // Only struct exposed
	LoggerLib = logger.LoggerLib
	RouteLib  = route.RouteLib
	UtilLib   = util.UtilLib
	ViewLib   = view.ViewLib
)

var (
	Conf   = conf.Conf
	Db     = db.Db
	Rdb    = db.Rdb
	Env    = env.Env
	Log    = logger.Logger.Log
	Logger = logger.Logger
	Route  = route.Route
	Util   = util.Util
	View   = view.View
)
