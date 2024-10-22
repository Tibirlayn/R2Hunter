package parm

type MaterialDraw []struct {
	// MaterialDrawIndex
	MDIDIndex        int64   `json:"MDIDIndex" gorm:"column:MDID; not null"`
	MDRDIndex        int64   `json:"MDRDIndex" gorm:"column:MDRD; not null"`
	MResType         int     `json:"MResType" gorm:"column:mResType; not null"`
	MMaxResCnt       int     `json:"MMaxResCnt" gorm:"column:mMaxResCnt; not null"`
	MSuccess         float64 `json:"MSuccess" gorm:"column:mSuccess"`
	MDesc            string  `json:"MDesc" gorm:"column:mDesc; not null; size:500"`
	MAddQuestionMark int16   `json:"MAddQuestionMark" gorm:"column:mAddQuestionMark; not null"`
	MDescRus         string  `json:"MDescRus" gorm:"column:mDescRus; size:2000"`

	// MaterialDrawMaterial
	MSeqMaterial int   `json:"MSeqMaterial" gorm:"column:mSeq; not null"`
	MDIDMaterial int64 `json:"MDIDMaterial" gorm:"column:MDID; not null"`
	IIDMaterial  int   `json:"IIDMaterial" gorm:"column:IID; not null"`
	MCntMaterial int   `json:"MCntMaterial" gorm:"column:mCnt; not null"`

	// MaterialDrawResult
	MSeqResult  int     `json:"MSeqResult" gorm:"column:mSeq; not null"`
	MDRDResult  int64   `json:"MDRDResult" gorm:"column:MDRD; not null"`
	IIDResult   int     `json:"IIDResult" gorm:"column:IID; not null"`
	MPerOrRate  float64 `json:"MPerOrRate" gorm:"column:mPerOrRate"`
	MItemStatus int8    `json:"MItemStatus" gorm:"column:mItemStatus; not null"`
	MCntResult  int     `json:"MCntResult" gorm:"column:mCnt; not null"`
	MBinding    string  `json:"MBinding" gorm:"column:mBinding; not null"`
	MEffTime    int     `json:"MEffTime" gorm:"column:mEffTime; not null"`
	MValTime    int16   `json:"MValTime" gorm:"column:mValTime; not null"`
	MResource   int     `json:"MResource" gorm:"column:mResource; not null"`
	MAddGroup   int8    `json:"MAddGroup" gorm:"column:mAddGroup; not null"`

	// Item
	IName string `json:"IName" gorm:"column:IName;size:40"`
	INameRes string `json:"INameRes" gorm:"column:IName;size:40"`

	// ItemResource
	RIDRes       int    `json:"RIDRes" gorm:"column:RID;not null;primaryKey"`
	ROwnerIDRes  int    `json:"ROwnerIDRes" gorm:"column:ROwnerID"`
	RTypeRes     int    `json:"RTypeRes" gorm:"column:RType"`
	RFileNameRes string `json:"RFileNameRes" gorm:"column:RFileName"`
	RPosXRes     int    `json:"RPosXRes" gorm:"column:RPosX"`
	RPosYRes     int    `json:"RPosYRes" gorm:"column:RPosY"`

	RIDMat       int    `json:"RIDMat" gorm:"column:RID;not null;primaryKey"`
	ROwnerIDMat  int    `json:"ROwnerIDMat" gorm:"column:ROwnerID"`
	RTypeMat     int    `json:"RTypeMat" gorm:"column:RType"`
	RFileNameMat string `json:"RFileNameMat" gorm:"column:RFileName"`
	RPosXMat     int    `json:"RPosXMat" gorm:"column:RPosX"`
	RPosYMat     int    `json:"RPosYMat" gorm:"column:RPosY"`
}
