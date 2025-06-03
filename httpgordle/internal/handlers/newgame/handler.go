package newgame

import (
	"encoding/json"
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/session"
	"log"
	"net/http"
)

type Adder interface {
	Add(game session.Game) error
}

func Handler(adder Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		game, err := createGame()
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("faile to write response: %s", err)
		}
	}
}

func createGame() (session.Game, error) {
	return session.Game{}, nil
}
