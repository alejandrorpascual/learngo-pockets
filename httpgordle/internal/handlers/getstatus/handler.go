package getstatus

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
		return
	}
	log.Printf("retrieve status of game with id: %v", id)

	apiGame := api.GameResponse{ID: id}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write responde: %s", err)
	}
}
