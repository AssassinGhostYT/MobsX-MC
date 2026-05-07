package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"github.com/AssassinGhostYT/MobsX-MC"
	"math/rand/v2"
)

type WanderBehavior struct {
	Radius    int
	navigator *mobsx.Navigator
}

func NewWander(n *mobsx.Navigator, radius int) *WanderBehavior {
	return &WanderBehavior{
		Radius:    radius,
		navigator: n,
	}
}

func (w *WanderBehavior) Priority() int { return 1 }

func (w *WanderBehavior) CanRun(e api.Entity, world api.World) bool {
	return true
}

func (w *WanderBehavior) Run(e api.Entity, world api.World) {
	if w.navigator.Path.AtEnd() {
		pos := e.Position()
		target := mmath.Pos{
			int(pos[0]) + rand.IntN(w.Radius*2) - w.Radius,
			int(pos[1]),
			int(pos[2]) + rand.IntN(w.Radius*2) - w.Radius,
		}
		
		for y := 2; y >= -2; y-- {
			check := target
			check[1] += y
			if w.isWalkable(check, world) {
				w.navigator.SetTarget(check)
				break
			}
		}
	}
	
	w.navigator.Tick()
}

func (w *WanderBehavior) isWalkable(pos mmath.Pos, world api.World) bool {
	b := world.Block(pos)
	if b.Solid() {
		return false
	}
	below := world.Block(pos.Side(0))
	return below.Solid()
}
