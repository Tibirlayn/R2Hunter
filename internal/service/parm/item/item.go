package item

import (
	"fmt"
	"log/slog"
	"time"

	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	log *slog.Logger
	parmItemProvider ParmItemProvider
	auth *auth.Auth
	tokenTTL time.Duration 
}

type ParmItemProvider interface {
	BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.ItemBossDrop, error)
}

func New(log *slog.Logger, parmItemProvider ParmItemProvider, auth *auth.Auth, tokenTTL time.Duration) *Item{
	return &Item{
		log: log,
		parmItemProvider: parmItemProvider,
		auth: auth,
		tokenTTL: tokenTTL,
	} 
}

func (i *Item) BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.ItemBossDrop, error) {
	const op = "service.parm.item.BossDrop"

	// проверка на авторизацию 
	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return []queryParm.ItemBossDrop{}, err
	}

	res, err := i.parmItemProvider.BossDrop(ctx, name)
	if err != nil {
		return []queryParm.ItemBossDrop{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}