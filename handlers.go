package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/oleiade/lane.v1"
	"log"
	"net/http"
	"math/rand"
)

var taunts = []string {
	"Crush your enemies, see them driven before you, and hear the lamentations of their women!",
	"You've just been ERASED!!",
	"I'm the party pooper!",
	"Allow me to break the ~ice~.",
	"Who is your daddy, and what does he do??",
	"Hasta la vista, baby!!",
	"Talk to the hand.",
	"If it bleeds, we can kill it!",
	"I'll be back!",
}

func StartHandler(response http.ResponseWriter, request *http.Request) {
	start_request, _ := NewStartRequest(request)
	active_games[start_request.Game_id] = lane.NewDeque()
	response_data := BSResponse{
		"name":            "Snekenegger",
		"color":           "#0091BA",
		"taunt":           taunts[0],
		"head_type":       "tongue",
		"tail_type":       "small-rattle",
		"head_url":        "https://raw.githubusercontent.com/omazhary/chiseler-snake/personality/static/conanFace.png",
		"secondary_color": "#AA0F01",
	}
	json.NewEncoder(response).Encode(response_data)
	log.Println(fmt.Sprintf("Started game %d.", start_request.Game_id))
	log.Println(fmt.Sprintf("Total number of running games: %d.", len(active_games)))
}

func MoveHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Received move request.")
	world, _ := NewMoveRequest(request)
	tauntNum := rand.Intn(len(taunts))
	response_data := BSResponse{
		"move": Strategize(world),
		"taunt": taunts[tauntNum],
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
