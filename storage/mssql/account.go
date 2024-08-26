package mssql

import (
	"errors"
	"fmt"
	"strings"

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

	var userID string
	// TODO: убрать @ и . и записать как логин
	// причина есть две почты например artur1994@gmail.com и почта artur1994@mail.com в логин он запишет как artur1994 
	// не может быть у нас два логина с artur1994
	// по этому будет в ник добавлять например artur1994@gmail.com = artur1994gmailcom
	index := strings.Index(email, "@") 
	if index != -1 {
		userID = email[:index]
	} else {
		return 0, fmt.Errorf("%s, %w", op, storage.ErrValidEmail)
	}
 
	// Создаем нового пользователя
	member := account.Member{
		MUserId: userID,
		Email:   email,
		MUserPswd: password,
	}
	
	user := account.User{
		MUserId: userID,
		MUserPswd: password,
		MCertifiedKey: 2491,
	}

	tx := a.db.Begin() // Начинаем транзакцию
	result := tx.Where("email = ?", email).First(&member)
	if result.RowsAffected != 0 {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	// Создаем нового пользователя
	// Проверяем ошибки при создании пользователя
	if err := tx.Omit("superpwd", "cash", "tgzh", "uid", "klq", "ylq", "auth", "mSum", "isAdmin" , "isdl", "dlmoney", "registerIp", "country", "cashBack").
		Create(&member).Error; err != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrCreateUser)
	}

	if err := tx.Omit("mRegDate", "mUserAuth", "mUserNo", "mIp", "mLoginTm", "mLogoutTm", "mTotUseTm", "mWorldNo", "mDelDate", "mPcBangLv", "mSecKeyTableUse", "mUseMacro", "mIpEx", "mJoinCode", "mLoginChannelID", "mTired", "mChnSID", "mNewId", "mLoginSvrType", "mAccountGuid", "mNormalLimitTime", "mPcBangLimitTime", "mRegIp", "mIsMovingToBattleSvr").
		Create(&user).Error; err != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrCreateUser)
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrCreateUserCommit)
	}

	resid := int64(user.MUserNo)

	return resid, nil
}

func (a *AccountStorage) User(ctx *fiber.Ctx, email string) (account.Member, error) {
	const op = "storage.mssql.account.User"

	var user account.Member
	result := a.db.Where("email = ?", email).First(&user)
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