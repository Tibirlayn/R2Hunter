package billing

import "time"

// SysOrderList And ItemRes
type SolItemRes struct {
	MRegDate         time.Time
	MSysOrderID      int64
	MSysID           int64
	MUserNo          int
	MSvrNo           int16
	MItemID          int
	MCnt             int
	MAvailablePeriod int
	MPracticalPeriod int
	MStatus          uint8
	MReceiptDate     time.Time
	MReceiptPcNo     int
	MRecepitPcNm     string
	MBindingType     uint8
	MLimitedDate     time.Time
	MItemStatus      uint8
	IID              int
	IName            string
	RFileName        string
	RPosX            int
	RPosY            int
}
