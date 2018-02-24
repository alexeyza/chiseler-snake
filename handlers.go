package main

import (
	"encoding/json"
	"net/http"
)

type BSResponse map[string]interface{}

func StartHandler(response http.ResponseWriter, request *http.Request) {
	responseData := BSResponse{
		"name":            "Chiseler",
		"color":           "#6495ED",
		"taunt":           "You've just been ERASED!!",
		"head_type":       "sand-worm",
		"tail_type":       "fat-rattle",
		"head_url":        "",
		"secondary_color": "",
	}
	json.NewEncoder(response).Encode(responseData)
}

func MoveHandler(response http.ResponseWriter, request *http.Request) {
	Strategize(response, request)
}

func Strategize(response http.ResponseWriter, request *http.Request) {
	// move code goes here
}
