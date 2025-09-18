package rdb

import (
	"context"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// Shared global Redis clients
var sharedClients = make(map[string]*redis.Client)

// RdbStore wraps Redis client management and access
type RdbStore struct {
	prefix     string
	defaultCli string
	clients    map[string]*redis.Client
	client     *redis.Client
	ctx        context.Context

	mu   sync.RWMutex
	once sync.Once
}

// Global instance
var Rdb = &RdbStore{
	prefix:     "app",
	defaultCli: "redis",
	clients:    sharedClients,
	ctx:        context.Background(),
}

func (d *DbStore) initPre() {
	// Set global Redis and DB defaults

	Rdb.SetCtx(context.Background())
	Rdb.SetDefault(cfg.Db.RdbDefault)
	cfg.Db.RdbPrefix = util.Str.Fallback(cfg.Db.RdbPrefix,
		util.Str.IfNotEmptyElse(cfg.Org.Abbr, cfg.Org.Abbr+cfg.Build.Revision, cfg.Build.Revision))
	Rdb.SetPrefix(cfg.Db.RdbPrefix)
}

// New returns a new RdbStore instance with optional prefix/context
func (r *RdbStore) New(defaultCli string, opts ...any) *RdbStore {
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
		prefix:     prefix,
		defaultCli: defaultCli,
		client:     r.GetCli(defaultCli),
		clients:    sharedClients,
		ctx:        ctx,
	}
}

// SetCli registers a new Redis client
func (r *RdbStore) SetCli(name string, client *redis.Client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.clients[name]; exists {
		log.Warn().Msgf("Redis client '%s' already exists", name)
	}
	for n, c := range r.clients {
		if c == client {
			log.Warn().Msgf("Redis client already registered as '%s'", n)
			break
		}
	}

	r.clients[name] = client
	if r.client == nil || r.defaultCli == name || strings.TrimSpace(r.defaultCli) == "" {
		r.client = client
		r.defaultCli = name
	}
}

// GetCli returns a Redis client by name or default
func (r *RdbStore) GetCli(name ...string) *redis.Client {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := r.defaultCli
	if len(name) > 0 && strings.TrimSpace(name[0]) != "" {
		key = name[0]
	}
	return r.clients[key]
}

// SetDefault sets the default Redis client by name
func (r *RdbStore) SetDefault(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.clients) == 0 {
		r.defaultCli = name
		return
	}

	if cli, ok := r.clients[name]; ok {
		r.defaultCli = name
		r.client = cli
		log.Info().Msgf("Redis default: client set to '%s'", name)
	} else {
		log.Warn().Msgf("Redis default: client '%s' not found", name)
	}
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

// GetDefault returns default client name
func (r *RdbStore) GetDefault() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.defaultCli
}

// With returns a new RdbStore bound to the given client name
func (r *RdbStore) With(cliName string) *RdbStore {
	return r.New(cliName, r.prefix)
}
