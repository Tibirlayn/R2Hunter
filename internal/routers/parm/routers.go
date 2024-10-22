package parm

import "github.com/gofiber/fiber/v2"

type ParmHandler interface {
	BossDrop(ctx *fiber.Ctx) error
	ItemDDrop(ctx *fiber.Ctx) error
	UpdateItemDDrop(ctx *fiber.Ctx) error
	ItemsRess(ctx *fiber.Ctx) error
	Monster(ctx *fiber.Ctx) error

	ParmSvr(ctx *fiber.Ctx) error
	ParmSvrOp(ctx *fiber.Ctx) error
	UpdateParmSvrOp(ctx *fiber.Ctx) error
	MaterialDraw(ctx *fiber.Ctx) error
	ClearMaterialDraw(ctx *fiber.Ctx) error
	UpdateMaterialDrawIndex(ctx *fiber.Ctx) error
	UpdateMaterialDrawResult(ctx *fiber.Ctx) error
	DeleteMaterialDrawResult(ctx *fiber.Ctx) error
	SetMaterialDrawResult(ctx *fiber.Ctx) error 

	ItemSearch(ctx *fiber.Ctx) error

	QuestReward(ctx *fiber.Ctx) error
	SetQuestReward(ctx *fiber.Ctx) error
	UpdateQuestReward(ctx *fiber.Ctx) error
}

func NewRoutersParm(appf *fiber.App, api ParmHandler) {
	appf.Get("boss-drop", api.BossDrop)
	appf.Get("item-ddrop", api.ItemDDrop)
	appf.Put("update-item-ddrop", api.UpdateItemDDrop)
	appf.Get("item", api.ItemsRess)
	appf.Get("item-search", api.ItemSearch)
	appf.Get("monster", api.Monster)
	appf.Get("parm-svr", api.ParmSvr)
	appf.Get("parm-svr-op", api.ParmSvrOp)
	appf.Put("update-parm-svr-op", api.UpdateParmSvrOp)
	appf.Get("material-draw", api.MaterialDraw)
	appf.Delete("clear-material-draw", api.ClearMaterialDraw)
	appf.Put("update-mdi", api.UpdateMaterialDrawIndex)
	appf.Put("update-mdr", api.UpdateMaterialDrawResult)
	appf.Delete("delete-mdrd", api.DeleteMaterialDrawResult)
	appf.Post("add-mdr", api.SetMaterialDrawResult)

	appf.Get("quest-reward", api.QuestReward)
	appf.Post("add-quest-reward", api.SetQuestReward)
	appf.Put("update-quest-reward", api.UpdateQuestReward)
}