package main

import (
	"learngo-pockets/gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin, "hello", 4)
	g.Play()
}
