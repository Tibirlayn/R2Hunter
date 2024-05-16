package app

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/gofiber/fiber/v2"
)

func App(appf *fiber.App, cfg *config.Config) error {
	const op = "internal.app.App"

	if err := appf.Listen(cfg.HTTPServer.Address); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return nil
}