package sensor

import (
	"github.com/AssassinGhostYT/MobsX-MC"
	"math"
)

type PlayerSensor struct {
	Range    float64
	Detected []mobsx.Entity
}

func (s *PlayerSensor) Scan(e mobsx.Entity, w mobsx.World) bool {
	s.Detected = nil
	// In a real implementation, the World would provide a way to get entities.
	// For this library, we expect the implementation to provide them.
	return false
}
