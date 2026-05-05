package mobsx

// Sensor gathers information from the world.
type Sensor interface {
	Scan(e Entity, w World) bool
}

// Behavior defines a goal or action.
type Behavior interface {
	Priority() int
	CanRun(e Entity, w World) bool
	Run(e Entity, w World)
}

// MobsX is the main AI controller.
type MobsX struct {
	sensors   []Sensor
	behaviors []Behavior
	
	activeBehavior Behavior
}

func New() *MobsX {
	return &MobsX{}
}

func (m *MobsX) AddSensor(s Sensor) {
	m.sensors = append(m.sensors, s)
}

func (m *MobsX) AddBehavior(b Behavior) {
	m.behaviors = append(m.behaviors, b)
}

func (m *MobsX) Tick(e Entity, w World) {
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
