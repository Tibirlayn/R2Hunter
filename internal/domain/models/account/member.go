package account

import "database/sql"

// Member
// "mUserId", "mUserPswd", "superpwd", "cash", "email", "tgzh", "uid", "klq", "ylq", "auth", "mSum", "isAdmin" , "isdl", "dlmoney", "registerIp", "country", "cashBack"
type Member struct {
	MUserId    string  `json:"mUserId" gorm:"column:mUserId;size:50;not null;"`
	MUserPswd  string  `json:"mUserPswd" gorm:"column:mUserPswd;size:50;default:'null'"`
	Superpwd   string  `json:"superpwd" gorm:"column:superpwd;size:50;default:'null'"`
	Cash       float64 `json:"cash" gorm:"column:cash;default:0"`
	Email      string  `json:"email" gorm:"column:email;size:255;default:'null'" validate:"required,min=3,max=25"`
	Tgzh       string  `json:"tgzh" gorm:"column:tgzh;size:255;default:'null'"`
	Uid        int     `json:"uid" gorm:"column:uid;not null"`
	Klq        int     `json:"klq" gorm:"column:klq;default:0"`
	Ylq        int     `json:"ylq" gorm:"column:ylq;default:0"`
	Auth       int     `json:"auth" gorm:"column:auth;default:0"`
	MSum       string  `json:"mSum" gorm:"column:mSum;size:255;default:'null'"`
	IsAdmin    int     `json:"isAdmin" gorm:"column:isAdmin;default:0"`
	Isdl       int     `json:"isdl" gorm:"column:isdl;default:0"`
	Dlmoney    int     `json:"dlmoney" gorm:"column:dlmoney;default:0"`
	RegisterIp string  `json:"registerIp" gorm:"column:registerIp;size:255;default:'null'"`
	Country    string  `json:"country" gorm:"column:country;size:20;default:'null'"`
	CashBack   int     `json:"cashBack" gorm:"column:cashBack;default:0"`
}

func (Member) TableName() string {
	return "Member"
}


// алиас
type IntermediateMember struct {
	MUserId    sql.NullString  `json:"Member_mUserId" gorm:"column:Member_mUserId;size:50;not null;primaryKey;"` // mUserId
	MUserPswd  sql.NullString  `json:"Member_mUserPswd" gorm:"column:Member_mUserPswd;size:50"`                 // mUserPswd
	Superpwd   sql.NullString  `json:"superpwd" gorm:"column:superpwd;size:50"`
	Cash       sql.NullFloat64 `json:"cash" gorm:"column:cash"`
	Email      sql.NullString  `json:"Member_Email" gorm:"column:Member_Email;size:255" validate:"required,min=3,max=25"` // email
	Tgzh       sql.NullString  `json:"tgzh" gorm:"column:tgzh;size:255"`
	Uid        sql.NullInt64   `json:"uid" gorm:"column:uid;not null"`
	Klq        sql.NullInt64   `json:"klq" gorm:"column:klq"`
	Ylq        sql.NullInt64   `json:"ylq" gorm:"column:ylq"`
	Auth       sql.NullInt64   `json:"auth" gorm:"column:auth"`
	MSum       sql.NullString  `json:"mSum" gorm:"column:mSum;size:255"`
	IsAdmin    sql.NullInt64   `json:"isAdmin" gorm:"column:isAdmin"`
	Isdl       sql.NullInt64   `json:"isdl" gorm:"column:isdl"`
	Dlmoney    sql.NullInt64   `json:"dlmoney" gorm:"column:dlmoney"`
	RegisterIp sql.NullString  `json:"registerIp" gorm:"column:registerIp;size:255"`
	Country    sql.NullString  `json:"country" gorm:"column:country;size:20"`
	CashBack   sql.NullInt64   `json:"cashBack" gorm:"column:cashBack"`
}

func (IntermediateMember) TableName() string {
	return "Member"
}
