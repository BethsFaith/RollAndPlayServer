package apiserver

import (
	"RnpServer/internal/app/store"
	"RnpServer/internal/config"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
)

type APIServer struct {
	conf   *config.Config
	logger *slog.Logger
	router *mux.Router
	store  *store.Store
}

func New(conf *config.Config, log *slog.Logger) *APIServer {
	return &APIServer{
		conf:   conf,
		logger: log,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.conf.Address, s.router)
}

func (s *APIServer) Stop() {
	s.store.Close()
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.logger)

	if err := st.Open(s.conf.DbConnection); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// ...
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello")
	}
}
