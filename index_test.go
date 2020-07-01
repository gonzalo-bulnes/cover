package main

import (
	"testing"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/gonzalo-bulnes/cover/distance"
	"github.com/mandykoh/go-covertree"
)

func TestIndex(t *testing.T) {

	t.Run("known values", func(t *testing.T) {
		var words = corpus.New(
			"hello",
			"hullo",
			"allo?",
			"muelle",
			"anaconda",
		)

		tree := covertree.NewInMemoryTree(10.0, 3.0, distance.Levenshtein)

		inserted, err := index(tree, w...)
		if err != nil {
			t.Fatalf("Unexpected error when indexing: %v", err)
		}

		if expected := len(words); inserted != expected {
			t.Errorf("Expected %d insertions, got %d", expected, inserted)
		}
	})
}
