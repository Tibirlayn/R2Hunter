package mssql

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type StatisticsStorage struct {
	db *gorm.DB
}

func NewStatisticsStorage(cfg_db *config.ConfigDB) (*StatisticsStorage, error) {
	const op = "storage.mssql.statistics.New"
		parm := cfg_db.Statistics

		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		parm.User, parm.Password, parm.Server, parm.Port, parm.NameDB)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	return &StatisticsStorage{db: db}, nil
}

func (s *StatisticsStorage) Stop() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}