package parm

import (
	"fmt"
	"strconv"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
	qParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm"
	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	routersParm "github.com/Tibirlayn/R2Hunter/internal/routers/parm"
	"github.com/Tibirlayn/R2Hunter/pkg/lib/check"
	"github.com/gofiber/fiber/v2"
)

type Parm interface {
	BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error)
	ItemDDrop(ctx *fiber.Ctx, id int) (parm.DropItem, error)
	UpdateItemDDrop(ctx *fiber.Ctx, name parm.DropItem) (parm.DropItem, error)

	ItemsRess(ctx *fiber.Ctx, name string) (qParm.ItemRes, error)
	ItemSearch(ctx *fiber.Ctx, name string) (qParm.ItemSearch, error)
	Monster(ctx *fiber.Ctx, name string) ([]parm.Moster, error)

	ParmSvr(ctx *fiber.Ctx) ([]parm.ParmSvr, error)
	ParmSvrOp(ctx *fiber.Ctx, worldNo []int16) ([]parm.ParmSvrOp, error)
	UpdateParmSvrOp(ctx *fiber.Ctx, svrOp parm.ParmSvrOp) (string, error)

	MaterialDraw(ctx *fiber.Ctx, name string) (qParm.MaterialDraw, error)
	UpdateMaterialDrawResult(ctx *fiber.Ctx, mdi parm.MaterialDrawResult) (string, error)
	ClearMaterialDraw(ctx *fiber.Ctx, id int) (string, error)
	
	UpdateMaterialDrawIndex(ctx *fiber.Ctx, mdi parm.MaterialDrawIndex) (string, error)
	SetMaterialDrawIndex(ctx *fiber.Ctx, mdi parm.MaterialDrawIndex, mdm parm.MaterialDrawMaterial) (string, error)

	DeleteMaterialDrawResult(ctx *fiber.Ctx, seq int, mdrd int64, id int) (string, error)
	SetMaterialDrawResult(ctx *fiber.Ctx, mdr parm.MaterialDrawResult) (parm.MaterialDrawResult, error)

	QuestReward(ctx *fiber.Ctx, pageNumber int, limitCnt int) ([]qParm.QuestRewardRes, error)
	SetQuestReward(ctx *fiber.Ctx, qr parm.QuestReward) (string, error)
	// UpdateQuestReward(ctx *fiber.Ctx) (string, error)
	DeleteQuestReward(ctx *fiber.Ctx, qr parm.QuestReward) (string, error)
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

