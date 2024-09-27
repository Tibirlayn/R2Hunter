package parm

import (
	"fmt"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
	"log/slog"
	"time"

	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	qParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/gofiber/fiber/v2"
)

type Parm struct {
	log              *slog.Logger
	parmItemProvider ParmItemProvider
	auth             *auth.Auth
	tokenTTL         time.Duration
}

type ParmItemProvider interface {
	BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error)
	ItemDDrop(ctx *fiber.Ctx, id int) (parm.DropItem, error)
	UpdateItemDDrop(ctx *fiber.Ctx, name parm.DropItem) (parm.DropItem, error)

	ItemsRess(ctx *fiber.Ctx, name string) (qParm.ItemRes, error)
	ItemsRessbyID(ctx *fiber.Ctx, id []int) (qParm.ItemRes, error)
}

func New(log *slog.Logger, parmItemProvider ParmItemProvider, auth *auth.Auth, tokenTTL time.Duration) *Parm {
	return &Parm{
		log:              log,
		parmItemProvider: parmItemProvider,
		auth:             auth,
		tokenTTL:         tokenTTL,
	}
}

func (i *Parm) BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error) {
	const op = "service.parm.item.BossDrop"

	// проверка на авторизацию
	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return []queryParm.MonsterDrop{}, err
	}

	res, err := i.parmItemProvider.BossDrop(ctx, name)
	if err != nil {
		return []queryParm.MonsterDrop{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (i *Parm) ItemDDrop(ctx *fiber.Ctx, id int) (parm.DropItem, error) {
	const op = "service.parm.item.ItemDDrop"

	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return parm.DropItem{}, err
	}

	res, err := i.parmItemProvider.ItemDDrop(ctx, id)
	if err != nil {
		return parm.DropItem{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (i *Parm) UpdateItemDDrop(ctx *fiber.Ctx, name parm.DropItem) (parm.DropItem, error) {
	const op = "service.parm.item.UpdateItemDDrop"

	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return parm.DropItem{}, err
	}

	res, err := i.parmItemProvider.UpdateItemDDrop(ctx, name)
	if err != nil {
		return parm.DropItem{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (i *Parm) ItemsRess(ctx *fiber.Ctx, name string) (qParm.ItemRes, error) {
	const op = "service.parm.item.ItemsRess"

	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return qParm.ItemRes{}, err
	}

	res, err := i.parmItemProvider.ItemsRess(ctx, name)
	if err != nil {
		return qParm.ItemRes{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (i *Parm) ItemsRessbyID(ctx *fiber.Ctx, id []int) (qParm.ItemRes, error) {
	const op = "service.parm.item.ItemsRess"

	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return qParm.ItemRes{}, err
	}

	res, err := i.parmItemProvider.ItemsRessbyID(ctx, id)
	if err != nil {
		return qParm.ItemRes{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}