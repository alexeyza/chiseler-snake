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
