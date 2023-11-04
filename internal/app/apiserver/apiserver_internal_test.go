package apiserver

import (
	"RnpServer/internal/config"
	"RnpServer/internal/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApiServer_HandleHello(t *testing.T) {
	os.Setenv("CONFIG_PATH", "D:\\GO\\Projects\\RnpServer\\configs\\local.yaml")
	cfg := config.MustLoad()
	logger := log.SetupLogger(cfg.Env)

	s := New(cfg, logger)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "Hello")
}
