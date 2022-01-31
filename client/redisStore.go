package client

import (
	"encoding/json"

	"com.hook.cache/config"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RedisStore interface {
	SaveToHash(ctx context.Context, key string, field string, value interface{}) (bool, error)
	FetchFromHash(ctx context.Context, key string, field string, value interface{}) error
}

type RedisStoreImpl struct {
	redisClient redis.Cmdable
}

func NewRedisStoreImpl(appConfig config.AppConfig) RedisStore {
	client := config.NewRedisClient(appConfig)
	return &RedisStoreImpl{
		redisClient: client,
	}
}

func (r *RedisStoreImpl) SaveToHash(ctx context.Context, key string, field string, value interface{}) (bool, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	r.redisClient.HSet(ctx, key, field, val)
	return true, nil
}

func (r *RedisStoreImpl) FetchFromHash(ctx context.Context, key string, field string, value interface{}) error {
	val, err := r.redisClient.HGet(ctx, key, field).Bytes()
	err1 := json.Unmarshal(val, value)
	if err1 != nil {
		return err
	}
	return nil
}
