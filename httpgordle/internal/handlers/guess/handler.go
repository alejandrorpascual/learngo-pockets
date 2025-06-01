package guess

import (
	"encoding/json"
	"learngo-pockets/httpgordle/internal/api"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
	}

	guess := api.GuessRequest{}
	err := json.NewDecoder(r.Body).Decode(&guess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Thid was the guess: %v", guess)

	apiGamge := api.GameResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGamge)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}
