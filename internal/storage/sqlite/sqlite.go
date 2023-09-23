package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"spinner/internal/storage"
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
-- 		CREATE TABLE IF NOT EXISTS post(
-- 		    id INTEGER PRIMARY KEY,
-- 		    title TEXT NOT NULL,
-- 		    text TEXT,
-- 		    date_update DATA 
-- 		); 
		CREATE TABLE IF NOT EXISTS url(
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

func (s *Storage) SaveURL(urlToSave string, alias string) error {
	const fn = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare(`
		INSERT INTO url(url, alias) VALUES (?, ?)
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const fn = "storage.sqlite.GetURL"
	var url string

	sqlStatement := `SELECT url FROM url WHERE alias=$1 LIMIT 1;`
	row := s.db.QueryRow(sqlStatement, alias)

	switch err := row.Scan(&url); err {
	case sql.ErrNoRows:
		return "", storage.ErrURLNotFound
	case nil:
		return url, nil
	default:
		return "", fmt.Errorf("%s: %w", fn, err)
	}
}

// TODO: Delete
func (s *Storage) DeleteURL(alias string) error {
	const fn = "storage.sqlite.DeleteURL"

}
