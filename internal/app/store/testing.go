package store

import (
	"RnpServer/internal/log"
	"fmt"
	"golang.org/x/exp/slog"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	logger := log.SetupLogger("local")
	logger = logger.With(slog.String("env", "local"))

	s := New(logger)
	if err := s.Open(databaseURL); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
