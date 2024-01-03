package main

import (
	"RnpServer/internal/app/apiserver"
	"RnpServer/internal/config"
	"RnpServer/internal/log"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	if err := os.Setenv("CONFIG_PATH", "configs/local.yaml"); err != nil {
		panic(err)
	}
	cfg := config.MustLoad()

	logger := log.SetupLogger(cfg.Env)
	logger = logger.With(slog.String("env", cfg.Env))

	logger.Info("initializing data base", slog.String("db", cfg.DbConnection))
	logger.Info("initializing server", slog.String("address", cfg.Address))
	logger.Debug("logger debug mode enabled")

	if err := apiserver.Start(cfg, logger); err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}
