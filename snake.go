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
	obstructions = append(obstructions, state.You.Body.Data...)
	for _, opponent := range state.Snakes {
		obstructions = append(obstructions, opponent.Head())
		obstructions = append(obstructions, opponent.Body.Data...)
	}
	// Filter out moves based on obstructions
	for key, value := range moves {
		if value.InObstructions(obstructions) || value.IsOutOfMapBounds(state) {
			delete(moves, key)
		}
	}
	return moves
}
