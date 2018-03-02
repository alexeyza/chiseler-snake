package main

import (
	"fmt"
	"log"
	"math"
)

var movement_map = map[int]string{
	1: "up",
	2: "right",
	3: "down",
	4: "left",
}

func Strategize(world *MoveRequest) string {
	minimum_score := math.MaxFloat64
	candidate_move := 0
	current_state := NewBoardStateFromMoveRequest(world)
	moves := current_state.You.GeneratePossibleMoves(current_state)
	states := GeneratePossibleStates(current_state, moves, world, 4)
	for key, state := range states {
		if state.Score < minimum_score {
			minimum_score = state.Score
			candidate_move = key
		}
		log.Println(fmt.Sprintf("%v - %v", key, state.Score))
	}
	return movement_map[candidate_move]
}
func GeneratePossibleStates(current_state *BoardState, possible_moves map[int]*Point, move_req *MoveRequest, levels int) map[int]*BoardState {
	possible_states := make(map[int]*BoardState)
	other_snake_moves := make(map[string]int)
	for _, snake := range current_state.Snakes {
		opponent_moves := snake.GeneratePossibleMoves(current_state)
		opponent_move := 0
		for key, _ := range opponent_moves {
			opponent_move = key
			break
		}
		other_snake_moves[snake.Id] = opponent_move
	}
	for move, _ := range possible_moves {
		possible_states[move] = current_state.NewBoardStateFromBoardState(move, other_snake_moves, move_req)
	}
	return possible_states
}
