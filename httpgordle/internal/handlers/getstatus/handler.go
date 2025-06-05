package getstatus

import (
	"encoding/json"
	"errors"
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/repository"
	"learngo-pockets/httpgordle/internal/session"
	"log"
	"net/http"
)

type gameFinder interface {
	Find(id session.GameID) (session.Game, error)
}

func Handler(finder gameFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusBadRequest)
			return
		}
		log.Printf("retrieve status of game with id: %v", id)

		game, err := finder.Find(session.GameID(id))
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				http.Error(w, "this game does not exist yet", http.StatusNotFound)
				return
			}

			log.Printf("cannot fetch game %s: %s", id, err)
			http.Error(w, "faled to fetch game", http.StatusInternalServerError)
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write responde: %s", err)
		}
	}
}
