package main

type Snake struct {
	Body struct {
		Data []Point `json:"data"`
	} `json:"body"`
	Health int    `json:"health"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Taunt  string `json:"taunt"`
	Length int    `json:"length"`
}

func (snake Snake) Head() Point { return snake.Body.Data[0] }
func (snake Snake) GeneratePossibleMoves(state *BoardState) map[int]*Point {
	moves := make(map[int]*Point)
	var obstructions []Point
	// Figure out head location
	head := state.You.Head()
	// Figure out movement direction(s)
	moves[1] = &Point{head.X, head.Y - 1}
	moves[2] = &Point{head.X + 1, head.Y}
	moves[3] = &Point{head.X, head.Y + 1}
	moves[4] = &Point{head.X - 1, head.Y}
	// Figure out obstructions in closest squares
	obstructions = append(obstructions, snake.Body.Data...)
	for _, opponent := range state.Snakes {
		obstructions = append(obstructions, opponent.Head())
		obstructions = append(obstructions, opponent.Body.Data...)
	}
	// Filter out moves based on obstructions
	for key, value := range moves {
		if value.InArray(obstructions) || value.IsOutOfMapBounds(state) {
			delete(moves, key)
		}
	}
	return moves
}
func (snake *Snake) MoveInDirection(direction int, world *MoveRequest) {
	var previous_point Point
	has_eaten := false
	for index, point := range snake.Body.Data {
		if index == 0 {
			previous_point = point
			if point.InArray(world.Food.Data) {
				has_eaten = true
			}
			if direction == 1 {
				snake.Body.Data[index].Y -= 1
			} else if direction == 2 {
				snake.Body.Data[index].X += 1
			} else if direction == 3 {
				snake.Body.Data[index].Y += 1
			} else if direction == 4 {
				snake.Body.Data[index].X -= 1
			}
		} else {
			temp := previous_point
			previous_point = snake.Body.Data[index]
			snake.Body.Data[index] = temp
		}
		if has_eaten && (index+1) == len(snake.Body.Data) {
			snake.Body.Data = append(snake.Body.Data, previous_point)
		}
	}
}
func (snake Snake) Equals(other_snake Snake) bool {
	if snake.Id == other_snake.Id {
		return true
	}
	return false
}
