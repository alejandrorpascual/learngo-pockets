package main

import (
	"fmt"
	"learngo-pockets/httpgordle/internal/handlers"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":8081", handlers.Mux())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
