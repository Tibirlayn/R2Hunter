package mssql

import (
	"fmt"
	"strconv"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
	"github.com/Tibirlayn/R2Hunter/storage"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	qParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm"
	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type ParmStorage struct {
	db *gorm.DB
}

func NewParmStorage(cfg_db *config.ConfigDB) (*ParmStorage, error) {
	const op = "storage.mssql.parm.New"
	parm := cfg_db.Parm

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		parm.User, parm.Password, parm.Server, parm.Port, parm.NameDB)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ParmStorage{db: db}, nil
}

func (s *ParmStorage) Stop() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (p *ParmStorage) BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.MonsterDrop, error) {
	const op = "storage.mssql.parm.BossDrop"

	resIBD := []queryParm.MonsterDrop{}
	var query string

	number, err := strconv.Atoi(name)

	if err != nil {
		query = `
			SELECT DISTINCT
				mo.MName AS 'Boss',
				mo.MID AS 'BossID',
				tdg.DName AS 'NameGroup',
				it.IName AS 'NameItem',
				tpis.mDesc AS 'ItemDesc',
				CASE
					WHEN evo.mObjID IS NULL THEN 'Not event item'
					ELSE 'Event item'
				END AS 'Event',
			    ir.RFileName AS 'RFileName',
				ir.RPosX AS 'RPosX',
				ir.RPosY AS 'RPosY',
			    di.DDrop AS 'DDrop'
				FROM DT_Monster AS mo
				INNER JOIN DT_MonsterDrop AS md ON mo.MID = md.MID
				INNER JOIN DT_DropGroup AS dg ON md.DGroup = dg.DGroup
				INNER JOIN TP_DropGroup AS tdg ON md.DGroup = tdg.DGroup
				INNER JOIN DT_DropItem AS di ON dg.DDrop = di.DDrop
				INNER JOIN TP_ItemStatus AS tpis ON di.DStatus = tpis.mStatus
				INNER JOIN DT_Item AS it ON di.DItem = it.IID
				LEFT JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE (mo.MName = ? AND di.DIsEvent = 0) AND ir.RFileName != '0'
		`
		if err := p.db.Raw(query, name).Scan(&resIBD).Error; err != nil {
			return []queryParm.MonsterDrop{}, fmt.Errorf("%s, %w", op, err)
		}
	} else {
		query = `
			SELECT DISTINCT
				mo.MName AS 'Boss',
				mo.MID AS 'BossID',
				tdg.DName AS 'NameGroup',
				it.IName AS 'NameItem',
				tpis.mDesc AS 'ItemDesc',
				CASE
					WHEN evo.mObjID IS NULL THEN 'Not event item'
					ELSE 'Event item'
				END AS 'Event',
			    ir.RFileName AS 'RFileName',
				ir.RPosX AS 'RPosX',
				ir.RPosY AS 'RPosY',
			    di.DDrop AS 'DDrop'
				FROM DT_Monster AS mo
				INNER JOIN DT_MonsterDrop AS md ON mo.MID = md.MID
				INNER JOIN DT_DropGroup AS dg ON md.DGroup = dg.DGroup
				INNER JOIN TP_DropGroup AS tdg ON md.DGroup = tdg.DGroup
				INNER JOIN DT_DropItem AS di ON dg.DDrop = di.DDrop
				INNER JOIN TP_ItemStatus AS tpis ON di.DStatus = tpis.mStatus
				INNER JOIN DT_Item AS it ON di.DItem = it.IID
				LEFT JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE (mo.MID = ? AND di.DIsEvent = 0) AND ir.RFileName != '0'
		`
		if err := p.db.Raw(query, number).Scan(&resIBD).Error; err != nil {
			return []queryParm.MonsterDrop{}, fmt.Errorf("%s, %w", op, err)
		}
	}

	return resIBD, nil
}

func (p *ParmStorage) ItemDDrop(ctx *fiber.Ctx, id int) (parm.DropItem, error) {
	const op = "storage.mssql.parm.ItemDDrop"

	di := parm.DropItem{}

	if err := p.db.Where("DDrop = ?", id).First(&di).Error; err != nil {
		return parm.DropItem{}, fmt.Errorf("%s, %w", op, err)
	}

	return di, nil
}

