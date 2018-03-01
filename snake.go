package main

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

func (snake Snake) Head() Point { return snake.Body.Data[0] }
