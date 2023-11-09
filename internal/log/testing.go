package log

import "golang.org/x/exp/slog"

func TestLogger() *slog.Logger {
	env := "local"
	logger := SetupLogger(env)
	return logger.With(slog.String("env", env))
}
