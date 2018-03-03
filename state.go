package main

import "log"
import "math"
import "time"

type BoardState struct {
	BoardLength   int
	BoardWidth    int
	CandidateMove int
	Children      []*BoardState
	Food          []Point
	ParentState   *BoardState
	Score         float64
	Snakes        []Snake
	You           Snake
}

func NewBoardStateFromMoveRequest(move_req *MoveRequest) *BoardState {
	state := BoardState{
		BoardLength: move_req.Height,
		BoardWidth:  move_req.Width,
		Food:        move_req.Food.Data,
		Snakes:      move_req.Snakes.Data,
		You:         move_req.You,
		ParentState: nil,
	}
	return &state
}
func (state *BoardState) NewBoardStateFromBoardState(direction int, others map[string]int, move_req *MoveRequest) *BoardState {
	new_state := BoardState{
		BoardLength: state.BoardLength,
		BoardWidth:  state.BoardWidth,
		Food:        state.Food,
		ParentState: state,
		Snakes:      state.Snakes,
		You:         state.You,
	}
	var taken_spaces []Point
	for _, snake := range new_state.Snakes {
		if !snake.Equals(new_state.You) {
			snake.MoveInDirection(others[snake.Id], move_req)
		}
		taken_spaces = append(taken_spaces, snake.Body.Data...)
	}
	new_state.You.MoveInDirection(direction, move_req)
	taken_spaces = append(taken_spaces, new_state.You.Body.Data...)
	for _, taken_space := range taken_spaces {
		if taken_space.InArray(new_state.Food) {
			new_state.Food = taken_space.RemoveFromArray(new_state.Food)
		}
	}
	new_state.ScoreState(move_req, 1)
	state.Children = append(state.Children, &new_state)
	return &new_state
}
func (state *BoardState) GetDistanceFromClosestFood() (float64, Point) {
	minimum_distance := math.MaxFloat64
	var candidate_food Point
	for _, food_item := range state.Food {
		if food_item.GetDistance(state.You.Head()) < minimum_distance {
			minimum_distance = math.MaxFloat64 - food_item.GetDistance(state.You.Head())
			candidate_food = food_item
		}
	}
	return minimum_distance, candidate_food
}
func (state *BoardState) ScoreState(move_req *MoveRequest, levels int) float64 {
	log.Println(time.Now())
	if len(state.Children) == 0 || levels == 0 {
		score, _ := state.GetDistanceFromClosestFood()
		state.Score = score
		log.Println(time.Now())
		return state.Score
	} else {
		var child_scores []float64
		new_level := levels - 1
		for _, child := range state.Children {
			child_scores = append(child_scores, child.ScoreState(move_req, new_level))
			var minimum_child_score float64
			for _, child_score := range child_scores {
				if child_score < minimum_child_score {
					minimum_child_score = child_score
				}
			}
		}
		return state.Score
	}
}
