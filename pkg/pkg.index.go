package pkg

import (
	"xi/pkg/hook"
	"xi/pkg/util"
)

// Expose structs
type (
	Hook      = hook.Hook // Only struct exposed
	UtilLib   = util.UtilLib
)

// Expose Global instance
var (
	Util   = util.Util
)
