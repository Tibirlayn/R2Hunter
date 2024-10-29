package billing

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/billing"
	qBilling "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/billing"
	routersBilling "github.com/Tibirlayn/R2Hunter/internal/routers/billing"
	"github.com/gofiber/fiber/v2"
)

type Billing interface {
	//посмотреть все подарки
	SysOrderList(ctx *fiber.Ctx) ([]billing.SysOrderList, error) 
	//посмотреть в определенном аккаунте
	SysOrderListEmail(ctx *fiber.Ctx, email string) ([]qBilling.SolItemRes, error) 
	// выдать подарок персонажу
	SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error) 
	// выдать подарок всем персонажам
	SetSysOrderListAll(ctx *fiber.Ctx, gift billing.SysOrderList) error 
	// Магазин
	Shop(ctx *fiber.Ctx) ([]qBilling.ShopItemRes, error)
	DeleteShop(ctx *fiber.Ctx, goldItemID int) (string, error)
}

type ServiceBillingAPI struct {
	billing Billing
}

func RegisterBilling(RestAPI *fiber.App, billing Billing) {
	api := &ServiceBillingAPI{billing: billing}

	routersBilling.NewRoutersBilling(RestAPI, api)
}

func (b *ServiceBillingAPI) SetSysOrderList(ctx *fiber.Ctx) error {
	const 
	(op = "restapi.billing.SetSysOrderList"
	empty = "empty")

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	userNo, _ := strconv.Atoi(data["userNo"])
	// id, _ := strconv.ParseInt(data["id"], 10, 64)                      /* — ID сообщения от администратора */
	svr, _ := strconv.ParseInt(data["svr"], 10, 16)                      /* — Номер вашего сервера */
	itemid, _ := strconv.Atoi(data["itemid"])                            /* — Номер предмета (подарка) */
	cnt, _ := strconv.Atoi(data["cnt"])                                  /* — Количество */
	aperiod, _ := strconv.Atoi(data["aperiod"])                          /* — Доступный период (сколько будет лежать в подароках) */
	pperiod, _ := strconv.Atoi(data["pperiod"])                          /* — Практический период (количество времени которое будет у предмета после получения)*/
	binding, _ := strconv.ParseUint(data["binding"], 10, 8)              /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	limitedDate, _ := time.Parse("2006-01-02 15:04:05", data["limitedDate"]) /* — Ограниченная дата */
	status, _ := strconv.ParseUint(data["status"], 10, 8)                /* — Статус предмета */

	// TODO: сделать валидацию 

	gift := billing.SysOrderList{                          
		MUserNo: userNo,            
		MItemID: itemid,
		MSvrNo: int16(svr),
		MCnt: cnt,
		MAvailablePeriod: aperiod,
		MPracticalPeriod: pperiod,  
		MBindingType: uint8(binding),
		MLimitedDate: limitedDate,
		MItemStatus: uint8(status),
	}
	
	res, err := b.billing.SetSysOrderList(ctx, gift)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}

func (b *ServiceBillingAPI) SetSysOrderListAll(ctx *fiber.Ctx) error {
	const (
		op = "restapi.billing.SetSysOrderListAll"
		empty = "empty"
	)

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	// userNo, _ := strconv.Atoi(data["userNo"])
	// id, _ := strconv.ParseInt(data["id"], 10, 64)                      /* — ID сообщения от администратора */
	svr, _ := strconv.ParseInt(data["svr"], 10, 16)                      /* — Номер вашего сервера */
	itemid, _ := strconv.Atoi(data["itemid"])                            /* — Номер предмета (подарка) */
	cnt, _ := strconv.Atoi(data["cnt"])                                  /* — Количество */
	aperiod, _ := strconv.Atoi(data["aperiod"])                          /* — Доступный период (сколько будет лежать в подароках) */
	pperiod, _ := strconv.Atoi(data["pperiod"])                          /* — Практический период (количество времени которое будет у предмета после получения)*/
	binding, _ := strconv.ParseUint(data["binding"], 10, 8)              /* — Под замком предмет или нет (Нет = 0 | Да = 1) */
	limitedDate, _ := time.Parse("2006-01-02 15:04:05", data["limitedDate"]) /* — Ограниченная дата */
	status, _ := strconv.ParseUint(data["status"], 10, 8)                /* — Статус предмета */

	// TODO: сделать валидацию 

	gift := billing.SysOrderList{                                  
		MItemID: itemid,
		MSvrNo: int16(svr),
		MCnt: cnt,
		MAvailablePeriod: aperiod,
		MPracticalPeriod: pperiod,  
		MBindingType: uint8(binding),
		MLimitedDate: limitedDate,
		MItemStatus: uint8(status),
	}
	
	err := b.billing.SetSysOrderListAll(ctx, gift)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(fiber.Map{
		"message": "gifts sent",
	})
}

// выдать подарок персонажу
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

func (b *ServiceBillingAPI) Shop(ctx *fiber.Ctx) error {
	const (
		op = "restapi.billing.Shop"
		empty = "empty"
	)

	res, err := b.billing.Shop(ctx) 
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (b *ServiceBillingAPI) DeleteShop(ctx *fiber.Ctx) error {
	const (
		op = "restapi.billing.DeleteShop"
		empty = "empty"
	)

	goldItemID := ctx.QueryInt("GoldItemID")
	if goldItemID == 0 {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := b.billing.DeleteShop(ctx, goldItemID) 
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}