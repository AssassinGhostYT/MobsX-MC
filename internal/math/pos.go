package math

import "math"

// Pos represents a position in a 3D grid.
type Pos [3]int

func (p Pos) X() int { return p[0] }
func (p Pos) Y() int { return p[1] }
func (p Pos) Z() int { return p[2] }

// Side returns the position adjacent to this one on the given face.
func (p Pos) Side(face int) Pos {
	switch face {
	case 0: return Pos{p[0], p[1] - 1, p[2]} // Down
	case 1: return Pos{p[0], p[1] + 1, p[2]} // Up
	case 2: return Pos{p[0], p[1], p[2] + 1} // South
	case 3: return Pos{p[0], p[1], p[2] - 1} // North
	case 4: return Pos{p[0] + 1, p[1], p[2]} // East
	case 5: return Pos{p[0] - 1, p[1], p[2]} // West
	}
	return p
}

// Distance calculates the Euclidean distance between two positions.
func (p Pos) Distance(other Pos) float64 {
	return math.Sqrt(float64((p[0]-other[0])*(p[0]-other[0]) + (p[1]-other[1])*(p[1]-other[1]) + (p[2]-other[2])*(p[2]-other[2])))
}
