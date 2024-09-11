package conv

import (
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	query "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
)

func ConvMember(
		member account.IntermediateMember,
		user account.IntermediateUser,
		pc game.IntermediatePc,
		pcInv game.IntermediatePcInventory,
		pcState game.IntermediatePcState,
		pcStore game.IntermediatePcStore, 
		memberParm *query.MemberParm) {
	
	// conv
	convMember := toMemberResponse(member)
	convUser := toUserResponse(user)
	convPc := toPcResponse(pc)
	convPcInv := toPcInvResponse(pcInv)
	convPcState := toPcStateResponse(pcState)
	convPcStore := toPcStoreResponse(pcStore)

	// add array
	memberParm.Members = append(memberParm.Members, convMember)
	memberParm.Users = append(memberParm.Users, convUser)
	memberParm.Pcs = append(memberParm.Pcs, convPc)
	memberParm.PcInvs = append(memberParm.PcInvs, convPcInv)
	memberParm.PcStates = append(memberParm.PcStates, convPcState)
	memberParm.PcStores = append(memberParm.PcStores, convPcStore)

	// remove duplicate
	memberParm.Members =  removeDuplicateMembers(memberParm.Members)
	memberParm.Users =    removeDuplicateUsers(memberParm.Users)
   	memberParm.Pcs =      removeDuplicatePcs(memberParm.Pcs)
   	memberParm.PcInvs =   removeDuplicatePcInvs(memberParm.PcInvs)
   	memberParm.PcStates = removeDuplicatePcStates(memberParm.PcStates)
   	memberParm.PcStores = removeDuplicatePcStores(memberParm.PcStores)
}

// TODO: можно сделать один запрос (подумать)
func toMemberResponse(member account.IntermediateMember) account.Member {
	return account.Member{
		MUserId:    member.MUserId.String,     // string
		MUserPswd:  member.MUserPswd.String,   // string
		Superpwd:   member.Superpwd.String,    // string
		Cash:       member.Cash.Float64,       // float64
		Email:      member.Email.String,       // string
		Tgzh:       member.Tgzh.String,        // string
		Uid:        int(member.Uid.Int64),     // int
		Klq:        int(member.Klq.Int64),     // int
		Ylq:        int(member.Ylq.Int64),     // int
		Auth:       int(member.Auth.Int64),    // int
		MSum:       member.MSum.String,        // string
		IsAdmin:    int(member.IsAdmin.Int64), // int
		Isdl:       int(member.Isdl.Int64),    // int
		Dlmoney:    int(member.Dlmoney.Int64), // int
		RegisterIp: member.RegisterIp.String,  // string
		Country:    member.Country.String,     // string
		CashBack:   int(member.Cash.Float64),  // int
	}
}

func toUserResponse(user account.IntermediateUser) account.User {
	return account.User{
		MRegDate:             user.MRegDate,
		MUserAuth:            user.MUserAuth,
		MUserNo:              int(user.MUserNo.Int64),
		MUserId:              user.MUserId.String,
		MUserPswd:            user.MUserPswd.String,
		MCertifiedKey:        int(user.MCertifiedKey.Int64),
		MIp:                  user.MIp.String,
		MLoginTm:             user.MLoginTm,
		MLogoutTm:            user.MLogoutTm,
		MTotUseTm:            int(user.MTotUseTm.Int64),
		MWorldNo:             user.MWorldNo,
		MDelDate:             user.MDelDate,
		MPcBangLv:            int(user.MPcBangLv.Int64),
		MSecKeyTableUse:      user.MSecKeyTableUse,
		MUseMacro:            user.MUseMacro,
		MIpEX:                user.MIpEX.Int64,
		MJoinCode:            user.MJoinCode.String,
		MLoginChannelID:      user.MLoginChannelID.String,
		MTired:               user.MTired.String,
		MChnSID:              user.MChnSID.String,
		MNewId:               user.MNewId.Bool,
		MLoginSvrType:        user.MLoginSvrType,
		MAccountGuid:         int(user.MAccountGuid.Int64),
		MNormalLimitTime:     int(user.MNormalLimitTime.Int64),
		MPcBangLimitTime:     int(user.MPcBangLimitTime.Int64),
		MRegIp:               user.MRegIp.String,
		MIsMovingToBattleSvr: user.MIsMovingToBattleSvr.Bool,
	}
}

