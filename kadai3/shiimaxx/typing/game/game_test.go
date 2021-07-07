package game

import (
	"testing"
	"time"
)

func TestNewGame(t *testing.T) {
	var cases = []struct {
		name           string
		timeout        time.Duration
		numOfQuestions int
	}{
		{
			name:           "default",
			timeout:        time.Duration(60),
			numOfQuestions: 100,
		},
		{
			name:           "default",
			timeout:        time.Duration(0),
			numOfQuestions: 100,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			game := NewGame(c.timeout, c.numOfQuestions)
			if game.Timeout != c.timeout {
				t.Errorf("expected %d to eq %d", game.Timeout, c.timeout)
			}
			if len(game.Words) != c.numOfQuestions {
				t.Errorf("expected %d to eq %d", len(game.Words), c.numOfQuestions)
			}
		})
	}
}

func TestRun(t *testing.T) {
	var cases = []struct {
		name           string
		timeout        time.Duration
		numOfQuestions int
		actual         Result
	}{
		{
			name:           "default",
			timeout:        time.Duration(60),
			numOfQuestions: 100,
			actual:         Result{questionCount: 0, okCount: 0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			game := NewGame(c.timeout, c.numOfQuestions)
			questionCount, okCount := game.Run()
			if questionCount != c.actual.questionCount {
				t.Errorf("expected %d to eq %d", questionCount, c.actual.questionCount)
			}
			if okCount != c.actual.okCount {
				t.Errorf("expected %d to eq %d", okCount, c.actual.okCount)
			}
		})
	}
}
