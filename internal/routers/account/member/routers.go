package member

import "github.com/gofiber/fiber/v2"

type MemberHandler interface {
	Member(c *fiber.Ctx) error
}



func NewRoutersMember(appf *fiber.App, api MemberHandler) {
	appf.Get("/member", api.Member)
}
