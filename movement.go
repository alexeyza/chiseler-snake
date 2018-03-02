package main

import (
	"gopkg.in/oleiade/lane.v1"
	"math"
)

var movement_map = map[int]string{
	1: "up",
	2: "right",
	3: "down",
	4: "left",
}

// Adjust the health threshold and snake size that control when our snake needs to eat
var health_threshold = 50
var minimu_snake_size = 10

// This is the main movement method.
// Returns a string indication the next movement direction.
func Strategize(world *MoveRequest) string {

	// Get the location points of our snake's head, tail,
	// points near our tail, and our targed food location
	my_head_location := world.You.Head()
	my_tail_location := GetTail(world.You)
	near_tail_locations := GetValidAdjacentPoints(my_tail_location, world)
	food_location := FindFood(my_head_location, world)

	var path_map []int
	// first, check if we should aim for food
	if ShouldSearchForFood(world) {
		// find path towards the food
		path_map = ShortestPath(my_head_location, food_location, world)

		//if the path to food blocked, spin in place until path to food is clear
		if path_map == nil {
			for _, possible_target_destination := range near_tail_locations {
				path_map = ShortestPath(my_head_location, possible_target_destination, world)
				if path_map != nil {
					break
				}
			}
		}
		// if we don't need food, spin in place
	} else {
		for _, possible_target_destination := range near_tail_locations {
			path_map = ShortestPath(my_head_location, possible_target_destination, world)
			if path_map != nil {
				break
			}
		}
	}
	// if haven't found to either food or tail
	if path_map == nil {
		for i := 1; i < 5; i++ {
			if IsValidPointToMoveTo(GetNextPointBasedOnDirection(i, my_head_location), world) {
				path_map = []int{i}
			}
		}
		// if no valid point found, return up
		if path_map == nil {
			path_map = []int{1}
		}
	}
	return movement_map[path_map[0]]
}

// This method checks if a snake going to hit himself at the given point 'p'.
func IsGoingToHitHimselfAtPoint(p Point, snake Snake) bool {
	for _, body_points := range snake.Body.Data {
		if p.Equals(body_points) {
			return true
		}
	}
	return false
}

// This method checks if our snake may hit other snakes at the given point 'p'.
func IsGoingToHitOthersAtPoint(p Point, world *MoveRequest) bool {
	for _, enemy_snake := range world.Snakes.Data {
		// first, check if we hit any of the enemy snake bodies
		if IsGoingToHitHimselfAtPoint(p, enemy_snake) {
			return true
		}
		if IsRiskyPoint(p, world) {
			return true
		}
	}
	return false
}

// This method checks if the given point 'p' might be risky (e.g., close to other snakes.
// Note this doesn't check for collision, for that use the other method.
func IsRiskyPoint(p Point, world *MoveRequest) bool {
	for _, enemy_snake := range world.Snakes.Data {
		// Skip our snake in the list of all snakes
		if enemy_snake.Id == world.You.Id {
			continue
		}
		// next, check if we may hit an enemy snake's head,
		// and if that snake's health is higher than ours, mark this as invalid move
		for _, position_next_to_enemys_head := range GetAdjacentPoints(enemy_snake.Head()) {
			if position_next_to_enemys_head.Equals(p) && enemy_snake.Length >= world.You.Length {
				return true
			}
		}
	}
	return false
}

// This method checks if the given point 'p' is a valid move for our snake.
// it takes into account:
// - not being outside map bounds
// - not hitting himself
// - not hitting other snakes (or collide with snakes bigger than us)
func IsValidPointToMoveTo(p Point, world *MoveRequest) bool {
	if p.IsOutOfMapBounds(world) {
		return false
	}
	if IsGoingToHitHimselfAtPoint(p, world.You) {
		return false
	}
	if IsGoingToHitOthersAtPoint(p, world) {
		return false
	}
	return true
}

