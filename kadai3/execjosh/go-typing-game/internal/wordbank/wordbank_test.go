package wordbank_test

import (
	"testing"

	"github.com/gopherdojo/dojo1/kadai3/execjosh/go-typing-game/internal/wordbank"
)

func TestRandomWordProvider(t *testing.T) {
	seed := int64(1234567890)
	words := []string{
		"a", "b", "c", "d", "abc",
	}
	wb := wordbank.NewRandomWordBank(words, seed)

	expected := []string{
		"d", "b", "d", "b", "d",
	}

	for _, exp := range expected {
		if exp != wb.NextWord() {
			t.Fail()
		}
	}
}
