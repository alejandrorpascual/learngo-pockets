package handlers

import (
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
	"net/http"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)

	return mux
}