// This method generates a point coordinates based on a given point and a direction.
// For instance, give it current location and direction, it will return the next location.
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

// This method returns a location of food.
// If multiple foods exists, it will go for the closes one at the moment (but we may change this here).
func FindFood(location Point, world *MoveRequest) Point {
	closest_distance := math.MaxFloat64
	closest_food := world.Food.Data[0]
	for _, food_source := range world.Food.Data {
		dist := location.distance(food_source)
		if dist < closest_distance {
			closest_distance = dist
			closest_food = food_source
		}
	}
	return closest_food
}

// This method returns a path towards the given destination (returns an array of directions).
// If a path to the given destination was not found, returns nil
// This is a standard iterative BFS algorithm
func ShortestPath(source Point, destination Point, world *MoveRequest) []int {
	queue := lane.NewDeque()
	var parent Point
	visited := map[Point]bool{}
	plan_to_visit := map[Point]bool{}
	possible_directions := []int{1, 2, 3, 4}
	map_of_paths_to_any_point := map[Point][]int{}

	// start BFS by queuing the source point
	queue.Prepend(source)

	// while there are neighboring/adjacent points we haven't visited yet
	for !queue.Empty() {
		// pop the current element, and mark it as "not plan to visit through other nodes"
		parent, _ = queue.Pop().(Point)
		plan_to_visit[parent] = false

		// if we reached destination, stop
		if parent.Equals(destination) {
			// return the list of directions from source to destination
			// note: we actually have a list of directions from the source to any other point on the board
			return map_of_paths_to_any_point[destination]
		}

		// for every neighboring/adjacent point to the current point
		// it iterates over the four directions: up, down, left, right
		for _, next_move := range possible_directions {
			next_position := GetNextPointBasedOnDirection(next_move, parent)

			// if the neighbor is an invalid point (e.g., wall, other snake)
			if !IsValidPointToMoveTo(next_position, world) {
				continue
			}
			// if already visited this neighbor, skip it
			if visited[next_position] {
				continue
			}
			// If haven't seen this neighbor before, add it to "plan to visit"
			// and document the direction we'd need to take to reach it,
			// appending it to the directions we've already taken
			if !plan_to_visit[next_position] {
				queue.Prepend(next_position)
				plan_to_visit[next_position] = true
				map_of_paths_to_any_point[next_position] = append(map_of_paths_to_any_point[parent], next_move)
			}
		}
		// mark the current point as visited
		visited[parent] = true
	}

	return nil
}

// This method returns the tail point of a given snake.
func GetTail(snake Snake) Point {
	return snake.Body.Data[snake.Length-1]
}

// This method returns the adjacent points, based on the given point
// Does NOT check if they are valid points to move to.
func GetAdjacentPoints(point Point) []Point {
	output := []Point{
		Point{X: point.X + 1, Y: point.Y},
		Point{X: point.X, Y: point.Y + 1},
		Point{X: point.X - 1, Y: point.Y},
		Point{X: point.X, Y: point.Y - 1},
	}
	return output
}

// This method returns the adjacent points. Only returns valid points.
func GetValidAdjacentPoints(point Point, world *MoveRequest) []Point {
	var output []Point
	for _, adj_point := range GetAdjacentPoints(point) {
		if IsValidPointToMoveTo(adj_point, world) {
			output = append(output, adj_point)
		}
	}
	return output
}

// This method returns true if our snake should start looking for food.
func ShouldSearchForFood(world *MoveRequest) bool {

	my_lenght := world.You.Length

	// check if our length is lower than the other snakes, if yes, find food to grow
	for _, snake := range world.Snakes.Data {
		// skip our snake
		if snake.Id == world.You.Id {
			continue
		}
		if snake.Length >= my_lenght {
			return true
		}
	}

	// check if our health below a threshold, or if we haven't reached minimum snake size
	return world.You.Health < health_threshold || world.You.Length < minimu_snake_size
}
