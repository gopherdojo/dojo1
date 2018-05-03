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
	fmt.Println("> start typing game.")

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
	ch := input(os.Stdin)
	for {
		question := getRandWord()
		fmt.Print(">" + question)
		answer := <-ch
		if question == answer {
			fmt.Print("correct!")
		} else {
			fmt.Print("wrong!")
		}

		fmt.Println(<-ch)
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
