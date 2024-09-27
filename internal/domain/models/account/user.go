package account

import (
	"database/sql"
	"time"
)

// TblUser
// "mRegDate", "mUserAuth", "mUserNo", "mUserId", "mUserPswd", "mCertifiedKey", "mIp", "mLoginTm", "mLogoutTm", "mTotUseTm", "mWorldNo", "mDelDate", "mPcBangLv", "mSecKeyTableUse", "mUseMacro", "mIpEx", "mJoinCode", "mLoginChannelID", "mTired", "mChnSID", "mNewId", "mLoginSvrType", "mAccountGuid", "mNormalLimitTime", "mPcBangLimitTime", "mRegIp", "mIsMovingToBattleSvr"
type User struct {
	MRegDate             time.Time `json:"mRegDate" gorm:"column:mRegDate"`                                  // Дата регистрации
	MUserAuth            uint8     `json:"mUserAuth" gorm:"column:mUserAuth;default:0"`                                // Авторизация пользователя
	MUserNo              int       `json:"mUserNo" gorm:"column:mUserNo;primaryKey;default:0"`                         // Идентификатор пользователя
	MUserId              string    `json:"mUserId" gorm:"column:mUserId;default:'null'"`                                    // Уникальный номер пользователя
	MUserPswd            string    `json:"mUserPswd" gorm:"column:mUserPswd;default:'null'"`                                // Пароль пользователя
	MCertifiedKey        int       `json:"mCertifiedKey" gorm:"column:mCertifiedKey;default:0"`                        // Сертифицированный ключ
	MIp                  string    `json:"mIp" gorm:"column:mIp;default:'null'"`                                            // IP-адрес
	MLoginTm             time.Time `json:"mLoginTm" gorm:"not null;column:mLoginTm;default:1900-01-01 00:00:00.000"`                         // Время входа
	MLogoutTm            time.Time `json:"mLogoutTm" gorm:"not null;column:mLogoutTm;default:1900-01-01 00:00:00.000"`                       // Время выхода
	MTotUseTm            int       `json:"mTotUseTm" gorm:"column:mTotUseTm;default:0"`                                // Общее время использования
	MWorldNo             int16     `json:"mWorldNo" gorm:"column:mWorldNo;default:0"`                                  // Номер мира
	MDelDate             time.Time `json:"mDelDate" gorm:"column:mDelDate;default:1900-01-01 00:00:00.000"`                                  // Дата удаления
	MPcBangLv            int       `json:"mPcBangLv" gorm:"column:mPcBangLv;default:0"`                                // Уровень PcBang
	MSecKeyTableUse      uint8     `json:"mSecKeyTableUse" gorm:"column:mSecKeyTableUse;default:0"`                    // Использование таблицы SecKey
	MUseMacro            int16     `json:"mUseMacro" gorm:"column:mUseMacro;default:0"`                                // Использование макроса
	MIpEX                int64     `json:"mIpEx" gorm:"column:mIpEx;default:0"`                                        // Дополнительный IP-адрес
	MJoinCode            string    `json:"mJoinCode" gorm:"column:mJoinCode;default:'null'"`                                // Код приглашения
	MLoginChannelID      string    `json:"mLoginChannelID" gorm:"column:mLoginChannelID;default:'null'"`                    // Идентификатор канала входа
	MTired               string    `json:"mTired" gorm:"column:mTired;default:'null'"`                                      // Усталость
	MChnSID              string    `json:"mChnSID" gorm:"column:mChnSID;default:'null'"`                                    // SID канала
	MNewId               bool      `json:"mNewId" gorm:"column:mNewId;not null;default:0"`                                      // Новый идентификатор
	MLoginSvrType        uint8     `json:"mLoginSvrType" gorm:"column:mLoginSvrType;default:0"`                        // Тип сервера входа
	MAccountGuid         int       `json:"mAccountGuid" gorm:"column:mAccountGuid;default:0"`                          // GUID аккаунта
	MNormalLimitTime     int       `json:"mNormalLimitTime" gorm:"column:mNormalLimitTime;default:0"`                  // Лимит времени нормального использования
	MPcBangLimitTime     int       `json:"mPcBangLimitTime" gorm:"column:mPcBangLimitTime;default:0"`                  // Лимит времени использования PcBang
	MRegIp               string    `json:"mRegIp" gorm:"column:mRegIp;default:'null'"`                                      // IP-адрес при регистрации
	MIsMovingToBattleSvr bool      `json:"mIsMovingToBattleSvr" gorm:"column:mIsMovingToBattleSvr;not null;"` // Перемещается ли на боевой сервер
}

