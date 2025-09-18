package store

import (
	"context"
	"fmt"
	"strings"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/conf"
	"xi/pkg/lib/env"
	"xi/pkg/lib/util"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Init initializes DBs once
func (s *StoreService) Init() { s.once.Do(s.InitCore) }

func (s *StoreService) initPre() {
	conf.Conf.Init()

	// Set global Redis and DB defaults
	d.SetDefault(cfg.Db.DbDefault)
	Rdb.SetCtx(context.Background())
	Rdb.SetDefault(cfg.Db.RdbDefault)
	cfg.Db.RdbPrefix = util.Str.Fallback(cfg.Db.RdbPrefix,
		util.Str.IfNotEmptyElse(cfg.Org.Abbr, cfg.Org.Abbr+cfg.Build.Revision, cfg.Build.Revision))
	Rdb.SetPrefix(cfg.Db.RdbPrefix)
}
func (s *StoreService) logStat(err error, profile string) error {
	if err != nil {
		log.Error().Caller().Err(err).Str("profile", profile).Str("type", "MySQL").Msg("db connect")
		return nil
	}
	log.Info().Str("profile", profile).Str("type", "MySQL").Msg("db connected")
	return nil
}

// Initializes all DBs and Redis clients (forced)
func (s *StoreService) InitCore() {
	conf.Conf.Init()
	s.Hooks.RunPre()

	if cfg.Db.Conn == nil {
		log.Warn().Msgf("DB WRN: No connections were configured")
	}
	for profile, c := range cfg.Db.Conn {
		if !c.Enable {
			continue
		}

		// Fallback for DB credentials
		if c.User == "" {
			c.User = c.Db + "_u"
		}
		if c.Pass == "" {
			c.Pass = env.Env.Get("DB_PASS")
		}

		switch strings.ToLower(c.Driver) {
		case "mysql", "mariadb":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
				c.User, c.Pass, c.Host, c.Port, c.Db, c.Charset)
			dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Error().Caller().Err(err).Str("profile", profile).Str("type", "MySQL").Msg("db connect")
				continue
			}
			Db.SetCli(profile, dbConn)
			log.Info().Str("profile", profile).Str("type", "MySQL").Msg("db connected")

		case "sqlite":
			dbConn, err := gorm.Open(sqlite.Open(c.Db), &gorm.Config{})
			if err != nil {
				log.Error().Caller().Err(err).Str("profile", profile).Str("type", "SQLite").Msg("db connect")
				continue
			}
			Db.SetCli(profile, dbConn)
			log.Info().Str("profile", profile).Str("type", "SQLite").Msg("db connected")

		case "redis":
			rdb := redis.NewClient(&redis.Options{
				Addr:     c.Host + ":" + c.Port,
				Password: c.Pass,
				DB:       c.Rdb,
			})
			if err := rdb.Ping(context.Background()).Err(); err != nil {
				log.Error().Caller().Err(err).Str("profile", profile).Str("type", "Redis").Msg("db connect")
				continue
			}
			Rdb.SetCli(profile, rdb)
			log.Info().Str("profile", profile).Str("type", "Redis").Msg("db connected")

		default:
			log.Warn().Caller().Str("profile", profile).Str("driver", c.Driver).Msg("db unsupported driver")
		}
	}

	s.Hooks.RunPost()

}
