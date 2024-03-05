package repository

import (
	"database/sql"
	"github.com/kave08/url-shortener/repository/cache"
	"github.com/kave08/url-shortener/repository/database"


	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Database database.DatabaseInterface
	Cache    cache.CacheInterface
}

func NewRepository(mdb *sql.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Database: database.NewDatabase(mdb),
		Cache:    cache.NewCache(rdb),
	}
}
