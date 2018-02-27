package main

import (
	"math/rand"

	"gopkg.in/oleiade/lane.v1"
)

var movement_map = map[int]string{
	1: "up",
	2: "right",
	3: "down",
	4: "left",
}

var move_queue = lane.NewDeque()

func Strategize(world *MoveRequest) string {

	myHeadLocation := world.You.Head()
	foodLocation := world.Food.Data[0]

	if move_queue.Empty() {
		SimplePath(myHeadLocation, foodLocation)
	}
	next_move, _ := move_queue.Pop().(int)
	next_position := GetNextPointBasedOnDirection(next_move, myHeadLocation)

	for flag := IsValidPointToMoveTo(next_position, world); flag == false; flag = IsValidPointToMoveTo(next_position, world) {
		next_move = rand.Intn(4) + 1
		next_position = GetNextPointBasedOnDirection(next_move, myHeadLocation)
	}

	return movement_map[next_move]
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
	if IsGoingToHitOthersAtPoint(p, world) {
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

func SimplePath(source Point, destination Point) {
	xdimension := source.X - destination.X
	ydimension := source.Y - destination.Y

	if xdimension < 0 {
		for i := xdimension; i < 0; i++ {
			//fmt.Println("right")
			move_queue.Prepend(2)
		}
	} else {
		for i := 0; i < xdimension; i++ {
			//fmt.Println("left")
			move_queue.Prepend(4)
		}
	}
	if ydimension < 0 {
		for i := ydimension; i < 0; i++ {
			//fmt.Println("down")
			move_queue.Prepend(3)
		}
	} else {
		for i := 0; i < ydimension; i++ {
			//fmt.Println("up")
			move_queue.Prepend(1)
		}
	}
}
