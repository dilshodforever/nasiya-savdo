package storage

import (
	"context"
	"time"

	"github.com/dilshodforever/nasiya-savdo/config"
	e "github.com/dilshodforever/nasiya-savdo/email"
	"github.com/redis/go-redis/v9"
)

type InMemoryStorageI interface {
	Set(key, value string, exp time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	SaveToken(email string, token string, exp time.Duration) error
}

type storageRedis struct {
	client *redis.Client
}

func NewInMemoryStorage(rdb *redis.Client) InMemoryStorageI {
	return &storageRedis{
		client: rdb,
	}
}

func (r *storageRedis) Set(key, value string, exp time.Duration) error {
	err := r.client.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *storageRedis) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *storageRedis) Del(key string) error {
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *storageRedis) SaveToken(email string, token string, exp time.Duration) error {
	cnf := config.Load()
	st := e.SendEmailRequest{To: []string{email}, Type: "Verification email", Subject: "Vertification", Code: token}
	err := e.SendEmail(&cnf, &st)
	if err != nil {
		return err
	}
	return r.Set(email, token, exp)
}
