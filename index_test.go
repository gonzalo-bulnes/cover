package main

import (
	"testing"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/gonzalo-bulnes/cover/distance"
	"github.com/mandykoh/go-covertree"
)

var w = corpus.New(
	"hello",
	"hullo",
	"allo?",
	"muelle",
	"anaconda",
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

func BenchmarkIndex(b *testing.B) {
	benchmarkIndex(nil, b)
}

func benchmarkIndex(_ *covertree.Tree, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		tree := covertree.NewInMemoryTree(10.0, 3.0, distance.Levenshtein)
		b.StartTimer()
		index(tree, w...)
	}
}

func BenchmarkFindNearest(b *testing.B) {
	tree := covertree.NewInMemoryTree(10.0, 3.0, distance.Levenshtein)
	word := corpus.NewWord("hello")
	benchmarkFindNearest(tree, &word, b)
}

func benchmarkFindNearest(tree *covertree.Tree, word *corpus.Word, b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.FindNearest(word, maxResults, maxDistance)
	}
}
