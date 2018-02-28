package main

import (
	"encoding/json"
	"math"
	"net/http"
)

// This file includes all the data structures for our game board and game objects
//

type BSResponse map[string]interface{}
type StartRequest struct {
	Game_id int `json: "id"`
}

// This is the world object
type MoveRequest struct {
	Food struct {
		Data []Point `json:"data"`
	} `json:"food"`
	Id     int `json:"id"`
	Height int `json:"height"`
	Width  int `json:"width"`
	Turn   int `json:"turn"`
	Snakes struct {
		Data []Snake `json:"data"`
	} `json:"snakes"`
	You Snake `json:"you"`
}

// This is a single point with two coordinates: X and Y
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// This is the snake object, represents snakes in the game (including ours)
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

func NewStartRequest(req *http.Request) (*StartRequest, error) {
	decoded := StartRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}
func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	decoded := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

// This class method returns the head of this snake.
func (snake Snake) Head() Point {
	return snake.Body.Data[0]
}

// This class method checks if this point is outside the map bounds
func (p *Point) IsOutOfMapBounds(world *MoveRequest) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}
	if p.Y >= world.Height || p.X >= world.Width {
		return true
	}
	return false
}

// This class method checks if this point is equal (in values) to the given one
func (p *Point) Equals(q Point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}

// This class method returns distance between this point and the given point
func (p1 *Point) distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
