package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"github.com/AssassinGhostYT/MobsX-MC"
	"github.com/AssassinGhostYT/MobsX-MC/sensor"
	goMath "math"
)

// FollowBehavior makes an entity follow a player holding specific items.
type FollowBehavior struct {
	sensor    *sensor.PlayerSensor
	navigator *mobsx.Navigator
	
	// StopRange is the distance at which the entity will stop and look at the player.
	StopRange float64
	// Filter returns true if the held item should be followed.
	Filter func(name string, meta int16) bool
	
	target api.Entity
}

// NewFollow creates a new follow behavior.
func NewFollow(s *sensor.PlayerSensor, n *mobsx.Navigator, filter func(name string, meta int16) bool) *FollowBehavior {
	return &FollowBehavior{
		sensor:    s,
		navigator: n,
		StopRange: 2.5,
		Filter:    filter,
	}
}

// Priority returns 5 (Medium priority).
func (f *FollowBehavior) Priority() int { return 5 }

// CanRun returns true if a player with a valid item is detected.
func (f *FollowBehavior) CanRun(e api.Entity, world api.World) bool {
	f.target = nil
	for _, p := range f.sensor.Detected {
		name, meta := p.HeldItem()
		if f.Filter(name, meta) {
			f.target = p
			return true
		}
	}
	return false
}

// Run executes the follow logic.
func (f *FollowBehavior) Run(e api.Entity, world api.World) {
	if f.target == nil {
		return
	}

	targetPos := f.target.Position()
	currentPos := e.Position()

	dx := targetPos[0] - currentPos[0]
	dz := targetPos[2] - currentPos[2]
	dist := goMath.Sqrt(dx*dx + dz*dz)

	// Face the player
	angle := goMath.Atan2(dz, dx)
	yaw := float32(angle * 180 / goMath.Pi) - 90
	e.SetRotation(yaw, 0)

	if dist <= f.StopRange {
		// Just stop and look
		f.navigator.Path.Nodes = nil
		return
	}

	// Move towards the player
	tPos := mmath.Pos{int(goMath.Floor(targetPos[0])), int(goMath.Floor(targetPos[1])), int(goMath.Floor(targetPos[2]))}
	
	targetPosVec := mmath.Vec3{targetPos[0], targetPos[1], targetPos[2]}
	var distToPathEnd float64 = 999
	if len(f.navigator.Path.Nodes) > 0 {
		endNode := f.navigator.Path.Nodes[len(f.navigator.Path.Nodes)-1]
		distToPathEnd = targetPosVec.Distance(mmath.Vec3{float64(endNode.X()), float64(endNode.Y()), float64(endNode.Z())})
	}

	if f.navigator.Path.AtEnd() || distToPathEnd > 1.5 {
		f.navigator.SetTarget(tPos)
	}

	f.navigator.Tick()
}
