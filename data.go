package main

import (
	"encoding/json"
	"net/http"
)

type BSResponse map[string]interface{}
type StartRequest struct {
	Game_id int `json: "id"`
}
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
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
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
func (snake Snake) Head() Point { return snake.Body.Data[0] }
func (p *Point) IsOutOfMapBounds(world *MoveRequest) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}
	if p.Y >= world.Height || p.X >= world.Width {
		return true
	}
	return false
}
func (p *Point) Equals(q Point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}
