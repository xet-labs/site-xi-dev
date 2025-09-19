package store

import (
	"sync"
	"xi/pkg/lib/hook"
	"xi/pkg/service/store/db"
	"xi/pkg/service/store/rdb"
)

type StoreService struct {
	Hooks *hook.Hook
	mu    sync.RWMutex
	once  sync.Once
}

var Store = &StoreService{
	Hooks: &hook.Hook{},
}

type (
	DbStore  = db.DbStore
	RdbStore = rdb.RdbStore
)

var (
	Db  = db.Db
	Rdb = rdb.Rdb
)
