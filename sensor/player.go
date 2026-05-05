package sensor

import (
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"math"
)

// PlayerSensor is a detection system that searches for players and entities in the world.
type PlayerSensor struct {
	// Range is the distance in blocks the sensor can detect.
	Range    float64
	// Detected holds the list of entities found in the last scan.
	Detected []api.Entity
}

// Scan searches the world for entities within the sensor's range.
func (s *PlayerSensor) Scan(e api.Entity, w api.World) bool {
	s.Detected = nil
	pos := e.Position()
	
	for _, other := range w.Entities() {
		// Ignore self!
		if other.ID() == e.ID() {
			continue
		}
		
		oPos := other.Position()
		// Calculate 3D Euclidean distance
		dx := pos[0] - oPos[0]
		dy := pos[1] - oPos[1]
		dz := pos[2] - oPos[2]
		dist := math.Sqrt(dx*dx + dy*dy + dz*dz)
		
		if dist <= s.Range {
			s.Detected = append(s.Detected, other)
		}
	}
	
	// Sort by distance (closest first) could be added here for efficiency.
	return len(s.Detected) > 0
}
