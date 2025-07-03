package service

import (
	"database/sql"
	"github.com/Wucop228/ShortURLS/internal/cache"
	"github.com/Wucop228/ShortURLS/internal/model"
)

func GetOriginalURL(db *sql.DB, rdb *cache.RedisCache, short string) (string, error) {
	val, err := rdb.Get(short)
	if err == nil {
		return val, nil
	}

	url, err := model.GetURL(db, short)
	if err != nil {
		return "", err
	}

	_ = rdb.Set(short, url, rdb.TTL)

	return url, nil
}

func SetOriginalURL(db *sql.DB, rdb *cache.RedisCache, url model.URL) error {
	err := model.AddURL(db, url)
	if err != nil {
		return err
	}

	return rdb.Set(url.ShortUrl, url.BaseUrl, rdb.TTL)
}
