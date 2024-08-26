package query

import (
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
)

type MemberParm struct {
	Member    account.Member
	User      account.User
	UserAdmin account.UserAdmin
	Pc        game.Pc
	PcInv     game.PcInventory
	PcState   game.PcState
}