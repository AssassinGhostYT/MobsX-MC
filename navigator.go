package mobsx

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"github.com/AssassinGhostYT/MobsX-MC/pathfinding"
	goMath "math"
)

// Navigator handles path following and movement for an entity.
type Navigator struct {
	entity api.Entity
	world  api.World
	Finder *pathfinding.Finder

	Path  pathfinding.Path
	Speed float64
}

// NewNavigator creates a new navigator for the given entity.
func NewNavigator(e api.Entity, w api.World) *Navigator {
	return &Navigator{
		entity: e,
		world:  w,
		Finder: pathfinding.NewFinder(w),
		Speed:  0.1,
	}
}

// Sync updates the current world reference to ensure fresh transactions.
func (n *Navigator) Sync(w api.World) {
	n.world = w
	n.Finder.SetWorld(w)
}

// SetTarget calculates a new path to the target position.
func (n *Navigator) SetTarget(target mmath.Pos) bool {
	pos := n.entity.Position()
	start := mmath.Pos{int(goMath.Floor(pos[0])), int(goMath.Floor(pos[1])), int(goMath.Floor(pos[2]))}
	
	path, ok := n.Finder.FindPath(start, target)
	if ok {
		n.Path = path
	}
	return ok
}

// Tick moves the entity towards the next node in the path.
func (n *Navigator) Tick() {
	if n.Path.AtEnd() {
		return
	}

	targetNode := n.Path.Nodes[n.Path.Index]
	targetPos := [3]float64{float64(targetNode.X()) + 0.5, float64(targetNode.Y()), float64(targetNode.Z()) + 0.5}
	
	currentPos := n.entity.Position()
	dx := targetPos[0] - currentPos[0]
	dy := targetPos[1] - currentPos[1]
	dz := targetPos[2] - currentPos[2]
	dist := goMath.Sqrt(dx*dx + dz*dz)

	if dist < 0.2 {
		n.Path.Index++ 
		return
	}

	angle := goMath.Atan2(dz, dx)
	newX := currentPos[0] + goMath.Cos(angle)*n.Speed
	newZ := currentPos[2] + goMath.Sin(angle)*n.Speed
	
	// Si hay un bloque sólido en los pies, intentamos subirlo si el objetivo está arriba o al mismo nivel.
	if n.world.Block(checkPos).Solid() && dy >= -0.5 {
		n.entity.Jump()
	}

	n.entity.SetPosition([3]float64{newX, currentPos[1], newZ})
	
	yaw := float32(angle * 180 / goMath.Pi) - 90
	n.entity.SetRotation(yaw, 0)
}
