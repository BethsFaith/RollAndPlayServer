package main

import (
	"RnpServer/config"
	"RnpServer/db"
	"os"
)

func main() {
	os.Setenv("CONFIG_PATH", "config/local.yaml")
	cfg := config.MustLoad()

	dataBase := new(db.Common)

	err := dataBase.Start(cfg.DbConnection)
	if err != nil {
		panic(err)
	}

	defer dataBase.Close()
}
