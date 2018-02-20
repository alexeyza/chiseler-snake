package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to our BattleSnake server for 2018!!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/start", StartHandler)
	http.HandleFunc("/move", MoveHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	log.Println(fmt.Sprintf("Running server on port %s...", port))
	http.ListenAndServe(":"+port, nil)
}
