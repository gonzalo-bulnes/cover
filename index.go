package main

import (
	"fmt"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/mandykoh/go-covertree"
)

func index(tree *covertree.Tree, words ...corpus.Word) (inserted int, err error) {
	for i := range words {
		err := tree.Insert(&words[i])
		if err != nil {
			return inserted, fmt.Errorf("error inserting %s: %w", words[i], err)
		}
		inserted++
	}
	return
}
