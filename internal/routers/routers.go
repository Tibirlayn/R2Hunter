package routers

import (
	"github.com/Tibirlayn/R2Hunter/internal/restapi/account/auth"
	"github.com/gofiber/fiber/v2"
)

func NewRoutersAccount(appf *fiber.App, api *auth.ServerAPI) {
	appf.Post("/login", func(c *fiber.Ctx) error {
		
		return nil
	})
}