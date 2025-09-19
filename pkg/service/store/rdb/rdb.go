package rdb

import (
	"context"
	"errors"
	"strings"
	"sync"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/util"

	"github.com/redis/go-redis/v9"
)

// RdbStore wraps Redis Cli management and access
type RdbStore struct {
	Cli        *redis.Client
	CliProfile string
	clis       map[string]*redis.Client
	ctx        context.Context
	prefix     string

	mu sync.RWMutex
}

// Global instance
var Rdb = &RdbStore{
	prefix: "app",
	clis:   make(map[string]*redis.Client),
	ctx:    context.Background(),
}

// New returns a new RdbStore instance with optional prefix/context
func (r *RdbStore) New(cliProfile string, opts ...any) *RdbStore {
	prefix, ctx := r.prefix, r.ctx
	for _, opt := range opts {
		switch v := opt.(type) {
		case string:
			if s := strings.TrimSpace(v); s != "" {
				prefix = s
			}
		case context.Context:
			ctx = v
		}
	}

	return &RdbStore{
		Cli:        r.GetCli(cliProfile),
		CliProfile: cliProfile,
		clis:       make(map[string]*redis.Client),
		ctx:        ctx,
		prefix:     prefix,
	}
}

// AddCli registers a new Redis Cli
func (r *RdbStore) AddCli(cliProfile string, cli *redis.Client, opts ...any) error {
	if cli == nil {
		return errors.New("rdb cli is nil for profile '" + cliProfile + "'")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.clis[cliProfile] = cli

	// Set as global if this is the default profile OR if global isn't set yet
	if cfg.Store.Rdb.DefaultProfile == cliProfile || r.Cli == nil {
		r.Cli = cli
		r.CliProfile = cliProfile
	}

	if r.prefix == "" {
		Rdb.SetPrefix(util.Str.Fallback(cfg.Store.Rdb.Prefix,
			util.Str.IfNotEmptyElse(cfg.Org.Abbr, cfg.Org.Abbr+cfg.App.Build.Revision, cfg.App.Build.Revision)))
	}
	return nil
}

// Set sets the global Redis Cli by name
func (r *RdbStore) SetCli(cliProfile string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cli, ok := r.clis[cliProfile]; ok {
		r.Cli = cli
		r.CliProfile = cliProfile
	}
}

// GetCli returns a Redis Cli by name or default
func (r *RdbStore) GetCli(cliProfiles ...string) *redis.Client {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Try profiles in order
	for _, profile := range cliProfiles {
		if cli, ok := r.clis[profile]; ok && cli != nil {
			return cli
		}
	}

	// Fall back to global
	if r.Cli != nil {
		return r.Cli
	}

	// Nothing found
	return nil
}

// SetPrefix updates the Redis key prefix
func (r *RdbStore) SetPrefix(prefix string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.prefix = strings.TrimSpace(prefix)
}

// GetPrefix returns current Redis key prefix
func (r *RdbStore) GetPrefix() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.prefix
}

// SetCtx sets Redis context
func (r *RdbStore) SetCtx(ctx context.Context) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ctx = ctx
}

// GetCtx returns current context
func (r *RdbStore) GetCtx() context.Context {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.ctx
}

// With returns a new RdbStore bound to the given Cli name
func (r *RdbStore) With(cliName string) *RdbStore {
	return r.New(cliName, r.prefix)
}
