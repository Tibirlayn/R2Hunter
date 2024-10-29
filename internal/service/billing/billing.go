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
	// выдать подарок персонажу
	SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error)
	//выдать подарок всем персонажам
	SetSysOrderListAll(ctx *fiber.Ctx, gift billing.SysOrderList, userNo []int) error
	// Магазин
	Shop(ctx *fiber.Ctx) ([]billing.GoldItem, error)
	DeleteShop(ctx *fiber.Ctx, goldItemID int) (string, error)
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
	const op = "service.billing.SetSysOrderList"

	user, err := b.auth.ValidJWT(ctx, op)
	if err != nil {
		return billing.SysOrderList{}, err
	}

	gift.MSysID = user

	// TODO: Проверить на существование пользователя (нужно но мне лень)

	res, err := b.billingProvider.SetSysOrderList(ctx, gift)
	if err != nil {
		return gift, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (b *Billing) SetSysOrderListAll(ctx *fiber.Ctx, gift billing.SysOrderList) error {
	const op = "service.billing.SetSysOrderListAll"

	user, err := b.auth.ValidJWT(ctx, op)
	if err != nil {
		return err
	}

	gift.MSysID = user

	resAuth, err := b.member.UserLastLogin(ctx)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	if err := b.billingProvider.SetSysOrderListAll(ctx, gift, resAuth); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return nil 
}

func (b *Billing) Shop(ctx *fiber.Ctx) ([]qBilling.ShopItemRes, error) {
	const op = "service.billing.Shop"

	var ids []int
	var sirs []qBilling.ShopItemRes

	_, err := b.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	resBil, err := b.billingProvider.Shop(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	for _, n := range resBil{
		ids = append(ids, n.IID)
	}

	resItem, err := b.parm.ItemsRessbyID(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	for _, order := range resBil {
		var sir qBilling.ShopItemRes
		sir.GoldItemID        = order.GoldItemID       
		sir.GIID              = order.IID              
		sir.ItemName          = order.ItemName         
		sir.ItemImage         = order.ItemImage        
		sir.ItemDesc          = order.ItemDesc         
		sir.OriginalGoldPrice = order.OriginalGoldPrice
		sir.GoldPrice         = order.GoldPrice        
		sir.ItemCategory      = order.ItemCategory     
		sir.IsPackage         = order.IsPackage        
		sir.Status            = order.Status           
		sir.AvailablePeriod   = order.AvailablePeriod  
		sir.Count             = order.Count            
		sir.PracticalPeriod   = order.PracticalPeriod  
		sir.RegistDate        = order.RegistDate       
		sir.RegistAdmin       = order.RegistAdmin      
		sir.RegistIP          = order.RegistIP         
		sir.UpdateDate        = order.UpdateDate       
		sir.UpdateAdmin       = order.UpdateAdmin      
		sir.UpdateIP          = order.UpdateIP         
		sir.ItemNameRUS       = order.ItemNameRUS      
		sir.ItemDescRUS       = order.ItemDescRUS      

		for _, item := range resItem {
			if item.IID == order.IID {
				sir.IID               = item.IID      
				sir.IName             = item.IName    
				sir.RFileName         = item.RFileName
				sir.RPosX             = item.RPosX    
				sir.RPosY             = item.RPosY  
			}
		}

		sirs = append(sirs, sir)
	}

	return sirs, nil
}

func (b *Billing) DeleteShop(ctx *fiber.Ctx, goldItemID int) (string, error) {
	const op = "service.billing.DeleteShop"

	_, err := b.auth.ValidJWT(ctx, op)
	if err != nil {
		return "", err
	}

	res, err := b.billingProvider.DeleteShop(ctx, goldItemID)
	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}