package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client  *redis.Client
	timeout time.Duration
}

func NewRedis(addr string) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		PoolSize:     10,
		MinIdleConns: 3,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}
	return &Redis{
		client:  client,
		timeout: 3 * time.Second,
	}, nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}

func (r *Redis) Set(key string, value any, ttl time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	data, _ := json.Marshal(value)
	r.client.Set(ctx, key, data, ttl)
}

func (r *Redis) Get(key string) (any, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, false
	}
	var value any
	json.Unmarshal([]byte(data), &value)
	return value, true
}

func (r *Redis) Delete(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	r.client.Del(ctx, key)
}

func (r *Redis) Exists(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	count, _ := r.client.Exists(ctx, key).Result()
	return count > 0
}
