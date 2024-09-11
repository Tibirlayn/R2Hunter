package account

import (
	"database/sql"
	"time"
)

// TblUser
// "mRegDate", "mUserAuth", "mUserNo", "mUserId", "mUserPswd", "mCertifiedKey", "mIp", "mLoginTm", "mLogoutTm", "mTotUseTm", "mWorldNo", "mDelDate", "mPcBangLv", "mSecKeyTableUse", "mUseMacro", "mIpEx", "mJoinCode", "mLoginChannelID", "mTired", "mChnSID", "mNewId", "mLoginSvrType", "mAccountGuid", "mNormalLimitTime", "mPcBangLimitTime", "mRegIp", "mIsMovingToBattleSvr"
type User struct {
	MRegDate             time.Time `json:"mRegDate" gorm:"column:mRegDate"`                                  // Дата регистрации
	MUserAuth            uint8     `json:"mUserAuth" gorm:"column:mUserAuth"`                                // Авторизация пользователя
	MUserNo              int       `json:"mUserNo" gorm:"column:mUserNo;primaryKey"`                         // Идентификатор пользователя
	MUserId              string    `json:"mUserId" gorm:"column:mUserId"`                                    // Уникальный номер пользователя
	MUserPswd            string    `json:"mUserPswd" gorm:"column:mUserPswd"`                                // Пароль пользователя
	MCertifiedKey        int       `json:"mCertifiedKey" gorm:"column:mCertifiedKey"`                        // Сертифицированный ключ
	MIp                  string    `json:"mIp" gorm:"column:mIp"`                                            // IP-адрес
	MLoginTm             time.Time `json:"mLoginTm" gorm:"not null;column:mLoginTm"`                         // Время входа
	MLogoutTm            time.Time `json:"mLogoutTm" gorm:"not null;column:mLogoutTm"`                       // Время выхода
	MTotUseTm            int       `json:"mTotUseTm" gorm:"column:mTotUseTm"`                                // Общее время использования
	MWorldNo             int16     `json:"mWorldNo" gorm:"column:mWorldNo"`                                  // Номер мира
	MDelDate             time.Time `json:"mDelDate" gorm:"column:mDelDate"`                                  // Дата удаления
	MPcBangLv            int       `json:"mPcBangLv" gorm:"column:mPcBangLv"`                                // Уровень PcBang
	MSecKeyTableUse      uint8     `json:"mSecKeyTableUse" gorm:"column:mSecKeyTableUse"`                    // Использование таблицы SecKey
	MUseMacro            int16     `json:"mUseMacro" gorm:"column:mUseMacro"`                                // Использование макроса
	MIpEX                int64     `json:"mIpEx" gorm:"column:mIpEx"`                                        // Дополнительный IP-адрес
	MJoinCode            string    `json:"mJoinCode" gorm:"column:mJoinCode"`                                // Код приглашения
	MLoginChannelID      string    `json:"mLoginChannelID" gorm:"column:mLoginChannelID"`                    // Идентификатор канала входа
	MTired               string    `json:"mTired" gorm:"column:mTired"`                                      // Усталость
	MChnSID              string    `json:"mChnSID" gorm:"column:mChnSID"`                                    // SID канала
	MNewId               bool      `json:"mNewId" gorm:"column:mNewId"`                                      // Новый идентификатор
	MLoginSvrType        uint8     `json:"mLoginSvrType" gorm:"column:mLoginSvrType"`                        // Тип сервера входа
	MAccountGuid         int       `json:"mAccountGuid" gorm:"column:mAccountGuid"`                          // GUID аккаунта
	MNormalLimitTime     int       `json:"mNormalLimitTime" gorm:"column:mNormalLimitTime"`                  // Лимит времени нормального использования
	MPcBangLimitTime     int       `json:"mPcBangLimitTime" gorm:"column:mPcBangLimitTime"`                  // Лимит времени использования PcBang
	MRegIp               string    `json:"mRegIp" gorm:"column:mRegIp"`                                      // IP-адрес при регистрации
	MIsMovingToBattleSvr bool      `json:"mIsMovingToBattleSvr" gorm:"not null;column:mIsMovingToBattleSvr"` // Перемещается ли на боевой сервер
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
