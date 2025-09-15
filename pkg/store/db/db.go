package db

import (
	"context"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/conf"
	"xi/pkg/lib/util"

	"gorm.io/gorm"
)

type DbStore struct {
	defaultProfile string
	defaultCli     *gorm.DB
	clients        map[string]*gorm.DB
}

var Db = &DbStore{
	defaultProfile: "database",
	clients:        make(map[string]*gorm.DB),
}

func (d *DbStore) initPre() {
	conf.Conf.Init()

	// Set global Redis and DB defaults
	d.SetDefault(cfg.Db.DbDefault)
	Rdb.SetCtx(context.Background())
	Rdb.SetDefault(cfg.Db.RdbDefault)
	cfg.Db.RdbPrefix = util.Str.Fallback(cfg.Db.RdbPrefix,
		util.Str.IfNotEmptyElse(cfg.Org.Abbr, cfg.Org.Abbr+cfg.Build.Revision, cfg.Build.Revision))
	Rdb.SetPrefix(cfg.Db.RdbPrefix)
}
