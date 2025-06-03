package main

import (
	"fmt"
	"learngo-pockets/httpgordle/internal/handlers"
	"learngo-pockets/httpgordle/internal/repository"
	"net/http"
	"os"
)

func main() {
	db := repository.New()

	err := http.ListenAndServe(":8081", handlers.NewRouter(db))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
