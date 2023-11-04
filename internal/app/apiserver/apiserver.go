package apiserver

import (
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

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.conf.Address, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// ...
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello")
	}
}
