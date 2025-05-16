package main

import (
	"fmt"
	"learngo-pockets/gordle/gordle"
	"os"
)

func main() {
	corpus, err := gordle.ReadCorpus("./corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	g, err := gordle.New(corpus, gordle.WithReader(os.Stdin), gordle.WithMaxAttempts(4))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
		return
	}
	g.Play()

}
