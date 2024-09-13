package mssql

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	game "github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	query "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
	queryGame "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
	"github.com/Tibirlayn/R2Hunter/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type GameStorage struct {
	db *gorm.DB
	log *slog.Logger
}

func NewGameStorage(cfg_db *config.ConfigDB) (*GameStorage, error) {
	const op = "storage.mssql.game.New"
	parm := cfg_db.Game

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		parm.User, parm.Password, parm.Server, parm.Port, parm.NameDB)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GameStorage{db: db}, nil
}

func (s *GameStorage) Stop() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

// TODO: не используется
func (g *GameStorage) PcCard(ctx *fiber.Ctx, name string, pcID int64) ([]query.PcParm, error) {
	const op = "storage.mssql.game.PcCard" 

	var pcs []game.Pc
	if result := g.db.Preload("PcInventories").
	Where("pc.mNm = ? OR pc.mOwner = ?", name, pcID).
	Find(&pcs); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s, %w", op, storage.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s, %w", op, result.Error)
	}
	
	var pcParms []query.PcParm
    for _, pc := range pcs {
        pcParm := query.PcParm{
			Pc: game.Pc{
				MRegDate: pc.MRegDate,
				MOwner:   pc.MOwner,
				MSlot:    pc.MSlot,
				MNo:      pc.MNo,
				MNm:      pc.MNm,
				MClass:   pc.MClass,
				MSex:     pc.MSex,
				MHead:    pc.MHead,
				MFace:    pc.MFace,
				MBody:    pc.MBody,
				MHomeMapNo: pc.MHomeMapNo,
				MHomePosX:  pc.MHomePosX,
				MHomePosY:  pc.MHomePosY,
				MHomePosZ:  pc.MHomePosZ,
				MDelDate:   pc.MDelDate,
			},
        }
        pcParms = append(pcParms, pcParm)
    }

	return pcParms, nil
 }

func (g *GameStorage) PcTopLVL(ctx *fiber.Ctx) ([]queryGame.PcTopLVL, error) {
	const op = "storage.mssql.game.PcTopLVL" 

	var pcTopLVL []queryGame.PcTopLVL

	if err := g.db.Table("TblPc AS a").
		Select(`TOP 100 a.mNo AS ID, 
				CASE
					WHEN a.mClass = 0 THEN 'Рыцарь'
					WHEN a.mClass = 1 THEN 'Рейджер'
					WHEN a.mClass = 2 THEN 'Маг'
					WHEN a.mClass = 3 THEN 'Ассасин'
					WHEN a.mClass = 4 THEN 'Призыватель'
					ELSE 'неизвестный класс'
				END AS Class, 
				RTRIM(a.mNm) AS Name,
				b.mLevel AS Level, 
				b.mChaotic AS Chaotic, 
				b.mPkCnt AS PkCnt`).
		Joins("JOIN TblPcState AS b ON a.mNo = b.mNo").
		Where("a.mNo > ? AND LEFT(a.mNm, 1) <> ?", 0, ",").
		Order("b.mLevel DESC").
		Scan(&pcTopLVL).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return pcTopLVL, nil
}

func (g *GameStorage) PcTopByGold(ctx *fiber.Ctx) ([]queryGame.PcTopByGold, error) {
	const op = "storage.mssql.game.PcTopByGold" 

	var pcTopByGold []queryGame.PcTopByGold

	if err := g.db.Table("TblPc AS a").
	Select("TOP 100 a.mOwner AS MOwner, b.mSerialNo AS MSerialNo, RTRIM(a.mNm) AS Name, b.mPcNo AS MPcNo, b.mItemNo AS MItemNo, b.mCnt AS MCnt").
	Joins("INNER JOIN TblPcInventory AS b ON b.mPcNo = a.mNo").
	Where("b.mItemNo = ? AND LEFT (a.mNm, 1) <> ?", 409, ",").
	Order("b.mCnt DESC").
	Scan(&pcTopByGold).Error; err != nil {
	return nil, fmt.Errorf("%s, %w", op, err)
}

	return pcTopByGold, nil
}