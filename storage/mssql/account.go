package mssql

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
	"github.com/Tibirlayn/R2Hunter/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type AccountStorage struct {
	log *slog.Logger
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

func (a *AccountStorage) SaveUser(ctx *fiber.Ctx, member account.Member, appID int) (uid int64, err error) {
	const op = "storage.mssql.account.SaveUser"

	user := account.User{
		MUserId:       member.MUserId,
		MUserPswd:     member.MUserPswd,
		MCertifiedKey: appID,
	}

	tx := a.db.Begin() // Начинаем транзакцию
	resultMember := tx.Where("email = ? OR mUserId = ?", member.Email, member.MUserId).First(&member)
	if resultMember.RowsAffected != 0 {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	resultUser := tx.Where("mUserId = ?", user.MUserId).First(&user)
	if resultUser.RowsAffected != 0 {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	// Создаем нового пользователя
	// Проверяем ошибки при создании пользователя
	if err := tx.Omit("superpwd", "cash", "tgzh", "uid", "klq", "ylq", "auth", "mSum", "isAdmin", "isdl", "dlmoney", "registerIp", "country", "cashBack").
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

func (a *AccountStorage) User(ctx *fiber.Ctx, email string) (account.Member, account.User, error) {
	const op = "storage.mssql.account.User"

	var member account.Member
	resultMember := a.db.Where("email = ?", email).First(&member)

	var user account.User
	resultUser := a.db.Where("mUserId = ?", member.MUserId).First(&user)

	if errors.Is(resultMember.Error, gorm.ErrRecordNotFound) {
		return member, user, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	if resultMember.Error != nil {
		return member, user, fmt.Errorf("%s: %w", op, resultMember.Error)
	}

	if errors.Is(resultUser.Error, gorm.ErrRecordNotFound) {
		return member, user, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	if resultUser.Error != nil {
		return member, user, fmt.Errorf("%s: %w", op, resultUser.Error)
	}

	return member, user, nil
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

func (a *AccountStorage) Member(ctx *fiber.Ctx, name string) (mp query.MemberParm, err error) {
	const op = "storage.mssql.account.Member"

	var member account.Member
	var user account.User
	var userAdmin account.UserAdmin

	if resultMember := a.db.
	Where("email = ? OR mUserId = ?", name, name).
	Find(&member); resultMember.Error != nil {
		if errors.Is(resultMember.Error, gorm.ErrRecordNotFound) {
			return mp, fmt.Errorf("%s, %w", op, storage.ErrUserNotFound)
		} else {
			return mp, fmt.Errorf("%s, %w", op, resultMember.Error)
		}
	}

 	if resultUser := a.db.Where("mUserId = ?", member.MUserId).Find(&user); 
	resultUser.Error != nil {
		if errors.Is(resultUser.Error, gorm.ErrRecordNotFound) {
			return mp, fmt.Errorf("%s, %w", op, storage.ErrUserNotFound)
		} else {
			return mp, fmt.Errorf("%s, %w", op, resultUser.Error)
		}
	} 

	if resultUserAdmin := a.db.Where("mUserNo = ?", user.MUserNo).Find(&userAdmin); 
	resultUserAdmin.Error != nil {
		if errors.Is(resultUserAdmin.Error, gorm.ErrRecordNotFound) {
			a.log.Info("%s, %w", op, storage.ErrNotFound)
		} else {
			a.log.Info("%s, %w", op, resultUserAdmin.Error)
		}
	}

/* 	mp.Member = member
	mp.User = user
	mp.UserAdmin = userAdmin */

	return mp, nil
}

 /* func (a *AccountStorage) Member(ctx *fiber.Ctx, mp query.MemberParm) (query.MemberParm, error) {
	const op = "storage.mssql.account.Member"

	var errorsList []error

	// Первая часть: Поиск по Member
	resultMember := a.db.Table("Member m").
		Select("*").
		Joins("INNER JOIN TblUser u ON m.mUserId = u.mUserId").
		Joins("INNER JOIN TblUserAdmin ua ON u.mUserNo = ua.mUserNo").
		Where("m.email = ? OR m.mUserId = ?", mp.Member.Email, mp.Member.MUserId).
		Find(&mp)

	if resultMember.Error != nil {
		if errors.Is(resultMember.Error, gorm.ErrRecordNotFound) {
			a.log.Info("%s, %s", op, "record not found in Member")
		} else {
			a.log.Info("%s, %v", op, resultMember.Error)
		}
		errorsList = append(errorsList, resultMember.Error)
	}

	// Вторая часть: Поиск по имени персонажа
	resultNikname := a.db.Table("TblPc pc").
		Select("*").
		Joins("INNER JOIN PcState pcState ON pc.mNo = pcState.mNo").
		Joins("INNER JOIN TblPcInventory inventory ON pc.mNo = inventory.mPcNo").
		Joins("INNER JOIN PcStore store ON pc.mNo = store.mNo").
		Where("pc.mNm = ?", mp.Pc.MNm).
		Find(&mp)

	if resultNikname.Error != nil {
		if errors.Is(resultNikname.Error, gorm.ErrRecordNotFound) {
			a.log.Info("%s, %s", op, "record not found by nickname")
			
		} else {
			a.log.Info("%s, %v", op, resultNikname.Error)
		}
		errorsList = append(errorsList, resultNikname.Error)
	}

	// Если ошибок больше 2, вернуть пустоту и список ошибок
	if len(errorsList) >= 2 {
		return query.MemberParm{}, fmt.Errorf("%s, %v", op, errorsList)
	}

	return mp, nil

}  */


/* 
	resultMember := a.db.Table("Member m").
		Select("*").
		Joins("INNER JOIN TblUser u ON m.mUserId = u.mUserId").
		Joins("INNER JOIN TblUserAdmin ua ON u.mUserNo = ua.mUserNo").
		Where("m.email = ? OR m.mUserId = ?", mp.Member.Email, mp.Member.MUserId).Find(&mp); 
	if resultMember.Error != nil {
		if errors.Is(resultMember.Error, gorm.ErrRecordNotFound) {
			// тут записать ошибку, но ничего не возращать 
			query.MemberParm{}, fmt.Errorf("%s, %w", op, errors.New("record not found"))
		}
		// записать ошибку, но ничего не возращать 
		query.MemberParm{}, fmt.Errorf("%s, %w", op, resultMember.Error)
	} else {
		a.db.Table("TblPc pc").Select("*").
		Joins("INNER JOIN PcState pcState ON pc.mNo = pcState.mNo").
		Joins("INNER JOIN TblPcInventory inventory ON pc.mNo = inventory.mPcNo").
		Joins("INNER JOIN PcStore store ON u.mUserNo = store.mUserNo").
		Where("pc.mNm = ?", mp.Pc.MNm).Find(&mp);
	}

	// Поиск по имени персонажа
	resultNikname := a.db.Table("TblPc pc").Select("*").
		Joins("INNER JOIN PcState pcState ON pc.mNo = pcState.mNo").
		Joins("INNER JOIN TblPcInventory inventory ON pc.mNo = inventory.mPcNo").
		Joins("INNER JOIN PcStore store ON u.mUserNo = store.mUserNo").
		Where("pc.mNm = ?", mp.Pc.MNm)

	if resultNikname.Error != nil {
		if errors.Is(resultNikname.Error, gorm.ErrRecordNotFound) {
			// тут записать ошибку, но ничего не возращать
			query.MemberParm{}, fmt.Errorf("%s, %w", op, errors.New("record not found"))
		}
		// тут записать ошибку, но ничего не возращать
		query.MemberParm{}, fmt.Errorf("%s, %w", op, resultNikname.Error)
	} else {
		a.db.Table("TblUser u").
		Select("*").
		Joins("INNER JOIN TblUserAdmin ua ON u.mUserNo = ua.mUserNo").
		Joins("INNER JOIN Member m ON m.mUserId = u.mUserId").
		Where("m.mUserNo = ?", mp.Pc.MOwner).Find(&mp);
	}

	// если ошибок = 2 тогда возращаем пустоту и список ошибок  */