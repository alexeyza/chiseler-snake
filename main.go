package main

import (
	"fmt"
	"gopkg.in/oleiade/lane.v1"
	"io"
	"log"
	"net/http"
	"os"
)

var active_games map[int]*lane.Deque

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to our BattleSnake server for 2018!!")
}

func main() {
	active_games = make(map[int]*lane.Deque)
	http.HandleFunc("/", hello)
	http.HandleFunc("/start", StartHandler)
	http.HandleFunc("/move", MoveHandler)
	http.HandleFunc("/end", EndHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	log.Println(fmt.Sprintf("Running server on port %s...", port))
	http.ListenAndServe(":"+port, nil)
}
