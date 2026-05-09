package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"github.com/AssassinGhostYT/MobsX-MC"
	"github.com/AssassinGhostYT/MobsX-MC/sensor"
	goMath "math"
)

// TemptBehavior makes an entity follow a player holding specific items.
type TemptBehavior struct {
	sensor    *sensor.PlayerSensor
	navigator *mobsx.Navigator
	
	// StopRange is the distance at which the entity will stop and look at the player.
	StopRange float64
	// Filter returns true if the held item should be followed.
	Filter func(name string, meta int16) bool
	
	target api.Entity
}

// NewTempt creates a new tempt behavior.
func NewTempt(s *sensor.PlayerSensor, n *mobsx.Navigator, filter func(name string, meta int16) bool) *TemptBehavior {
	return &TemptBehavior{
		sensor:    s,
		navigator: n,
		StopRange: 2.5,
		Filter:    filter,
	}
}

// Priority returns 6 (Higher than wander).
func (t *TemptBehavior) Priority() int { return 6 }

// CanRun returns true if a player with a valid item is detected.
func (t *TemptBehavior) CanRun(e api.Entity, world api.World) bool {
	t.target = nil
	for _, p := range t.sensor.Detected {
		name, meta := p.HeldItem()
		if t.Filter(name, meta) {
			t.target = p
			return true
		}
	}
	return false
}

// Run executes the tempt follow logic.
func (t *TemptBehavior) Run(e api.Entity, world api.World) {
	if t.target == nil {
		return
	}

	targetPos := t.target.Position()
	currentPos := e.Position()

	dx := targetPos[0] - currentPos[0]
	dz := targetPos[2] - currentPos[2]
	dist := goMath.Sqrt(dx*dx + dz*dz)

	// Face the player
	angle := goMath.Atan2(dz, dx)
	yaw := float32(angle * 180 / goMath.Pi) - 90
	e.SetRotation(yaw, 0)

	if dist <= t.StopRange {
		// Just stop and look
		t.navigator.Path.Nodes = nil
		return
	}

	// Move towards the player
	tPos := mmath.Pos{int(goMath.Floor(targetPos[0])), int(goMath.Floor(targetPos[1])), int(goMath.Floor(targetPos[2]))}
	
	targetPosVec := mmath.Vec3{targetPos[0], targetPos[1], targetPos[2]}
	var distToPathEnd float64 = 999
	if len(t.navigator.Path.Nodes) > 0 {
		endNode := t.navigator.Path.Nodes[len(t.navigator.Path.Nodes)-1]
		distToPathEnd = targetPosVec.Distance(mmath.Vec3{float64(endNode.X()), float64(endNode.Y()), float64(endNode.Z())})
	}

	if t.navigator.Path.AtEnd() || distToPathEnd > 1.5 {
		t.navigator.SetTarget(tPos)
	}

	t.navigator.Tick()
}