func toPcResponse(pc game.IntermediatePc) game.Pc {
	return game.Pc{
		MRegDate:   pc.MRegDate,
		MOwner:     int(pc.MOwner.Int64),
		MSlot:      pc.MSlot,
		MNo:        int(pc.MNo.Int64),
		MNm:        pc.MNm,
		MClass:     pc.MClass,
		MSex:       pc.MSex,
		MHead:      pc.MHead,
		MFace:      pc.MFace,
		MBody:      pc.MBody,
		MHomeMapNo: int(pc.MHomeMapNo.Int64),
		MHomePosX:  pc.MHomePosX,
		MHomePosY:  pc.MHomePosY,
		MHomePosZ:  pc.MHomePosZ,
		MDelDate:   pc.MDelDate.Time,
	}
}

func toPcInvResponse(pcInv game.IntermediatePcInventory) game.PcInventory {
	return game.PcInventory{
		MRegDate:             pcInv.MRegDate,
		MSerialNo:            pcInv.MSerialNo.Int64,
		MPcNo:                int(pcInv.MPcNo.Int64),
		MItemNo:              int(pcInv.MItemNo.Int64),
		MEndDate:             pcInv.MEndDate,
		MIsConfirm:           pcInv.MIsConfirm,
		MStatus:              pcInv.MStatus,
		MCnt:                 int(pcInv.MCnt.Int64),
		MCntUse:              pcInv.MCntUse,
		MIsSeizure:           pcInv.MIsSeizure,
		MApplyAbnItemNo:      int(pcInv.MApplyAbnItemNo.Int64),
		MApplyAbnItemEndDate: pcInv.MApplyAbnItemEndDate,
		MOwner:               int(pcInv.MOwner.Int64),
		MPracticalPeriod:     int(pcInv.MPracticalPeriod.Int64),
		MBindingType:         pcInv.MBindingType,
		MRestoreCnt:          pcInv.MRestoreCnt,
		MHoleCount:           pcInv.MHoleCount,
	}
}

func toPcStateResponse(pcState game.IntermediatePcState) game.PcState {
	return game.PcState{
		MNo:                int(pcState.MNo.Int64),
		MLevel:             pcState.MLevel,
		MExp:               pcState.MExp.Int64,
		MHpAdd:             int(pcState.MHpAdd.Int64),
		MHp:                int(pcState.MHp.Int64),
		MMpAdd:             int(pcState.MMp.Int64),
		MMp:                int(pcState.MMp.Int64),
		MMapNo:             int(pcState.MMapNo.Int64),
		MPosX:              pcState.MPosX,
		MPosY:              pcState.MPosY,
		MPosZ:              pcState.MPosZ,
		MStomach:           pcState.MStomach,
		MIp:                pcState.MIp.String,
		MLoginTm:           pcState.MLoginTm.Time,
		MLogoutTm:          pcState.MLogoutTm.Time,
		MTotUseTm:          int(pcState.MTotUseTm.Int64),
		MPkCnt:             int(pcState.MPkCnt.Int64),
		MChaotic:           int(pcState.MChaotic.Int64),
		MDiscipleJoinCount: int(pcState.MDiscipleJoinCount.Int64),
		MPartyMemCntLevel:  int(pcState.MPartyMemCntLevel.Int64),
		MLostExp:           pcState.MLostExp.Int64,
		MIsLetterLimit:     pcState.MIsLetterLimit.Bool,
		MFlag:              pcState.MFlag,
		MIsPreventItemDrop: pcState.MIsPreventItemDrop.Bool,
		MSkillTreePoint:    pcState.MSkillTreePoint,
		MRestExpGuild:      pcState.MRestExpGuild.Int64,
		MRestExpActivate:   pcState.MRestExpActivate.Int64,
		MRestExpDeactivate: pcState.MRestExpDeactivate.Int64,
		MQMCnt:             pcState.MQMCnt,
		MGuildQMCnt:        pcState.MGuildQMCnt,
		MFierceCnt:         pcState.MFierceCnt,
		MBossCnt:           pcState.MBossCnt,
	}
}