func (i *ServiceParmAPI) Monster(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.Monster"
		empty = "empty"
	)

	var name = ctx.Query("name")

	if name == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := i.parm.Monster(ctx, name)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) ParmSvr(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.Monster"
	)

	res, err := p.parm.ParmSvr(ctx)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) ItemSearch(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.ItemSearch"
		empty = "empty"
	)

	name := ctx.Query("name")
	if name == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := p.parm.ItemSearch(ctx, name)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) ParmSvrOp(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.ItemSearch"
	)

	// TODO: Сделать чтоб бодгружал с конфига, смотрит есть ли в конфиге записи 
	// если запись в конфиги не пуста берет мир с конфига

	worlds, err := p.parm.ParmSvr(ctx)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	var wn []int16
	for _, world := range worlds {
		if world.MWorldNo != 0 {
			wn = append(wn, world.MWorldNo)
		}
	}

	res, err := p.parm.ParmSvrOp(ctx, wn)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) UpdateParmSvrOp(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.UpdateParmSvrOp"
		empty = "empty"
	)

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	svrNo, _ := strconv.Atoi(data["mSvrNo"])
	opNo, _ := strconv.Atoi(data["mOpNo"])
	isSetup, _ := strconv.ParseBool(data["mIsSetup"]) 
	opValue1, _ := strconv.ParseFloat(data["mOpValue1"], 64)   
	opValue2, _ := strconv.ParseFloat(data["mOpValue2"], 64)   
	opValue3, _ := strconv.ParseFloat(data["mOpValue3"], 64)   
	opValue4, _ := strconv.ParseFloat(data["mOpValue4"], 64)   
	opValue5, _ := strconv.ParseFloat(data["mOpValue5"], 64)   
	opValue6, _ := strconv.ParseFloat(data["mOpValue6"], 64)   
	opValue7, _ := strconv.ParseFloat(data["mOpValue7"], 64)   
	opValue8, _ := strconv.ParseFloat(data["mOpValue8"], 64)   
	opValue9, _ := strconv.ParseFloat(data["mOpValue9"], 64)   
	opValue10, _ := strconv.ParseFloat(data["mOpValue10"], 64)
	opValue11, _ := strconv.ParseFloat(data["mOpValue11"], 64)
	opValue12, _ := strconv.ParseFloat(data["mOpValue12"], 64)
	opValue13, _ := strconv.ParseFloat(data["mOpValue13"], 64)
	opValue14, _ := strconv.ParseFloat(data["mOpValue14"], 64)
	opValue15, _ := strconv.ParseFloat(data["mOpValue15"], 64)
	opValue16, _ := strconv.ParseFloat(data["mOpValue16"], 64)
	opValue17, _ := strconv.ParseFloat(data["mOpValue17"], 64)
	opValue18, _ := strconv.ParseFloat(data["mOpValue18"], 64)
	opValue19, _ := strconv.ParseFloat(data["mOpValue19"], 64)
	opValue20, _ := strconv.ParseFloat(data["mOpValue20"], 64)
	opValue21, _ := strconv.ParseFloat(data["mOpValue21"], 64)
	opValue22, _ := strconv.ParseFloat(data["mOpValue22"], 64)
	opValue23, _ := strconv.ParseFloat(data["mOpValue23"], 64)
	opValue24, _ := strconv.ParseFloat(data["mOpValue24"], 64)
	opValue25, _ := strconv.ParseFloat(data["mOpValue25"], 64)
	opValue26, _ := strconv.ParseFloat(data["mOpValue26"], 64)
	opValue27, _ := strconv.ParseFloat(data["mOpValue27"], 64)
	opValue28, _ := strconv.ParseFloat(data["mOpValue28"], 64)
	opValue29, _ := strconv.ParseFloat(data["mOpValue29"], 64)
	opValue30, _ := strconv.ParseFloat(data["mOpValue30"], 64)

	svrOp := parm.ParmSvrOp{
		MSvrNo     : int16(svrNo),
		MOpNo      : int(opNo),
		MIsSetup   : isSetup,
		MOpValue1  : opValue1,
		MOpValue2  : opValue2,
		MOpValue3  : opValue3,
		MOpValue4  : opValue4,
		MOpValue5  : opValue5,
		MOpValue6  : opValue6,
		MOpValue7  : opValue7,
		MOpValue8  : opValue8,
		MOpValue9  : opValue9,
		MOpValue10 : opValue10,
		MOpValue11 : opValue11,
		MOpValue12 : opValue12,
		MOpValue13 : opValue13,
		MOpValue14 : opValue14,
		MOpValue15 : opValue15,
		MOpValue16 : opValue16,
		MOpValue17 : opValue17,
		MOpValue18 : opValue18,
		MOpValue19 : opValue19,
		MOpValue20 : opValue20,
		MOpValue21 : opValue21,
		MOpValue22 : opValue22,
		MOpValue23 : opValue23,
		MOpValue24 : opValue24,
		MOpValue25 : opValue25,
		MOpValue26 : opValue26,
		MOpValue27 : opValue27,
		MOpValue28 : opValue28,
		MOpValue29 : opValue29,
		MOpValue30 : opValue30,
	}

	res, err := p.parm.UpdateParmSvrOp(ctx, svrOp)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(fiber.Map{
		"response: ": res,
	})
}

func (p *ServiceParmAPI) MaterialDraw(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.UpdateParmSvrOp"
		empty = "empty"
	)

	nameQuery := ctx.Query("name")
	if nameQuery == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := p.parm.MaterialDraw(ctx, nameQuery)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) ClearMaterialDraw(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.ClearMaterialDraw"
		empty = "empty"
	)

	id := ctx.QueryInt("id")
	if id == 0 {
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := p.parm.ClearMaterialDraw(ctx, id)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) UpdateMaterialDrawIndex(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.UpdateMaterialDrawIndex"
		empty = "empty"
	)

	mdi := parm.MaterialDrawIndex{}

	if err := ctx.BodyParser(&mdi); err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	res, err := p.parm.UpdateMaterialDrawIndex(ctx, mdi)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) UpdateMaterialDrawResult(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.UpdateMaterialDrawResult"
		empty = "empty"
	)

