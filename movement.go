package main

import (
	"fmt"
	"log"
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
	log.Println(fmt.Sprintf("Returned %v possible moves...", len(moves)))
	log.Println(moves)
	keys := make([]int, 0, len(moves))
	for k := range moves {
		keys = append(keys, k)
	}
	random_key := rand.Intn(len(keys))
	log.Println(keys)
	next_move := keys[random_key]
	log.Println(fmt.Printf("Next Move: %v", movement_map[next_move]))
	return movement_map[next_move]
}
