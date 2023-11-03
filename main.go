package main

import (
	"RnpServer/internal/config"
	"RnpServer/internal/db"
	"RnpServer/internal/log"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	os.Setenv("CONFIG_PATH", "config/local.yaml")
	cfg := config.MustLoad()

	logger := log.SetupLogger(cfg.Env)
	logger = logger.With(slog.String("env", cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении

	logger.Info("initializing server", slog.String("address", cfg.Address)) // Помимо сообщения выведем параметр с адресом
	logger.Debug("logger debug mode enabled")

	dataBase := new(db.Common)

	err := dataBase.Start(cfg.DbConnection)
	if err != nil {
		panic(err)
	}

	defer dataBase.Close()
}
