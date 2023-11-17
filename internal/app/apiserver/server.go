package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

type server struct {
	router       *mux.Router
	logger       *slog.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store, log *slog.Logger) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       log,
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

// ServeHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	// /private/***
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/who-am-i", s.handleWhoami()).Methods("GET")

	private.HandleFunc("/skills", s.handleSkillCreate()).Methods("POST")
	private.HandleFunc("/skills", s.handleSkillUpdate()).Methods("PUT")
	private.HandleFunc("/skills", s.handleSkillDelete()).Methods("DELETE")

	private.HandleFunc("/skill-categories", s.handleSkillCategoryCreate()).Methods("POST")
	private.HandleFunc("/skill-categories", s.handleSkillCategoryUpdate()).Methods("PUT")
	private.HandleFunc("/skill-categories", s.handleSkillCategoryDelete()).Methods("DELETE")

	private.HandleFunc("/races", s.handleRaceCreate()).Methods("POST")
	private.HandleFunc("/races", s.handleRaceUpdate()).Methods("PUT")
	private.HandleFunc("/races", s.handleRaceDelete()).Methods("DELETE")
	private.HandleFunc("/races/bonuses", s.handleRaceBonuses()).Methods("GET")
	private.HandleFunc("/races/bonuses", s.handleRaceBonusCreate()).Methods("POST")
	private.HandleFunc("/races/bonuses", s.handleRaceBonusUpdate()).Methods("PUT")
	private.HandleFunc("/races/bonuses", s.handleRaceBonusDelete()).Methods("DELETE")

	private.HandleFunc("/actions", s.handleActionCreate()).Methods("POST")
	private.HandleFunc("/actions", s.handleActionUpdate()).Methods("PUT")
	private.HandleFunc("/actions", s.handleActionDelete()).Methods("DELETE")

	private.HandleFunc("/classes", s.handleClassCreate()).Methods("POST")
	private.HandleFunc("/classes", s.handleClassUpdate()).Methods("PUT")
	private.HandleFunc("/classes", s.handleClassDelete()).Methods("DELETE")

	private.HandleFunc("/classes/bonuses", s.handleClassBonuses()).Methods("GET")
	private.HandleFunc("/classes/bonuses", s.handleClassBonusCreate()).Methods("POST")
	private.HandleFunc("/classes/bonuses", s.handleClassBonusUpdate()).Methods("PUT")
	private.HandleFunc("/classes/bonuses", s.handleClassBonusDelete()).Methods("DELETE")
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.With(
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("request_id", r.Context().Value(ctxKeyRequestID).(string)),
		)

		logger.Info(fmt.Sprintf("started %s %s", r.Method, r.RequestURI))

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Info(fmt.Sprintf(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		))
	})
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, http.StatusUnauthorized, ErrorNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, http.StatusUnauthorized, ErrorNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, http.StatusUnauthorized, ErrorIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
