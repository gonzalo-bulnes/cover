package main

import (
	"fmt"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/gonzalo-bulnes/cover/distance"
	"github.com/mandykoh/go-covertree"
)

const (
	// basis is the logarithmic base for determining the coverage of nodes at each level of the tree.
	basis = 10.0
	// rootDistance is the minimum expected distance between root nodes. New nodes that exceed this distance will be created as additional roots.
	rootDistance = 3.0
	// maxDistance is the maximum distnce acceptable in a result set.
	maxDistance = 6.0
	// maxResults is the maximum number of results acceptable in a result set.
	maxResults = 30
)

var words = corpus.New(
	"hello",
	"hullo",
	"allo?",
	"muelle",
	"anaconda",
)

func main() {
	tree := covertree.NewInMemoryTree(basis, rootDistance, distance.Levenshtein)

	fmt.Printf("\nIndexing phase.\n\n")
	for _, w := range words {
		x := w // copy is required
		err := tree.Insert(&x)
		if err != nil {
			fmt.Printf("Error inserting '%s': %v\n", w.String(), err)
		}
		fmt.Printf("Inserted '%+v'\n", &w)
	}

	fmt.Printf("\nQuerying phase.\n\n")
	w := corpus.NewWord("hello")
	fmt.Printf("Finding the %d nearest words that are closer than %f from '%s'\n", maxResults, maxDistance, w)
	results, err := tree.FindNearest(&w, maxResults, maxDistance)
	if err != nil {
		fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
	}

	fmt.Printf("Results %+v\n", results)
}
