package main

import (
	"RnpServer/internal/app/apiserver"
	"RnpServer/internal/config"
	"RnpServer/internal/log"
	"RnpServer/internal/store"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	os.Setenv("CONFIG_PATH", "configs/local.yaml")
	cfg := config.MustLoad()

	logger := log.SetupLogger(cfg.Env)
	logger = logger.With(slog.String("env", cfg.Env))

	logger.Info("initializing data base", slog.String("db", cfg.DbConnection))
	logger.Info("initializing server", slog.String("address", cfg.Address))
	logger.Debug("logger debug mode enabled")

	dataBase := store.New()
	err := dataBase.Open(cfg.DbConnection)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	} else {
		logger.Info("starting data base")
	}
	defer dataBase.Close()

	server := apiserver.New(cfg, logger)
	if err := server.Start(); err != nil {
		logger.Error(err.Error())
	}
}
