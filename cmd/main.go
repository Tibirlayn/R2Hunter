package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tibirlayn/R2Hunter/internal/app"
	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/logger"
)

func main() {
	// инициализировать конфиг: cleanenv
	cfg, cfgdb := config.MustLoad()

	// инициализировать логгер: log/slog
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")
	

	// инициализировать приложение (app):
	application := app.New(log, cfg.HTTPServer.Address, cfgdb, cfg.TokenTTL);

	go func() {
		application.RestApi.MustLoad()
	}()

	// Настройка канала для получения сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала
	sign := <-quit
	log.Info("stopping application", slog.String("signal", sign.String()))

	application.RestApi.Stop()
	log.Info("Server gracefully stopped")
}