package mssql

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/storage"
	query "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
	game "github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
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

 


/* 	Pc          []game.Pc
	PcState     []game.PcState
	PcInventory []game.PcInventory
	PcStore     []game.PcStore */

/* 	result := g.db.Table("TblPc AS pc").
	Select("pc.*, pcState.*, inventory.*").
	Joins("INNER JOIN TblPcState AS pcState ON pc.mNo = pcState.mNo").
	Joins("INNER JOIN TblPcInventory AS inventory ON pc.mNo = inventory.mPcNo").
	Where("pc.mNm = ? OR pc.mOwner = ?", name, pcID).
	Find(&pc) 

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return pc, fmt.Errorf("%s, %w", op, storage.ErrUserNotFound)
		}
		return pc, fmt.Errorf("%s, %w", op, result.Error)
	} */


/*  	if resultPc := g.db.Where("mNm = ? OR mOwner = ?", name, pcID).Find(&pc.Pc); resultPc.Error != nil {
		if errors.Is(resultPc.Error, gorm.ErrRecordNotFound) {
			return pc, fmt.Errorf("%s, %w", op, storage.ErrUserNotFound)
		} else {
			return pc, fmt.Errorf("%s, %w", op, resultPc.Error)
		}
	}

	if resultPcState := g.db.Where("mNo = ?", pc.Pc.MNo).Find(&pc.PcState); resultPcState.Error != nil {
		if errors.Is(resultPcState.Error, gorm.ErrRecordNotFound) {
			g.log.Info("%s, %s", op, storage.ErrNotFound)
		} else {
			g.log.Info("%s, %v", op, resultPcState.Error)
		}
	}

	if resultPcInventory := g.db.Where("mPcNo = ?", pc.Pc.MNo).Find(&pc.PcInv); resultPcInventory.Error != nil {
		if errors.Is(resultPcInventory.Error, gorm.ErrRecordNotFound) {
			g.log.Info("%s, %s", op, storage.ErrNotFound)
		} else {
			g.log.Info("%s, %v", op, resultPcInventory.Error)
		}
	}

	if resultPcStore := g.db.Where("mUserNo = ?", pc.Pc.MOwner).Find(&pc.PcStore); resultPcStore.Error != nil {
		if errors.Is(resultPcStore.Error, gorm.ErrRecordNotFound) {
			g.log.Info("%s, %s", op, storage.ErrNotFound)
		} else {
			g.log.Info("%s, %v", op, resultPcStore.Error)
		}
	}  */
