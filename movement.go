package main

import "math/rand"

var movement_map = map[int]string{
	1: "up",
	2: "right",
	3: "down",
	4: "left",
}

func strategize() string {
	return movement_map[rand.Intn(4)+1]
}
