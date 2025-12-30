package store

import (
	"sync"
	"xi/pkg/hook"
	"xi/internal/infra/store/db"
	"xi/internal/infra/store/rdb"
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
