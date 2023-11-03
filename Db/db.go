package Db

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

func (object *Common) Select(queryStr string) error {
	rows, err := object.db.Query(queryStr)
	if err != nil {
		return err
	}

	err = rows.Close()
	return err
}

func (object *Common) Update(queryStr string) (int64, error) {
	result, err := object.db.Exec(queryStr)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
