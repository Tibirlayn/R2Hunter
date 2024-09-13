package game

import (
	"database/sql"
	"time"
)

// TblPc
type Pc struct {
	MRegDate   time.Time `json:"mRegDate" gorm:"column:mRegDate"`
	MOwner     int       `json:"mOwner" gorm:"column:mOwner"`      // аккаунт персонажа
	MSlot      int8      `json:"mSlot" gorm:"column:mSlot"`        // слот персонажа в меню выбара [0, 1, 2]
	MNo        int       `json:"mNo" gorm:"column:mNo;primaryKey"` // id персонажа
	MNm        string    `json:"mNm" gorm:"column:mNm;size:12"`    // имя персонажа
	MClass     int8      `json:"mClass" gorm:"column:mClass"`      // класс персонажа [ 0 - рыцарь, 1 - рейджер, 2 - маг, 3 - ассасин, 4 - призыватель ]
	MSex       int8      `json:"mSex" gorm:"column:mSex"`          // пол персонажа
	MHead      int8      `json:"mHead" gorm:"column:mHead"`        // волосы персонажа
	MFace      int8      `json:"mFace" gorm:"column:mFace"`        // лицо персонажа
	MBody      int8      `json:"mBody" gorm:"column:mBody"`        // тело
	MHomeMapNo int       `json:"mHomeMapNo" gorm:"column:mHomeMapNo"`
	MHomePosX  float64   `json:"mHomePosX" gorm:"column:mHomePosX"` // координаты расположения персонажа
	MHomePosY  float64   `json:"mHomePosY" gorm:"column:mHomePosY"` // координаты расположения персонажа
	MHomePosZ  float64   `json:"mHomePosZ" gorm:"column:mHomePosZ"` // координаты расположения персонажа
	MDelDate   time.Time `json:"mDelDate" gorm:"column:mDelDate"`   // если стоит дата персонаж удален

	// PcInventories []PcInventory `gorm:"foreignKey:MPcNo;references:MNo"`
	// PcStates      []PcState     `gorm:"foreignKey:MNo;references:MNo"`
}

func (Pc) TableName() string {
	return "TblPc"
}

type IntermediatePc struct {
	MRegDate   time.Time     `json:"Pc_mRegDate" gorm:"column:Pc_mRegDate"` // mRegDate,
	MOwner     sql.NullInt64 `json:"mOwner" gorm:"column:mOwner"`
	MSlot      int8          `json:"mSlot" gorm:"column:mSlot"`
	MNo        sql.NullInt64 `json:"Pc_mNo" gorm:"column:Pc_mNo;primaryKey"` // mNo
	MNm        string        `json:"mNm" gorm:"column:mNm;size:12"`
	MClass     int8          `json:"mClass" gorm:"column:mClass"`
	MSex       int8          `json:"mSex" gorm:"column:mSex"`
	MHead      int8          `json:"mHead" gorm:"column:mHead"`
	MFace      int8          `json:"mFace" gorm:"column:mFace"`
	MBody      int8          `json:"mBody" gorm:"column:mBody"`
	MHomeMapNo sql.NullInt64 `json:"mHomeMapNo" gorm:"column:mHomeMapNo"`
	MHomePosX  float64       `json:"mHomePosX" gorm:"column:mHomePosX"`
	MHomePosY  float64       `json:"mHomePosY" gorm:"column:mHomePosY"`
	MHomePosZ  float64       `json:"mHomePosZ" gorm:"column:mHomePosZ"`
	MDelDate   sql.NullTime  `json:"Pc_mDelDate" gorm:"column:Pc_mDelDate"` // mDelDate
}

func (IntermediatePc) TableName() string {
	return "TblPc"
}

