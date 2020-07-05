package main

import (
	"fmt"
	"testing"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/gonzalo-bulnes/cover/distance"
	"github.com/gonzalo-bulnes/cover/file"
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

	words, err := file.NewCorpus()
	if err != nil {
		fmt.Printf("Error when importing file: %v", err)
	}

	t.Run("known values", func(t *testing.T) {
		tree := covertree.NewInMemoryTree(10.0, 3.0, distance.Levenshtein)

		inserted, err := index(tree, words...)
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
	benchmarkFindNearest(tree, word, b)
}

func benchmarkFindNearest(tree *covertree.Tree, word *corpus.Word, b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.FindNearest(word, maxResults, maxDistance)
	}
}
