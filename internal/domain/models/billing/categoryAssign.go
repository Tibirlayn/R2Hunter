package billing

import "time"

// TBLCategoryAssign
type CategoryAssign struct {
	CategoryID  int16     `json:"CategoryID" gorm:"column:CategoryID;not null"`			  	// smallint
	GoldItemID  int64     `json:"GoldItemID" gorm:"column:GoldItemID;not null;primaryKey"`	// bigint
	Status      string    `json:"Status" gorm:"column:Status;not null"`			  			// nchar	1
	OrderNO     int16     `json:"OrderNO" gorm:"column:OrderNO;not null"`			  		// smallint
	RegistDate  time.Time `json:"RegistDate" gorm:"column:RegistDate;not null"`			  	// datetime
	RegistAdmin string    `json:"RegistAdmin" gorm:"column:RegistAdmin"`			  		// nvarchar	20
	RegistIP    string    `json:"RegistIP" gorm:"column:RegistIP"`					  		// nvarchar	19
	UpdateDate  time.Time `json:"UpdateDate" gorm:"column:UpdateDate"`			  			// datetime
	UpdateAdmin string    `json:"UpdateAdmin" gorm:"column:UpdateAdmin"`			  		// nvarchar	20
	UpdateIP    string    `json:"UpdateIP" gorm:"column:UpdateIP"`				  			// nvarchar	19
}

func (CategoryAssign) TableName() string {
	return "TBLCategoryAssign"
}
			