func (p *ParmStorage) UpdateItemDDrop(ctx *fiber.Ctx, name parm.DropItem) (parm.DropItem, error) {
	const op = "storage.mssql.parm.UpdateItemDDrop"

	tx := p.db.Begin()

	// Проверка наличия записи по условию DDrop
	var existingItem parm.DropItem
	result := tx.Where("DDrop = ?", name.DDrop).First(&existingItem)

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		if result.RowsAffected == 0 {
			return parm.DropItem{}, fmt.Errorf("%s: запись с DDrop = %v не найдена", op, name.DDrop)
		}
		return parm.DropItem{}, fmt.Errorf("%s, %w", op, result.Error)
	}

	if err := tx.Model(&existingItem).Where("DDrop = ?", name.DDrop).Updates(name).Error; err != nil {
		tx.Rollback()
		return parm.DropItem{}, fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return parm.DropItem{}, fmt.Errorf("%s, %w", op, err)
	}

	return name, nil
}

func (p *ParmStorage) ItemsRess(ctx *fiber.Ctx, name string) (qParm.ItemRes, error) {
	const op = "storage.mssql.parm.ItemsRess"

	var itemRes qParm.ItemRes

	if err := p.db.Table("DT_Item it").
	Select("it.IID AS IID, it.IName AS IName, ir.RFileName AS RFileName, ir.RPosX AS RPosX, ir.RPosY AS RPosY").
	Joins("INNER JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID").
	Where("it.IName LIKE '%' + ? + '%' AND ir.RFileName != '0' ", name).
	Scan(&itemRes).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return itemRes, nil
}

func (p *ParmStorage) ItemsRessbyID(ctx *fiber.Ctx, id []int) (qParm.ItemRes, error) {
	const op = "storage.mssql.parm.ItemsRess"

	var itemRes qParm.ItemRes

	if err := p.db.Table("DT_Item it").
	Select("it.IID AS IID, it.IName AS IName, ir.RFileName AS RFileName, ir.RPosX AS RPosX, ir.RPosY AS RPosY").
	Joins("INNER JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID").
	Where("it.IID IN (?) AND ir.RFileName != '0' ", id).
	Scan(&itemRes).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return itemRes, nil
}

