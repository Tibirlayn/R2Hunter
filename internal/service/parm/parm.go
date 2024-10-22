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
	parmMonsterProvider ParmMonsterProvider
	parmProvider ParmProvider
	auth             *auth.Auth
	tokenTTL         time.Duration
}

type ParmItemProvider interface {
	BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error)
	ItemDDrop(ctx *fiber.Ctx, id int) (parm.DropItem, error)
	UpdateItemDDrop(ctx *fiber.Ctx, name parm.DropItem) (parm.DropItem, error)

	ItemsRess(ctx *fiber.Ctx, name string) (qParm.ItemRes, error)
	ItemsRessbyID(ctx *fiber.Ctx, id []int) (qParm.ItemRes, error)
	ItemSearch(ctx *fiber.Ctx, name string) (qParm.ItemSearch, error)
}

type ParmMonsterProvider interface {
	Monster(ctx *fiber.Ctx, name string) ([]parm.Moster, error)
}

type ParmProvider interface {
	ParmSvr(ctx *fiber.Ctx) ([]parm.ParmSvr, error)
	ParmSvrOp(ctx *fiber.Ctx, worldNo []int16) ([]parm.ParmSvrOp, error)
	UpdateParmSvrOp(ctx *fiber.Ctx, svrOp parm.ParmSvrOp) (string, error)

	MaterialDraw(ctx *fiber.Ctx, name string) (qParm.MaterialDraw, error)
	ClearMaterialDraw(ctx *fiber.Ctx, id int) (string, error)
	UpdateMaterialDrawIndex(ctx *fiber.Ctx, mdi parm.MaterialDrawIndex) (string, error)
	
	UpdateMaterialDrawResult(ctx *fiber.Ctx, mdi parm.MaterialDrawResult) (string, error)
	DeleteMaterialDrawResult(ctx *fiber.Ctx, seq int, mdrd int64, id int) (string, error)
	SetMaterialDrawResult(ctx *fiber.Ctx, mdr parm.MaterialDrawResult) (parm.MaterialDrawResult, error)

	QuestReward(ctx *fiber.Ctx, pageNumber int, limitCnt int) ([]qParm.QuestRewardRes, error)
}

func New(log *slog.Logger, parmItemProvider ParmItemProvider, parmMonsterProvider ParmMonsterProvider, parmProvider ParmProvider, auth *auth.Auth, tokenTTL time.Duration) *Parm {
	return &Parm{
		log:              log,
		parmItemProvider: parmItemProvider,
		parmMonsterProvider: parmMonsterProvider,
		parmProvider: parmProvider,
		auth:             auth,
		tokenTTL:         tokenTTL,
	}
}

func (i *Parm) BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error) {
	const op = "service.parm.BossDrop"

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
	const op = "service.parm.ItemDDrop"

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
	const op = "service.parm.UpdateItemDDrop"

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
	const op = "service.parm.ItemsRess"

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
	const op = "service.parm.ItemsRess"

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

func (i *Parm) Monster(ctx *fiber.Ctx, name string) ([]parm.Moster, error) {
	const op = "service.parm.Monster"

	_, err := i.auth.ValidJWT(ctx, op)
	if err != nil {
		return []parm.Moster{}, err
	}

	res, err := i.parmMonsterProvider.Monster(ctx, name) 
	if err != nil {
		return []parm.Moster{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (p *Parm) ParmSvr(ctx *fiber.Ctx) ([]parm.ParmSvr, error) {
	const op = "service.parm.ParmSvr"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := p.parmProvider.ParmSvr(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (p *Parm) ItemSearch(ctx *fiber.Ctx, name string) (qParm.ItemSearch, error) {
	const op = "service.parm.ItemSearch"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := p.parmItemProvider.ItemSearch(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (p *Parm) ParmSvrOp(ctx *fiber.Ctx, worldNo []int16) ([]parm.ParmSvrOp, error) {
	const op = "service.parm.ParmSvrOp"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := p.parmProvider.ParmSvrOp(ctx, worldNo)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil 
}

func (p *Parm) UpdateParmSvrOp(ctx *fiber.Ctx, svrOp parm.ParmSvrOp) (string, error) {
	const op = "service.parm.SetParmSvrOp"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return "", err
	}

	res, err := p.parmProvider.UpdateParmSvrOp(ctx, svrOp)
	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return res, nil 
}

func (p *Parm) MaterialDraw(ctx *fiber.Ctx, name string) (qParm.MaterialDraw, error) {
	const op = "service.parm.MaterialDraw"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := p.parmProvider.MaterialDraw(ctx, name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Parm) ClearMaterialDraw(ctx *fiber.Ctx, id int) (string, error) {
	const op = "service.parm.ClearMaterialDraw"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return "", err
	}

	res, err := p.parmProvider.ClearMaterialDraw(ctx, id)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *Parm) UpdateMaterialDrawIndex(ctx *fiber.Ctx, mdi parm.MaterialDrawIndex) (string, error) {
	const op = "service.parm.UpdateMaterialDrawIndex"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return "", err
	}

	res, err := p.parmProvider.UpdateMaterialDrawIndex(ctx, mdi)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *Parm) UpdateMaterialDrawResult(ctx *fiber.Ctx, mdi parm.MaterialDrawResult) (string, error) {
	const op = "service.parm.UpdateMaterialDrawResult"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return "", err
	}

	res, err := p.parmProvider.UpdateMaterialDrawResult(ctx, mdi)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *Parm) DeleteMaterialDrawResult(ctx *fiber.Ctx, seq int, mdrd int64, id int) (string, error) {
	const op = "service.parm.DeleteMaterialDrawResult"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return "", err
	}

	res, err := p.parmProvider.DeleteMaterialDrawResult(ctx, seq, mdrd, id)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *Parm) SetMaterialDrawResult(ctx *fiber.Ctx, mdr parm.MaterialDrawResult) (parm.MaterialDrawResult, error) {
	const op = "service.parm.SetMaterialDrawResult"

	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return parm.MaterialDrawResult{}, err
	}

	res, err := p.parmProvider.SetMaterialDrawResult(ctx, mdr)
	if err != nil {
		return parm.MaterialDrawResult{}, err
	}

	return res, nil
}

func (p *Parm) QuestReward(ctx *fiber.Ctx, pageNumber int, limitCnt int) ([]qParm.QuestRewardRes, error) {
	const op = "service.parm.QuestReward"
	
	_, err := p.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := p.parmProvider.QuestReward(ctx, pageNumber, limitCnt)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}