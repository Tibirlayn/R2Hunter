package auth

import (
	"github.com/Tibirlayn/R2Hunter/internal/routers"
	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Login(ctx *fiber.Ctx, email string, password string, appID int) (token string, err error)
	RegisterNewUser(ctx *fiber.Ctx, email string, password string) (userID int64, err error)
}

type ServerAPI struct {
	auth Auth
}

func Register(RestAPI *fiber.App, auth Auth) {
	api := &ServerAPI{auth: auth}

	routers.NewRoutersAccount(RestAPI, api)
}

func (s *ServerAPI) Login(c *fiber.Ctx) (token string, err error) {
	const op = "restapi.account.auth.service.Login"

	return
}

func (s *ServerAPI) Register(c *fiber.Ctx) (userID int64, err error) {
	const op = "restapi.account.auth.service.Register"

	return
}

