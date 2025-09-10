package cfg

import (
	"xi/pkg/cfg/static"
	model_config "xi/internal/app/model/config"

	"github.com/knadh/koanf/v2"
)

// Runtime config (mutable)
var Config = &model_config.Config{}
var Raw *koanf.Koanf

// Direct pointers for convenience
var (
	Api = &Config.Api
	App = &Config.App
	Org = &Config.Org
	Db  = &Config.Db
	Web = &Config.Web
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

func All() model_config.Config { return *Config }

// Set replaces the entire config (except Build, which stays static)
func Set(cfg model_config.Config) {
	cfg.Build = Build // enforce static build info
	*Config   = cfg
}

// Update merges in a new config but keeps Build static
func Update(cfg model_config.Config) {
	SetupRelease(&cfg)
	
	*Config = cfg
	Api = &Config.Api
	App = &Config.App
	Org = &Config.Org
	Db 	= &Config.Db
	Web = &Config.Web
}

func SetupRelease(cfg *model_config.Config) {
	cfg.Build = Build

	if Build.Mode == "release" && cfg.App.Mode == "test" {
		cfg.App.Mode = Build.Mode
	}
}

// Update cfg.Raw *koanf.Koanf
func RUpdate(raw *koanf.Koanf) {
	Raw = raw
}

// Alias to *koanf.Koanf.Get()
func RGet(path string) any {
    return Raw.Get(path)
}

func RAll() (map[string]any, error) {
	var cfgRaw map[string]any
	if err := Raw.Unmarshal("", &cfgRaw); err != nil {
		return map[string]any{}, err
	}
	return cfgRaw, nil
}