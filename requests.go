package main

import (
	"encoding/json"
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
