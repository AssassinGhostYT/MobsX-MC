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
	finder *pathfinding.Finder

	Path pathfinding.Path
	Speed float64
}

// NewNavigator creates a new navigator for the given entity.
func NewNavigator(e api.Entity, w api.World) *Navigator {
	return &Navigator{
		entity: e,
		world:  w,
		finder: pathfinding.NewFinder(w),
		Speed:  0.1,
	}
}

// Sync updates the current world reference to ensure fresh transactions.
func (n *Navigator) Sync(w api.World) {
	n.world = w
	n.finder.SetWorld(w)
}

// SetTarget calculates a new path to the target position.
func (n *Navigator) SetTarget(target mmath.Pos) bool {
	pos := n.entity.Position()
	start := mmath.Pos{int(goMath.Floor(pos[0])), int(goMath.Floor(pos[1])), int(goMath.Floor(pos[2]))}
	
	path, ok := n.finder.FindPath(start, target)
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
	
	newY := currentPos[1]
	if dy > 0 && dy <= 1.2 {
		newY += 0.3 
	} else if dy < 0 {
		newY -= 0.1 
	}

	n.entity.SetPosition([3]float64{newX, newY, newZ})
	
	yaw := float32(angle * 180 / goMath.Pi) - 90
	n.entity.SetRotation(yaw, 0)
}
