package services

import (
	"xi/internal/config"
	"xi/internal/infra/logger"
	"xi/internal/infra/store"
)

// Init Core Libs
// xi/pkg/lib.* are designed so self init on method calls but adding them here ensures they are called once
func InitPre() {
	logger.Logger.Init()
	config.Config.Init()
}
func InitCore() {
	store.Store.Init()
}
func InitPost() {
}

func Init() {
	InitPre()
	InitCore()
	InitPost()
}
