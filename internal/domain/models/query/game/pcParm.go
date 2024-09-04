package game

import "github.com/Tibirlayn/R2Hunter/internal/domain/models/game"

type PcParm struct {
	Pc      game.Pc
	PcInv   game.PcInventory
	PcState game.PcState
	PcStore game.PcStore

}