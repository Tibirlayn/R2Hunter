package main

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/logger"
	"github.com/Tibirlayn/R2Hunter/storage/mssql"
)

func main() {
	// инициализировать конфиг: cleanenv
	cfg, cfgdb := config.MustLoad()

	// инициализировать логгер: log/slog
	loggers := logger.SetupLogger(cfg.Env)
	loggers.Info("starting application")

	// инициализировать СУБД: MS SQL
	storage, err := mssql.New(cfgdb)
	if err != nil {
		panic(err)
	}

	fmt.Println(storage)

	// инициализировать роутер: chi

	// инициализировать приложение (app):
}