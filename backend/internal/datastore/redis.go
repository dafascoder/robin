package datastore

import (
	"backend/internal/config"
	logging "backend/internal/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisStoreInterface interface {
	Get(ctx context.Context, key string) (*string, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Del(ctx context.Context, key string) error
}

type RedisStore struct {
	redisClient *redis.Client
}

func InitRedis() (*RedisStore, error) {
	opt, err := redis.ParseURL(config.Env.RedisUrl)
	if err != nil {
		logging.Logger.LogFatal().Msgf("Fail to connect to redis: %v", err)
		return nil, err
	}

	rdb := redis.NewClient(opt)

	return &RedisStore{
		redisClient: rdb,
	}, nil
}

func (store *RedisStore) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := store.redisClient.Set(ctx, key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (store *RedisStore) Get(ctx context.Context, key string) (*string, error) {
	hash, err := store.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &hash, nil
}

func (store *RedisStore) Del(ctx context.Context, key string) error {
	_, err := store.redisClient.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
