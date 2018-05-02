package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
		question := getWord()
		fmt.Print(">" + question)
		answer := <-ch
		if question == answer {
			fmt.Print("correct!")
		} else {
			fmt.Print("bad!")
		}

		fmt.Println(<-ch)
	}
}

func getWord() string {
	return "dog"
}
