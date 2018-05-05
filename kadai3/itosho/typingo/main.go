package main

import (
	"bufio"
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

func main() {
	fmt.Println("Start typingo game!")
	fmt.Println("Time limit is one minute.")
	fmt.Println("Are you ready? [Y/n]")

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	if s == []byte("Y")[0] || s == []byte("y")[0] {
		fmt.Println("Ready Go!")
	} else if s == []byte("N")[0] || s == []byte("n")[0] {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

	questionCnt := 0
	correctCnt := 0
	wrongCnt := 0

	t := time.After(10 * time.Second)
	ch := input(os.Stdin)
	for {
		select {
		case <-t:
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
