package main

import (
	"testing"

	"github.com/gonzalo-bulnes/cover/distance"
)

var config = distance.Config{
	MinDistance:             3,
	NoPrefix:                true,
	NoSameFirstThreeLetters: true,
}

func BenchmarkReferenceImplementation(b *testing.B) {
	benchmarkReferenceImplementation(nil, b)
}

func benchmarkReferenceImplementation(_, b *testing.B) {
	for n := 0; n < b.N; n++ {
		reference(config)
	}
}

func BenchmarkTwoCoverTrees(b *testing.B) {
	benchmarkTwoCoverTrees(nil, b)
}

func benchmarkTwoCoverTrees(_, b *testing.B) {
	for n := 0; n < b.N; n++ {
		usingCoverTree(config, true)
	}
}

func BenchmarkSingleCoverTree(b *testing.B) {
	benchmarkSingleCoverTree(nil, b)
}

func benchmarkSingleCoverTree(_, b *testing.B) {
	for n := 0; n < b.N; n++ {
		usingCoverTree(config, false)
	}
}
