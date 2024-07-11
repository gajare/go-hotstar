package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CricketMatch struct {
	MatchID string `json:"match_id"`
	Score   string `json:"score"`
}

var matches []CricketMatch

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/produce", produceHandler).Methods("GET")
	http.Handle("/", r)
	log.Println("Producer service is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func produceHandler(w http.ResponseWriter, r *http.Request) {
	match := CricketMatch{
		MatchID: "1",
		Score:   "TeamA 250/3",
	}
	matches = append(matches, match)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}
