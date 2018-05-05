package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

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

func game(ctx context.Context) bool {
	correctCnt := 0
	wrongCnt := 0

	questions, err := getQuestions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Get Questions Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	ch := input(os.Stdin)
	for {
		question := getRandQuestion(questions)
		fmt.Println("-----")
		fmt.Println(question)
		fmt.Print("> ")

		select {
		case <-ctx.Done():
			fmt.Println("\n-----")
			fmt.Println("Time Over!")
			fmt.Println(fmt.Sprintf("Score: %d/%d", correctCnt, correctCnt+wrongCnt))
			return true
		case answer := <-ch:
			if question == answer {
				correctCnt++
				fmt.Println("Correct!")
			} else {
				wrongCnt++
				fmt.Println("Wrong!")
			}
		}
	}
}
