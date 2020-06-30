package main

import (
	"fmt"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/mandykoh/go-covertree"
)

func index(tree *covertree.Tree, words []corpus.Word) (inserted int, err error) {
	for _, w := range words {
		x := w // copy is required
		err := tree.Insert(&x)
		if err != nil {
			return inserted, fmt.Errorf("error inserting %s: %w", w, err)
		}
		inserted++
	}
	return
}
