package apiserver

import (
	"RnpServer/internal/app/store/teststore"
	"RnpServer/internal/log"
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

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New(), logger)
	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
