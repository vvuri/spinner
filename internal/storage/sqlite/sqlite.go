package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const fn = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	// пока без использования миграций, создание напрямую
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS post(
		    id INTEGER PRIMARY KEY,
		    title TEXT NOT NULL,
		    text TEXT,
		    date_update DATA 
		); 
		CREATE TABLE IF NOT EXISTS post(
		    id INTEGER PRIMARY KEY,
		    alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
			);
		CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: db}, nil
}
