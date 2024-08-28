package pc

import (
	"github.com/gofiber/fiber/v2"
)

type PcHandler interface {
	PcCard(c *fiber.Ctx) error
}

func NewRoutersPc(appf *fiber.App, api PcHandler) {
	appf.Get("/pc", api.PcCard)
}