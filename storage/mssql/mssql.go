package mssql

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Storage struct {
	dbMap map[string]*gorm.DB
}

func New(cfg_db *config.ConfigDB) (*Storage, error) {
	const op = "storage.mssql.storage.New"

	// создаем массив где string (название бд), а *gorm.DB (Данные: User, Password и т.д.)
	dbMap := make(map[string]*gorm.DB)

	for name, dbConfigParm := range map[string]config.ConfigParm{
		"Account": cfg_db.Account,
		"Battle": cfg_db.Battle,
		"Billing": cfg_db.Billing,
		"Game": cfg_db.Game,
		"Logs": cfg_db.Logs,
		"Parm": cfg_db.Parm,
		"Statistics": cfg_db.Statistics,
	} {
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		dbConfigParm.User, dbConfigParm.Password, dbConfigParm.Server, dbConfigParm.Port, dbConfigParm.NameDB)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		dbMap[name] = db
	}

	return &Storage{dbMap: dbMap}, nil
}
