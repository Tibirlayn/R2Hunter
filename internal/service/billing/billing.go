package billing

import (
	"fmt"
	"log/slog"
	"time"

	qBilling "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/billing"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/billing"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/member"
	"github.com/Tibirlayn/R2Hunter/internal/service/parm"
	"github.com/gofiber/fiber/v2"
)

type Billing struct {
	log             *slog.Logger
	billingProvider BillingProvider
	member          *member.Member
	parm            *parm.Parm
	auth            *auth.Auth
	tokenTTL        time.Duration
}

type BillingProvider interface {
	//посмотреть все подарки
	SysOrderList(ctx *fiber.Ctx) ([]billing.SysOrderList, error)
	//посмотреть в определенном аккаунте
	SysOrderListEmail(ctx *fiber.Ctx, id int) ([]billing.SysOrderList, error)
	//выдать подарок всем персонажам
	SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error)
}

func New(log *slog.Logger, billingProvider BillingProvider, member *member.Member, parm *parm.Parm, auth *auth.Auth, tokenTTL time.Duration) *Billing {
	return &Billing{
		log:             log,
		billingProvider: billingProvider,
		member:          member,
		parm:            parm,
		auth:            auth,
		tokenTTL:        tokenTTL,
	}
}

func (b *Billing) SysOrderList(ctx *fiber.Ctx) ([]billing.SysOrderList, error) {
	const op = "service.billing.SysOrderList"

	_, err := b.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := b.billingProvider.SysOrderList(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (b *Billing) SysOrderListEmail(ctx *fiber.Ctx, email string) ([]qBilling.SolItemRes, error) {
	const op = "service.billing.SysOrderListEmail"
	var ids []int
	var solItems []qBilling.SolItemRes

	_, err := b.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	member, err := b.member.MemberBil(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	res, err := b.billingProvider.SysOrderListEmail(ctx, member.MUserNo)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	for _, n := range res{
		ids = append(ids, n.MItemID)
	}

	ires, err := b.parm.ItemsRessbyID(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	for _, order := range res {
		var solItem qBilling.SolItemRes
		solItem.MRegDate = order.MRegDate 
		solItem.MSysOrderID = order.MSysOrderID 
		solItem.MSysID = order.MSysID 
		solItem.MUserNo = order.MUserNo 
		solItem.MSvrNo = order.MSvrNo 
		solItem.MItemID = order.MItemID 
		solItem.MCnt = order.MCnt 
		solItem.MAvailablePeriod = order.MAvailablePeriod
		solItem.MPracticalPeriod = order.MPracticalPeriod
		solItem.MStatus = order.MStatus 
		solItem.MReceiptDate = order.MReceiptDate 
		solItem.MReceiptPcNo = order.MReceiptPcNo 
		solItem.MRecepitPcNm = order.MRecepitPcNm 
		solItem.MBindingType = order.MBindingType 
		solItem.MLimitedDate = order.MLimitedDate 
		solItem.MItemStatus = order.MItemStatus 

		for _, item := range ires {
			if item.IID == order.MItemID {
				solItem.IID        = item.IID      
				solItem.IName      = item.IName    
				solItem.RFileName  = item.RFileName
				solItem.RPosX      = item.RPosX    
				solItem.RPosY      = item.RPosY    
			}
		}

		solItems = append(solItems, solItem)
	}

	return solItems, nil
}

func (b *Billing) SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error) {

	return billing.SysOrderList{}, nil
}
