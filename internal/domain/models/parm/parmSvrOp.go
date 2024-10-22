package parm

type ParmSvrOp struct{
	MSvrNo     int16   `gorm:"column:mSvrNo" json:"mSvrNo"`         // smallint
	MOpNo      int     `gorm:"column:mOpNo" json:"mOpNo"`           // int
	MIsSetup   bool    `gorm:"column:mIsSetup" json:"mIsSetup"`     // bit
	MOpValue1  float64 `gorm:"column:mOpValue1" json:"mOpValue1"`   // float(53)
	MOpValue2  float64 `gorm:"column:mOpValue2" json:"mOpValue2"`   // float(53)
	MOpValue3  float64 `gorm:"column:mOpValue3" json:"mOpValue3"`   // float(53)
	MOpValue4  float64 `gorm:"column:mOpValue4" json:"mOpValue4"`   // float(53)
	MOpValue5  float64 `gorm:"column:mOpValue5" json:"mOpValue5"`   // float(53)
	MOpValue6  float64 `gorm:"column:mOpValue6" json:"mOpValue6"`   // float(53)
	MOpValue7  float64 `gorm:"column:mOpValue7" json:"mOpValue7"`   // float(53)
	MOpValue8  float64 `gorm:"column:mOpValue8" json:"mOpValue8"`   // float(53)
	MOpValue9  float64 `gorm:"column:mOpValue9" json:"mOpValue9"`   // float(53)
	MOpValue10 float64 `gorm:"column:mOpValue10" json:"mOpValue10"` // float(53)
	MOpValue11 float64 `gorm:"column:mOpValue11" json:"mOpValue11"` // float(53)
	MOpValue12 float64 `gorm:"column:mOpValue12" json:"mOpValue12"` // float(53)
	MOpValue13 float64 `gorm:"column:mOpValue13" json:"mOpValue13"` // float(53)
	MOpValue14 float64 `gorm:"column:mOpValue14" json:"mOpValue14"` // float(53)
	MOpValue15 float64 `gorm:"column:mOpValue15" json:"mOpValue15"` // float(53)
	MOpValue16 float64 `gorm:"column:mOpValue16" json:"mOpValue16"` // float(53)
	MOpValue17 float64 `gorm:"column:mOpValue17" json:"mOpValue17"` // float(53)
	MOpValue18 float64 `gorm:"column:mOpValue18" json:"mOpValue18"` // float(53)
	MOpValue19 float64 `gorm:"column:mOpValue19" json:"mOpValue19"` // float(53)
	MOpValue20 float64 `gorm:"column:mOpValue20" json:"mOpValue20"` // float(53)
	MOpValue21 float64 `gorm:"column:mOpValue21" json:"mOpValue21"` // float(53)
	MOpValue22 float64 `gorm:"column:mOpValue22" json:"mOpValue22"` // float(53)
	MOpValue23 float64 `gorm:"column:mOpValue23" json:"mOpValue23"` // float(53)
	MOpValue24 float64 `gorm:"column:mOpValue24" json:"mOpValue24"` // float(53)
	MOpValue25 float64 `gorm:"column:mOpValue25" json:"mOpValue25"` // float(53)
	MOpValue26 float64 `gorm:"column:mOpValue26" json:"mOpValue26"` // float(53)
	MOpValue27 float64 `gorm:"column:mOpValue27" json:"mOpValue27"` // float(53)
	MOpValue28 float64 `gorm:"column:mOpValue28" json:"mOpValue28"` // float(53)
	MOpValue29 float64 `gorm:"column:mOpValue29" json:"mOpValue29"` // float(53)
	MOpValue30 float64 `gorm:"column:mOpValue30" json:"mOpValue30"` // float(53)
}

func (ParmSvrOp) TableName() string {
	return "TblParmSvrOp"
}
