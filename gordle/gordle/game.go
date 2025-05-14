package gordle

import "fmt"

// Game holds all the information we need to play a game of gordle.
type Game struct {
}

// New returns a game, which can be used to Play!
func New() *Game {
	g := &Game{}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess:\n")
}
