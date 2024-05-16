package routers

import (
	"github.com/Tibirlayn/R2Hunter/internal/http-server/handlers"
	"github.com/gofiber/fiber/v2"
)

func New(appf *fiber.App) {
	handlers.SetupRouter(appf)
}