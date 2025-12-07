package config

import (
	"xi/pkg/lib/hook"
	confHook "xi/pkg/service/config/hooks"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}
