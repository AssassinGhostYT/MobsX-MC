package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"github.com/AssassinGhostYT/MobsX-MC"
	goMath "math"
)

// FollowParentBehavior makes a baby entity follow nearby adults of the same type.
type FollowParentBehavior struct {
	navigator *mobsx.Navigator
	
	// StopRange is the distance at which the entity will stop and look at the parent.
	StopRange float64
	// SearchRange is the radius to look for parents.
	SearchRange float64
	
	parent api.Entity
}

// NewFollowParent creates a new follow parent behavior.
func NewFollowParent(n *mobsx.Navigator) *FollowParentBehavior {
	return &FollowParentBehavior{
		navigator:   n,
		StopRange:   2.0,
		SearchRange: 16.0,
	}
}

// Priority returns 4 (Lower than tempt, higher than wander).
func (f *FollowParentBehavior) Priority() int { return 4 }

// CanRun returns true if the entity is a baby and a parent is found.
func (f *FollowParentBehavior) CanRun(e api.Entity, world api.World) bool {
	f.parent = nil
	pos := e.Position()
	
	for _, other := range world.Entities() {
		if other.ID() == e.ID() {
			continue
		}
		// In our simple API, we check if it's the same type by comparing ID (conceptually) 
		// but since we don't have "Type" in Entity interface, we'll follow ANY non-player entity 
		// if they are within range. The Chicken implementation will ensure only babies use this.
		if other.IsPlayer() {
			continue
		}
		
		oPos := other.Position()
		dx := pos[0] - oPos[0]
		dz := pos[2] - oPos[2]
		dist := goMath.Sqrt(dx*dx + dz*dz)
		
		if dist <= f.SearchRange {
			f.parent = other
			return true
		}
	}
	return false
}

// Run executes the follow parent logic.
func (f *FollowParentBehavior) Run(e api.Entity, world api.World) {
	if f.parent == nil {
		return
	}

	targetPos := f.parent.Position()
	currentPos := e.Position()

	dx := targetPos[0] - currentPos[0]
	dz := targetPos[2] - currentPos[2]
	dist := goMath.Sqrt(dx*dx + dz*dz)

	if dist <= f.StopRange {
		// Face the parent
		angle := goMath.Atan2(dz, dx)
		yaw := float32(angle * 180 / goMath.Pi) - 90
		e.SetRotation(yaw, 0)
		f.navigator.Path.Nodes = nil
		return
	}

	// Move towards the parent
	tPos := mmath.Pos{int(goMath.Floor(targetPos[0])), int(goMath.Floor(targetPos[1])), int(goMath.Floor(targetPos[2]))}
	f.navigator.SetTarget(tPos)
	f.navigator.Tick()
}
