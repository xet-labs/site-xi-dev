package lib

import (
	"xi/pkg/conf"
	"xi/pkg/db"
	"xi/pkg/env"
	"xi/pkg/hook"
	"xi/pkg/logger"
	"xi/pkg/route"
	"xi/pkg/util"
	"xi/pkg/web"
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
	WebLib   = web.WebLib
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
	Web   = web.Web
)
