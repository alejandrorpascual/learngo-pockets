package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader      *bufio.Reader
	corpus      []string
	solution    []rune
	maxAttempts int
}

// New returns a game, which can be used to Play!
func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	solution := pickWord(corpus)

	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g, nil
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}

		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb)
	}

	fmt.Printf("😞 You've lost! The solution was: %s.\n", string(g.solution))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerAnswer, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerAnswer))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// errInvalidWordLength is returned when the guess has the wrong number of
// characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't  have the same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution"+
			" have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	for i, char := range guess {
		if char == solution[i] {
			result[i] = correctPosition
			used[i] = true
			continue
		}
	}

	for i, char := range guess {
		if result[i] != absentCharacter {
			// char has already been marked, ignore it.
			continue
		}

		for j, targetChar := range solution {
			if used[j] {
				// The letter of the solution is already assigned to a
				// letter of the guess.
				// Skip to the next letter of the solution
				continue
			}

			if char == targetChar {
				result[i] = wrongPosition
				used[j] = true
				// skip to the next letter of the guess.
				break
			}

		}

	}

	return result
}
