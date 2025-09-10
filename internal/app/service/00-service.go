package service

import (
	"xi/pkg"
)

// xi/pkg.* are designed so self init on method calls but adding them here ensures they are called once
func Init() {
	// Init Core Libs
	lib.Logger.Init()
	lib.Env.Init()
	lib.Conf.Init()
	lib.Db.Init()
	
	Stats.MemD(60*5)
}
