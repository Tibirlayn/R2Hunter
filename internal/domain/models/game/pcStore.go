package game

import "time"

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
