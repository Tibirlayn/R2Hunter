package main

import (
	"fmt"
	"os"

	"github.com/Tibirlayn/R2Hunter/internal/app"
	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/http-server/routers"
	"github.com/Tibirlayn/R2Hunter/internal/logger"
	"github.com/Tibirlayn/R2Hunter/internal/logger/sl"
	"github.com/Tibirlayn/R2Hunter/storage/mssql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// инициализировать конфиг: cleanenv
	cfg, cfgdb := config.MustLoad()

	// инициализировать логгер: log/slog
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")

	// инициализировать СУБД: MS SQL
	storage, err := mssql.New(cfgdb)
	if err != nil {
		log.Error("faild to init storage", sl.Err(err))
		os.Exit(1)
	}

	fmt.Println(storage)

	// инициализировать роутер: chi
/* 	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
 	router.Use(mwLogger.New(log))  логирование запросов подробнее узнать github.com/Tibirlayn/R2Hunter/internal/http-server/middleware/logger
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
 */

	// инициализировать роутер: fiber
	appf := fiber.New()

	appf.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	routers.New(appf)

	// инициализировать приложение (app):
	if err := app.App(appf, cfg); err != nil {
		log.Error("faild to init Listen", sl.Err(err))
	}
}