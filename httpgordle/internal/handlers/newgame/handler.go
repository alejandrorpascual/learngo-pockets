package newgame

import (
	"encoding/json"
	"learngo-pockets/httpgordle/internal/api"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := api.GameResponse{}
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("faile to write response: %s", err)
	}
}
