package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"math/rand/v2"
)

// PanicBehavior makes an entity run randomly at high speed when panicking.
type PanicBehavior struct {
	navigator *mobsx.Navigator
}

// NewPanic creates a new panic behavior.
func NewPanic(n *mobsx.Navigator) *PanicBehavior {
	return &PanicBehavior{navigator: n}
}

// Priority returns 100 (Highest priority).
func (p *PanicBehavior) Priority() int { return 100 }

// CanRun returns true if the entity is panicking.
func (p *PanicBehavior) CanRun(e api.Entity, world api.World) bool {
	return e.Panicking()
}

// Run executes the fleeing logic.
func (p *PanicBehavior) Run(e api.Entity, world api.World) {
	if p.navigator.Path.AtEnd() {
		pos := e.Position()
		// Higher radius for panic
		target := mmath.Pos{
			int(pos[0]) + rand.IntN(16) - 8,
			int(pos[1]),
			int(pos[2]) + rand.IntN(16) - 8,
		}
		p.navigator.SetTarget(target)
	}
	
	// Fast speed during panic
	originalSpeed := p.navigator.Speed
	p.navigator.Speed = 0.4
	p.navigator.Tick()
	p.navigator.Speed = originalSpeed
}
