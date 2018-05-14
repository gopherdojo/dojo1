package word_test

import (
	"testing"

	"github.com/xlune/dojo1/kadai3/xlune/001/word"
)

func TestIssue(t *testing.T) {
	word.ClearHistory()
	c := word.CountTotal()
	words := []string{}
	for i := 0; i < c; i++ {
		str, err := word.Issue()
		if err != nil {
			t.Fatal("return value invalid")
		}
		for _, w := range words {
			if w == str {
				t.Fatal("duplicate words")
			}
		}
		words = append(words, str)
	}
	_, err := word.Issue()
	if err == nil {
		t.Fatal("need error")
	}
}

func TestGetLatest(t *testing.T) {
	word.ClearHistory()
	if word.GetLatest() != "" {
		t.Fatal("need empty string")
	}
	str, err := word.Issue()
	if err != nil {
		t.Fatal("return value invalid")
	}
	if str != word.GetLatest() {
		t.Fatal("return value not match")
	}
}

func TestCountHistory(t *testing.T) {
	word.ClearHistory()
	if word.CountHistory() != 0 {
		t.Fatal("need count 0")
	}
	_, err := word.Issue()
	if err != nil {
		t.Fatal("return value invalid")
	}
	if word.CountHistory() != 1 {
		t.Fatal("need count 1")
	}
}

func TestCheckLatest(t *testing.T) {
	word.ClearHistory()
	if word.CheckLatest("") {
		t.Fatal("need return false")
	}
	str, err := word.Issue()
	if err != nil {
		t.Fatal("return value invalid")
	}
	if !word.CheckLatest(str) {
		t.Fatal("need return true")
	}
	if word.CheckLatest("hoge") {
		t.Fatal("need return false")
	}
}
