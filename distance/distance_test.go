package distance

import (
	"math"
	"testing"

	"github.com/gonzalo-bulnes/cover/corpus"
)

func TestDistance(t *testing.T) {

	testcases := []struct {
		message  string
		a        corpus.Word
		b        corpus.Word
		expected float64
	}{
		{
			message:  "a is an exact prefix of b",
			a:        corpus.NewWord("agua"),
			b:        corpus.NewWord("aguacero"),
			expected: 100,
		},
		{
			message:  "a and b are 3 edits apart",
			a:        corpus.NewWord("agua"),
			b:        corpus.NewWord("tregua"),
			expected: 3,
		},
		{
			message:  "a and b are more than 3 edits apart",
			a:        corpus.NewWord("agua"),
			b:        corpus.NewWord("abrelatas"),
			expected: 7,
		},
		{
			message:  "a and b are less than 3 edits apart",
			a:        corpus.NewWord("cono"),
			b:        corpus.NewWord("coro"),
			expected: 100,
		},
		{
			message:  "a and b start by the same three letters",
			a:        corpus.NewWord("deslinde"),
			b:        corpus.NewWord("desayuno"),
			expected: 100,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.message, func(t *testing.T) {
			if distance := Distance(&tc.a, &tc.b); math.Abs(distance-tc.expected) > epsilon {
				t.Errorf("%s: expected %f, got %f", tc.message, tc.expected, distance)
			}
		})
	}
}
