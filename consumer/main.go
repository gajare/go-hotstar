package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CricketMatch struct {
	MatchID string `json:"match_id"`
	Score   string `json:"score"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/consume", consumeHandler).Methods("GET")
	http.Handle("/", r)
	log.Println("Consumer service is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func consumeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://producer:8080/produce")
	if err != nil {
		http.Error(w, "Failed to fetch data from producer", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var matches []CricketMatch
	if err := json.Unmarshal(body, &matches); err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}
