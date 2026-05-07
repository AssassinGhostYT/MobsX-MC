package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"math/rand/v2"
	"slices"
)

// InfestBehavior makes an entity hide inside stone-like blocks.
type InfestBehavior struct {
	// InfestChance is the probability (0-1) to infest a block each tick.
	InfestChance float64
}

var infestableBlocks = []string{
	"minecraft:stone",
	"minecraft:cobblestone",
	"minecraft:deepslate",
	"minecraft:stone_bricks",
	"minecraft:mossy_stone_bricks",
	"minecraft:cracked_stone_bricks",
	"minecraft:chiseled_stone_bricks",
}

func (i *InfestBehavior) Priority() int { return 2 }

func (i *InfestBehavior) CanRun(e api.Entity, w api.World) bool {
	return rand.Float64() < i.InfestChance
}

func (i *InfestBehavior) Run(e api.Entity, w api.World) {
	pos := e.Position()
	center := mmath.Pos{int(pos[0]), int(pos[1]), int(pos[2])}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				check := mmath.Pos{center[0] + x, center[1] + y, center[2] + z}
				if i.isInfestable(w.Block(check)) {
					e.HideInBlock(check)
					return
				}
			}
		}
	}
}

func (i *InfestBehavior) isInfestable(b api.Block) bool {
	return slices.Contains(infestableBlocks, b.Name())
}
