package conf

import (
	confHook "xi/pkg/conf/hooks"
	"xi/pkg/hook"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}

