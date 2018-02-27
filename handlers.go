package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/oleiade/lane.v1"
	"log"
	"net/http"
)

func StartHandler(response http.ResponseWriter, request *http.Request) {
	start_request, _ := NewStartRequest(request)
	active_games[start_request.Game_id] = lane.NewDeque()
	response_data := BSResponse{
		"name":            "Chiseler",
		"color":           "#0A3F71",
		"taunt":           "You've just been ERASED!!",
		"head_type":       "tongue",
		"tail_type":       "small-rattle",
		"head_url":        "https://raw.githubusercontent.com/omazhary/chiseler-snake/master/static/head.png",
		"secondary_color": "#FEB23B",
	}
	json.NewEncoder(response).Encode(response_data)
	log.Println(fmt.Sprintf("Started game %d.", start_request.Game_id))
	log.Println(fmt.Sprintf("Total number of running games: %d.", len(active_games)))
}

func MoveHandler(response http.ResponseWriter, request *http.Request) {
	world, _ := NewMoveRequest(request)
	response_data := BSResponse{
		"move": Strategize(world),
	}
	json.NewEncoder(response).Encode(response_data)
}