/* 	MSeq := ctx.QueryInt("MSeq")
	MDRD := ctx.QueryInt("MDRD")
	IID := ctx.QueryInt("IID")
	MPerOrRate := ctx.QueryFloat("MPerOrRate")
	MItemStatus:= ctx.QueryInt("MItemStatus")
	MCnt := ctx.QueryInt("MCnt")
	MBinding := ctx.QueryInt("MBinding")
	MEffTime := ctx.QueryInt("MEffTime")
	MValTime := ctx.QueryInt("MValTime")
	MResource := ctx.QueryInt("MResource")
	MAddGroup := ctx.QueryInt("MAddGroup")

	mdi := parm.MaterialDrawResult{
		MSeq        : MSeq,
		MDRD        : int64(MDRD),
		IID         : IID,
		MPerOrRate  : MPerOrRate,
		MItemStatus : int8(MItemStatus),
		MCnt        : MCnt ,
		MBinding    : MBinding ,
		MEffTime    : MEffTime ,
		MValTime    : int16(MValTime) ,
		MResource   : MResource,
		MAddGroup   : int8(MAddGroup),
	} */

	mdr := parm.MaterialDrawResult{}

	if err := ctx.BodyParser(&mdr); err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	res, err := p.parm.UpdateMaterialDrawResult(ctx, mdr)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) DeleteMaterialDrawResult(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.DeleteMaterialDrawResult"
		empty = "empty"
	)

	seq := ctx.QueryInt("seq")
	mdrd := int64(ctx.QueryInt("mdrd"))
	id := ctx.QueryInt("id")
	
	if seq == 0 || mdrd == 0 || id == 0{
		return fmt.Errorf("%s, %s", op, empty)
	}

	res, err := p.parm.DeleteMaterialDrawResult(ctx, seq, mdrd, id)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) SetMaterialDrawResult(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.SetMaterialDrawResult"
		empty = "empty"
	)

	mdr := parm.MaterialDrawResult{}

	if err := ctx.BodyParser(&mdr); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}	

	res, err := p.parm.SetMaterialDrawResult(ctx, mdr)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) QuestReward(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.QuestReward"
		empty = "empty"
	)

	pageNumber, limitCnt, err := check.CheckPageAndLimit(ctx)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	res, err := p.parm.QuestReward(ctx, pageNumber, limitCnt)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) SetQuestReward(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.SetQuestReward"
		empty = "empty"
	) 

	qr := parm.QuestReward{}

	if err := ctx.BodyParser(&qr); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	res, err := p.parm.SetQuestReward(ctx, qr)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) DeleteQuestReward(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.DeleteQuestReward"
		empty = "empty"
	) 

	qr := parm.QuestReward{}

	// if err := ctx.BodyParser(&qr); err != nil {
	// 	return fmt.Errorf("%s, %w", op, err)
	// }

	qr.MRewardNo = ctx.QueryInt("mRewardNo")
	qr.MExp = int64(ctx.QueryInt("MExp"))
	qr.MID = ctx.QueryInt("mID")     
	qr.MCnt = ctx.QueryInt("mCnt")    
	qr.MBinding = int8(ctx.QueryInt("mBinding"))
	qr.MStatus = int8(ctx.QueryInt("mStatus") )
	qr.MEffTime = ctx.QueryInt("mEffTime")
	qr.MValTime = ctx.QueryInt("mValTime")

	res, err := p.parm.DeleteQuestReward(ctx, qr)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}

func (p *ServiceParmAPI) UpdateQuestReward(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.UpdateQuestReward"
		empty = "empty"
	) 

	

	return ctx.JSON(nil)
}

func (p *ServiceParmAPI) SetMaterialDrawIndex(ctx *fiber.Ctx) error {
	const (
		op = "service.parm.SetMaterialDrawIndex"
		empty = "empty"
	)

	type requestBody struct {
		MDI parm.MaterialDrawIndex    `json:"mdi"`
		MDM parm.MaterialDrawMaterial `json:"mdm"`
	}

	body := requestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	res, err := p.parm.SetMaterialDrawIndex(ctx, body.MDI, body.MDM)
	if err != nil {
		return fmt.Errorf("%s, %s", op, err)
	}

	return ctx.JSON(res)
}