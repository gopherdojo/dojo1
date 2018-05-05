package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
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
	questionCnt := 0
	correctCnt := 0
	wrongCnt := 0

	ch := input(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time over.")
			fmt.Println("-")
			fmt.Println(fmt.Sprintf("result: %d/%d", correctCnt, questionCnt))
			return
		default:
			questionCnt++

			question := getRandWord()
			fmt.Println("-")
			fmt.Println(question)
			answer := <-ch

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

func getRandWord() string {
	words := [...]string{
		"dog",
		"cat",
		"elephant",
		"lion",
		"bird",
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(words))

	return words[i]
}