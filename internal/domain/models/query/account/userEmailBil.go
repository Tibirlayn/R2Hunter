package query

type UserEmailBil struct {
	MUserNo              int       `json:"mUserNo" gorm:"column:mUserNo;primaryKey;default:0"`                         // Идентификатор пользователя
	MUserId              string    `json:"mUserId" gorm:"column:mUserId;default:'null'"`                                    // Уникальный номер пользователя
}

func (UserEmailBil) TableName() string {
	return "TblUser"
}
