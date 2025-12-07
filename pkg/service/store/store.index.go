package store

import (
	"sync"
	"xi/pkg/lib/hook"
	"xi/pkg/service/store/db"
	"xi/pkg/service/store/rdb"
)

type (
	StoreService struct {
		Hooks *hook.Hook
		once  sync.Once
	}

	DbStore  = db.DbStore
	RdbStore = rdb.RdbStore
)

var (
	Store = &StoreService{
		Hooks: &hook.Hook{},
	}
	Db  = db.Db
	Rdb = rdb.Rdb
)
