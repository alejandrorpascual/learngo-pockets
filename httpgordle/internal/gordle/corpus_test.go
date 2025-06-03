package gordle_test

import (
	"errors"
	"learngo-pockets/httpgordle/internal/gordle"
	"testing"
)

func TestParseCorpus(t *testing.T) {
	words, err := gordle.ParseCorpus()
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	const wordsInEnglishCorpus = 34
	if len(words) != wordsInEnglishCorpus {
		t.Errorf("expected %d word, got %d", wordsInEnglishCorpus, len(words))
	}
}

func TestParseCorpus_emptyCorpus(t *testing.T) {
	gordle.OverrideCorpus("")
	words, err := gordle.ParseCorpus()
	if !errors.Is(err, gordle.ErrEmptyCorpus) {
		t.Errorf("expected error %s, got %s", gordle.ErrEmptyCorpus, err)
	}

	if len(words) != 0 {
		t.Errorf("expected words (empty corpus), got %d", len(words))
	}
}
