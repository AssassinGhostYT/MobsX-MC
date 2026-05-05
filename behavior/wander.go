package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC"
	"github.com/AssassinGhostYT/MobsX-MC/internal/math"
)

type WanderBehavior struct {
	Radius int
}

func (WanderBehavior) Priority() int { return 1 }

func (WanderBehavior) CanRun(e mobsx.Entity, w mobsx.World) bool { return true }

func (WanderBehavior) Run(e mobsx.Entity, w mobsx.World) {
	// Pick random spot and move.
}
