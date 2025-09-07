package conf

import (
	confHook "xi/internal/app/lib/conf/hooks"
	"xi/internal/app/lib/hook"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}

