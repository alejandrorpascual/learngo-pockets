package handlers

import (
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/handlers/getstatus"
	"learngo-pockets/httpgordle/internal/handlers/guess"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
	"learngo-pockets/httpgordle/internal/repository"
	"net/http"
)

func NewRouter(db *repository.GameRepository) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))
	mux.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db))
	mux.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))

	return mux
}
