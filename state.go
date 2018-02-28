package main

import "log"

type BoardState struct {
	BoardLength int
	BoardWidth  int
	Food        []Point
	Snakes      []Snake
	You         Snake
}

func NewBoardStateFromMoveRequest(move_req *MoveRequest) *BoardState {
	state := BoardState{
		BoardLength: move_req.Height,
		BoardWidth:  move_req.Width,
		Food:        move_req.Food.Data,
		Snakes:      move_req.Snakes.Data,
		You:         move_req.You,
	}
	return &state
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
	log.Println(obstructions)
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

func InObstructions(point Point, obstructions []Point) bool {
	for _, obstruction := range obstructions {
		if point.Equals(obstruction) {
			return true
		}
	}
	return false
}
