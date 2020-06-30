package corpus

// New returns a new corpus of words.
func New(words ...string) (corpus []Word) {
	corpus = make([]Word, len(words))

	for i, w := range words {
		corpus[i] = NewWord(w)
	}
	return
}
