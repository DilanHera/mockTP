package store

import (
	"database/sql"
	"fmt"
	"strings"
)

type ApiInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Resp     string `json:"resp"`
	State    string `json:"state"`
	HttpCode int    `json:"http_code"`
}

type ApiInfoStore struct {
	db *sql.DB
}

func NewApiInfoStore(db *sql.DB) *ApiInfoStore {
	db.Exec(`
	  CREATE TABLE IF NOT EXISTS app_info (
	    id INTEGER PRIMARY KEY,
	    name TEXT UNIQUE,
	    resp TEXT,
	    state TEXT,
	    http_code INTEGER
	  )
	`)
	return &ApiInfoStore{db: db}
}

func (s *ApiInfoStore) GetAll() (*[]ApiInfo, error) {
	var resp []ApiInfo
	rows, err := s.db.Query("SELECT id, name, resp, state, http_code FROM app_info")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r ApiInfo
		if err := rows.Scan(&r.ID, &r.Name, &r.Resp, &r.State, &r.HttpCode); err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}
	return &resp, nil
}

func (s *ApiInfoStore) Get(name string) (*ApiInfo, error) {
	var resp ApiInfo
	err := s.db.QueryRow("SELECT id, name, resp, state, http_code FROM app_info WHERE name = ?", name).Scan(&resp.ID, &resp.Name, &resp.Resp, &resp.State, &resp.HttpCode)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ApiInfoStore) GetMany(names []string) (*[]ApiInfo, error) {
	var resp []ApiInfo
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
		"SELECT id, name, resp, state, http_code FROM app_info WHERE name IN (%s)",
		strings.Join(placeholders, ","),
	)
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r ApiInfo
		if err := rows.Scan(&r.ID, &r.Name, &r.Resp, &r.State, &r.HttpCode); err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}
	return &resp, nil
}

func (s *ApiInfoStore) Create(name string, resp string, state string, httpCode int) (*ApiInfo, error) {
	_, err := s.db.Exec("INSERT OR IGNORE INTO app_info (name, resp, state, http_code) VALUES (?, ?, ?, ?)", name, resp, state, httpCode)
	if err != nil {
		return nil, err
	}
	return s.Get(name)
}

func (s *ApiInfoStore) UpdateState(name string, state string) error {
	_, err := s.db.Exec("UPDATE app_info SET state = ? WHERE name = ?", state, name)
	return err
}

func (s *ApiInfoStore) UpdateResp(name string, resp string) error {
	_, err := s.db.Exec("UPDATE app_info SET resp = ? WHERE name = ?", resp, name)
	return err
}

func (s *ApiInfoStore) UpdateHttpCode(name string, httpCode int) error {
	_, err := s.db.Exec("UPDATE app_info SET http_code = ? WHERE name = ?", httpCode, name)
	return err
}

func (s *ApiInfoStore) Delete(name string) error {
	_, err := s.db.Exec("DELETE FROM app_info WHERE name = ?", name)
	return err
}
