package main

import "math"

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
func (point Point) InArray(array []Point) bool {
	for _, element := range array {
		if point.Equals(element) {
			return true
		}
	}
	return false
}
func (point Point) RemoveFromArray(array []Point) []Point {
	index := -1
	for i, value := range array {
		if point.Equals(value) {
			index = i
			break
		}
	}
	array[len(array)-1], array[index] = array[index], array[len(array)-1]
	return array[:len(array)-1]
}
func (point Point) GetDistance(other_point Point) float64 {
	return math.Sqrt(math.Pow(float64(point.X-other_point.X), 2.0) + math.Pow(float64(point.Y-other_point.Y), 2.0))
}
