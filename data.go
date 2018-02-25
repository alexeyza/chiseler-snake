package main

import (
	"encoding/json"
	//"errors"
	"net/http"
	//	"strings"
)

// type GameStartRequest struct {
// 	GameId string `json:"game_id"`
// 	Height int    `json:"height"`
// 	Width  int    `json:"width"`
// }

// type GameStartResponse struct {
// 	Color          string  `json:"color"`
// 	HeadUrl        *string `json:"head_url,omitempty"`
// 	Name           string  `json:"name"`
// 	Taunt          *string `json:"taunt,omitempty"`
// 	HeadType       string  `json:"head_type"`
// 	TailType       string  `json:"tail_type"`
// 	SecondaryColor string  `json:"secondary_color"`
// }

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

// type MoveResponse struct {
// 	Move  string  `json:"move"`
// 	Taunt *string `json:"taunt,omitempty"`
// }

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

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	decoded := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

// func NewGameStartRequest(req *http.Request) (*GameStartRequest, error) {
// 	decoded := GameStartRequest{}
// 	err := json.NewDecoder(req.Body).Decode(&decoded)
// 	return &decoded, err
// }

// func (snake Snake) Head() Point { return snake.Coords[0] }

// Decode [number, number] JSON array into a Point
// func (point *Point) UnmarshalJSON(data []byte) error {
// 	var coords []int
// 	json.Unmarshal(data, &coords)
// 	if len(coords) != 2 {
// 		return errors.New("Bad set of coordinates: " + string(data))
// 	}
// 	*point = Point{X: coords[0], Y: coords[1]}
// 	return nil
// }
