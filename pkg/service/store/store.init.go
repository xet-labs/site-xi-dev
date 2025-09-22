package store

import (
	"context"
	"fmt"
	"strings"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/config"
	"xi/pkg/lib/env"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Init initializes DBs once
func (s *StoreService) Init() { s.once.Do(s.InitCore) }

// Initializes all DBs and Redis clients (forced)
func (s *StoreService) InitCore() {
    config.Config.Init()
    s.Hooks.RunPre()

    // --- Initialize SQL DBs ---
    for profile, c := range cfg.Store.Db.Conn {
        if !c.Enable {
            continue
        }

        // Fallbacks for DB credentials
        if c.User == "" {
            c.User = c.Db + "_u"
        }
        if c.Pass == "" {
            c.Pass = env.Env.Get("DB_PASS")
        }
        if c.Charset == "" {
            c.Charset = "utf8mb4"
        }

        switch strings.ToLower(c.Driver) {
        case "mysql", "mariadb":
            dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
                c.User, c.Pass, c.Host, c.Port, c.Db, c.Charset)
            dbCli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
            if err != nil {
                log.Error().Caller().Err(err).Str("profile", profile).Str("type", "MySQL").Msg("db connect")
                continue
            }
            Db.AddCli(profile, dbCli)
            log.Info().Str("profile", profile).Str("type", "MySQL").Msg("db connected")

        case "sqlite":
            dbCli, err := gorm.Open(sqlite.Open(c.Db), &gorm.Config{})
            if err != nil {
                log.Error().Caller().Err(err).Str("profile", profile).Str("type", "SQLite").Msg("db connect")
                continue
            }
            Db.AddCli(profile, dbCli)
            log.Info().Str("profile", profile).Str("type", "SQLite").Msg("db connected")

        default:
            log.Warn().Caller().Str("profile", profile).Str("driver", c.Driver).Msg("db unsupported driver")
        }
    }
	if Db.RawCli() == nil {
		log.Warn().Msg("no active database connections")
	}

    // --- Initialize Redis clients ---
    for profile, c := range cfg.Store.Rdb.Conn {
        if !c.Enable {
            continue
        }

        rdbCli := redis.NewClient(&redis.Options{
            Addr:     c.Host + ":" + c.Port,
            Password: c.Pass,
            DB:       c.Rdb,
        })
        if err := rdbCli.Ping(context.Background()).Err(); err != nil {
            log.Error().Caller().Err(err).Str("profile", profile).Str("type", "Redis").Msg("db connect")
            continue
        }
        Rdb.AddCli(profile, rdbCli)
        log.Info().Str("profile", profile).Str("type", "Redis").Msg("db connected")
    }

    s.Hooks.RunPost()
}
