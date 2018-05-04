package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
	"context"
)

var m sync.Mutex
var count int
var chars = []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij", "ijk", "jkl", "klm", "lmn"}

func init() {
	count = 0
}

func playGame(ctx context.Context) error {
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	errChan := make(chan error)
	question := make(chan string)
	answer := make(chan string)
	for {
		go func() {
			errChan <- sendQuestion(child, question)
		}()
		go func() {
			errChan <- sendAnswer(child, answer)
		}()
		questionChar := <-question
		fmt.Printf("%s\n>", questionChar)
		answerChar := <-answer
		if questionChar == answerChar {
			fmt.Print("correct!\n\n")
			m.Lock()
			count += 1
			m.Unlock()
		} else {
			fmt.Print("wrong...\n\n")
		}

		select {
		case err := <-errChan:
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func sendAnswer(ctx context.Context, answer chan string) error {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		answer<- s.Text()
		break
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func sendQuestion(ctx context.Context, question chan string) error {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(chars))
	question<- chars[index]

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func main() {
	parent, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	errChan := make(chan error)
	go func() {
		errChan <-playGame(parent)
	}()
	select {
	case err := <-errChan:
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case <-parent.Done():
		fmt.Print("\ntime up!\n")
		fmt.Printf("correct count: %d\n", count)
	}
}
