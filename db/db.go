package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Common struct {
	db *sql.DB
}

func (object *Common) Start(connStr string) error {
	var err error

	object.db, err = sql.Open("postgres", connStr)

	return err
}

func (object *Common) Close() {
	err := object.db.Close()
	if err != nil {
		panic(err)
	}
}

func (object *Common) Create(queryStr string, parameters ...any) (sql.Result, error) {
	return object.db.Exec(queryStr, parameters...)
}

func (object *Common) Select(queryStr string, parameters ...any) (*sql.Rows, error) {
	rows, err := object.db.Query(queryStr, parameters...)
	if err != nil {
		return rows, err
	}

	err = rows.Close()

	return rows, err
}

func (object *Common) Update(queryStr string, parameters ...any) (sql.Result, error) {
	return object.db.Exec(queryStr, parameters...)
}

func (object *Common) Delete(queryStr string, parameters ...any) (sql.Result, error) {
	return object.db.Exec(queryStr, parameters)
}
