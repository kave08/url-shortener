package database

import (
	"context"
	"database/sql"
	"time"
)

const (
	InsertUrlQuery = "INSERT INTO url_shortener (short_url, long_url, created_at) VALUES(?, ?, ?)"
	GetUrlQuery    = "SELECT long_url FROM url_shortener WHERE short_url = ? LIMIT 1"
)

type DatabaseInterface interface {
	InsertUrl(ctx context.Context, shortUrl, longUrl string) error
	GetUrl(ctx context.Context, shortUrl string) (string, error)
}

type Database struct {
	mdb *sql.DB
}

func NewDatabase(mdb *sql.DB) DatabaseInterface {
	return &Database{
		mdb: mdb,
	}
}

func (d *Database) InsertUrl(ctx context.Context, shortUrl, longUrl string) error {
	_, err := d.mdb.ExecContext(ctx, InsertUrlQuery, shortUrl, longUrl, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetUrl(ctx context.Context, shortUrl string) (string, error) {
	var longUrl string

	err := d.mdb.QueryRowContext(ctx, GetUrlQuery, shortUrl).Scan(&longUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}
