package store

import (
	"sync"
	"xi/pkg/store/db"
	"xi/pkg/store/rdb"
)

type StoreService struct {
	mu         sync.RWMutex
	once       sync.Once
}

var Store = &StoreService{}


type (
	DbStore  = db.DbStore
	RdbStore = rdb.RdbStore
)

var (
	Db  = db.Db
	Rdb = rdb.Rdb
)
