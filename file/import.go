package file

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/gonzalo-bulnes/cover/corpus"
)

func NewCorpus() (words []corpus.Word, err error) {
	file, err := os.Open(filepath.Join("file", "testdata", "corpus.txt"))
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// words = []corpus.Word{}
	for scanner.Scan() {
		words = append(words, corpus.NewWord(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return
}
