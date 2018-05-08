package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
)

// Game game settings
type Game struct {
	Timeout time.Duration
	Words   []string
}

// Result game result
type Result struct {
	questionCount int
	okCount       int
}

// NewGame constractor for Game
func NewGame(timeout time.Duration, numOfQuestions int) *Game {
	g := new(Game)
	g.Timeout = timeout
	for i := 0; i < numOfQuestions; i++ {
		g.Words = append(g.Words, randomdata.Adjective())
	}
	return g
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func question(ctx context.Context, words []string, resultCh chan Result) {
	ch := input(os.Stdin)

	questionCount := 0
	okCount := 0
QUESTION_LOOP:
	for _, word := range words {
		fmt.Printf("question %d: %s\n", questionCount+1, word)
		fmt.Print("> ")

		select {
		case v, ok := <-ch:
			if ok {
				if word == v {
					okCount++
				} else {
					fmt.Println("miss")
				}
				questionCount++
			} else {
				break QUESTION_LOOP
			}
		case <-ctx.Done():
			fmt.Print("\nTimeup\n\n")
			break QUESTION_LOOP
		}
	}
	result := Result{
		questionCount: questionCount,
		okCount:       okCount,
	}
	resultCh <- result
}

// Run run game
func (g *Game) Run() (int, int) {
	ctx, cancel := context.WithTimeout(context.Background(), g.Timeout*time.Second)
	defer cancel()

	resultCh := make(chan Result)
	go func() {
		question(ctx, g.Words, resultCh)
	}()
	result := <-resultCh

	return result.questionCount, result.okCount
}
