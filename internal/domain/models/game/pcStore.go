package game

import (
	"database/sql"
	"time"
)

// TblPcStore
type PcStore struct {
	MRegDate             time.Time `json:"mRegDate" gorm:"column:mRegDate;not null"`
	MSerialNo            int64     `json:"mSerialNo" grom:"column:mSerialNo;not null;primaryKey"`
	MUserNo              int       `json:"mUserNo" gorm:"column:mUserNo;not null"`
	MItemNo              int       `json:"mItemNo" gorm:"column:mItemNo;not null"`
	MEndDate             time.Time `json:"mEndDate" gorm:"column:mEndDate;not null"`
	MIsConfirm           bool      `json:"mIsConfirm" gorm:"column:mIsConfirm;not null"`
	MStatus              int8      `json:"mStatus" gorm:"column:mStatus;not null"`
	MCnt                 int       `json:"mCnt" gorm:"column:mCnt"`
	MCntUse              int16     `json:"mCntUse" gorm:"column:mCntUse;not null"`
	MIsSeizure           bool      `json:"mIsSeizure" gorm:"column:mIsSeizure;not null"`
	MApplyAbnItemNo      int       `json:"mApplyAbnItemNo" gorm:"column:mApplyAbnItemNo;not null"`
	MApplyAbnItemEndDate time.Time `json:"mApplyAbnItemEndDate" gorm:"column:mApplyAbnItemEndDate;not null"`
	MOwner               int       `json:"mOwner" gorm:"column:mOwner;not null"`
	MPracticalPeriod     int       `json:"mPracticalPeriod" gorm:"column:mPracticalPeriod;not null"`
	MBindingType         int8      `json:"mBindingType" gorm:"column:mBindingType;not null"`
	MRestoreCnt          int8      `json:"mRestoreCnt" gorm:"column:mRestoreCnt;not null"`
	MHoleCount           int8      `json:"mHoleCount" gorm:"column:mHoleCount;not null"`
}

func (PcStore) TableName() string {
	return "TblPcStore"
}

type IntermediatePcStore struct {
	MRegDate             sql.NullTime  `json:"PcStore_mRegDate" gorm:"column:PcStore_mRegDate;not null"`
	MSerialNo            sql.NullInt64 `json:"PcStore_mSerialNo" grom:"column:PcStore_mSerialNo;not null;primaryKey"`
	MUserNo              sql.NullInt64 `json:"PcStore_mUserNo" gorm:"column:PcStore_mUserNo;not null"`
	MItemNo              sql.NullInt64 `json:"PcStore_mItemNo" gorm:"column:PcStore_mItemNo;not null"`
	MEndDate             sql.NullTime  `json:"PcStore_mEndDate" gorm:"column:PcStore_mEndDate;not null"`
	MIsConfirm           sql.NullBool  `json:"PcStore_mIsConfirm" gorm:"column:PcStore_mIsConfirm;not null"`
	MStatus              sql.NullInt16 `json:"PcStore_mStatus" gorm:"column:PcStore_mStatus;not null"`
	MCnt                 sql.NullInt64 `json:"PcStore_mCnt" gorm:"column:PcStore_mCnt"`
	MCntUse              sql.NullInt16 `json:"PcStore_mCntUse" gorm:"column:PcStore_mCntUse;not null"`
	MIsSeizure           sql.NullBool  `json:"PcStore_mIsSeizure" gorm:"column:PcStore_mIsSeizure;not null"`
	MApplyAbnItemNo      sql.NullInt64 `json:"PcStore_mApplyAbnItemNo" gorm:"column:PcStore_mApplyAbnItemNo;not null"`
	MApplyAbnItemEndDate sql.NullTime  `json:"PcStore_mApplyAbnItemEndDate" gorm:"column:PcStore_mApplyAbnItemEndDate;not null"`
	MOwner               sql.NullInt64 `json:"PcStore_mOwner" gorm:"column:PcStore_mOwner;not null"`
	MPracticalPeriod     sql.NullInt64 `json:"PcStore_mPracticalPeriod" gorm:"column:PcStore_mPracticalPeriod;not null"`
	MBindingType         sql.NullInt16          `json:"PcStore_mBindingType" gorm:"column:PcStore_mBindingType;not null"`
	MRestoreCnt          sql.NullInt16          `json:"PcStore_mRestoreCnt" gorm:"column:PcStore_mRestoreCnt;not null"`
	MHoleCount           sql.NullInt16          `json:"PcStore_mHoleCount" gorm:"column:PcStore_mHoleCount;not null"`
}

func (IntermediatePcStore) TableName() string {
	return "TblPcStore"
}
