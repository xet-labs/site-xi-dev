package cfg

import (
	model_config "xi/internal/app/model/config"
	"xi/pkg/lib/cfg/static"

	"github.com/knadh/koanf/v2"
)

// Static BuildConf (never changes at runtime)
var Build = model_config.BuildConf{
	Date:     static.BuildDate,
	Name:     static.BuildName,
	Revision: static.BuildRevision,
	Version:  static.BuildVersion,
	Mode:     static.BuildMode,
}

// Runtime config (mutable)
var Config = &model_config.Config{}
var Raw *koanf.Koanf


// Direct pointers for convenience
var (
	Api   = &Config.Api
	App   = &Config.App
	Org   = &Config.Org
	Store = &Config.Store
	Web   = &Config.Web
)

// Get returns current runtime config
func Get() *model_config.Config { return Config }

func All() model_config.Config { return *Config }

// Set replaces the entire config (except Build, which stays static)
func Set(cfg model_config.Config) {
	cfg.App.Build = Build // enforce static build info
	*Config = cfg
}

// Update merges in a new config but keeps Build static
func Update(cfg model_config.Config) {
	SetupRelease(&cfg)

	*Config = cfg

	Api = &Config.Api
	App = &Config.App
	Org = &Config.Org
	Store = &Config.Store
	Web = &Config.Web
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
