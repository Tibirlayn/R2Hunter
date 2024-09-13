package mssql

import (
	"fmt"
	"strconv"

	"github.com/Tibirlayn/R2Hunter/internal/config"
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

func (p *ParmStorage) BossDrop(ctx *fiber.Ctx, name string) ([]queryParm.ItemBossDrop, error) {
	const op = "storage.mssql.parm.BossDrop"

	resIBD := []queryParm.ItemBossDrop{}
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
				END AS 'Event'
				FROM DT_Monster AS mo
				INNER JOIN DT_MonsterDrop AS md ON mo.MID = md.MID
				INNER JOIN DT_DropGroup AS dg ON md.DGroup = dg.DGroup
				INNER JOIN TP_DropGroup AS tdg ON md.DGroup = tdg.DGroup
				INNER JOIN DT_DropItem AS di ON dg.DDrop = di.DDrop
				INNER JOIN TP_ItemStatus AS tpis ON di.DStatus = tpis.mStatus
				INNER JOIN DT_Item AS it ON di.DItem = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE mo.MName = ? AND di.DIsEvent = 0
		`
		if err := p.db.Raw(query, name).Scan(&resIBD).Error; err != nil {
			return []queryParm.ItemBossDrop{}, fmt.Errorf("%s, %w", op, err)
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
				END AS 'Event'
				FROM DT_Monster AS mo
				INNER JOIN DT_MonsterDrop AS md ON mo.MID = md.MID
				INNER JOIN DT_DropGroup AS dg ON md.DGroup = dg.DGroup
				INNER JOIN TP_DropGroup AS tdg ON md.DGroup = tdg.DGroup
				INNER JOIN DT_DropItem AS di ON dg.DDrop = di.DDrop
				INNER JOIN TP_ItemStatus AS tpis ON di.DStatus = tpis.mStatus
				INNER JOIN DT_Item AS it ON di.DItem = it.IID
				LEFT JOIN TblEventObj AS evo ON it.IID = evo.mObjID
				WHERE mo.MID = ? AND di.DIsEvent = 0
		`
		if err := p.db.Raw(query, number).Scan(&resIBD).Error; err != nil {
			return []queryParm.ItemBossDrop{}, fmt.Errorf("%s, %w", op, err)
		}
	}

	return resIBD, nil
}