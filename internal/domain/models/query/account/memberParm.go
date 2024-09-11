package query

import (
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	// gamePcParm "github.com/Tibirlayn/R2Hunter/internal/domain/models/query/game"
)

type MemberParm struct {
	Members  []account.Member
	Users    []account.User
	Pcs      []game.Pc
	PcInvs   []game.PcInventory
	PcStates []game.PcState
	PcStores []game.PcStore
	// UserAdmin []account.UserAdmin
	// PcCards   []gamePcParm.PcParm
}

type IntermediateMemberParm struct {
	InteMembers  []account.IntermediateMember
	InteUsers    []account.IntermediateUser
	IntePcs      []game.IntermediatePc
	IntePcInvs   []game.IntermediatePcInventory
	IntePcStates []game.IntermediatePcState
	IntePcStores []game.IntermediatePcStore
	// UserAdmin []account.UserAdmin
	// PcCards   []gamePcParm.PcParm
}
