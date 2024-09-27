package query

import (
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/parm"
)

type MemberParm struct {
	Members       []account.Member
	Users         []account.User
	Pcs           []game.Pc
	PcInvs        []game.PcInventory
	PcStates      []game.PcState
	PcStores      []game.PcStore
	ItemResources []parm.ItemResource
	Items         []parm.Item
}

type IntermediateMemberParm struct {
	InteMembers      []account.IntermediateMember
	InteUsers        []account.IntermediateUser
	IntePcs          []game.IntermediatePc
	IntePcInvs       []game.IntermediatePcInventory
	IntePcStates     []game.IntermediatePcState
	IntePcStores     []game.IntermediatePcStore
	InteItemResource []parm.IntermediateItemResource
	Items            []parm.IntermediateItem
}

type MemberPcItem struct {
	Members       []account.Member
	Users         []account.User
	Pcs           []game.Pc
	PcInvs        []game.PcInventory
	PcStates      []game.PcState
	PcStores      []game.PcStore
	ItemResources []parm.ItemResource
	Items         []parm.Item
}

type IntermediateMemberPcItem struct {
	InteMembers      []account.IntermediateMember
	InteUsers        []account.IntermediateUser
	IntePcs          []game.IntermediatePc
	IntePcInvs       []game.IntermediatePcInventory
	IntePcStates     []game.IntermediatePcState
	IntePcStores     []game.IntermediatePcStore
	InteItemResource []parm.IntermediateItemResource
	Items            []parm.IntermediateItem
}