func (p *ParmStorage) Monster(ctx *fiber.Ctx, name string) ([]parm.Moster, error) {
	const op = "storage.mssql.parm.Monster"

	monster := []parm.Moster{}

	if err := p.db.Where("MName LIKE '%' + ? + '%' ", name).Find(&monster).Error; err != nil {
		return []parm.Moster{}, fmt.Errorf("%s, %w", op, err)
	}


/* 	if err := p.db.Table("DT_Monster m").
	Select("m.MID AS MID, m.MName AS MName, m.MLevel AS MLevel, m.MClass AS MClass, m.MExp AS MExp, m.MHIT AS MHIT, m.MMinD AS MMinD, m.MMaxD AS MMaxD, m.MAttackRateOrg AS MAttackRateOrg, m.MMoveRateOrg AS MMoveRateOrg, m.MAttackRateNew AS MAttackRateNew, m.MMoveRateNew AS MMoveRateNew, m.MHP AS MHP, m.MMP AS MMP, m.MMoveRange AS MMoveRange, m.MGbjType AS MGbjType, m.MRaceType AS MRaceType, m.MAiType AS MAiType, m.MCastingDelay AS MCastingDelay, m.MChaotic AS MChaotic, m.MSameRace1 AS MSameRace1, m.MSameRace2 AS MSameRace2, m.MSameRace3 AS MSameRace3, m.MSameRace4 AS MSameRace4, m.MSightRange AS MSightRange, m.MAttackRange AS MAttackRange, m.MSkillRange AS MSkillRange, m.MBodySize AS MBodySize, m.MDetectTransF AS MDetectTransF, m.MDetectTransP AS MDetectTransP, m.MDetectChao AS MDetectChao, m.MAiEx AS MAiEx, m.MScale AS MScale, m.MIsResistTransF AS MIsResistTransF, m.MIsEvent AS MIsEvent, m.MIsTest AS MIsTest, m.MHPNew AS MHPNew, m.MMPNew AS MMPNew, m.MBuyMerchanID AS MBuyMerchanID, m.MSellMerchanID AS MSellMerchanID, m.MChargeMerchanID AS MChargeMerchanID, m.MTransformWeight AS MTransformWeight, m.MNationOp AS MNationOp, m.MHPRegen AS MHPRegen, m.MMPRegen AS MMPRegen, m.IContentsLv AS IContentsLv, m.MIsEventTest AS MIsEventTest, m.MIsShowHp AS MIsShowHp, m.MSupportType AS MSupportType, m.MVolitionOfHonor AS MVolitionOfHonor, m.MWMapIconType AS MWMapIconType, m.MIsAmpliableTermOfValidity AS MIsAmpliableTermOfValidity, m.MAttackType AS MAttackType, m.MTransType AS MTransType, m.MDPV AS MDPV, m.MMPV AS MMPV, m.MRPV AS MRPV, m.MDDV AS MDDV, m.MMDV AS MMDV, m.MRDV AS MRDV, m.MSubDDWhenCritical AS MSubDDWhenCritical, m.MEnemySubCriticalHit AS MEnemySubCriticalHit, m.MEventQuest AS MEventQuest, m.MEScale AS MEScale").	
	Where("MName LIKE '%' + ? + '%' ", name).
	Scan(&monster).Error; err != nil {
		return parm.Moster{}, fmt.Errorf("%s, %w", op, err)
	} */

	return monster, nil
}

func (p *ParmStorage) ParmSvr(ctx *fiber.Ctx) ([]parm.ParmSvr, error) {
	const op = "storage.mssql.parm.Monster"

	parmSvr := []parm.ParmSvr{}
	if err := p.db.Where("MIsValid = 1").Find(&parmSvr).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return parmSvr, nil
}

func (p *ParmStorage) ParmSvrOp(ctx *fiber.Ctx, worldNo []int16) ([]parm.ParmSvrOp, error) {
	const op = "storage.mssql.parm.Monster"

	parmSvrOp := []parm.ParmSvrOp{}
	if err := p.db.Where("mSvrNo IN ?", worldNo).Find(&parmSvrOp).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return parmSvrOp, nil
}

func (p *ParmStorage) UpdateParmSvrOp(ctx *fiber.Ctx, svrOp parm.ParmSvrOp) (string, error) {
	const op = "storage.mssql.parm.UpdateParmSvrOp"

	svr := parm.ParmSvrOp{}
	tx := p.db.Begin()
	
	result := tx.Where("mSvrNo = ? AND mOpNo = ?", svrOp.MSvrNo, svrOp.MOpNo).First(&svr)

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		if result.RowsAffected == 0 {
			return "", fmt.Errorf("%s: запись с MSvrNo = %v mOpNo = %v не найдена", op, svrOp.MSvrNo, svrOp.MOpNo)
		}
		return "", fmt.Errorf("%s, %w", op, result.Error)
	}

	if err := tx.Model(&svr).Where("mSvrNo = ? AND mOpNo = ?", svrOp.MSvrNo, svrOp.MOpNo).Updates(svrOp).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "changes saved", nil
}

