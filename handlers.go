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
		"color":           "#0091BA",
		"taunt":           "You've just been ERASED!!",
		"head_type":       "tongue",
		"tail_type":       "small-rattle",
		"head_url":        "https://raw.githubusercontent.com/omazhary/chiseler-snake/master/static/head.png",
		"secondary_color": "#DD002F",
	}
	json.NewEncoder(response).Encode(response_data)
	log.Println(fmt.Sprintf("Started game %d.", start_request.Game_id))
	log.Println(fmt.Sprintf("Total number of running games: %d.", len(active_games)))
}

func MoveHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Received move request.")
	world, _ := NewMoveRequest(request)
	response_data := BSResponse{
		"move": Strategize(world),
	}
	json.NewEncoder(response).Encode(response_data)
	log.Println("Responded to move request.")
}

func EndHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Received end request.")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("200 - Game Over!"))
	log.Println("Responded to end request.")
}
