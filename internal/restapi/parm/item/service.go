package item

import (
	"fmt"

	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	routersParm "github.com/Tibirlayn/R2Hunter/internal/routers/parm/item"
	"github.com/gofiber/fiber/v2"
)

type Item interface {
	BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.ItemBossDrop, error)
}

type ServiceParmAPI struct {
	item Item
}

func RegisterParm(RestAPI *fiber.App, item Item) {
	api := &ServiceParmAPI{item: item}

	routersParm.NewRoutersPc(RestAPI, api)
}

func (i *ServiceParmAPI) BossDrop(ctx *fiber.Ctx) error {
	const (
		op = "restapi.parm.item.BossDrop"
		empty = "empty"
	)

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	if data["name"] == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := i.item.BossDrop(ctx, data["name"])
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}