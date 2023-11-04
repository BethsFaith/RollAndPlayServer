package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store/teststore"
	"RnpServer/internal/log"
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	env := "local"
	logger := log.SetupLogger(env)
	logger = logger.With(slog.String("env", env))

	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")), logger)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	env := "local"
	logger := log.SetupLogger(env)
	logger = logger.With(slog.String("env", env))

	u := model.TestUser(t)
	store := teststore.New()
	s := newServer(store, sessions.NewCookieStore([]byte("secret")), logger)

	if err := store.User().Create(u); err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req := httptest.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
