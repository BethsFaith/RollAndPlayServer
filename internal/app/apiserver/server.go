package apiserver

import (
	"RnpServer/internal/app/store"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *slog.Logger
	store  store.Store
}

func newServer(store store.Store, log *slog.Logger) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: log,
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
