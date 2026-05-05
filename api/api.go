package api

import (
	"github.com/AssassinGhostYT/MobsX-MC/internal/math"
)

// World represents a 3D world that can be queried for block and entity information.
type World interface {
	// Block returns the block at the given position.
	Block(pos math.Pos) Block
	// Entities returns all entities in the world.
	Entities() []Entity
}

// Block represents a single block in the world.
type Block interface {
	// Name returns the identifier of the block (e.g., "minecraft:stone").
	Name() string
	// Solid returns true if the block is solid.
	Solid() bool
	// Passable returns true if an entity can walk through this block.
	Passable() bool
}

// Entity represents any living being in the world.
type Entity interface {
	// Position returns the current coordinates of the entity.
	Position() [3]float64
	// SetPosition updates the entity's position.
	SetPosition(pos [3]float64)
	// Rotation returns the rotation (yaw, pitch) of the entity.
	Rotation() [2]float32
	// SetRotation updates the entity's rotation.
	SetRotation(yaw, pitch float32)
	// ID returns a unique identifier for the entity.
	ID() int64
	// HideInBlock hides the entity inside a block at the given position.
	HideInBlock(pos math.Pos)
	// AlertOthers signals other entities of the same type within a range.
	AlertOthers(rangeX, rangeY, rangeZ int)
}
