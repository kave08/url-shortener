package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func initializeMySQL() (*sql.DB, error) {
	d, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&collation=utf8_unicode_ci&loc=%s&parseTime=true",
			Cfg.Mysql.Username,
			Cfg.Mysql.Password,
			Cfg.Mysql.Host,
			Cfg.Mysql.Port,
			Cfg.Mysql.DBName,
			url.QueryEscape(time.Local.String()),
		),
	)
	if err != nil {
		log.Panicln(err)
	}

	d.SetMaxOpenConns(Cfg.Mysql.MaxOpenConnections)
	d.SetMaxIdleConns(Cfg.Mysql.MaxIdleConnections)
	d.SetConnMaxLifetime(5 * time.Minute)

	if err := d.Ping(); err != nil {
		return nil, err
	}

	return d, nil
}
