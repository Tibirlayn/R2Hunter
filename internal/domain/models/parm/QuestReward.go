package parm

type QuestReward struct {
	MRewardNo int   `json:"mRewardNo" gorm:"column:mRewardNo; not null"`
	MExp      int64 `json:"MExp" gorm:"column:MExp; not null"`
	MID       int   `json:"mID" gorm:"column:mID; not null"`
	MCnt      int   `json:"mCnt" gorm:"column:mCnt; not null"`
	MBinding  int8  `json:"mBinding" gorm:"column:mBinding; not null"`
	MStatus   int8  `json:"mStatus" gorm:"column:mStatus; not null"`
	MEffTime  int   `json:"mEffTime" gorm:"column:mEffTime; not null"`
	MValTime  int   `json:"mValTime" gorm:"column:mValTime; not null"`
}

func (QuestReward) TableName() string {
	return "TblQuestReward"
}
