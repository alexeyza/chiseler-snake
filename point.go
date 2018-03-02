package main

import "math"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Point) IsOutOfMapBounds(world *MoveRequest) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}
	if p.Y >= world.Height || p.X >= world.Width {
		return true
	}
	return false
}
func (p *Point) Equals(q Point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}
func (point Point) InObstructions(obstructions []Point) bool {
	for _, obstruction := range obstructions {
		if point.Equals(obstruction) {
			return true
		}
	}
	return false
}

// This class method returns distance between this point and the given point
func (p1 *Point) distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
