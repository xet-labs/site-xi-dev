package rdb

import (
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func (r *RdbStore) key(k string) string {
	return r.prefix + ":" + k
}

func (r *RdbStore) Set(k string, v any, ttl time.Duration) error {
	if r.client == nil {
		return redis.ErrClosed
	}

	if err := r.client.Set(r.ctx, r.key(k), v, ttl).Err(); err != nil {
		log.Warn().Err(err).Str("key", k).Msg("rdb SET")
		return err
	}
	return nil
}

func (r *RdbStore) Get(k string) (string, error) {
	if r.client == nil {
		return "", redis.ErrClosed
	}

	v, err := r.client.Get(r.ctx, r.key(k)).Result()
	if err != nil {
		log.Warn().Err(err).Str("key", k).Msg("rdb GET")
	}
	return v, err
}

func (r *RdbStore) GetBytes(k string) ([]byte, error) {
	if r.client == nil {
		return nil, redis.ErrClosed
	}
	return r.client.Get(r.ctx, r.key(k)).Bytes()
}

func (r *RdbStore) SetJson(k string, v any, ttl time.Duration) error {
	val, err := json.Marshal(v)
	if err != nil {
		log.Warn().Err(err).Str("key", k).Msg("rdb SetJson")
		return err
	}
	return r.Set(k, val, ttl)
}

func (r *RdbStore) GetJson(k string, out any) error {
	v, err := r.GetBytes(k)
	if err != nil {
		return err
	}
	return json.Unmarshal(v, out)
}

func (r *RdbStore) Del(keys ...string) error {
	if r.client == nil {
		return redis.ErrClosed
	}
	if len(keys) == 0 {
		return nil
	}

	var redisKeys []string
	for _, k := range keys {
		redisKeys = append(redisKeys, r.key(k))
	}

	return r.client.Del(r.ctx, redisKeys...).Err()
}

func (r *RdbStore) Exists(k string) (bool, error) {
	if r.client == nil {
		return false, redis.ErrClosed
	}
	n, err := r.client.Exists(r.ctx, r.key(k)).Result()
	return n > 0, err
}

func (r *RdbStore) Keys(pattern string) ([]string, error) {
	if r.client == nil {
		return nil, redis.ErrClosed
	}
	return r.client.Keys(r.ctx, r.key(pattern)).Result()
}

func (r *RdbStore) FlushAll() error {
	if r.client == nil {
		return redis.ErrClosed
	}
	return r.client.FlushAll(r.ctx).Err()
}
