package config

import (
	"xi/pkg/hook"
	confHook "xi/internal/config/hooks"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}
