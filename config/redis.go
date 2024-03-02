package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func initializeRedisConnection() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     Cfg.Redis.Address,
		Password: Cfg.Redis.Password,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
