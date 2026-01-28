package store

import (
	"sync"
	"xi/internal/infra/store/db"
	"xi/internal/infra/store/rdb"
	"xi/pkg/hook"
)

type (
	StoreService struct {
		HookPre *hook.Hook[struct{}, struct{}]
		once    sync.Once
	}

	DbStore  = db.DbStore
	RdbStore = rdb.RdbStore
)

var (
	Store = &StoreService{
		HookPre: hook.New[struct{}, struct{}](),
	}
	Db  = db.Db
	Rdb = rdb.Rdb
)
