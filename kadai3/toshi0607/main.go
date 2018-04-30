package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
)

func main() {
	ch := input(os.Stdin)

	questions := []string{"a", "b", "c", "d", "e", "f"}

	for _, v := range questions {
		fmt.Printf("type %s\n", v)
		for {
			fmt.Print(">")
			if <-ch == v {
				fmt.Println("ええで")
				break
			} else {
				fmt.Println("あかんで")
				fmt.Printf("type %s\n", v)
			}
		}
	}
	fmt.Println("おしまい")
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
