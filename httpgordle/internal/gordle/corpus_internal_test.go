package gordle

import (
	"slices"
	"testing"
)

func TestPickWord(t *testing.T) {
	words := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word, err := PickRandomWord(words)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if !inCorpus(words, word) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}

func inCorpus(words []string, word string) bool {
	return slices.Contains(words, word)
}

func OverrideCorpus(newCorpus string) {
	corpus = newCorpus
}