func (User) TableName() string {
	return "TblUser"
}

// алиас
type IntermediateUser struct {
	MRegDate             time.Time      `json:"User_mRegDate" gorm:"column:User_mRegDate"` // mRegDate
	MUserAuth            uint8          `json:"mUserAuth" gorm:"column:mUserAuth"`
	MUserNo              sql.NullInt64  `json:"User_mUserNo" gorm:"column:User_mUserNo;primaryKey"` // mUserNo
	MUserId              sql.NullString `json:"User_mUserId" gorm:"column:User_mUserId"`            // mUserId
	MUserPswd            sql.NullString `json:"User_mUserPswd" gorm:"column:User_mUserPswd"`        // mUserPswd
	MCertifiedKey        sql.NullInt64  `json:"mCertifiedKey" gorm:"column:mCertifiedKey"`
	MIp                  sql.NullString `json:"User_mIp" gorm:"column:User_mIp"`                    // mIp
	MLoginTm             time.Time      `json:"User_mLoginTm" gorm:"not null;column:User_mLoginTm"` // mLoginTm
	MLogoutTm            time.Time      `json:"mLogoutTm" gorm:"not null;column:mLogoutTm"`
	MTotUseTm            sql.NullInt64  `json:"mTotUseTm" gorm:"column:mTotUseTm"`
	MWorldNo             int16          `json:"mWorldNo" gorm:"column:mWorldNo"`
	MDelDate             time.Time      `json:"User_mDelDate" gorm:"column:User_mDelDate"` // mDelDate
	MPcBangLv            sql.NullInt64  `json:"mPcBangLv" gorm:"column:mPcBangLv"`
	MSecKeyTableUse      uint8          `json:"mSecKeyTableUse" gorm:"column:mSecKeyTableUse"`
	MUseMacro            int16          `json:"mUseMacro" gorm:"column:mUseMacro"`
	MIpEX                sql.NullInt64  `json:"mIpEx" gorm:"column:mIpEx"`
	MJoinCode            sql.NullString `json:"mJoinCode" gorm:"column:mJoinCode"`
	MLoginChannelID      sql.NullString `json:"mLoginChannelID" gorm:"column:mLoginChannelID"`
	MTired               sql.NullString `json:"mTired" gorm:"column:mTired"`
	MChnSID              sql.NullString `json:"mChnSID" gorm:"column:mChnSID"`
	MNewId               sql.NullBool   `json:"mNewId" gorm:"column:mNewId"`
	MLoginSvrType        uint8          `json:"mLoginSvrType" gorm:"column:mLoginSvrType"`
	MAccountGuid         sql.NullInt64  `json:"mAccountGuid" gorm:"column:mAccountGuid"`
	MNormalLimitTime     sql.NullInt64  `json:"mNormalLimitTime" gorm:"column:mNormalLimitTime"`
	MPcBangLimitTime     sql.NullInt64  `json:"mPcBangLimitTime" gorm:"column:mPcBangLimitTime"`
	MRegIp               sql.NullString `json:"mRegIp" gorm:"column:mRegIp"`
	MIsMovingToBattleSvr sql.NullBool   `json:"mIsMovingToBattleSvr" gorm:"not null;column:mIsMovingToBattleSvr"`
}

func (IntermediateUser) TableName() string {
	return "TblUser"
}