func toPcStoreResponse(pcStore game.IntermediatePcStore) game.PcStore {
	return game.PcStore{
		MRegDate:             pcStore.MRegDate.Time,
		MSerialNo:            pcStore.MSerialNo.Int64,
		MUserNo:              int(pcStore.MUserNo.Int64),
		MItemNo:              int(pcStore.MItemNo.Int64),
		MEndDate:             pcStore.MEndDate.Time,
		MIsConfirm:           pcStore.MIsConfirm.Bool,
		MStatus:              int8(pcStore.MStatus.Int16),
		MCnt:                 int(pcStore.MCnt.Int64),
		MCntUse:              pcStore.MCntUse.Int16,
		MIsSeizure:           pcStore.MIsSeizure.Bool,
		MApplyAbnItemNo:      int(pcStore.MApplyAbnItemNo.Int64),
		MApplyAbnItemEndDate: pcStore.MApplyAbnItemEndDate.Time,
		MOwner:               int(pcStore.MOwner.Int64),
		MPracticalPeriod:     int(pcStore.MPracticalPeriod.Int64),
		MBindingType:         int8(pcStore.MBindingType.Int16),
		MRestoreCnt:          int8(pcStore.MRestoreCnt.Int16),
		MHoleCount:           int8(pcStore.MHoleCount.Int16),
	}
}

func removeDuplicateMembers(members []account.Member) []account.Member {
	seen := make(map[string]bool)
	uniqueMembers := []account.Member{}

	for _, member := range members {
		if _, exists := seen[member.MUserId]; !exists {
			seen[member.MUserId] = true
			uniqueMembers = append(uniqueMembers, member)
		}
	}

	return uniqueMembers
}

func removeDuplicateUsers(users []account.User) []account.User {
	seen := make(map[string]bool)
	uniqueUsers := []account.User{}

	for _, user := range users {
		if _, exists := seen[user.MUserId]; !exists {
			seen[user.MUserId] = true
			uniqueUsers = append(uniqueUsers, user)
		}
	}

	return uniqueUsers
}

func removeDuplicatePcs(pcs []game.Pc) []game.Pc {
	seen := make(map[int]bool)
	unique := []game.Pc{}

	for _, pc := range pcs {
		if _, exists := seen[pc.MNo]; !exists {
			seen[pc.MNo] = true
			unique = append(unique, pc)
		}
	}

	return unique
}

func removeDuplicatePcInvs(pcInvs []game.PcInventory) []game.PcInventory {
	seen := make(map[int64]bool)
	unique := []game.PcInventory{}

	for _, pcInv := range pcInvs {
		if _, exists := seen[pcInv.MSerialNo]; !exists {
			seen[pcInv.MSerialNo] = true
			unique = append(unique, pcInv)
		}
	}

	return unique
}

func removeDuplicatePcStates(pcStates []game.PcState) []game.PcState {
	seen := make(map[int]bool)
	unique := []game.PcState{}

	for _, pcState := range pcStates {
		if _, exists := seen[pcState.MNo]; !exists {
			seen[pcState.MNo] = true
			unique = append(unique, pcState)
		}
	}

	return unique
}

func removeDuplicatePcStores(pcStores []game.PcStore) []game.PcStore {
	seen := make(map[int64]bool)
	unique := []game.PcStore{}

	for _, pcStore := range pcStores {
		if _, exists := seen[pcStore.MSerialNo]; !exists {
			seen[pcStore.MSerialNo] = true
			unique = append(unique, pcStore)
		}
	}

	return unique
}
