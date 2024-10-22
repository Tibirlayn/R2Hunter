package parm

import (
    "time"
)

// ServerInfo представляет информацию о сервере
type ParmSvr struct {
    MRegDate        time.Time `json:"mRegDate" gorm:"column:mRegDate"`                  // Дата регистрации
    MIsValid        bool      `json:"mIsValid" gorm:"column:mIsValid"`                  // Флаг валидности (bit)
    MDispOrder      int16     `json:"mDispOrder" gorm:"column:mDispOrder"`              // Порядок отображения (smallint)
    MSvrNo          int16     `json:"mSvrNo" gorm:"column:mSvrNo"`                      // Номер сервера (smallint)
    MType           uint8     `json:"mType" gorm:"column:mType"`                        // Тип сервера (tinyint)
    MMajorIp        string    `json:"mMajorIp" gorm:"column:mMajorIp"`                  // Основной IP (char)
    MMinorIp        string    `json:"mMinorIp" gorm:"column:mMinorIp"`                  // Вторичный IP (char)
    MWorldNo        int16     `json:"mWorldNo" gorm:"column:mWorldNo"`                  // Номер мира (smallint)
    MMajorVer       int       `json:"mMajorVer" gorm:"column:mMajorVer"`                // Основная версия (int)
    MMinorVer       int       `json:"mMinorVer" gorm:"column:mMinorVer"`                // Вспомогательная версия (int)
    MThdWkCnt       int16     `json:"mThdWkCnt" gorm:"column:mThdWkCnt"`                // Количество рабочих потоков (smallint)
    MThdTmCnt       int16     `json:"mThdTmCnt" gorm:"column:mThdTmCnt"`                // Количество потоков по времени (smallint)
    MThdDbCnt       int16     `json:"mThdDbCnt" gorm:"column:mThdDbCnt"`                // Количество потоков для БД (smallint)
    MThdLogCnt      int16     `json:"mThdLogCnt" gorm:"column:mThdLogCnt"`              // Количество потоков для логов (smallint)
    MSessionCnt     int16     `json:"mSessionCnt" gorm:"column:mSessionCnt"`            // Количество сессий (smallint)
    MSendCnt        int       `json:"mSendCnt" gorm:"column:mSendCnt"`                  // Количество отправок (int)
    MTcpPort        int16     `json:"mTcpPort" gorm:"column:mTcpPort"`                  // TCP порт (smallint)
    MUdpPort        int16     `json:"mUdpPort" gorm:"column:mUdpPort"`                  // UDP порт (smallint)
    MDesc           string    `json:"mDesc" gorm:"column:mDesc"`                        // Описание (char)
    MIsSiege        bool      `json:"mIsSiege" gorm:"column:mIsSiege"`                  // Флаг осады (bit)
    MEvtStx         time.Time `json:"mEvtStx" gorm:"column:mEvtStx"`                    // Время начала события (datetime)
    MEvtEtx         time.Time `json:"mEvtEtx" gorm:"column:mEvtEtx"`                    // Время окончания события (datetime)
    MIsCheckRSC     bool      `json:"mIsCheckRSC" gorm:"column:mIsCheckRSC"`            // Проверка RSC (bit)
    MNationID       uint8     `json:"mNationID" gorm:"column:mNationID"`                // Идентификатор нации (tinyint)
    MBullet         int       `json:"mBullet" gorm:"column:mBullet"`                    // Количество пуль (int)
    MSupportType    uint8     `json:"mSupportType" gorm:"column:mSupportType"`          // Тип поддержки (tinyint)
    MLastSupportDate time.Time `json:"mLastSupportDate" gorm:"column:mLastSupportDate"`  // Дата последней поддержки (smalldatetime)
    MSvrInfo        uint8     `json:"mSvrInfo" gorm:"column:mSvrInfo"`                  // Информация о сервере (tinyint)
    MIsInputDlg     bool      `json:"mIsInputDlg" gorm:"column:mIsInputDlg"`            // Флаг ввода (bit)
    MSmallSendCnt   int       `json:"mSmallSendCnt" gorm:"column:mSmallSendCnt"`        // Количество малых отправок (int)
}

func (ParmSvr) TableName() string {
	return "TblParmSvr"
}