package game

import (
	"database/sql"
	"time"
)

type PcState struct {
	MNo                int       `json:"mNo" gorm:"column:mNo;not null;primaryKey"`
	MLevel             int16     `json:"mLevel" gorm:"column:mLevel;not null"`
	MExp               int64     `json:"mExp" gorm:"column:mExp;not null"`
	MHpAdd             int       `json:"mHpAdd" gorm:"column:mHpAdd;not null"`
	MHp                int       `json:"mHp" gorm:"column:mHp;not null"`
	MMpAdd             int       `json:"mMpAdd" gorm:"column:mMpAdd;not null"`
	MMp                int       `json:"mMp" gorm:"column:mMp;not null"`
	MMapNo             int       `json:"mMapNo" gorm:"column:mMapNo;not null"`
	MPosX              float32   `json:"mPosX" gorm:"column:mPosX;not null"`
	MPosY              float32   `json:"mPosY" gorm:"column:mPosY;not null"`
	MPosZ              float32   `json:"mPosZ" gorm:"column:mPosZ;not null"`
	MStomach           int16     `json:"mStomach" gorm:"column:mStomach;not null"`
	MIp                string    `json:"mIp" gorm:"column:mIp"`
	MLoginTm           time.Time `json:"mLoginTm" gorm:"column:mLoginTm"`
	MLogoutTm          time.Time `json:"mLogoutTm" gorm:"column:mLogoutTm"`
	MTotUseTm          int       `json:"mTotUseTm" gorm:"column:mTotUseTm;not null"`
	MPkCnt             int       `json:"mPkCnt" gorm:"column:mPkCnt;not null"`
	MChaotic           int       `json:"mChaotic" gorm:"column:mChaotic;not null"`
	MDiscipleJoinCount int       `json:"mDiscipleJoinCount" gorm:"column:mDiscipleJoinCount;not null"`
	MPartyMemCntLevel  int       `json:"mPartyMemCntLevel" gorm:"column:mPartyMemCntLevel;not null"`
	MLostExp           int64     `json:"mLostExp" gorm:"column:mLostExp;not null"`
	MIsLetterLimit     bool      `json:"mIsLetterLimit" gorm:"column:mIsLetterLimit;not null"`
	MFlag              int16     `json:"mFlag" gorm:"column:mFlag;not null"`
	MIsPreventItemDrop bool      `json:"mIsPreventItemDrop" gorm:"column:mIsPreventItemDrop"`
	MSkillTreePoint    int16     `json:"mSkillTreePoint" gorm:"column:mSkillTreePoint;not null"`
	MRestExpGuild      int64     `json:"mRestExpGuild" gorm:"column:mRestExpGuild;not null"`
	MRestExpActivate   int64     `json:"mRestExpActivate" gorm:"column:mRestExpActivate;not null"`
	MRestExpDeactivate int64     `json:"mRestExpDeactivate" gorm:"column:mRestExpDeactivate;not null"`
	MQMCnt             int16     `json:"mQMCnt" gorm:"column:mQMCnt;not null"`
	MGuildQMCnt        int16     `json:"mGuildQMCnt" gorm:"column:mGuildQMCnt;not null"`
	MFierceCnt         int16     `json:"mFierceCnt" gorm:"column:mFierceCnt;not null"`
	MBossCnt           int16     `json:"mBossCnt" gorm:"column:mBossCnt;not null"`

	// Pc Pc `gorm:"foreignKey:MPcNo;references:MNo"` // Добавлено для явного обозначения связи, опционально
}

func (PcState) TableName() string {
	return "TblPcState"
}

type IntermediatePcState struct {
	MNo                sql.NullInt64  `json:"PcState_mNo" gorm:"column:PcState_mNo;not null;primaryKey"`
	MLevel             int16          `json:"mLevel" gorm:"column:mLevel;not null"`
	MExp               sql.NullInt64  `json:"mExp" gorm:"column:mExp;not null"`
	MHpAdd             sql.NullInt64  `json:"mHpAdd" gorm:"column:mHpAdd;not null"`
	MHp                sql.NullInt64  `json:"mHp" gorm:"column:mHp;not null"`
	MMpAdd             sql.NullInt64  `json:"mMpAdd" gorm:"column:mMpAdd;not null"`
	MMp                sql.NullInt64  `json:"mMp" gorm:"column:mMp;not null"`
	MMapNo             sql.NullInt64  `json:"mMapNo" gorm:"column:mMapNo;not null"`
	MPosX              float32        `json:"mPosX" gorm:"column:mPosX;not null"`
	MPosY              float32        `json:"mPosY" gorm:"column:mPosY;not null"`
	MPosZ              float32        `json:"mPosZ" gorm:"column:mPosZ;not null"`
	MStomach           int16          `json:"mStomach" gorm:"column:mStomach;not null"`
	MIp                sql.NullString `json:"PcState_mIp" gorm:"column:PcState_mIp"`
	MLoginTm           sql.NullTime   `json:"PcState_mLoginTm" gorm:"column:PcState_mLoginTm"`
	MLogoutTm          sql.NullTime   `json:"PcState_mLogoutTm" gorm:"column:PcState_mLogoutTm"`
	MTotUseTm          sql.NullInt64  `json:"PcState_mTotUseTm" gorm:"column:PcState_mTotUseTm;not null"`
	MPkCnt             sql.NullInt64  `json:"mPkCnt" gorm:"column:mPkCnt;not null"`
	MChaotic           sql.NullInt64  `json:"mChaotic" gorm:"column:mChaotic;not null"`
	MDiscipleJoinCount sql.NullInt64  `json:"mDiscipleJoinCount" gorm:"column:mDiscipleJoinCount;not null"`
	MPartyMemCntLevel  sql.NullInt64  `json:"mPartyMemCntLevel" gorm:"column:mPartyMemCntLevel;not null"`
	MLostExp           sql.NullInt64  `json:"mLostExp" gorm:"column:mLostExp;not null"`
	MIsLetterLimit     sql.NullBool   `json:"mIsLetterLimit" gorm:"column:mIsLetterLimit;not null"`
	MFlag              int16          `json:"mFlag" gorm:"column:mFlag;not null"`
	MIsPreventItemDrop sql.NullBool   `json:"mIsPreventItemDrop" gorm:"column:mIsPreventItemDrop"`
	MSkillTreePoint    int16          `json:"mSkillTreePoint" gorm:"column:mSkillTreePoint;not null"`
	MRestExpGuild      sql.NullInt64  `json:"mRestExpGuild" gorm:"column:mRestExpGuild;not null"`
	MRestExpActivate   sql.NullInt64  `json:"mRestExpActivate" gorm:"column:mRestExpActivate;not null"`
	MRestExpDeactivate sql.NullInt64  `json:"mRestExpDeactivate" gorm:"column:mRestExpDeactivate;not null"`
	MQMCnt             int16          `json:"mQMCnt" gorm:"column:mQMCnt;not null"`
	MGuildQMCnt        int16          `json:"mGuildQMCnt" gorm:"column:mGuildQMCnt;not null"`
	MFierceCnt         int16          `json:"mFierceCnt" gorm:"column:mFierceCnt;not null"`
	MBossCnt           int16          `json:"mBossCnt" gorm:"column:mBossCnt;not null"`
}

func (IntermediatePcState) TableName() string {
	return "TblPcState"
}
