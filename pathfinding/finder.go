package pathfinding

import (
	"container/heap"
	"github.com/AssassinGhostYT/MobsX-MC/api"
	"github.com/AssassinGhostYT/MobsX-MC/mmath"
)

type Node struct {
	Pos    mmath.Pos
	Parent *Node
	G, H   float64
}

func (n *Node) F() float64 { return n.G + n.H }

type Path struct {
	Nodes []mmath.Pos
	Index int
}

func (p *Path) Next() (mmath.Pos, bool) {
	if p.Index >= len(p.Nodes) {
		return mmath.Pos{}, false
	}
	pos := p.Nodes[p.Index]
	p.Index++
	return pos, true
}

func (p *Path) AtEnd() bool {
	return p.Index >= len(p.Nodes)
}

type Finder struct {
	w      api.World
	Height int
}

func NewFinder(w api.World) *Finder {
	return &Finder{w: w, Height: 2}
}

func (f *Finder) SetWorld(w api.World) {
	f.w = w
}

func (f *Finder) FindPath(start, end mmath.Pos) (Path, bool) {
	openSet := &priorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &Node{Pos: start, G: 0, H: start.Distance(end)})

	closedSet := make(map[mmath.Pos]struct{})
	iterations := 0

	for openSet.Len() > 0 {
		iterations++
		if iterations > 1000 {
			return Path{}, false
		}
		
		current := heap.Pop(openSet).(*Node)
		if current.Pos == end {
			return f.reconstructPath(current), true
		}
		closedSet[current.Pos] = struct{}{}
		
		for i := 0; i < 6; i++ {
			neighborPos := current.Pos.Side(i)
			if _, ok := closedSet[neighborPos]; ok {
				continue
			}
			
			// Si el vecino es sólido, intentamos subir (si no es el movimiento hacia arriba/abajo ya)
			if i > 1 && !f.isWalkable(neighborPos) {
				up := neighborPos.Side(1)
				if f.isWalkable(up) && !f.w.Block(current.Pos.Side(1).Add(0, f.Height-1, 0)).Solid() {
					neighborPos = up
				} else {
					// Si no es transitable, probamos si es un escalón hacia abajo
					down := neighborPos.Side(0)
					if f.isWalkable(down) {
						neighborPos = down
					} else {
						continue
					}
				}
			} else if i <= 1 && !f.isWalkable(neighborPos) {
				continue
			}

			gScore := current.G + 1
			neighborNode := &Node{
				Pos:    neighborPos,
				Parent: current,
				G:      gScore,
				H:      neighborPos.Distance(end),
			}
			heap.Push(openSet, neighborNode)
		}
	}
	return Path{}, false
}

func (f *Finder) isWalkable(pos mmath.Pos) bool {
	b := f.w.Block(pos)
	if b.Solid() {
		return false
	}
	for h := 1; h < f.Height; h++ {
		if f.w.Block(pos.Add(0, h, 0)).Solid() {
			return false
		}
	}
	
	// Cliff avoidance: Check depth (max 3 blocks drop for chickens)
	foundGround := false
	for d := 1; d <= 4; d++ {
		if f.w.Block(pos.Add(0, -d, 0)).Solid() {
			foundGround = true
			break
		}
	}
	return foundGround
}

func (f *Finder) reconstructPath(endNode *Node) Path {
	nodes := []mmath.Pos{}
	curr := endNode
	for curr != nil {
		nodes = append([]mmath.Pos{curr.Pos}, nodes...)
		curr = curr.Parent
	}
	return Path{Nodes: nodes}
}

type priorityQueue []*Node
func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].F() < pq[j].F() }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any) { *pq = append(*pq, x.(*Node)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
