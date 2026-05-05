package mobsx

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
)

// Sensor gathers information from the world.
type Sensor interface {
	Scan(e api.Entity, w api.World) bool
}

// Behavior defines a goal or action.
type Behavior interface {
	Priority() int
	CanRun(e api.Entity, w api.World) bool
	Run(e api.Entity, w api.World)
}

// Brain is the main AI controller.
type Brain struct {
	sensors   []Sensor
	behaviors []Behavior
	
	activeBehavior Behavior
}

func NewBrain() *Brain {
	return &Brain{}
}

func (m *Brain) AddSensor(s Sensor) {
	m.sensors = append(m.sensors, s)
}

func (m *Brain) AddBehavior(b Behavior) {
	m.behaviors = append(m.behaviors, b)
}

func (m *Brain) Tick(e api.Entity, w api.World) {
	for _, s := range m.sensors {
		s.Scan(e, w)
	}

	var best Behavior
	for _, b := range m.behaviors {
		if b.CanRun(e, w) {
			if best == nil || b.Priority() > best.Priority() {
				best = b
			}
		}
	}

	if best != nil {
		m.activeBehavior = best
		best.Run(e, w)
	}
}
