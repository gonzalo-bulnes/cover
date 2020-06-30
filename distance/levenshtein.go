package distance

import (
	"fmt"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/mandykoh/go-covertree"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// Levenshtein returns the edit distance between two words.
var Levenshtein covertree.DistanceFunc = func(a, b interface{}) float64 {
	w1 := a.(*corpus.Word)
	w2 := b.(*corpus.Word)

	s1 := w1.String()
	s2 := w2.String()

	distance := levenshtein.DistanceForStrings([]rune(s1), []rune(s2), levenshtein.DefaultOptionsWithSub)
	fmt.Printf("Distance between '%s' and '%s' computed as %f\n", s1, s2, float64(distance))

	return float64(distance)
}
