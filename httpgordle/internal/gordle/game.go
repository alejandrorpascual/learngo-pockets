package gordle

import (
	"fmt"
	"os"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	solution []rune
}

// New returns a game, which can be used to Play!
func New(solution string) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrEmptyCorpus
	}

	g := &Game{
		solution: splitToUppercaseCharacters(solution),
	}

	return g, nil
}

const (
	//
	ErrInvalidGuessLength = gameError("invalid guess length")
)

func (g *Game) Play(guess string) (Feedback, error) {
	err := g.validateGuess(guess)
	if err != nil {
		return Feedback{}, fmt.Errorf("this guess is not the correct length: %w", err)
	}

	characthers := splitToUppercaseCharacters(guess)
	fb := computeFeedback(characthers, g.solution)

	return fb, nil
}

func (g *Game) ShowAnswer() string {
	return string(g.solution)
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess string) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), ErrInvalidGuessLength)
	}

	return nil
}

func computeFeedback(guess, solution []rune) Feedback {
	result := make(Feedback, len(guess))
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
