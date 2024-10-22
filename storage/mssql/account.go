package mssql

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
	query "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
	"github.com/Tibirlayn/R2Hunter/pkg/lib/conv"
	"github.com/Tibirlayn/R2Hunter/storage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type AccountStorage struct {
	log *slog.Logger
	db  *gorm.DB
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

	memberParm := query.MemberParm{
		Members: []account.Member{},
		Users: []account.User{},
		Pcs: []game.Pc{},
		PcInvs: []game.PcInventory{},
		PcStates: []game.PcState{},
		PcStores: []game.PcStore{},
		ItemResources: []parm.ItemResource{},
		Items: []parm.Item{},
	}

	rows, err := a.db.Raw("SELECT * FROM dbo.UspGetMemberUserAll(@Login)", sql.Named("Login", name)).Rows()
	if err != nil {
		a.log.Info("%s, %w", op, err)
		return mp, fmt.Errorf("%s, %w", op, err)
	}

	for rows.Next() {
		var member account.IntermediateMember
		var user account.IntermediateUser
		var pc game.IntermediatePc
		var pcInv game.IntermediatePcInventory
		var pcState game.IntermediatePcState
		var pcStore game.IntermediatePcStore
		var item parm.IntermediateItem
		var itemRes parm.IntermediateItemResource

		err := rows.Scan(
			&user.MRegDate,&user.MUserAuth,&user.MUserNo,&user.MUserId,&user.MUserPswd,&user.MCertifiedKey,
			&user.MIp,&user.MLoginTm,&user.MLogoutTm,&user.MTotUseTm,&user.MWorldNo,&user.MDelDate,
			&user.MPcBangLv,&user.MSecKeyTableUse,&user.MUseMacro,&user.MIpEX,&user.MJoinCode,
			&user.MLoginChannelID,&user.MTired,&user.MChnSID,&user.MNewId,&user.MLoginSvrType,
			&user.MAccountGuid,&user.MNormalLimitTime,&user.MPcBangLimitTime,&user.MRegIp,&user.MIsMovingToBattleSvr,
			
			&member.MUserId,&member.MUserPswd,&member.Superpwd,&member.Cash,&member.Email,&member.Tgzh,
			&member.Uid,&member.Klq,&member.Ylq,&member.Auth,&member.MSum,&member.IsAdmin,&member.Isdl,
			&member.Dlmoney,&member.RegisterIp,&member.Country,&member.CashBack,
			
			&pc.MRegDate,&pc.MOwner,&pc.MSlot,&pc.MNo,&pc.MNm,&pc.MClass,&pc.MSex,&pc.MHead,&pc.MFace,
			&pc.MBody,&pc.MHomeMapNo,&pc.MHomePosX,&pc.MHomePosY,&pc.MHomePosZ,&pc.MDelDate,
			
			&pcState.MNo,&pcState.MLevel,&pcState.MExp,&pcState.MHpAdd,&pcState.MHp,&pcState.MMpAdd,&pcState.MMp,
			&pcState.MMapNo,&pcState.MPosX,&pcState.MPosY,&pcState.MPosZ,&pcState.MStomach,&pcState.MIp,
			&pcState.MLoginTm,&pcState.MLogoutTm,&pcState.MTotUseTm,&pcState.MPkCnt,&pcState.MChaotic,
			&pcState.MDiscipleJoinCount,&pcState.MPartyMemCntLevel,&pcState.MLostExp,&pcState.MIsLetterLimit,
			&pcState.MFlag,&pcState.MIsPreventItemDrop,&pcState.MSkillTreePoint,&pcState.MRestExpGuild,
			&pcState.MRestExpActivate,&pcState.MRestExpDeactivate,&pcState.MQMCnt,
			&pcState.MGuildQMCnt,&pcState.MFierceCnt,&pcState.MBossCnt,
			
			&pcInv.MRegDate,&pcInv.MSerialNo,&pcInv.MPcNo,&pcInv.MItemNo,&pcInv.MEndDate,&pcInv.MIsConfirm,
			&pcInv.MStatus,&pcInv.MCnt,&pcInv.MCntUse,&pcInv.MIsSeizure,&pcInv.MApplyAbnItemNo,
			&pcInv.MApplyAbnItemEndDate,&pcInv.MOwner,&pcInv.MPracticalPeriod,&pcInv.MBindingType,
			&pcInv.MRestoreCnt,&pcInv.MHoleCount,
			
			&pcStore.MRegDate,&pcStore.MSerialNo,&pcStore.MUserNo,&pcStore.MItemNo,&pcStore.MEndDate,
			&pcStore.MIsConfirm,&pcStore.MStatus,&pcStore.MCnt,&pcStore.MCntUse,&pcStore.MIsSeizure,
			&pcStore.MApplyAbnItemNo,&pcStore.MApplyAbnItemEndDate,&pcStore.MOwner,
			&pcStore.MPracticalPeriod,&pcStore.MBindingType,&pcStore.MRestoreCnt,&pcStore.MHoleCount,
		
			&itemRes.RID, &itemRes.ROwnerID, &itemRes.RType, &itemRes.RFileName, &itemRes.RPosX, &itemRes.RPosY,

			&item.IID, &item.IName, &item.IType, &item.ILevel, &item.IDHIT, &item.IDDD,
			&item.IRHIT, &item.IRDD, &item.IMHIT, &item.IMDD, &item.IHPPlus, &item.IMPPlus, &item.ISTR,
			&item.IDEX, &item.IINT, &item.IMaxStack, &item.IWeight, &item.IUseType, &item.IUseNum, &item.IRecycle,
			&item.IHPRegen, &item.IMPRegen, &item.IAttackRate, &item.IMoveRate, &item.ICritical, &item.ITermOfValidity, &item.ITermOfValidityMi,
			&item.IDesc, &item.IStatus, &item.IFakeID, &item.IFakeName, &item.IUseMsg, &item.IRange, &item.IUseClass,
			&item.IDropEffect, &item.IUseLevel, &item.IUseEternal, &item.IUseDelay, &item.IUseInAttack, &item.IIsEvent, &item.IIsIndict, &item.IAddWeight,
			&item.ISubType, &item.IIsCharge, &item.INationOp, &item.IPShopItemType, &item.IQuestNo, &item.IIsTest, &item.IQuestNeedCnt,
			&item.IContentsLv, &item.IIsConfirm, &item.IIsSealable, &item.IAddDDWhenCritical, &item.MSealRemovalNeedCnt, &item.MIsPracticalPeriod, &item.MIsReceiveTown,
			&item.IIsReinforceDestroy, &item.IAddPotionRestore, &item.IAddMaxHpWhenTransform, &item.IAddMaxMpWhenTransform, &item.IAddAttackRateWhenTransform, &item.IAddMoveRateWhenTransform, 
			&item.ISupportType, &item.ITermOfValidityLv, &item.MIsUseableUTGWSvr, &item.IAddShortAttackRange, &item.IAddLongAttackRange, &item.IWeaponPoisonType, &item.IDPV,
			&item.IMPV, &item.IRPV, &item.IDDV, &item.IMDV, &item.IRDV, &item.IHDPV, &item.IHMPV, &item.IHRPV,
			&item.IHDDV, &item.IHMDV, &item.IHRDV, &item.ISubDDWhenCritical, &item.IGetItemFeedback, &item.IEnemySubCriticalHit,
			&item.IIsPartyDrop, &item.IMaxBeadHoleCount, &item.ISubTypeOption, &item.MIsDeleteArenaSvr)
		if err != nil {
			return mp, fmt.Errorf("%s, %w", op, err)
		}

		conv.ConvMember(member, user, pc, pcInv, pcState, pcStore, item, itemRes, &memberParm)
	}

	return memberParm, nil
}


