package cfg

import (
	"xi/internal/app/lib/cfg/static"
	model_config "xi/internal/app/model/config"
)

// Runtime config (mutable)
var Config = &model_config.Config{}

// Direct pointers for convenience
var (
	Api  = &Config.Api
	App  = &Config.App
	Org  = &Config.Org
	Db   = &Config.Db
	View = &Config.View
)

// Static BuildConf (never changes at runtime)
var Build = model_config.BuildConf{
	Date:     static.BuildDate,
	Name:     static.BuildName,
	Revision: static.BuildRevision,
	Version:  static.BuildVersion,
	Mode:     static.BuildMode,
}

// Get returns current runtime config
func Get() *model_config.Config { return Config }

// Set replaces the entire config (except Build, which stays static)
func Set(cfg model_config.Config) {
	cfg.Build = Build // enforce static build info
	*Config = cfg
}

// Update merges in a new config but keeps Build static
func Update(cfg model_config.Config) {
	SetupRelease(&cfg)

	*Config = cfg
	Api = &Config.Api
	App = &Config.App
	Org = &Config.Org
	Db = &Config.Db
	View = &Config.View
}

func SetupRelease(cfg *model_config.Config) {
	cfg.Build = Build

	if Build.Mode == "release" && cfg.App.Mode == "test"{
		cfg.App.Mode = Build.Mode
	}
}
