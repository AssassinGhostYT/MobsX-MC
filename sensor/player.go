package sensor

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"math"
)

type PlayerSensor struct {
	Range    float64
	Detected []api.Entity
}

func (s *PlayerSensor) Scan(e api.Entity, w api.World) bool {
	s.Detected = nil
	pos := e.Position()
	for _, other := range w.Entities() {
		if other.ID() == e.ID() {
			continue
		}
		oPos := other.Position()
		dist := math.Sqrt((pos[0]-oPos[0])*(pos[0]-oPos[0]) + (pos[1]-oPos[1])*(pos[1]-oPos[1]) + (pos[2]-oPos[2])*(pos[2]-oPos[2]))
		if dist <= s.Range {
			s.Detected = append(s.Detected, other)
		}
	}
	return len(s.Detected) > 0
}
