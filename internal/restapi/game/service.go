package game

import (
	"fmt"

	//"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	queryGame "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
	routersGame "github.com/Tibirlayn/R2Hunter/internal/routers/game"
	"github.com/gofiber/fiber/v2"
)

type Pc interface {
	PcCard(ctx *fiber.Ctx, nikname string, pcID int64) (pc []queryGame.PcParm, err error)
	PcTopLVL(ctx *fiber.Ctx) (pc []queryGame.PcTopLVL, err error) // Запрос на просмотр ТОП 100 игроков по уровню:
	PcTopByGold(ctx *fiber.Ctx) (pc []queryGame.PcTopByGold, err error) // Запрос на просмотр ТОП 100 игроков по количеству золота:
	// TODO: сделать, чтоб с одного запроса выдавал PcTopLVL и PcTopByGold
}

type ServiceGameAPI struct {
	pc Pc
}

func RegisterGame(RestAPI *fiber.App, pc Pc) {
	api := &ServiceGameAPI{pc: pc}

	routersGame.NewRoutersPc(RestAPI, api)
}

func (g *ServiceGameAPI) PcCard(ctx *fiber.Ctx) error {
	const op = "restapi.game.pc.PcCard"

	return nil
} 

// Запрос на просмотр ТОП 100 игроков по уровню:
func (g *ServiceGameAPI) PcTopLVL(ctx *fiber.Ctx) error {
	const op = "restapi.game.pc.PcTopLVL"

	res, err := g.pc.PcTopLVL(ctx)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

// Запрос на просмотр ТОП 100 игроков по количеству золота:
func (g *ServiceGameAPI) PcTopByGold(ctx *fiber.Ctx) error {
	const op = "restapi.game.pc.PcTopByGold"

	res, err := g.pc.PcTopByGold(ctx)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}