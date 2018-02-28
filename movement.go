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

	current_state := NewBoardStateFromMoveRequest(world)
	moves := current_state.You.GeneratePossibleMoves(current_state)
	keys := make([]int, 0, len(moves))
	for k := range moves {
		keys = append(keys, k)
	}
	random_key := rand.Intn(len(keys))
	next_move := keys[random_key]
	return movement_map[next_move]
}

func (snake Snake) GeneratePossibleMoves(state *BoardState) map[int]Point {
	moves := make(map[int]Point)
	var obstructions []Point
	// Figure out head location
	head := state.You.Head()
	// Figure out movement direction(s)
	moves[1] = Point{head.X, head.Y - 1}
	moves[2] = Point{head.X + 1, head.Y}
	moves[3] = Point{head.X, head.Y + 1}
	moves[4] = Point{head.X - 1, head.Y}
	// Figure out obstructions in closest squares
	obstructions = append(obstructions, state.You.Body.Data...)
	for _, opponent := range state.Snakes {
		obstructions = append(obstructions, opponent.Head())
		obstructions = append(obstructions, opponent.Body.Data...)
	}
	// Filter out moves based on obstructions
	for key, value := range moves {
		if InObstructions(moves[key], obstructions) {
			delete(moves, key)
		}
		// Filter out moves based on border
		if value.X >= state.BoardWidth || value.X < 0 || value.Y >= state.BoardLength || value.Y < 0 {
			delete(moves, key)
		}
	}
	return moves
}
