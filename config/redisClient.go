package config

import (
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(appConfig AppConfig) redis.Cmdable {
	options := &redis.Options{
		Addr:         appConfig.GetHost(),
		DialTimeout:  time.Duration(appConfig.GetDialTimeout()),
		ReadTimeout:  time.Duration(appConfig.GetDialTimeout()),
		WriteTimeout: time.Duration(appConfig.GetDialTimeout()),
	}
	return redis.NewClient(options)
}
