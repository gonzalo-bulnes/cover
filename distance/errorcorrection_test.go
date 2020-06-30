package distance

import (
	"testing"

	"github.com/gonzalo-bulnes/cover/corpus"
)

func TestDistanceForErrorCorrection(t *testing.T) {

	t.Run("does no explode when words are identical", func(t *testing.T) {
		w := corpus.NewWord("hello")

		if affinity := ForErrorCorrection(&w, &w); affinity < 1.0 {
			t.Errorf("Expected affinity to always be inferior to 1.0, got %f", affinity)
		}
	})

	t.Run("decreases as edit distance increases", func(t *testing.T) {
		w1 := corpus.NewWord("hello")
		w2 := corpus.NewWord("hallo")
		w3 := corpus.NewWord("bonjour")

		ed12 := Levenshtein(&w1, &w2)
		ed13 := Levenshtein(&w1, &w3)

		ec12 := ForErrorCorrection(&w1, &w2)
		ec13 := ForErrorCorrection(&w1, &w3)

		if ed13 > ed12 && ec13 >= ec12 {
			t.Error("Expected distance for error corection to decrease as edit distance increased.")
		}
	})

}
