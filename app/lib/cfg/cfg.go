package cfg

import (
	model_config "xi/app/model/config"
	"xi/app/lib/cfg/static"
)

// Runtime config (mutable)
var Config = &model_config.Config{}

// Direct pointers for convenience
var (
	Api   = &Config.Api
	App   = &Config.App
	Org = &Config.Org
	Db    = &Config.Db
	View  = &Config.View
)

// Static BuildConf (never changes at runtime)
var Build = model_config.BuildConf{
	Date:     static.BuildDate,
	Name:     static.BuildName,
	Revision: static.BuildRevision,
	Version:  static.BuildVersion,
}

// Get returns current runtime config
func Get() *model_config.Config { return Config }

// Set replaces the entire config (except Build, which stays static)
func Set(cfg model_config.Config) {
	cfg.Build = Build         // enforce static build info
	*Config = cfg
}

// Update merges in a new config but keeps Build static
func Update(cfg model_config.Config) {
	cfg.Build = Build
	*Config = cfg
	Api   = &Config.Api
	App   = &Config.App
	Org = &Config.Org
	Db    = &Config.Db
	View  = &Config.View
}