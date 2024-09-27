package parm

import (
	"fmt"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	qParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm"
	routersParm "github.com/Tibirlayn/R2Hunter/internal/routers/parm"
	"github.com/gofiber/fiber/v2"
)

type Parm interface {
	BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error)
	ItemDDrop(ctx *fiber.Ctx, id int) (parm.DropItem, error)
	UpdateItemDDrop(ctx *fiber.Ctx, name parm.DropItem) (parm.DropItem, error)

	ItemsRess(ctx *fiber.Ctx, name string) (qParm.ItemRes, error)
}


type ServiceParmAPI struct {
	parm Parm
}

func RegisterParm(RestAPI *fiber.App, parm Parm) {
	api := &ServiceParmAPI{parm: parm}

	routersParm.NewRoutersParm(RestAPI, api)
}

func (i *ServiceParmAPI) BossDrop(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.parm.item.BossDrop"
		empty = "empty"
	)

	// Получаем параметр name из строки запроса
	name := ctx.Query("name")

	if name == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := i.parm.BossDrop(ctx, name)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (i *ServiceParmAPI) ItemDDrop(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.parm.item.ItemDDrop"
		empty = "empty"
	)

	id := ctx.QueryInt("id", 0)

	if id == 0 {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := i.parm.ItemDDrop(ctx, id)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (i *ServiceParmAPI) ItemsRess(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.parm.item.BossDrop"
		empty = "empty"
	)

	// Получаем параметр name из строки запроса
	name := ctx.Query("name")

	if name == "" {
		return fmt.Errorf("%s, %s", op, empty)
	} 


	res, err := i.parm.ItemsRess(ctx, name)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (i *ServiceParmAPI) UpdateItemDDrop(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.parm.item.UpdateItemDDrop"
		empty = "empty"
	)

	var di parm.DropItem

	if err := ctx.BodyParser(&di); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	res, err := i.parm.UpdateItemDDrop(ctx, di)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

