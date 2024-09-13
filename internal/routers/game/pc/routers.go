package pc

import (
	"github.com/gofiber/fiber/v2"
)

type PcHandler interface {
	PcCard(c *fiber.Ctx) error
	PcTopLVL(ctx *fiber.Ctx) error
	PcTopByGold(ctx *fiber.Ctx) error
}

func NewRoutersPc(appf *fiber.App, api PcHandler) {
	appf.Get("/pc", api.PcCard)
	appf.Get("/pc-top-lvl", api.PcTopLVL)
	appf.Get("/pc-top-gold", api.PcTopByGold)
}