package gordle

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"strings"
)

const (
	ErrEmptyCorpus    = corpusError("corpus is empty")
	ErrPickRandomWord = corpusError("failed to pick a random word")
)

//go:embed corpus/english.txt
var corpus string

func ParseCorpus() ([]string, error) {
	words := strings.Fields(corpus)

	if len(words) == 0 {
		return nil, ErrEmptyCorpus
	}

	return words, nil
}

func PickRandomWord(corpus []string) (string, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(corpus))))
	if err != nil {
		return "", fmt.Errorf("failed to rand index (%s): %w", err, ErrPickRandomWord)
	}

	return corpus[index.Int64()], nil
}
