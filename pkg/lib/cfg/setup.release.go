package cfg

import model_config "xi/internal/app/model/config"

func SetupRelease(cfg *model_config.Config) {
	cfg.App.Build = Build

	if Build.Mode == "release" {
		if cfg.App.Mode == "test" {
			cfg.App.Mode = Build.Mode // allow other modes in prod except test !!
		}
		cfg.Api.SecureCookies = true
	}
}