func (p *ParmStorage) ItemSearch(ctx *fiber.Ctx, name string) (qParm.ItemSearch, error) {
	const op = "storage.mssql.parm.ItemSearch"

	items := qParm.ItemSearch{}
	if err := p.db.Table("DT_Item it").
	Select("it.IID AS IID, it.IName AS IName, it.IType AS IType, it.ILevel AS ILevel, it.IDHIT AS IDHIT, it.IDDD AS IDDD, it.IRHIT AS IRHIT, it.IRDD AS IRDD, it.IMHIT AS IMHIT, it.IMDD AS IMDD, it.IHPPlus AS IHPPlus, it.IMPPlus AS IMPPlus, it.ISTR AS ISTR, it.IDEX AS IDEX, it.IINT AS IINT, it.IMaxStack AS IMaxStack, it.IWeight AS IWeight, it.IUseType AS IUseType, it.IUseNum AS IUseNum, it.IRecycle AS IRecycle, it.IHPRegen AS IHPRegen, it.IMPRegen AS IMPRegen, it.IAttackRate AS IAttackRate, it.IMoveRate AS IMoveRate, it.ICritical AS ICritical, it.ITermOfValidity AS ITermOfValidity, it.ITermOfValidityMi AS ITermOfValidityMi, it.IDesc AS IDesc, it.IStatus AS IStatus, it.IFakeID AS IFakeID, it.IFakeName AS IFakeName, it.IUseMsg AS IUseMsg, it.IRange AS IRange, it.IUseClass AS IUseClass, it.IDropEffect AS IDropEffect, it.IUseLevel AS IUseLevel, it.IUseEternal AS IUseEternal, it.IUseDelay AS IUseDelay, it.IUseInAttack AS IUseInAttack, it.IIsEvent AS IIsEvent, it.IIsIndict AS IIsIndict, it.IAddWeight AS IAddWeight, it.ISubType AS ISubType, it.IIsCharge AS IIsCharge, it.INationOp AS INationOp, it.IPShopItemType AS IPShopItemType, it.IQuestNo AS IQuestNo, it.IIsTest AS IIsTest, it.IQuestNeedCnt AS IQuestNeedCnt, it.IContentsLv AS IContentsLv, it.IIsConfirm AS IIsConfirm, it.IIsSealable AS IIsSealable, it.IAddDDWhenCritical AS IAddDDWhenCritical, it.mSealRemovalNeedCnt AS MSealRemovalNeedCnt, it.mIsPracticalPeriod AS MIsPracticalPeriod, it.mIsReceiveTown AS MIsReceiveTown, it.IIsReinforceDestroy AS IIsReinforceDestroy, it.IAddPotionRestore AS IAddPotionRestore, it.IAddMaxHpWhenTransform AS IAddMaxHpWhenTransfo, it.IAddMaxMpWhenTransform AS IAddMaxMpWhenTransfo, it.IAddAttackRateWhenTransform AS IAddAttackRateWhenTr, it.IAddMoveRateWhenTransform AS IAddMoveRateWhenTran, it.ISupportType AS ISupportType, it.ITermOfValidityLv AS ITermOfValidityLv, it.mIsUseableUTGWSvr AS MIsUseableUTGWSvr, it.IAddShortAttackRange AS IAddShortAttackRange, it.IAddLongAttackRange AS IAddLongAttackRange, it.IWeaponPoisonType AS IWeaponPoisonType, it.IDPV AS IDPV, it.IMPV AS IMPV, it.IRPV AS IRPV, it.IDDV AS IDDV, it.IMDV AS IMDV, it.IRDV AS IRDV, it.IHDPV AS IHDPV, it.IHMPV AS IHMPV, it.IHRPV AS IHRPV, it.IHDDV AS IHDDV, it.IHMDV AS IHMDV, it.IHRDV AS IHRDV, it.ISubDDWhenCritical AS ISubDDWhenCritical, it.IGetItemFeedback AS IGetItemFeedback, it.IEnemySubCriticalHit AS IEnemySubCriticalHit, it.IIsPartyDrop AS IIsPartyDrop, it.IMaxBeadHoleCount AS IMaxBeadHoleCount, it.ISubTypeOption AS ISubTypeOption, it.mIsDeleteArenaSvr AS MIsDeleteArenaSvr, ir.RID AS RID, ir.ROwnerID  AS ROwnerID, ir.RType AS RType, ir.RFileName AS RFileName, ir.RPosX AS RPosX, ir.RPosY AS RPosY ").
		Joins("INNER JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID").
		Where("it.IName LIKE '%' + ? + '%' AND ir.RFileName != '0' AND (RPosX != 0 AND RPosY != 0)", name).
		Scan(&items).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return items, nil
}

