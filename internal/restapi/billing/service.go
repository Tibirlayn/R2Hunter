package billing

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/billing"
	routersBilling "github.com/Tibirlayn/R2Hunter/internal/routers/billing"
	qBilling "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/billing"
	"github.com/gofiber/fiber/v2"
)

type Billing interface {
	//посмотреть все подарки
	SysOrderList(ctx *fiber.Ctx) ([]billing.SysOrderList, error) 
	//посмотреть в определенном аккаунте
	SysOrderListEmail(ctx *fiber.Ctx, email string) ([]qBilling.SolItemRes, error) 
	//выдать подарок всем персонажам
	SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error) 
}

type ServiceBillingAPI struct {
	billing Billing
}

func RegisterBilling(RestAPI *fiber.App, billing Billing) {
	api := &ServiceBillingAPI{billing: billing}

	routersBilling.NewRoutersBilling(RestAPI, api)
}

func (b *ServiceBillingAPI) SetSysOrderList(ctx *fiber.Ctx) error {
	const op = "restapi.billing.SetSysOrderList"

	// var data map[string]string
	// if err := ctx.BodyParser(&data); err != nil {
	// 	return fmt.Errorf("%s, %w", op, err)
	// }

	

	// id, _ := strconv.ParseInt(data["id"], 10, 64)                        /* — ID сообщения от администратора */
	// svr, _ := strconv.ParseInt(data["svr"], 10, 16)                      /* — Номер вашего сервера */
	// itemid, _ := strconv.Atoi(data["itemid"])                            /* — Номер предмета (подарка) */
	// cnt, _ := strconv.Atoi(data["cnt"])                                  /* — Количество */
	// aperiod, _ := strconv.Atoi(data["aperiod"])                          /* — Доступный период (сколько будет лежать в подароках) */
	// pperiod, _ := strconv.Atoi(data["pperiod"])                          /* — Практический период (количество времени которое будет у предмета после получения)*/
	// binding, _ := strconv.ParseUint(data["binding"], 10, 8)              /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	// limitedDate, _ := time.Parse("2006-01-02 15:04:05", data["limited"]) /* — Ограниченная дата */
	// status, _ := strconv.ParseUint(data["status"], 10, 8)                /* — Статус предмета */

	// if err := ctx.BodyParser(&data); err != nil {
	// 	return err
	// }

	return ctx.JSON(nil)
}

func (b *ServiceBillingAPI) SysOrderList(ctx *fiber.Ctx) error {
	const op = "restapi.billing.SysOrderList"

	res, err := b.billing.SysOrderList(ctx)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (b *ServiceBillingAPI) SysOrderListEmail(ctx *fiber.Ctx) error {
	const (
		op = "restapi.billing.SysOrderListEmail"
		empty = "empty"
	)

	name := ctx.Query("name")
	if name == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := b.billing.SysOrderListEmail(ctx, name)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}