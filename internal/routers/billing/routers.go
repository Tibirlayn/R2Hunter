package billing

import (
	"github.com/gofiber/fiber/v2"
)

type BillingHandler interface {
	SysOrderList(ctx *fiber.Ctx) error
	SysOrderListEmail(ctx *fiber.Ctx) error
	SetSysOrderList(ctx *fiber.Ctx) error
	

}

func NewRoutersBilling(appf *fiber.App, api BillingHandler) {
	appf.Get("gift", api.SysOrderList)
	appf.Get("gift-email", api.SysOrderListEmail)
	appf.Post("gift-all", api.SetSysOrderList)
}
