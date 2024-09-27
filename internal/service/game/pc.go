package pc

import (
	"fmt"
	"log/slog"
	"time"

	// "github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	queryGame "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"

	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/gofiber/fiber/v2"
)

type Pc struct {
	log *slog.Logger
	gamPcProvider GamePcProvider
	auth *auth.Auth
	tokenTTL time.Duration
}

type GamePcProvider interface {
	PcCard(ctx *fiber.Ctx, name string, pcID int64) ([]queryGame.PcParm, error)

	PcTopLVL(ctx *fiber.Ctx) ([]queryGame.PcTopLVL, error)
	PcTopByGold(ctx *fiber.Ctx) ([]queryGame.PcTopByGold, error)
}

func New(log *slog.Logger, gamPcProvider GamePcProvider, auth *auth.Auth, tokenTTL time.Duration) *Pc {
	return &Pc{
		log: log,
		gamPcProvider: gamPcProvider,
		auth: auth,
		tokenTTL: tokenTTL,
	}
}

func (g *Pc) PcCard(ctx *fiber.Ctx, nikname string, pcID int64) (pc []queryGame.PcParm, err error) {
	const op = "service.game.pc.PcCard"

	pcParm, err := g.gamPcProvider.PcCard(ctx, nikname, pcID)
	if err != nil {
		return pc, fmt.Errorf("%s, %w", op, err)
	}

	return pcParm, err
}

// Запрос на просмотр ТОП 100 игроков по уровню:
func (g *Pc) PcTopLVL(ctx *fiber.Ctx) ([]queryGame.PcTopLVL, error) {
	const op = "service.game.pc.PcTopLVL"

	// проверка на авторизацию 
	_, err := g.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := g.gamPcProvider.PcTopLVL(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (g *Pc) PcTopByGold(ctx *fiber.Ctx) ([]queryGame.PcTopByGold, error) {
	const op = "service.game.pc.PcTopByGold"

	// проверка на авторизацию 
	_, err := g.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := g.gamPcProvider.PcTopByGold(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}
