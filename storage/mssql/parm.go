package mssql

import (
	"fmt"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
	"strconv"

	"github.com/Tibirlayn/R2Hunter/internal/config"
	queryParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm/item"
	qParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/parm"
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
				INNER JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE mo.MName = ? AND di.DIsEvent = 0
		`
		if err := p.db.Raw(query, name).Scan(&resIBD).Error; err != nil {
			return []queryParm.MonsterDrop{}, fmt.Errorf("%s, %w", op, err)
		}
	} else {
		query = `
			SELECT DISTINCT
				mo.MName AS 'Boss',
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
				INNER JOIN DT_ItemResource ir ON ir.ROwnerID = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE mo.MID = ? AND di.DIsEvent = 0
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