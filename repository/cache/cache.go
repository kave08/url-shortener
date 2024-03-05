package cache

import (
	"context"

	"github.com/kave08/url-shortener/config"
	"github.com/redis/go-redis/v9"
)

type CacheInterface interface {
	InsertCache(ctx context.Context, shortUrl, longUrl string) error
	GetCache(ctx context.Context, shortUrl string) (string, error)
}

type Cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) CacheInterface {
	return &Cache{
		rdb: rdb,
	}
}

func (c *Cache) InsertCache(ctx context.Context, shortUrl, longUrl string) error {
	err := c.rdb.Set(ctx, shortUrl, longUrl, config.Cfg.Redis.TTL).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) GetCache(ctx context.Context, shortUrl string) (string, error) {
	longUrl, err := c.rdb.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", err
	}

	return longUrl, nil
}
