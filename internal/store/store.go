package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

type Store struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(l *slog.Logger) *Store {
	return &Store{
		logger: l,
	}
}

func (s *Store) Open(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db
	s.logger.Info("starting data base")

	return err
}

func (s *Store) Close() {
	err := s.db.Close()
	if err != nil {
		panic(err)
	}
}

func (s *Store) Create(queryStr string, parameters ...any) (sql.Result, error) {
	return s.db.Exec(queryStr, parameters...)
}

func (s *Store) Select(queryStr string, parameters ...any) (*sql.Rows, error) {
	rows, err := s.db.Query(queryStr, parameters...)
	if err != nil {
		return rows, err
	}

	err = rows.Close()

	return rows, err
}

func (s *Store) Update(queryStr string, parameters ...any) (sql.Result, error) {
	return s.db.Exec(queryStr, parameters...)
}

func (s *Store) Delete(queryStr string, parameters ...any) (sql.Result, error) {
	return s.db.Exec(queryStr, parameters)
}
