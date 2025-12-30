package config

import (
	confHook "xi/internal/config/hooks"
	"xi/pkg/hook"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}
