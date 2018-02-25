package main

import (
	"encoding/json"
	"net/http"
)

type BSResponse map[string]interface{}

func StartHandler(response http.ResponseWriter, request *http.Request) {
	responseData := BSResponse{
		"name":            "Chiseler",
		"color":           "#0A3F71",
		"taunt":           "You've just been ERASED!!",
		"head_type":       "tongue",
		"tail_type":       "small-rattle",
		"head_url":        "",
		"secondary_color": "#FEB23B",
	}
	json.NewEncoder(response).Encode(responseData)
}

func MoveHandler(response http.ResponseWriter, request *http.Request) {
	world, _ := NewMoveRequest(request)
	responseData := BSResponse{
		"move": Strategize(world),
	}
	json.NewEncoder(response).Encode(responseData)
}
