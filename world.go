package mobsx

import "github.com/AssassinGhostYT/MobsX-MC/internal/math"

// World represents a 3D world that can be queried for block information.
type World interface {
	// Block returns the block at the given position.
	Block(pos math.Pos) Block
}

// Block represents a single block in the world.
type Block interface {
	// Solid returns true if the block is solid.
	Solid() bool
}

// Entity represents any living being in the world.
type Entity interface {
	// Position returns the current coordinates of the entity.
	Position() [3]float64
	// ID returns a unique identifier for the entity.
	ID() int64
}
