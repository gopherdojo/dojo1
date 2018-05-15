package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"wordtype/wordlist"
)

var (
	statusCode   = 0
	quizCount    = 0
	correctCount = 0
	timeLimit    int
)

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err)
	statusCode = 1
}

func setQuiz(word string) {
	quizCount++
	fmt.Println("  " + word)
	fmt.Print("> ")
}

func checkAnswer(input, answer string) {
	if input == answer {
		correctCount++
	}
}

func startGame() {
	bc := context.Background()
	t := time.Duration(timeLimit) * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan bool)
	for {
		for word := range wordlist.Words {
			setQuiz(word)
			go func() {
				ch <- scanner.Scan()
			}()
			select {
			case ok := <-ch:
				if ok {
					checkAnswer(scanner.Text(), word)
				} else {
					if err := scanner.Err(); err != nil {
						handleError(err)
					}
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}
}

func printResult() {
	fmt.Println("")
	fmt.Printf("(正答数/問題数)は、(%v/%v)でした", correctCount, quizCount)
}

func init() {
	flag.IntVar(&timeLimit, "t", 10, "Time limit (sec)")
	flag.Parse()
}

func main() {
	startGame()
	printResult()
	os.Exit(statusCode)
}
