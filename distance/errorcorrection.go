package distance

import "github.com/mandykoh/go-covertree"

// ForErrorCorrection returns a distance that decreases as error correction becomes easier.
var ForErrorCorrection covertree.DistanceFunc = func(a, b interface{}) float64 {
	return 1.0 / (1 + Levenshtein(a, b))
}

// MaxEditDistanceForErrorCorrection returns a distance that ensures words
// have a larger edit distance than minEditDistsance.
func MaxEditDistanceForErrorCorrection(minEditDistance int) float64 {
	return 1.0 / (1 + float64(minEditDistance))
}
