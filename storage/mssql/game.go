package mssql

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
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

func (g *GameStorage) PcCard(ctx *fiber.Ctx, name string) (game.PcParm, error) {
	const op = "storage.mssql.game.PcCard" 

	var pc game.PcParm

	if resultPc := g.db.Where("mNm = ?", pc.Pc.MNm).First(&pc.Pc); resultPc.Error != nil {
		if errors.Is(resultPc.Error, gorm.ErrRecordNotFound) {
			return pc, fmt.Errorf("%s, %w", op, storage.ErrAppNotFound)
		} else {
			return pc, fmt.Errorf("%s, %w", op, resultPc.Error)
		}
	}

	if resultPcState := g.db.Where("mNo = ?", pc.Pc.MNo).Find(&pc.PcState); resultPcState.Error != nil {
		if errors.Is(resultPcState.Error, gorm.ErrRecordNotFound) {
			g.log.Info("%s, %s", op, storage.ErrAppNotFound)
		} else {
			g.log.Info("%s, %v", op, resultPcState.Error)
		}
	}

	if resultPcInventory := g.db.Where("mPcNo = ?", pc.Pc.MNo).Find(&pc.PcInv); resultPcInventory.Error != nil {
		if errors.Is(resultPcInventory.Error, gorm.ErrRecordNotFound) {
			g.log.Info("%s, %s", op, storage.ErrAppNotFound)
		} else {
			g.log.Info("%s, %v", op, resultPcInventory.Error)
		}
	}

	if resultPcStore := g.db.Where("mUserNo = ?", pc.Pc.MOwner).Find(&pc.PcStore); resultPcStore.Error != nil {
		if errors.Is(resultPcStore.Error, gorm.ErrRecordNotFound) {
			g.log.Info("%s, %s", op, storage.ErrAppNotFound)
		} else {
			g.log.Info("%s, %v", op, resultPcStore.Error)
		}
	}

	return pc, nil

 }