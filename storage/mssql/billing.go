package mssql

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/billing"
	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/gofiber/fiber/v2"
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

func (b *BillingStorage) SysOrderList(ctx *fiber.Ctx) ([]billing.SysOrderList, error)  {
	const op = "storage.mssql.billing.SysOrderList"

	var bill []billing.SysOrderList
	if err := b.db.Find(&bill).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return bill, nil
}

func (b *BillingStorage) SysOrderListEmail(ctx *fiber.Ctx, id int) ([]billing.SysOrderList, error) {
	const op = "storage.mssql.billing.SysOrderListEmail"

	bil := []billing.SysOrderList{}

	if err := b.db.Where("mUserNo = ?", id).Find(&bil).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return bil, nil
}

func (b *BillingStorage) SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error) {
	const op = "storage.mssql.billing.SetSysOrderList"

	return billing.SysOrderList{}, nil
}