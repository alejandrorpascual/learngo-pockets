package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

// New returns a game, which can be used to Play!
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerAnswer, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := []rune(string(playerAnswer))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution!\nExpected %d characters, got %d.\n", solutionLength, len(guess))
		} else {
			return guess
		}
	}
}
