package handler

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type InMemoryStorageI interface {
	Set(key, value string, exp time.Duration) error
	Get(key string) (string, error)
}

type storageRedis struct {
	client *redis.Client
}

// NewInMemoryStorage creates a new instance of Redis storage.
func NewInMemoryStorage(rdb *redis.Client) InMemoryStorageI {
	return &storageRedis{
		client: rdb,
	}
}

// Set stores a key-value pair with expiration in Redis.
func (r *storageRedis) Set(key, value string, exp time.Duration) error {
	err := r.client.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the value associated with a key from Redis.
func (r *storageRedis) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
