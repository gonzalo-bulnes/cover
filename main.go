package main

import (
	"fmt"
	"math"

	"github.com/mandykoh/go-covertree"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

const (
	Basis        = 10.0
	RootDistance = 3.0
)

var corpus = []Word{
	{string: "hello"},
	{string: "hullo"},
	{string: "allo?"},
	{string: "muelle"},
	{string: "anaconda"},
}

type Point struct {
	X float64
	Y float64
}

type Word struct {
	string
}

func (w *Word) String() string {
	return w.string
}

func New(w string) *Word {
	return &Word{
		string: w,
	}
}

func euclidianDistance(a, b interface{}) float64 {
	p1 := a.(*Point)
	p2 := b.(*Point)

	distX := p1.X - p2.X
	distY := p1.Y - p2.Y
	fmt.Printf("Distance between '%v' and '%v'\n", p1, p2)

	return math.Sqrt(distX*distX + distY*distY)
}

func levenshteinDistance(a, b interface{}) float64 {
	w1 := a.(*Word)
	w2 := b.(*Word)

	s1 := w1.String()
	s2 := w2.String()

	distance := levenshtein.DistanceForStrings([]rune(s1), []rune(s2), levenshtein.DefaultOptionsWithSub)
	fmt.Printf("Distance between '%s' and '%s' computed as %f\n", s1, s2, float64(distance))

	return float64(distance)
}

func main() {
	tree := covertree.NewInMemoryTree(Basis, RootDistance, levenshteinDistance)

	fmt.Printf("\nIndexing phase.\n\n")
	for _, w := range corpus {
		x := w // copy is required
		err := tree.Insert(&x)
		if err != nil {
			fmt.Printf("Error inserting '%s': %v\n", w.String(), err)
		}
		fmt.Printf("Inserted '%+v'\n", &w)
	}

	fmt.Printf("\nQuerying phase.\n\n")
	w := Word{string: "hello"}
	n := 30
	maxDistance := 6.0
	fmt.Printf("Finding the %d nearest words that are closer than %f from '%s'\n", n, maxDistance, w)
	results, err := tree.FindNearest(&w, n, maxDistance)
	if err != nil {
		fmt.Printf("Error finding nearest to '%s': %v\n", w, err)
	}

	fmt.Printf("Results %+v\n", results)
}
