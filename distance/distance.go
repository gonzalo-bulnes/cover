package distance

import (
	"fmt"
	"os"
	"strings"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/mandykoh/go-covertree"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var print = os.Getenv("PRINT") != ""

const Max = 100.0

// Distance returns a distance between two words, suitable to compose wordlists
// for memorable passphrases.
//
// Returns 100.0 when a word is an exact prefix of the other.
// Returns 100.0 when the edit distance is less than 3.
// Returns the edit distance otherwise.
var Distance covertree.DistanceFunc = func(a, b interface{}) float64 {
	w1 := a.(*corpus.Word)
	w2 := b.(*corpus.Word)

	s1 := w1.String()
	s2 := w2.String()

	distance := levenshtein.DistanceForStrings([]rune(s1), []rune(s2), levenshtein.DefaultOptionsWithSub)
	if distance < 3 {
		if print {
			fmt.Printf("Rejecting – insufficient edit distance '%s' '%s'\n", s1, s2)
		}
		distance = Max
	}
	if strings.HasPrefix(s1, s2) || strings.HasPrefix(s2, s1) {
		if print {
			fmt.Printf("Rejecting – exact prefix '%s' '%s'\n", s1, s2)
		}
		distance = Max
	}
	if strings.HasPrefix(s1, s2[:3]) {
		if print {
			fmt.Printf("Rejecting – same first three letters '%s' '%s'\n", s1, s2)
		}
		distance = Max
	}
	if print {
		fmt.Printf("Distance between '%s' and '%s' computed as %f\n", s1, s2, float64(distance))
	}
	return float64(distance)
}
