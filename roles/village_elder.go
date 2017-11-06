package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type VillageElder struct {
	*BaseRole
	extraLife bool
}

func (v *VillageElder) Name() string {
	return VILLAGE_ELDER
}

func (v *VillageElder) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
	if v.extraLife && e.ID() == v.ID() && e.WerewolfAttack() {
		v.extraLife = false
		e.PreventDeath()
	}
}