/* type MemberUser struct {
	Members []account.Member
	Users   []account.User
}

type PcInvState struct {
	Pcs      []game.Pc
	PcInvs   []game.PcInventory
	PcStates []game.PcState
	PcStores []game.PcStore
}

type ItemRes struct {
	ItemResources []parm.ItemResource
	Items         []parm.Item
}

type IntermediateMemberUser []struct {
	MemberMUserId            string
	MemberMUserPswd          string
	MemberSuperpwd           string
	MemberCash               float64
	MemberEmail              string
	MemberTgzh               string
	MemberUid                int
	MemberKlq                int
	MemberYlq                int
	MemberAuth               int
	MemberMSum               string
	MemberIsAdmin            int
	MemberIsdl               int
	MemberDlmoney            int
	MemberRegisterIp         string
	MemberCountry            string
	MemberCashBack           int
	UserMRegDate             time.Time
	UserMUserAuth            uint8
	UserMUserNo              int
	UserMUserId              string
	UserMUserPswd            string
	UserMCertifiedKey        int
	UserMIp                  string
	UserMLoginTm             time.Time
	UserMLogoutTm            time.Time
	UserMTotUseTm            int
	UserMWorldNo             int16
	UserMDelDate             time.Time
	UserMPcBangLv            int
	UserMSecKeyTableUse      uint8
	UserMUseMacro            int16
	UserMIpEX                int64
	UserMJoinCode            string
	UserMLoginChannelID      string
	UserMTired               string
	UserMChnSID              string
	UserMNewId               bool
	UserMLoginSvrType        uint8
	UserMAccountGuid         int
	UserMNormalLimitTime     int
	UserMPcBangLimitTime     int
	UserMRegIp               string
	UserMIsMovingToBattleSvr bool
}

type IntermediatePcInvState []struct {
	PcMRegDate                      time.Time
	PcMOwner                        int
	PcMSlot                         int8
	PcMNo                           int
	PcMNm                           string
	PcMClass                        int8
	PcMSex                          int8
	PcMHead                         int8
	PcMFace                         int8
	PcMBody                         int8
	PcMHomeMapNo                    int
	PcMHomePosX                     float64
	PcMHomePosY                     float64
	PcMHomePosZ                     float64
	PcMDelDate                      time.Time
	PcInventoryMRegDate             time.Time
	PcInventoryMSerialNo            int64
	PcInventoryMPcNo                int
	PcInventoryMItemNo              int
	PcInventoryMEndDate             time.Time
	PcInventoryMIsConfirm           bool
	PcInventoryMStatus              int8
	PcInventoryMCnt                 int
	PcInventoryMCntUse              int16
	PcInventoryMIsSeizure           bool
	PcInventoryMApplyAbnItemNo      int
	PcInventoryMApplyAbnItemEndDate time.Time
	PcInventoryMOwner               int
	PcInventoryMPracticalPeriod     int
	PcInventoryMBindingType         int8
	PcInventoryMRestoreCnt          int8
	PcInventoryMHoleCount           int8
	PcStateMNo                      int
	PcStateMLevel                   int16
	PcStateMExp                     int64
	PcStateMHpAdd                   int
	PcStateMHp                      int
	PcStateMMpAdd                   int
	PcStateMMp                      int
	PcStateMMapNo                   int
	PcStateMPosX                    float32
	PcStateMPosY                    float32
	PcStateMPosZ                    float32
	PcStateMStomach                 int16
	PcStateMIp                      string
	PcStateMLoginTm                 time.Time
	PcStateMLogoutTm                time.Time
	PcStateMTotUseTm                int
	PcStateMPkCnt                   int
	PcStateMChaotic                 int
	PcStateMDiscipleJoinCount       int
	PcStateMPartyMemCntLevel        int
	PcStateMLostExp                 int64
	PcStateMIsLetterLimit           bool
	PcStateMFlag                    int16
	PcStateMIsPreventItemDrop       bool
	PcStateMSkillTreePoint          int16
	PcStateMRestExpGuild            int64
	PcStateMRestExpActivate         int64
	PcStateMRestExpDeactivate       int64
	PcStateMQMCnt                   int16
	PcStateMGuildQMCnt              int16
	PcStateMFierceCnt               int16
	PcStateMBossCnt                 int16
	PcStoreMRegDate                 time.Time
	PcStoreMSerialNo                int64
	PcStoreMUserNo                  int
	PcStoreMItemNo                  int
	PcStoreMEndDate                 time.Time
	PcStoreMIsConfirm               bool
	PcStoreMStatus                  int8
	PcStoreMCnt                     int
	PcStoreMCntUse                  int16
	PcStoreMIsSeizure               bool
	PcStoreMApplyAbnItemNo          int
	PcStoreMApplyAbnItemEndDate     time.Time
	PcStoreMOwner                   int
	PcStoreMPracticalPeriod         int
	PcStoreMBindingType             int8
	PcStoreMRestoreCnt              int8
	PcStoreMHoleCount               int8
}

type IntermediateItemRes []struct {
	ItemResourceRID                 int
	ItemResourceROwnerID            int
	ItemResourceRType               int
	ItemResourceRFileName           string
	ItemResourceRPosX               int
	ItemResourceRPosY               int
	ItemIID                         int
	ItemIName                       string
	ItemIType                       int
	ItemILevel                      uint8
	ItemIDHIT                       int16
	ItemIDDD                        string
	ItemIRHIT                       int16
	ItemIRDD                        string
	ItemIMHIT                       int16
	ItemIMDD                        string
	ItemIHPPlus                     int16
	ItemIMPPlus                     int16
	ItemISTR                        int16
	ItemIDEX                        int16
	ItemIINT                        int16
	ItemIMaxStack                   int
	ItemIWeight                     int16
	ItemIUseType                    int
	ItemIUseNum                     int
	ItemIRecycle                    int
	ItemIHPRegen                    uint8
	ItemIMPRegen                    uint8
	ItemIAttackRate                 uint8
	ItemIMoveRate                   uint8
	ItemICritical                   uint8
	ItemITermOfValidity             int16
	ItemITermOfValidityMi           int16
	ItemIDesc                       string
	ItemIStatus                     uint8
	ItemIFakeID                     int
	ItemIFakeName                   string
	ItemIUseMsg                     string
	ItemIRange                      int16
	ItemIUseClass                   uint8
	ItemIDropEffect                 int
	ItemIUseLevel                   int16
	ItemIUseEternal                 uint8
	ItemIUseDelay                   int
	ItemIUseInAttack                uint8
	ItemIIsEvent                    bool
	ItemIIsIndict                   bool
	ItemIAddWeight                  int16
	ItemISubType                    int16
	ItemIIsCharge                   bool
	ItemINationOp                   int64
	ItemIPShopItemType              uint8
	ItemIQuestNo                    int
	ItemIIsTest                     bool
	ItemIQuestNeedCnt               uint8
	ItemIContentsLv                 uint8
	ItemIIsConfirm                  bool
	ItemIIsSealable                 bool
	ItemIAddDDWhenCritical          int16
	ItemMSealRemovalNeedCnt         uint8
	ItemMIsPracticalPeriod          bool
	ItemMIsReceiveTown              bool
	ItemIIsReinforceDestroy         bool
	ItemIAddPotionRestore           int16
	ItemIAddMaxHpWhenTransform      int16
	ItemIAddMaxMpWhenTransform      int16
	ItemIAddAttackRateWhenTransform int16
	ItemIAddMoveRateWhenTransform   int16
	ItemISupportType                uint8
	ItemITermOfValidityLv           int16
	ItemMIsUseableUTGWSvr           bool
	ItemIAddShortAttackRange        int16
	ItemIAddLongAttackRange         int16
	ItemIWeaponPoisonType           int16
	ItemIDPV                        int16
	ItemIMPV                        int16
	ItemIRPV                        int16
	ItemIDDV                        int16
	ItemIMDV                        int16
	ItemIRDV                        int16
	ItemIHDPV                       int16
	ItemIHMPV                       int16
	ItemIHRPV                       int16
	ItemIHDDV                       int16
	ItemIHMDV                       int16
	ItemIHRDV                       int16
	ItemISubDDWhenCritical          int16
	ItemIGetItemFeedback            int16
	ItemIEnemySubCriticalHit        int16
	ItemIIsPartyDrop                bool
	ItemIMaxBeadHoleCount           uint8
	ItemISubTypeOption              int
	ItemMIsDeleteArenaSvr           bool
}
*/
