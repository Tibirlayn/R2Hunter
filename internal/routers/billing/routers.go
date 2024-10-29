package billing

import (
	"github.com/gofiber/fiber/v2"
)

type BillingHandler interface {
	SysOrderList(ctx *fiber.Ctx) error // выдать все данные о падарках на всех персонажа (не очень нужна вещь)
	SysOrderListEmail(ctx *fiber.Ctx) error // получить данные о подарках у персонажа
	SetSysOrderList(ctx *fiber.Ctx) error // выдать подарок одному персонажу
	SetSysOrderListAll(ctx *fiber.Ctx) error // выдать подарки всем персонажам
	Shop(ctx *fiber.Ctx) error // получить список предметов в магазине
	DeleteShop(ctx *fiber.Ctx) error 
	
}

func NewRoutersBilling(appf *fiber.App, api BillingHandler) {
	appf.Get("gift", api.SysOrderList)
	appf.Get("gift-email", api.SysOrderListEmail)
	appf.Post("add-gift", api.SetSysOrderList)
	appf.Post("add-gift-all", api.SetSysOrderListAll)
	appf.Get("shop", api.Shop)
	appf.Delete("delete-shop", api.DeleteShop)
}
