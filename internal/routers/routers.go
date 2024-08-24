package routers

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
    Register(c *fiber.Ctx) error
}

func NewRoutersAccount(appf *fiber.App, api AuthHandler) {
	appf.Post("/login", api.Login)
    appf.Post("/register", api.Register)
}