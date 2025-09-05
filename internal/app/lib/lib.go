package lib

import (
	"xi/internal/app/lib/conf"
	"xi/internal/app/lib/db"
	"xi/internal/app/lib/env"
	"xi/internal/app/lib/hook"
	"xi/internal/app/lib/logger"
	"xi/internal/app/lib/route"
	"xi/internal/app/lib/util"
	"xi/internal/app/lib/view"
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
