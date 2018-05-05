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

func Start(ctx context.Context) {
	questions, err := getQuestions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Get Questions Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	questionCnt := 0
	correctCnt := 0
	wrongCnt := 0

	ch := input(os.Stdin)
	for {
		question := getRandQuestion(questions)
		fmt.Println("-")
		fmt.Println(question)

		select {
		case <-ctx.Done():
			fmt.Println("\n-")
			fmt.Println("time over.")
			fmt.Println(fmt.Sprintf("result: %d/%d", correctCnt, questionCnt))
			return
		case answer := <-ch:
			questionCnt++
			if question == answer {
				correctCnt++
				fmt.Println("correct!")
			} else {
				wrongCnt++
				fmt.Println("wrong!")
			}
		}
	}
}
