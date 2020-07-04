package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gonzalo-bulnes/cover/distance"
	"github.com/gonzalo-bulnes/cover/file"
	"github.com/mandykoh/go-covertree"
)

var (
	// basis is the logarithmic base for determining the coverage of nodes at each level of the tree.
	basis = distance.Max
	// maxDistance is the maximum distance acceptable in a result set.
	maxDistance = 7.0
	// rootDistance is the minimum expected distance between root nodes. New nodes that exceed this distance will be created as additional roots.
	rootDistance = maxDistance
	// maxResults is the maximum number of results acceptable in a result set.
	maxResults = 10000
)

var print = os.Getenv("PRINT") != ""

func main() {

	if m := os.Getenv("MAX_DISTANCE"); m != "" {
		f, err := strconv.ParseFloat(m, 64)
		if err != nil {
			fmt.Printf("Invalid MAX_DISTANCE %s: %v", m, err)
		} else {
			maxDistance = f
		}
	}

	words, err := file.NewCorpus()
	if err != nil {
		fmt.Printf("Error when importing file: %v", err)
	}

	tree := covertree.NewInMemoryTree(basis, rootDistance, distance.Distance)

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

	if len(words) == 0 {
		return
	}

	w := words[0]
	if print {
		fmt.Printf("Finding the %d nearest words that are closer than %f from '%s'\n", maxResults, maxDistance, w)
	}
	results, err := tree.FindNearest(&w, maxResults, maxDistance)
	if err != nil {
		fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
	}

	if print {
		fmt.Printf("Results %+v\n", results)
	}
	fmt.Printf("%d results\n", len(results))
}
