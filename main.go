package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gonzalo-bulnes/cover/corpus"
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

	inputTree := covertree.NewInMemoryTree(basis, rootDistance, distance.Distance)

	if print {
		fmt.Printf("\nIndexing phase.\n\n")
	}
	inserted, err := index(inputTree, words...)
	if err != nil {
		fmt.Printf("Error while indexing: %v\n", err)
	}
	if print {
		fmt.Printf("Inserted %d words\n", inserted)
	}

	if len(words) == 0 {
		return
	}

	w := words[0]
	if print {
		fmt.Printf("Finding the %d nearest words that are closer than %f from '%s'\n", maxResults, maxDistance, w)
	}

	if print {
		fmt.Printf("Querying candidates from %s.\n\n", w)
	}

	candidates, err := inputTree.FindNearest(&w, maxResults, maxDistance)
	if err != nil {
		fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
	}

	if print {
		fmt.Printf("Candidates %+v\n", candidates)
	}
	fmt.Printf("%d candidates\n", len(candidates))

	outputTree := covertree.NewInMemoryTree(basis, rootDistance, distance.Distance)

	var nSelected int

	// insert the query word
	inserted, err = index(outputTree, w)
	if err != nil {
		fmt.Printf("Error while indexing: %v\n", err)
	}
	if inserted != 1 {
		fmt.Printf("Error while indexing: %s was not inserted\n", w)
	}
	nSelected += inserted

	for _, c := range candidates {
		candidate := c.Item.(*corpus.Word)
		compatibleWords, err := outputTree.FindNearest(candidate, maxResults, maxDistance)
		if err != nil {
			fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
		}

		if len(compatibleWords) == nSelected {
			inserted, err := index(outputTree, *candidate)
			if err != nil {
				fmt.Printf("Error while indexing: %v\n", err)
			}
			if inserted != 1 {
				fmt.Printf("Error while indexing: %s was not inserted\n", candidate)
			}

			nSelected += inserted
		}
	}

	if print {
		fmt.Printf("Querying wordlist from %s.\n\n", w)
	}

	wordlist, err := outputTree.FindNearest(&w, maxResults, maxDistance)
	if err != nil {
		fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
	}

	if print {
		fmt.Printf("Potential wordlist %+v\n", wordlist)
	}
	fmt.Printf("%d words\n", nSelected)

	fmt.Printf("Any subset of this wordlist conforms to all the constraints, there are %d words to pick from.\n", nSelected)
	if nSelected >= 1296 {
		fmt.Println("There are enough words to compose a short list for use with 4 dice.")
	}
	if nSelected >= 7776 {
		fmt.Println("There are enough words to compose a long list for use with 5 dice.")
	}
}
