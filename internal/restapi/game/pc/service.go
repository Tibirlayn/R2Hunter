package game

import (
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
	routersGame "github.com/Tibirlayn/R2Hunter/internal/routers/game/pc"
	"github.com/gofiber/fiber/v2"
)

type Pc interface {
	PcCard(ctx *fiber.Ctx, nikname string) (pc game.PcParm, err error)
}

type ServiceGameAPI struct {
	pc Pc
}

func RegisterGame(RestAPI *fiber.App, pc Pc) {
	api := &ServiceGameAPI{pc: pc}

	routersGame.NewRoutersPc(RestAPI, api)
}

func (g *ServiceGameAPI) PcCard(ctx *fiber.Ctx) error {

	return nil
} 
