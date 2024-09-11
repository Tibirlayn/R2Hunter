package game

import (
	"database/sql"
	"time"
)

// TblPcInventory
type PcInventory struct {
	MRegDate             time.Time `json:"mRegDate" gorm:"column:mRegDate;not null"`
	MSerialNo            int64     `json:"mSerialNo" gorm:"column:mSerialNo;not null;primaryKey"`
	MPcNo                int       `json:"mPcNo" gorm:"column:mPcNo;not null"`
	MItemNo              int       `json:"mItemNo" gorm:"column:mItemNo;not null"`
	MEndDate             time.Time `json:"mEndDate" gorm:"column:mEndDate;not null"`
	MIsConfirm           bool      `json:"mIsConfirm" gorm:"column:mIsConfirm;not null"`
	MStatus              int8      `json:"mStatus" gorm:"column:mStatus;not null"`
	MCnt                 int       `json:"mCnt" gorm:"column:mCnt;not null"`
	MCntUse              int16     `json:"mCntUse" gorm:"column:mCntUse"`
	MIsSeizure           bool      `json:"mIsSeizure" gorm:"column:mIsSeizure;not null"`
	MApplyAbnItemNo      int       `json:"mApplyAbnItemNo" gorm:"column:mApplyAbnItemNo;not null"`
	MApplyAbnItemEndDate time.Time `json:"mApplyAbnItemEndDate" gorm:"column:mApplyAbnItemEndDate;not null"`
	MOwner               int       `json:"mOwner" gorm:"column:mOwner;not null"`
	MPracticalPeriod     int       `json:"mPracticalPeriod" gorm:"column:mPracticalPeriod;not null"`
	MBindingType         int8      `json:"mBindingType" gorm:"column:mBindingType;not null"`
	MRestoreCnt          int8      `json:"mRestoreCnt" gorm:"column:mRestoreCnt;not null"`
	MHoleCount           int8      `json:"mHoleCount" gorm:"column:mHoleCount;not null"`

	// Pc Pc `gorm:"foreignKey:MPcNo;references:MNo"` // Добавлено для явного обозначения связи, опционально
}

func (PcInventory) TableName() string {
	return "TblPcInventory"
}

type IntermediatePcInventory struct {
	MRegDate             time.Time     `json:"PcInv_mRegDate" gorm:"column:PcInv_mRegDate;not null"`
	MSerialNo            sql.NullInt64 `json:"PcInv_mSerialNo" gorm:"column:PcInv_mSerialNo;not null;primaryKey"`
	MPcNo                sql.NullInt64 `json:"PcInv_mPcNo" gorm:"column:PcInv_mPcNo;not null"`
	MItemNo              sql.NullInt64 `json:"PcInv_mItemNo" gorm:"column:PcInv_mItemNo;not null"`
	MEndDate             time.Time     `json:"PcInv_mEndDate" gorm:"column:PcInv_mEndDate;not null"`
	MIsConfirm           bool          `json:"PcInv_mIsConfirm" gorm:"column:PcInv_mIsConfirm;not null"`
	MStatus              int8          `json:"PcInv_mStatus" gorm:"column:PcInv_mStatus;not null"`
	MCnt                 sql.NullInt64 `json:"PcInv_mCnt" gorm:"column:PcInv_mCnt;not null"`
	MCntUse              int16         `json:"PcInv_mCntUse" gorm:"column:PcInv_mCntUse"`
	MIsSeizure           bool          `json:"PcInv_mIsSeizure" gorm:"column:PcInv_mIsSeizure;not null"`
	MApplyAbnItemNo      sql.NullInt64 `json:"PcInv_mApplyAbnItemNo" gorm:"column:PcInv_mApplyAbnItemNo;not null"`
	MApplyAbnItemEndDate time.Time     `json:"PcInv_mApplyAbnItemEndDate" gorm:"column:PcInv_mApplyAbnItemEndDate;not null"`
	MOwner               sql.NullInt64 `json:"PcInv_mOwner" gorm:"column:PcInv_mOwner;not null"`
	MPracticalPeriod     sql.NullInt64 `json:"PcInv_mPracticalPeriod" gorm:"column:PcInv_mPracticalPeriod;not null"`
	MBindingType         int8          `json:"PcInv_mBindingType" gorm:"column:PcInv_mBindingType;not null"`
	MRestoreCnt          int8          `json:"PcInv_mRestoreCnt" gorm:"column:PcInv_mRestoreCnt;not null"`
	MHoleCount           int8          `json:"PcInv_mHoleCount" gorm:"column:PcInv_mHoleCount;not null"`
}

func (IntermediatePcInventory) TableName() string {
	return "TblPcInventory"
}
