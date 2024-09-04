package pc

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
	"github.com/gofiber/fiber/v2"
)

type Pc struct {
	log *slog.Logger
	gamPcProvider GamePcProvider
	tokenTTL time.Duration
}

type GamePcProvider interface {
	PcCard(ctx *fiber.Ctx, name string, pcID int64) ([]game.PcParm, error)
}

func New(log *slog.Logger, gamPcProvider GamePcProvider, tokenTTL time.Duration) *Pc {
	return &Pc{
		log: log,
		gamPcProvider: gamPcProvider,
		tokenTTL: tokenTTL,
	}
}

func (g *Pc) PcCard(ctx *fiber.Ctx, nikname string, pcID int64) (pc []game.PcParm, err error) {
	const op = "service.game.pc.PcCard"

	pcParm, err := g.gamPcProvider.PcCard(ctx, nikname, pcID)
	if err != nil {
		return pc, fmt.Errorf("%s, %w", op, err)
	}

	return pcParm, err
}