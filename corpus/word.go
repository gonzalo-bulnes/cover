package corpus

// Word represents a word from a corpus.
type Word struct {
	string
}

func (w *Word) String() string {
	return w.string
}

// New returns a new word.
func NewWord(w string) *Word {
	return &Word{
		string: w,
	}
}
