package billing

import "time"

type ShopItemRes struct {
	GoldItemID        int64
	GIID              int
	ItemName          string
	ItemImage         string
	ItemDesc          string
	OriginalGoldPrice int
	GoldPrice         int
	ItemCategory      int16
	IsPackage         string
	Status            string
	AvailablePeriod   int
	Count             int
	PracticalPeriod   int
	RegistDate        time.Time
	RegistAdmin       string
	RegistIP          string
	UpdateDate        string
	UpdateAdmin       string
	UpdateIP          string
	ItemNameRUS       string
	ItemDescRUS       string
	IID               int
	IName             string
	RFileName         string
	RPosX             int
	RPosY             int
}
