package mssql

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/billing"
	"github.com/Tibirlayn/R2Hunter/storage"
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

	// надо сделать массив 
	// получать например массив пользователей и выдать список 

	if err := b.db.Where("mUserNo = ?", id).Find(&bil).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return bil, nil
}

// Добавить подарок одному пользователю
func (b *BillingStorage) SetSysOrderList(ctx *fiber.Ctx, gift billing.SysOrderList) (billing.SysOrderList, error) {
	const op = "storage.mssql.billing.SetSysOrderList"

	tx := b.db.Begin()
	if err := tx.Omit("mRegDate", "mReceiptDate", "mReceiptPcNo", "mRecepitPcNm").Create(&gift).Error; err != nil {
		tx.Rollback()
		return billing.SysOrderList{}, fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return billing.SysOrderList{}, fmt.Errorf("%s, %w", op, storage.ErrCommit)
	}
	
	return gift, nil
}

// SELECT mUserNo, mLoginTm FROM [dbo].[TblUser] Where mLoginTm >= DATEADD(DAY, -30, GETDATE());
// Добавить подарок всем пользователям last30Days
func (b *BillingStorage) SetSysOrderListAll(ctx *fiber.Ctx, gift billing.SysOrderList, userNo []int) error {
	const op = "storage.mssql.billing.SetSysOrderListAll"

	tx := b.db.Begin()
	// Проходим по каждому пользователю в массиве userNo
	for _, UserID := range userNo {
		// Для каждого пользователя создаем новую запись подарка
		gift.MUserNo = UserID // Заменяем MSysID на текущий userID

		// Сохраняем подарок в базу данных
		if err := tx.Omit("mRegDate", "mReceiptDate", "mReceiptPcNo", "mRecepitPcNm", "MSysOrderID").Create(&gift).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s, %w", op, err)
		}
	}

	// После того как все записи добавлены, выполняем commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("%s, %w", op, storage.ErrCommit)
	}
	
	return nil
}

func (b *BillingStorage) Shop(ctx *fiber.Ctx) ([]billing.GoldItem, error) {
	const op = "storage.mssql.billing.Shop"

	gi := []billing.GoldItem{}

	if err := b.db.Find(&gi).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return gi, nil
}

func (b *BillingStorage) DeleteShop(ctx *fiber.Ctx, goldItemID int) (string, error) {
	const op = "storage.mssql.billing.DeleteShop"

	tx := b.db.Begin()
	ca := billing.CategoryAssign{}
	gi := billing.GoldItem{}
	giss := billing.GoldItemSupportSvr{}

	resCA := tx.Where("GoldItemID = ?", goldItemID).First(&ca)
	resGI := tx.Where("GoldItemID = ?", goldItemID).First(&gi)
	resGISS := tx.Where("GoldItemID = ?", goldItemID).First(&giss)

	if resCA.Error != nil {
		if resCA.RowsAffected == 0 {
			tx.Rollback() // Откат транзакции при ошибке
			return "", fmt.Errorf("%s: запись с условиями %d не найдена", op, goldItemID)
		}
		tx.Rollback() // Откат транзакции при ошибке
		return "", fmt.Errorf("%s, %w", op, resCA.Error)
	}

	if resGI.Error != nil {
		if resGI.RowsAffected == 0 {
			tx.Rollback() // Откат транзакции при ошибке
			return "", fmt.Errorf("%s: запись с условиями %d не найдена", op, goldItemID)
		}
		tx.Rollback() // Откат транзакции при ошибке
		return "", fmt.Errorf("%s, %w", op, resGI.Error)
	}

	if resGISS.Error != nil {
		if resGISS.RowsAffected == 0 {
			tx.Rollback() // Откат транзакции при ошибке
			return "", fmt.Errorf("%s: запись с условиями %d не найдена", op, goldItemID)
		}
		tx.Rollback() // Откат транзакции при ошибке
		return "", fmt.Errorf("%s, %w", op, resGISS.Error)
	}

	if err := tx.Where("GoldItemID = ?", goldItemID).Delete(&ca).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Where("GoldItemID = ?", goldItemID).Delete(&gi).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Where("GoldItemID = ?", goldItemID).Delete(&giss).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}
	
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data DeleteMaterialDrawResul", nil
}