func (p *ParmStorage) MaterialDraw(ctx *fiber.Ctx, name string) (qParm.MaterialDraw, error) {
	const op = "storage.mssql.parm.MaterialDraw"

	md := qParm.MaterialDraw{}
	if err := p.db.Table("TblMaterialDrawResult AS a").
	Select("a.mSeq AS MSeqResult, a.MDRD AS MDRDResult, a.IID AS IIDResult, a.mPerOrRate AS MPerOrRate, a.mItemStatus AS MItemStatus, a.mCnt AS MCntResult, CASE a.mBinding WHEN 0 THEN 'NO' WHEN 1 THEN 'YES' END AS MBinding, a.mEffTime AS MEffTime, a.mValTime AS MValTime, a.mResource AS MResource, a.mAddGroup AS MAddGroup, " +
	"a2.MDID AS MDIDIndex, a2.MDRD AS MDRDIndex, a2.mResType AS MResType, a2.mMaxResCnt AS MMaxResCnt, a2.mSuccess AS MSuccess, a2.mDesc AS MDesc, a2.mAddQuestionMark AS MAddQuestionMark, a2.mDescRus AS MDescRus, " +
	"a3.mSeq AS MSeqMaterial, a3.MDID AS MDIDMaterial, a3.IID AS IIDMaterial, a3.mCnt AS MCntMaterial, " +
	"b.IName AS INameRes, b2.IName AS IName, ira.RID AS RIDRes, ira.ROwnerID AS ROwnerIDRes, ira.RType AS RTypeRes, ira.RFileName AS RFileNameRes, ira.RPosX AS RPosXRes, ira.RPosY AS RPosYRes, ira3.RID AS RIDMat, ira3.ROwnerID AS ROwnerIDMat, ira3.RType AS RTypeMat, ira3.RFileName AS RFileNameMat, ira3.RPosX AS RPosXMat, ira3.RPosY AS RPosYMat").
	Joins("LEFT OUTER JOIN TblMaterialDrawIndex AS a2 ON a2.MDRD = a.MDRD").
	Joins("LEFT OUTER JOIN TblMaterialDrawMaterial AS a3 ON a3.MDID = a2.MDID").
	Joins("LEFT OUTER JOIN DT_Item AS b ON b.IID = a.IID").
	Joins("LEFT OUTER JOIN DT_Item AS b2 ON b2.IID = a3.IID").
	Joins("LEFT JOIN DT_ItemResource AS ira ON ira.ROwnerID =  a.IID").
	Joins("LEFT JOIN DT_ItemResource AS ira3 ON ira3.ROwnerID =  a3.IID").
	Where("ira.RFileName != '0' AND ira3.RFileName != '0' AND b2.IName LIKE '%' + ? + '%'", name).
	Scan(&md).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return md, nil
}

func (p *ParmStorage) ClearMaterialDraw(ctx *fiber.Ctx, id int) (string, error) {
	const op = "storage.mssql.parm.ClearMaterialDraw"

	if err := p.db.Where("MDID = ?", id).First(&parm.MaterialDrawMaterial{}).Error; err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	tx := p.db.Begin()

	if err := tx.Where("MDRD = ?", id).Delete(&parm.MaterialDrawResult{}).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Where("MDID = ?", id).Delete(&parm.MaterialDrawMaterial{}).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Where("MDID = ?", id).Delete(&parm.MaterialDrawIndex{}).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}	

	return "Data deleted ClearMaterialDraw", nil
}

func (p *ParmStorage) UpdateMaterialDrawIndex(ctx *fiber.Ctx, mdi parm.MaterialDrawIndex) (string, error) {
	const op = "storage.mssql.parm.UpdateMaterialDrawIndex"

	mdiRes := parm.MaterialDrawIndex{}

	tx := p.db.Begin()
	result := tx.Where("MDID = ?", mdi.MDRD).First(&mdiRes)

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		if result.RowsAffected == 0 {
			return "", fmt.Errorf("%s: запись %v не найдена", op, mdi.MDID)
		}
		return "", fmt.Errorf("%s, %w", op, result.Error)
	}

	if err := tx.Model(&mdiRes).Where("MDID = ?", mdi.MDRD).Updates(mdi).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data update UpdateMaterialDrawIndex", nil
}

