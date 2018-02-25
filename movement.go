package main

import (
	"math/rand"
)

var movement_map = map[int]string{
	1: "up",
	2: "right",
	3: "down",
	4: "left",
}

func Strategize(world *MoveRequest) string {

	myHeadLocation := world.You.Head()
	direction := FindDirection(world)
	nextPosition := GetNextPointBasedOnDirection(direction, myHeadLocation)
	for flag := false; flag == false; flag = IsValidPointToMoveTo(nextPosition, world) {
		direction = FindDirection(world)
		nextPosition = GetNextPointBasedOnDirection(direction, myHeadLocation)
	}
	return movement_map[direction]
}

func IsGoingToHitHimselfAtPoint(p Point, world *MoveRequest) bool {
	for _, bodyPoints := range world.You.Body.Data {
		if p.Equals(bodyPoints) {
			return true
		}
	}
	return false
}

func IsGoingToHitOthersAtPoint(p Point, world *MoveRequest) bool {
	for _, snake := range world.Snakes.Data {
		for _, bodyPoints := range snake.Body.Data {
			if p.Equals(bodyPoints) {
				return true
			}
		}
	}
	return false
}

func IsValidPointToMoveTo(p Point, world *MoveRequest) bool {
	if p.IsOutOfMapBounds(world) {
		return false
	}
	if IsGoingToHitHimselfAtPoint(p, world) {
		return false
	}
	return true
}

func GetNextPointBasedOnDirection(direction int, currentPoint Point) Point {
	nextPoint := currentPoint
	switch direction {
	case 1: //up
		nextPoint.Y = nextPoint.Y - 1
	case 2: //right
		nextPoint.X++
	case 3: // down
		nextPoint.Y = nextPoint.Y + 1
	case 4: // left
		nextPoint.X = nextPoint.X - 1
	default:
		nextPoint.X = nextPoint.X - 1
	}
	return nextPoint
}

func FindDirection(world *MoveRequest) int {
	return rand.Intn(4) + 1
}
