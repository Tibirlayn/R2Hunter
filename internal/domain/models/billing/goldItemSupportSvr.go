package billing

// TBLGoldItemSupportSvr
type GoldItemSupportSvr struct {
	GoldItemID int64 `json:"GoldItemID" gorm:"column:GoldItemID;not null"`
	MSvrNo     int16 `json:"mSvrNo" gorm:"column:mSvrNo;not null"`
}

func (GoldItemSupportSvr) TableName() string {
	return "TBLGoldItemSupportSvr"
}