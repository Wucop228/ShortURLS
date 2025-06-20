package model

import (
	"database/sql"
)

type URL struct {
	ID       int
	BaseUrl  string
	ShortUrl string
}

func AddURL(db *sql.DB, url URL) error {
	query := "INSERT INTO urls(base_url, short_url) VALUES($1, $2)"
	_, err := db.Exec(query, url.BaseUrl, url.ShortUrl)
	return err
}

func GetLastID(db *sql.DB) (int, error) {
	query := "SELECT id FROM urls ORDER BY id DESC LIMIT 1"
	row := db.QueryRow(query)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func GetURL(db *sql.DB, short string) (string, error) {
	query := "SELECT base_url FROM urls WHERE short_url=$1"
	row := db.QueryRow(query, short)

	var base string
	err := row.Scan(&base)
	return base, err
}
