package member

import "github.com/gofiber/fiber/v2"

type MemberHandler interface {
	Member(ctx *fiber.Ctx) error
	MemberAll(ctx *fiber.Ctx) error
	UserSearch(ctx *fiber.Ctx) error
}



func NewRoutersMember(appf *fiber.App, api MemberHandler) {
	appf.Get("/member", api.Member)
	appf.Get("/member-all", api.MemberAll)
	appf.Get("/user-search", api.UserSearch)
}
