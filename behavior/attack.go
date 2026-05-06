package behavior

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
	"github.com/AssassinGhostYT/MobsX-MC"
	"github.com/AssassinGhostYT/MobsX-MC/sensor"
	goMath "math"
)

// AttackBehavior makes an entity chase and attack a detected player.
type AttackBehavior struct {
	sensor    *sensor.PlayerSensor
	navigator *mobsx.Navigator
	
	lastStrike int64 // Tick of the last attack
}

// NewAttack creates a new attack behavior.
func NewAttack(s *sensor.PlayerSensor, n *mobsx.Navigator) *AttackBehavior {
	return &AttackBehavior{
		sensor:    s,
		navigator: n,
	}
}

// Priority returns 10 (High priority).
func (a *AttackBehavior) Priority() int { return 10 }

// CanRun returns true if the sensor has detected a player.
func (a *AttackBehavior) CanRun(e api.Entity, world api.World) bool {
	return len(a.sensor.Detected) > 0
}

// Run executes the chase and attack logic.
func (a *AttackBehavior) Run(e api.Entity, world api.World) {
	if len(a.sensor.Detected) == 0 {
		return
	}

	// Focus on the first detected player (closest)
	target := a.sensor.Detected[0]
	targetPos := target.Position()
	currentPos := e.Position()

	dx := targetPos[0] - currentPos[0]
	dy := targetPos[1] - currentPos[1]
	dz := targetPos[2] - currentPos[2]
	dist := goMath.Sqrt(dx*dx + dy*dy + dz*dz)

	// If within range (1.5 blocks), stop and attack
	if dist <= 1.5 {
		// Face the target
		angle := goMath.Atan2(dz, dx)
		yaw := float32(angle * 180 / goMath.Pi) - 90
		e.SetRotation(yaw, 0)
		
		// Trigger attack (this logic depends on the server implementation)
		// For now, we clear the path to stop movement.
		a.navigator.Path.Nodes = nil
		return
	}

	// Otherwise, chase the player using the navigator
	tPos := mmath.Pos{int(goMath.Floor(targetPos[0])), int(goMath.Floor(targetPos[1])), int(goMath.Floor(targetPos[2]))}
	
	// Recalculate path only if target moved or we have no path
	if a.navigator.Path.AtEnd() {
		a.navigator.SetTarget(tPos)
	}
	
	a.navigator.Tick()
}
