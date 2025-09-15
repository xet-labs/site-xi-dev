package conf

import (
	confHook "xi/pkg/lib/conf/hooks"
	"xi/pkg/lib/hook"
)

var PostHooks = []hook.HookFn{
	confHook.ViewPagesSetup,
}
