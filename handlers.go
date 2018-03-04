package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/oleiade/lane.v1"
	"log"
	"math/rand"
	"net/http"
)

var CurrentTaunt string

func getTaunt(world *MoveRequest) string {
	snakeData := world.Snakes.Data
	tauntNumber := rand.Intn(8)

	enemySnake := snakeData[tauntNumber % len(snakeData)]

	if(enemySnake.Id == world.You.Id) {
		tauntNumber += 1
		enemySnake = snakeData[tauntNumber % len(snakeData)]
	}

	enemySnakeName := enemySnake.Name

	var taunts = []string {
		"Crush your enemies, see them driven before you... etc",
		enemySnakeName + ", you've just been ERASED!!",
		"I'm the party pooper!",
		enemySnakeName + " - who is your daddy, and what does he do??",
		"Hasta la vista, baby!!",
		"Talk to the hand, " + enemySnakeName,
		"If " + enemySnakeName + " bleeds, we can kill it!",
		"I'll be back!",
	}

	return taunts[tauntNumber]
}

func StartHandler(response http.ResponseWriter, request *http.Request) {
	start_request, _ := NewStartRequest(request)
	active_games[start_request.Game_id] = lane.NewDeque()
	response_data := BSResponse{
		"name":            "Snekkenegger",
		"color":           "#AA0F01",
		"taunt":           "Allow me to break the ~ice~.",
		"head_type":       "tongue",
		"tail_type":       "small-rattle",
		"head_url":        "https://raw.githubusercontent.com/omazhary/chiseler-snake/personality/static/conanFace.png",
		"secondary_color": "#000000",
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
	if world.Turn % 10 == 0 {
		CurrentTaunt = getTaunt(world)
	}
	response_data["taunt"] = CurrentTaunt
	json.NewEncoder(response).Encode(response_data)
	log.Println("Responded to move request.")
}

func EndHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Received end request.")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("200 - Game Over!"))
	log.Println("Responded to end request.")
}
