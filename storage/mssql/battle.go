package mssql

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type BattleStorage struct {
	db *gorm.DB
}

func NewBattleStorage(cfg_db *config.ConfigDB) (*BattleStorage, error) {
	const op = "storage.mssql.battle.New"
		parm := cfg_db.Battle

		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		parm.User, parm.Password, parm.Server, parm.Port, parm.NameDB)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	return &BattleStorage{db: db}, nil
}