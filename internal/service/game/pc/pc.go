package pc

import (
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
	PcCard(ctx *fiber.Ctx, name string) (game.PcParm, error)
}

func New(log *slog.Logger, gamPcProvider GamePcProvider, tokenTTL time.Duration) *Pc {
	return &Pc{
		log: log,
		gamPcProvider: gamPcProvider,
		tokenTTL: tokenTTL,
	}
}

func (g *Pc) PcCard(ctx *fiber.Ctx, nikname string) (pc game.PcParm, err error) {

	return pc, err
}