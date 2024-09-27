package parm

import "database/sql"

// DT_ItemResource
type ItemResource struct {
	RID       int    `json:"RID" gorm:"column:RID;not null;primaryKey"`
	ROwnerID  int    `json:"ROwnerID" gorm:"column:ROwnerID"`
	RType     int    `json:"RType" gorm:"column:RType"`
	RFileName string `json:"RFileName" gorm:"column:RFileName"`
	RPosX     int    `json:"RPosX" gorm:"column:RPosX"`
	RPosY     int    `json:"RPosY" gorm:"column:RPosY"`
}

func (ItemResource) TableName() string {
	return "DT_ItemResource"
}

type IntermediateItemResource struct {
	RID       sql.NullInt64  `json:"RID"   gorm:"column:ItemResource_RID;not null;primaryKey"`
	ROwnerID  sql.NullInt64  `json:"ROwnerID"  gorm:"column:ItemResource_ROwnerID"`
	RType     sql.NullInt64  `json:"RType"   gorm:"column:ItemResource_RType"`
	RFileName sql.NullString `json:"RFileName"  gorm:"column:ItemResource_RFileName"`
	RPosX     sql.NullInt64  `json:"RPosX"   gorm:"column:ItemResource_RPosX"`
	RPosY     sql.NullInt64  `json:"RPosY"   gorm:"column:ItemResource_RPosY"`
}
