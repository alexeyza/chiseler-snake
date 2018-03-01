package main

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
