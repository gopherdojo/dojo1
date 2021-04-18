package wordbank

import (
	"math/rand"
)

// A WordProvider provides words
type WordProvider interface {
	NextWord() string
}

type randomWordBank struct {
	words []string
	rng   *rand.Rand
}

// NewRandomWordBank creates a new random word provider
func NewRandomWordBank(words []string, seed int64) WordProvider {
	return &randomWordBank{
		words: words,
		rng:   rand.New(rand.NewSource(seed)),
	}
}

// NextWord implements WordProvider by returning a random word
func (rwp *randomWordBank) NextWord() string {
	numWords := len(rwp.words)
	return rwp.words[rwp.rng.Intn(numWords)]
}
