package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

func answer(r io.Reader) <-chan string {
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

func play(ctx context.Context) error {
	correctCnt := 0
	wrongCnt := 0

	questions, err := getQuestions()
	if err != nil {
		return err
	}

	ch := answer(os.Stdin)
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
			return nil
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
