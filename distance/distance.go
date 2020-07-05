package distance

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gonzalo-bulnes/cover/corpus"
	"github.com/mandykoh/go-covertree"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

type Config struct {
	MinDistance             int
	NoPrefix                bool
	NoSameFirstThreeLetters bool
}

func NewConfig() Config {
	var minDistance int
	if m := os.Getenv("MIN_DISTANCE"); m != "" {
		f, err := strconv.Atoi(m)
		if err != nil {
			fmt.Printf("Invalid MIN_DISTANCE %s: %v", m, err)
		} else {
			minDistance = f
		}
	}
	noPrefix := os.Getenv("NO_PREFIX") != ""
	noSameFirstThreeLetters := os.Getenv("NO_SAME_FIRST_THREE_LETTERS") != ""

	return Config{
		MinDistance:             minDistance,
		NoPrefix:                noPrefix,
		NoSameFirstThreeLetters: noSameFirstThreeLetters,
	}
}

var print = os.Getenv("PRINT") != ""

const Max = 100.0

// DistanceWithOptions returns a distance between two words, suitable to compose wordlists
// for the puprose of creating memorable passphrases.
//
// Returns 100.0 when a word is an exact prefix of the other (config.NoPrefix).
// Returns 100.0 when a word starts with the same three letters as any other word (config.NoSameFirstThreeLetters).
// Returns 100.0 when the edit distance is less than config.MinDistance.
// Returns the edit distance otherwise.
func DistanceWithOptions(config Config) covertree.DistanceFunc {
	return func(a, b interface{}) float64 {
		w1 := a.(*corpus.Word)
		w2 := b.(*corpus.Word)

		s1 := w1.String()
		s2 := w2.String()

		distance := levenshtein.DistanceForStrings([]rune(s1), []rune(s2), levenshtein.DefaultOptionsWithSub)
		if distance < config.MinDistance {
			if print {
				fmt.Printf("Rejecting – insufficient edit distance '%s' '%s'\n", s1, s2)
			}
			distance = Max
		}
		if config.NoPrefix && (strings.HasPrefix(s1, s2) || strings.HasPrefix(s2, s1)) {
			if print {
				fmt.Printf("Rejecting – exact prefix '%s' '%s'\n", s1, s2)
			}
			distance = Max
		}
		if config.NoSameFirstThreeLetters && strings.HasPrefix(s1, s2[:3]) {
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
}