func (p *ParmStorage) UpdateMaterialDrawResult(ctx *fiber.Ctx, mdr parm.MaterialDrawResult) (string, error) {
	const op = "storage.mssql.parm.UpdateMaterialDrawResult"

	mdiRes := parm.MaterialDrawResult{}

	tx := p.db.Begin()
	result := tx.Where("mSeq = ?", mdr.MSeq).First(&mdiRes)

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		if result.RowsAffected == 0 {
			return "", fmt.Errorf("%s: запись %v не найдена", op, mdr.MSeq)
		}
		return "", fmt.Errorf("%s, %w", op, result.Error)
	}

	if err := tx.Model(&mdiRes).Where("mSeq = ?", mdr.MSeq).Updates(mdr).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data update UpdateMaterialDrawResult", nil
}

func (p *ParmStorage) DeleteMaterialDrawResult(ctx *fiber.Ctx, seq int, mdrd int64, id int) (string, error) {
	const op = "storage.mssql.parm.DeleteMaterialDrawResult"

	firstMDR := parm.MaterialDrawResult{}

	tx := p.db.Begin()
	result := tx.Where("MSeq = ? AND MDRD = ? AND IID = ?", seq, mdrd, id).First(&firstMDR)

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		if result.RowsAffected == 0 {
			return "", fmt.Errorf("%s: запись с условиями MSeq = %v, MDRD = %v, IID = %v не найдена", op, seq, mdrd, id)
		}
		return "", fmt.Errorf("%s, %w", op, result.Error)
	}

	// Удаляем запись
	if err := tx.Where("MSeq = ? AND MDRD = ? AND IID = ?", seq, mdrd, id).Delete(&firstMDR).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}
	
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data DeleteMaterialDrawResult", nil
} 

func (p *ParmStorage) SetMaterialDrawResult(ctx *fiber.Ctx, mdr parm.MaterialDrawResult) (parm.MaterialDrawResult, error) {
	const op = "storage.mssql.parm.SetMaterialDrawResult"

	tx := p.db.Begin()

	var lastMSeq int 
	tx.Raw("SELECT COALESCE(MAX(mSeq), 0) FROM TblMaterialDrawResult").Scan(&lastMSeq)
	mdr.MSeq = lastMSeq + 1

	result := tx.Where("mSeq = ?", mdr.MSeq).Find(&parm.MaterialDrawResult{}); 
	
	if result.RowsAffected >= 1 {
		tx.Rollback()
		return parm.MaterialDrawResult{}, fmt.Errorf("%s: запись %s существует", op, storage.ErrExists)
	}

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		return parm.MaterialDrawResult{}, fmt.Errorf("%s, %w", op, result.Error)
	}

	if err := tx.Create(mdr).Error; err != nil {
		return parm.MaterialDrawResult{}, fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		return parm.MaterialDrawResult{}, fmt.Errorf("%s, %w", op, err)
	}

	return parm.MaterialDrawResult{}, nil
}

func (p *ParmStorage) QuestReward(ctx *fiber.Ctx, pageNumber int, limitCnt int) ([]qParm.QuestRewardRes, error) {
	const op = "storage.mssql.parm.QuestReward"

	offset :=  (pageNumber - 1) * limitCnt

	qr := []qParm.QuestRewardRes{}
    if err := p.db.
        Table("TblQuestReward AS qr").
        Select("ir.ROwnerID, ir.RType, ir.RFileName, ir.RPosX, ir.RPosY,qr.mRewardNo AS MRewardNo, qr.mExp AS MExp, qr.mID AS MID, qr.mCnt AS MCnt, qr.mBinding AS MBinding, qr.mStatus AS MStatus, qr.mEffTime AS MEffTime, qr.mValTime AS MValTime, i.IName").
        Joins("LEFT JOIN DT_Item i ON i.IID = qr.mID").
        Joins("LEFT JOIN DT_ItemResource ir ON ir.ROwnerID = i.IID").
        // Where("ir.RFileName != ?", "0").
		Order("qr.mRewardNo ASC").
        Offset(offset).
        Limit(limitCnt).
        Scan(&qr).Error; err != nil {
        return nil, fmt.Errorf("%s, %w", op, err)
    }
	
	return qr, nil
}

