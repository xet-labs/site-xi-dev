package config

import (
	confHook "xi/pkg/lib/config/hooks"
	"xi/pkg/lib/hook"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}
