package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New() *Store {
	return &Store{}
}

func (object *Store) Open(connStr string) error {
	var err error

	object.db, err = sql.Open("postgres", connStr)

	return err
}

func (object *Store) Close() {
	err := object.db.Close()
	if err != nil {
		panic(err)
	}
}

func (object *Store) Create(queryStr string, parameters ...any) (sql.Result, error) {
	return object.db.Exec(queryStr, parameters...)
}

func (object *Store) Select(queryStr string, parameters ...any) (*sql.Rows, error) {
	rows, err := object.db.Query(queryStr, parameters...)
	if err != nil {
		return rows, err
	}

	err = rows.Close()

	return rows, err
}

func (object *Store) Update(queryStr string, parameters ...any) (sql.Result, error) {
	return object.db.Exec(queryStr, parameters...)
}

func (object *Store) Delete(queryStr string, parameters ...any) (sql.Result, error) {
	return object.db.Exec(queryStr, parameters)
}
