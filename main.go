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

	config := distance.NewConfig()

	// wordlist := usingCoverTree(config, true)
	// wordlist := usingCoverTree(config, false)
	wordlist := reference(config)

	n := len(wordlist)
	fmt.Printf("Any subset of this wordlist conforms to all the constraints, there are %d words to pick from.\n", n)
	if n >= 1296 {
		fmt.Println("There are enough words to compose a short list for use with 4 dice.")
	}
	if n >= 7776 {
		fmt.Println("There are enough words to compose a long list for use with 5 dice.")
	}
}

func reference(config distance.Config) []*corpus.Word {
	words, err := file.NewCorpus()
	if err != nil {
		fmt.Printf("Error when importing file: %v", err)
	}

	wordlist := []*corpus.Word{}
	for i, candidate := range words {

		if i == 0 {
			wordlist = append(wordlist, candidate)
			continue
		}

		suitable := true
		for _, word := range wordlist {
			if distance.DistanceWithOptions(config)(word, candidate) > maxDistance {
				suitable = false
				break
			}
		}
		if suitable {
			wordlist = append(wordlist, candidate)
		}
	}

	return wordlist
}

func usingCoverTree(config distance.Config, useInputCoverTree bool) []*corpus.Word {
	words, err := file.NewCorpus()
	if err != nil {
		fmt.Printf("Error when importing file: %v", err)
	}

	if useInputCoverTree {
		candidates, err := selectCandidates(config, words)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		wordlist, err := selectWords(config, candidates)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		return wordlist
	}

	wordlist, err := selectWords(config, words)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return wordlist
}

func selectCandidates(config distance.Config, words []*corpus.Word) ([]*corpus.Word, error) {
	if len(words) == 0 {
		return words, nil
	}

	inputTree := covertree.NewInMemoryTree(basis, rootDistance, distance.DistanceWithOptions(config))

	_, err := index(inputTree, words...)
	if err != nil {
		return nil, fmt.Errorf("input cover tree: indexing words: %w", err)
	}

	query := words[0]
	results, err := inputTree.FindNearest(query, maxResults, maxDistance)
	if err != nil {
		return nil, fmt.Errorf("input cover tree: query '%s': %w", query, err)
	}

	return wordsFrom(query, results), nil
}

func wordsFrom(query *corpus.Word, results []covertree.ItemWithDistance) []*corpus.Word {
	n := len(results) + 1 // assumes the query is not in the results (exclusive distance function)
	words := make([]*corpus.Word, n)
	words[0] = query
	for i, item := range results {
		words[i+1] = item.Item.(*corpus.Word)
	}
	return words
}

func selectWords(config distance.Config, candidates []*corpus.Word) ([]*corpus.Word, error) {
	if len(candidates) == 0 {
		return candidates, nil
	}

	outputTree := covertree.NewInMemoryTree(basis, rootDistance, distance.DistanceWithOptions(config))

	var n int
	for _, candidate := range candidates {
		compatibleWords, err := outputTree.FindNearest(candidate, maxResults, maxDistance)
		if err != nil {
			return nil, fmt.Errorf("output cover tree: query '%s': %w", candidate, err)
		}

		if len(compatibleWords) == n {
			_, err := index(outputTree, candidate)
			if err != nil {
				return nil, fmt.Errorf("output cover tree: indexing word '%s': %w", *candidate, err)
			}
			n++
		}
	}

	query := candidates[0]
	results, err := outputTree.FindNearest(query, maxResults, maxDistance)
	if err != nil {
		return nil, fmt.Errorf("output cover tree: query '%s': %w", query, err)
	}

	return wordsFrom(query, results), nil
}