func (p *ParmStorage) SetQuestReward(ctx *fiber.Ctx, qr parm.QuestReward) (string, error) {
	const op = "storage.mssql.parm.SetQuestReward"
	
	tx := p.db.Begin()

	if err := tx.Create(&qr).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data created", nil
}

func (p *ParmStorage) DeleteQuestReward(ctx *fiber.Ctx, qr parm.QuestReward) (string, error) {
	const op = "storage.mssql.parm.DeleteQuestReward"
	
	tx := p.db.Begin()
	resQR := parm.QuestReward{}

	result := tx.Where("mRewardNo = ? AND MExp = ? AND mID = ? AND mCnt = ? AND mBinding = ? AND mStatus = ? AND mEffTime = ? AND mValTime = ?", 
	qr.MRewardNo, qr.MExp, qr.MID, qr.MCnt, qr.MBinding, qr.MStatus, qr.MEffTime, qr.MValTime).First(&resQR)

	// Проверка на ошибку выполнения запроса
	if result.Error != nil {
		if result.RowsAffected == 0 {
			tx.Rollback() // Откат транзакции при ошибке
			return "", fmt.Errorf("%s: запись с условиями %+v не найдена", op, qr)
		}
		tx.Rollback() // Откат транзакции при ошибке
		return "", fmt.Errorf("%s, %w", op, result.Error)
	}

	if err := tx.Where("mRewardNo = ? AND MExp = ? AND mID = ? AND mCnt = ? AND mBinding = ? AND mStatus = ? AND mEffTime = ? AND mValTime = ?", 
	qr.MRewardNo, qr.MExp, qr.MID, qr.MCnt, qr.MBinding, qr.MStatus, qr.MEffTime, qr.MValTime).Delete(&qr).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data created", nil
}

func (p *ParmStorage) SetMaterialDrawIndex(ctx *fiber.Ctx, mdi parm.MaterialDrawIndex, mdm parm.MaterialDrawMaterial) (string, error) {
	const op = "storage.mssql.parm.SetMaterialDrawIndex"

	tx := p.db.Begin()

	var lastMDID int64 
	tx.Raw("SELECT COALESCE(MAX(MDID), 0) FROM TblMaterialDrawIndex").Scan(&lastMDID)
	mdi.MDID = lastMDID + 1
	mdi.MDRD = lastMDID + 1

	var lastMSeq int
	tx.Raw("SELECT COALESCE(MAX(mSeq), 0) FROM TblMaterialDrawMaterial").Scan(&lastMSeq)
	mdm.MSeq = lastMSeq + 1
	mdm.MDID = lastMDID + 1

	resultMDI := tx.Where("MDID = ? OR MDRD = ?", mdi.MDID, mdi.MDRD).Find(&parm.MaterialDrawIndex{}); 
	resultMDM := tx.Where("mSeq = ?", mdm.MSeq).Find(&parm.MaterialDrawMaterial{}); 

	if resultMDI.RowsAffected >= 1 {
		tx.Rollback()
		return "", fmt.Errorf("%s: запись %s существует", op, storage.ErrExists)
	}

	// Проверка на ошибку выполнения запроса
	if resultMDI.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		return "", fmt.Errorf("%s, %w", op, resultMDI.Error)
	}

	if resultMDM.RowsAffected >= 1 {
		tx.Rollback()
		return "", fmt.Errorf("%s: запись %s существует", op, storage.ErrExists)
	}

	// Проверка на ошибку выполнения запроса
	if resultMDM.Error != nil {
		tx.Rollback() // Откат транзакции при ошибке
		return "", fmt.Errorf("%s, %w", op, resultMDM.Error)
	}

	if err := tx.Create(mdi).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Create(mdm).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return "Data created", nil
}