package restapi

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	authRestAPI "github.com/Tibirlayn/R2Hunter/internal/restapi/account/auth"
	membRestAPI "github.com/Tibirlayn/R2Hunter/internal/restapi/account/member"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	log *slog.Logger
	appf *fiber.App
	address string
}

func New(log *slog.Logger, authService authRestAPI.Auth, memberService membRestAPI.Member, address string) *App {
	// инициализировать роутер: fiber
	appFiber := fiber.New()
	authRestAPI.Register(appFiber, authService)
	membRestAPI.RegisterMember(appFiber, memberService)

	return &App{
		log: log,
		appf: appFiber,
		address: address,
	}
}

func (a *App) MustLoad() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "restapi.Run"

	if err :=  a.appf.Listen(a.address); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	// необходимо почитать и переместить в другую папку middleware
	a.appf.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	a.log.Info("rest api server started", slog.String("addr", a.address))


	return nil
}

func (a *App) Stop() {
	const op = "restapi.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping rest api server", slog.String("addr", a.address))

	// Контекст с таймаутом для Graceful Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.appf.ShutdownWithContext(ctx); err != nil {
		log.Error("Error shutting down server:", err)
	}
}