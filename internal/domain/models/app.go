package models

type App struct {
	ID     int    `json:"id" gorm:"column:id;not null;primaryKey"`
	Name   string `json:"name" gorm:"column:name;size:255;not null"`
	Secret string `json:"secret" gorm:"column:secret;size:255;not null"`
}

func (App) TableName() string {
	return "apps"
}
