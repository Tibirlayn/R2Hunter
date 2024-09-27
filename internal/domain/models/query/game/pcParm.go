package game

import "github.com/Tibirlayn/R2Hunter/internal/domain/models/game"

type PcParm struct {
	Pc      game.Pc
	PcInv   game.PcInventory
	PcState game.PcState
	PcStore game.PcStore
}

type PcTopLVL struct {
	ID      int    `json:"id" gorm:"column:mNo"`
	Class   string `json:"class" gorm:"column:mClass"`
	Name    string `json:"name" gorm:"column:mNm"`
	Level   int    `json:"level" gorm:"column:mLevel"`
	Chaotic int    `json:"chaotic" gorm:"column:mChaotic"`
	PkCnt   int    `json:"pkCnt" gorm:"column:mPkCnt"`
}

type PcTopByGold struct {
	MOwner    int    `json:"id_account" gorm:"column:mOwner"` // аккаунт персонажа
	MSerialNo int64  `json:"serial_no" gorm:"column:mSerialNo"`
	MPcNo     int    `json:"id_pc" gorm:"column:mPcNo"`   // id персонажа
	MItemNo   int    `json:"item" gorm:"column:mItemNo"` // id предмета
	Name      string `json:"name" gorm:"column:mNm"`     // имя персонажа
	MCnt      int    `json:"Cnt" gorm:"column:mCnt"`    // кол-во
}
