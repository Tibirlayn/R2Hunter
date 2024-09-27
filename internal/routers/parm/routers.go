package parm

import "github.com/gofiber/fiber/v2"

type ParmHandler interface {
	BossDrop(ctx *fiber.Ctx) error
	ItemDDrop(ctx *fiber.Ctx) error
	UpdateItemDDrop(ctx *fiber.Ctx) error
	ItemsRess(ctx *fiber.Ctx) error
}

func NewRoutersParm(appf *fiber.App, api ParmHandler) {
	appf.Get("boss-drop", api.BossDrop)
	appf.Get("item-ddrop", api.ItemDDrop)
	appf.Put("update-item-ddrop", api.UpdateItemDDrop)
	appf.Get("item", api.ItemsRess)
}
