package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/kave08/url-shortener/repository"
)

const (
	ErrRedisNil = "redis: nil"
)

type Logics struct {
	repos *repository.Repository
}

func NewLogics(repos *repository.Repository) *Logics {
	return &Logics{
		repos: repos,
	}
}

func (l *Logics) ShortenUrl(ctx context.Context, longUrl string) (string, error) {
	hash := md5.Sum([]byte(longUrl))
	encUrl := hex.EncodeToString(hash[:])

	err := l.repos.Database.InsertUrl(ctx, encUrl, longUrl)
	if err != nil {
		return "", err
	}

	return encUrl, nil
}

func (l *Logics) ResolveUrl(ctx context.Context, shortUrl string) (string, error) {
	var longUrl string

	longUrl, err := l.repos.Cache.GetCache(ctx, shortUrl)
	if err != nil {
		if err.Error() == ErrRedisNil {
			longUrl, err := l.repos.Database.GetUrl(ctx, shortUrl)
			if err != nil {
				return "", err
			}
			err = l.repos.Cache.InsertCache(ctx, shortUrl, longUrl)
			if err != nil {
				return "", err
			}

			return longUrl, nil
		}

		return "", err
	}

	return longUrl, nil
}
