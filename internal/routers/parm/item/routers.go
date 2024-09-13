package item

import "github.com/gofiber/fiber/v2"

type ItemHandler interface {
	BossDrop(ctx *fiber.Ctx) error
}

func NewRoutersPc(appf *fiber.App, api ItemHandler) {
	appf.Get("boss-drop", api.BossDrop)
}