func (a *AccountStorage) MemberAll(ctx *fiber.Ctx, name string) (query.MemberPcItem, error) {
	const op = "storage.mssql.account.MemberAll"

	memberPcItem := query.MemberPcItem{
		Members: []account.Member{},
		Users: []account.User{},
		Pcs: []game.Pc{},
		PcInvs: []game.PcInventory{},
		PcStates: []game.PcState{},
		PcStores: []game.PcStore{},
		ItemResources: []parm.ItemResource{},
		Items: []parm.Item{},
	}

	rows, err := a.db.Raw("SELECT * FROM dbo.UspGetMemberUserAll(@Login)", sql.Named("Login", name)).Rows()
	if err != nil {
		a.log.Info("%s, %w", op, err)
		return query.MemberPcItem{}, fmt.Errorf("%s, %w", op, err)
	}

	for rows.Next() {
		var member account.IntermediateMember
		var user account.IntermediateUser
		var pc game.IntermediatePc
		var pcInv game.IntermediatePcInventory
		var pcState game.IntermediatePcState
		var pcStore game.IntermediatePcStore
		var item parm.IntermediateItem
		var itemRes parm.IntermediateItemResource

		err := rows.Scan(
			&user.MRegDate,&user.MUserAuth,&user.MUserNo,&user.MUserId,&user.MUserPswd,&user.MCertifiedKey,
			&user.MIp,&user.MLoginTm,&user.MLogoutTm,&user.MTotUseTm,&user.MWorldNo,&user.MDelDate,
			&user.MPcBangLv,&user.MSecKeyTableUse,&user.MUseMacro,&user.MIpEX,&user.MJoinCode,
			&user.MLoginChannelID,&user.MTired,&user.MChnSID,&user.MNewId,&user.MLoginSvrType,
			&user.MAccountGuid,&user.MNormalLimitTime,&user.MPcBangLimitTime,&user.MRegIp,&user.MIsMovingToBattleSvr,
			
			&member.MUserId,&member.MUserPswd,&member.Superpwd,&member.Cash,&member.Email,&member.Tgzh,
			&member.Uid,&member.Klq,&member.Ylq,&member.Auth,&member.MSum,&member.IsAdmin,&member.Isdl,
			&member.Dlmoney,&member.RegisterIp,&member.Country,&member.CashBack,
			
			&pc.MRegDate,&pc.MOwner,&pc.MSlot,&pc.MNo,&pc.MNm,&pc.MClass,&pc.MSex,&pc.MHead,&pc.MFace,
			&pc.MBody,&pc.MHomeMapNo,&pc.MHomePosX,&pc.MHomePosY,&pc.MHomePosZ,&pc.MDelDate,
			
			&pcState.MNo,&pcState.MLevel,&pcState.MExp,&pcState.MHpAdd,&pcState.MHp,&pcState.MMpAdd,&pcState.MMp,
			&pcState.MMapNo,&pcState.MPosX,&pcState.MPosY,&pcState.MPosZ,&pcState.MStomach,&pcState.MIp,
			&pcState.MLoginTm,&pcState.MLogoutTm,&pcState.MTotUseTm,&pcState.MPkCnt,&pcState.MChaotic,
			&pcState.MDiscipleJoinCount,&pcState.MPartyMemCntLevel,&pcState.MLostExp,&pcState.MIsLetterLimit,
			&pcState.MFlag,&pcState.MIsPreventItemDrop,&pcState.MSkillTreePoint,&pcState.MRestExpGuild,
			&pcState.MRestExpActivate,&pcState.MRestExpDeactivate,&pcState.MQMCnt,
			&pcState.MGuildQMCnt,&pcState.MFierceCnt,&pcState.MBossCnt,
			
			&pcInv.MRegDate,&pcInv.MSerialNo,&pcInv.MPcNo,&pcInv.MItemNo,&pcInv.MEndDate,&pcInv.MIsConfirm,
			&pcInv.MStatus,&pcInv.MCnt,&pcInv.MCntUse,&pcInv.MIsSeizure,&pcInv.MApplyAbnItemNo,
			&pcInv.MApplyAbnItemEndDate,&pcInv.MOwner,&pcInv.MPracticalPeriod,&pcInv.MBindingType,
			&pcInv.MRestoreCnt,&pcInv.MHoleCount,
			
			&pcStore.MRegDate,&pcStore.MSerialNo,&pcStore.MUserNo,&pcStore.MItemNo,&pcStore.MEndDate,
			&pcStore.MIsConfirm,&pcStore.MStatus,&pcStore.MCnt,&pcStore.MCntUse,&pcStore.MIsSeizure,
			&pcStore.MApplyAbnItemNo,&pcStore.MApplyAbnItemEndDate,&pcStore.MOwner,
			&pcStore.MPracticalPeriod,&pcStore.MBindingType,&pcStore.MRestoreCnt,&pcStore.MHoleCount,
		
			&itemRes.RID, &itemRes.ROwnerID, &itemRes.RType, &itemRes.RFileName, &itemRes.RPosX, &itemRes.RPosY,

			&item.IID, &item.IName, &item.IType, &item.ILevel, &item.IDHIT, &item.IDDD,
			&item.IRHIT, &item.IRDD, &item.IMHIT, &item.IMDD, &item.IHPPlus, &item.IMPPlus, &item.ISTR,
			&item.IDEX, &item.IINT, &item.IMaxStack, &item.IWeight, &item.IUseType, &item.IUseNum, &item.IRecycle,
			&item.IHPRegen, &item.IMPRegen, &item.IAttackRate, &item.IMoveRate, &item.ICritical, &item.ITermOfValidity, &item.ITermOfValidityMi,
			&item.IDesc, &item.IStatus, &item.IFakeID, &item.IFakeName, &item.IUseMsg, &item.IRange, &item.IUseClass,
			&item.IDropEffect, &item.IUseLevel, &item.IUseEternal, &item.IUseDelay, &item.IUseInAttack, &item.IIsEvent, &item.IIsIndict, &item.IAddWeight,
			&item.ISubType, &item.IIsCharge, &item.INationOp, &item.IPShopItemType, &item.IQuestNo, &item.IIsTest, &item.IQuestNeedCnt,
			&item.IContentsLv, &item.IIsConfirm, &item.IIsSealable, &item.IAddDDWhenCritical, &item.MSealRemovalNeedCnt, &item.MIsPracticalPeriod, &item.MIsReceiveTown,
			&item.IIsReinforceDestroy, &item.IAddPotionRestore, &item.IAddMaxHpWhenTransform, &item.IAddMaxMpWhenTransform, &item.IAddAttackRateWhenTransform, &item.IAddMoveRateWhenTransform, 
			&item.ISupportType, &item.ITermOfValidityLv, &item.MIsUseableUTGWSvr, &item.IAddShortAttackRange, &item.IAddLongAttackRange, &item.IWeaponPoisonType, &item.IDPV,
			&item.IMPV, &item.IRPV, &item.IDDV, &item.IMDV, &item.IRDV, &item.IHDPV, &item.IHMPV, &item.IHRPV,
			&item.IHDDV, &item.IHMDV, &item.IHRDV, &item.ISubDDWhenCritical, &item.IGetItemFeedback, &item.IEnemySubCriticalHit,
			&item.IIsPartyDrop, &item.IMaxBeadHoleCount, &item.ISubTypeOption, &item.MIsDeleteArenaSvr)
		if err != nil {
			return query.MemberPcItem{}, fmt.Errorf("%s, %w", op, err)
		}

		conv.ConvMemberAll(member, user, pc, pcInv, pcState, pcStore, item, itemRes, &memberPcItem)

	}

	return memberPcItem, nil
}

func (a *AccountStorage) MemberBil(ctx *fiber.Ctx, email string) (account.User, error) {
	const op = "storage.mssql.account.MemberBil"

	user := account.User{}
	if err := a.db.Where("mUserId = ?", email).First(&user).Error; err != nil {
		return account.User{}, fmt.Errorf("%s, %w", op, err)
	}

	return user, nil
}

func (a *AccountStorage) UserSearch(ctx *fiber.Ctx, name string) ([]account.User, error) {
	const op = "storage.mssql.account.User"

	user := []account.User{}
	if err := a.db.Where("MUserId LIKE '%' + ? + '%'", name).Find(&user).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return user, nil
}

func (a *AccountStorage) UserLastLogin(ctx *fiber.Ctx) ([]int, error) {
	const op = "storage.mssql.account.User"

	user := []account.User{}
	nowTime := time.Now()
	last30Days := nowTime.AddDate(0, 0, -30)
	userNo := []int{}
	
	if err := a.db.Model(&user).Where("mLoginTm >= ?", last30Days).Pluck("MUserNo", &userNo).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return userNo, nil
}