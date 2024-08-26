package mssql

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type BillingStorage struct {
	db *gorm.DB
}

func NewBillingStorage(cfg_db *config.ConfigDB) (*BillingStorage, error) {
	const op = "storage.mssql.billing.New"
		parm := cfg_db.Billing

		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		parm.User, parm.Password, parm.Server, parm.Port, parm.NameDB)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	return &BillingStorage{db: db}, nil
}

func (s *BillingStorage) Stop() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}