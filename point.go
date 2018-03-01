package main

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Point) IsOutOfMapBounds(state *BoardState) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}
	if p.Y >= state.BoardLength || p.X >= state.BoardWidth {
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
