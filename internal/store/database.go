package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./mock.db")
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}

	fmt.Println("Connected to Database...")
	return db, nil
}
