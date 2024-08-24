package mssql

import (
	"errors"
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type AccountStorage struct {
	db *gorm.DB
}

func NewAccountStorage(cfg_db *config.ConfigDB) (*AccountStorage, error) {
	const op = "storage.mssql.account.New"
		parm := cfg_db.Account

		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		parm.User, parm.Password, parm.Server, parm.Port, parm.NameDB)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

	return &AccountStorage{db: db}, nil
}

func (s *AccountStorage) Stop() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (a *AccountStorage) SaveUser(ctx *fiber.Ctx, email string, password string) (uid int64, err error) {
	const op = "storage.mssql.account.SaveUser"

	// Создаем нового пользователя
	newUser := account.Member{
		Email:   email,
		MUserPswd: password,
	}

	tx := a.db.Begin() // Начинаем транзакцию
	result := tx.Where("nphone = ?", email).Find(&newUser)
	if result.RowsAffected != 0 {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	// Создаем нового пользователя
	// Проверяем ошибки при создании пользователя
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrCreateUser)
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrCreateUserCommit)
	}

	// resid := newUser.MUserId

	return 1, nil
}

func (a *AccountStorage) User(ctx *fiber.Ctx, email string) (account.Member, error) {
	const op = "storage.mssql.account.User"

	var user account.Member
	result := a.db.Where("email = ?", email).Find(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	if result.Error != nil {
		return user, fmt.Errorf("%s: %w", op, result.Error)
	}

	return user, nil
}

func (a *AccountStorage) App(ctx *fiber.Ctx, appID int) (models.App, error) {
		const op = "storage.mssql.account.App"

	var app models.App

	result := a.db.First(&app, appID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return app, fmt.Errorf("%s, %w", op, storage.ErrAppNotFound)
	}

	if result.Error != nil {
		return app, fmt.Errorf("%s, %w", op, result.Error)
	}

	return app, nil
}