package guess

import (
	"encoding/json"
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
	}

	guessR := api.GuessRequest{}
	err := json.NewDecoder(r.Body).Decode(&guessR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game := guess(id, guessR)

	apiGame := api.ToGameResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func guess(id string, guessR api.GuessRequest) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
