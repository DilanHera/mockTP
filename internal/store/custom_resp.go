package store

import (
	"database/sql"
	"fmt"
	"strings"
)

type CustomResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Resp string `json:"resp"`
}

type CustomResponseStore struct {
	db *sql.DB
}

func NewCustomResponseStore(db *sql.DB) *CustomResponseStore {
	db.Exec(`
	  CREATE TABLE IF NOT EXISTS custom_responses (
	    id INTEGER PRIMARY KEY,
	    name TEXT UNIQUE,
	    resp TEXT
	  )
	`)
	return &CustomResponseStore{db: db}
}

func (s *CustomResponseStore) Get(name string) (*CustomResponse, error) {
	var resp CustomResponse
	err := s.db.QueryRow("SELECT id, name, resp FROM custom_responses WHERE name = ?", name).Scan(&resp.ID, &resp.Name, &resp.Resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CustomResponseStore) GetMany(names []string) (*[]CustomResponse, error) {
	var resp []CustomResponse
	if len(names) == 0 {
		return &resp, nil
	}

	placeholders := make([]string, len(names))
	args := make([]any, len(names))

	for i, name := range names {
		placeholders[i] = "?"
		args[i] = name
	}

	query := fmt.Sprintf(
		"SELECT id, name, resp FROM custom_responses WHERE name IN (%s)",
		strings.Join(placeholders, ","),
	)
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r CustomResponse
		if err := rows.Scan(&r.ID, &r.Name, &r.Resp); err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}
	return &resp, nil
}

func (s *CustomResponseStore) CreateOrUpdate(name string, resp string) (*CustomResponse, error) {
	_, err := s.db.Exec(`
		INSERT INTO custom_responses (name, resp)
		VALUES (?, ?)
		ON CONFLICT(name) DO UPDATE SET resp = excluded.resp
	`, name, resp)
	if err != nil {
		return nil, err
	}
	return s.Get(name)
}

func (s *CustomResponseStore) Delete(name string) error {
	_, err := s.db.Exec("DELETE FROM custom_responses WHERE name = ?", name)
	return err
}
