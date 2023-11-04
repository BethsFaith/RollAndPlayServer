package apiserver

import (
	"RnpServer/internal/app/store/sqlstore"
	"RnpServer/internal/config"
	"database/sql"
	"github.com/gorilla/sessions"
	"golang.org/x/exp/slog"
	"net/http"
)

func Start(config *config.Config, log *slog.Logger) error {
	db, err := newDB(config.DbConnection)

	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(db)

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := newServer(store, sessionStore, log)

	return http.ListenAndServe(config.Address, s)
}

func newDB(dbConnection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConnection)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
