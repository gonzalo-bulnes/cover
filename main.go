package main

import (
	"fmt"
	"os"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/gonzalo-bulnes/cover/distance"
	"github.com/gonzalo-bulnes/cover/file"
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

var print = os.Getenv("PRINT") != ""

func main() {
	words, err := file.NewCorpus()
	if err != nil {
		fmt.Printf("Error when importing file: %v", err)
	}

	tree := covertree.NewInMemoryTree(basis, rootDistance, distance.ForErrorCorrection)

	if print {
		fmt.Printf("\nIndexing phase.\n\n")
	}
	inserted, err := index(tree, words...)
	if err != nil {
		fmt.Printf("Error while indexing: %v\n", err)
	}
	if print {
		fmt.Printf("Inserted %d words\n", inserted)
	}

	if print {
		fmt.Printf("\nQuerying phase.\n\n")
	}
	w := corpus.NewWord("hello")
	if print {
		fmt.Printf("Finding the %d nearest words that are closer than %f from '%s'\n", maxResults, distance.MaxEditDistanceForErrorCorrection(3), w)
	}
	results, err := tree.FindNearest(&w, maxResults, maxDistance)
	if err != nil {
		fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
	}

	if print {
		fmt.Printf("Results %+v\n", results)
	}
}
