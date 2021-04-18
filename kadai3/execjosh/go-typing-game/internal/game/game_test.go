package game_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/gopherdojo/dojo1/kadai3/execjosh/go-typing-game/internal/game"
)

func TestRun(t *testing.T) {
	input := bytes.NewBufferString("abc\nbcd\nc\n\n")
	output := new(bytes.Buffer)
	wb := wordBank{}
	timeout := 1 * time.Second

	stats := game.Run(input, output, &wb, timeout)

	actualOutput := output.String()
	expectedOutput := "abc\n✅\nabc\n❌\nabc\n❌\nabc\n❌\nabc\n"
	if actualOutput != expectedOutput {
		t.Fail()
	}

	if stats.SuccessCount() != 1 {
		t.Fail()
	}

	if stats.FailureCount() != 3 {
		t.Fail()
	}
}

type wordBank struct {
}

func (wb *wordBank) NextWord() string {
	return "abc"
}
