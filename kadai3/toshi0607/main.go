package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"time"
)

func main() {
	ch := input(os.Stdin)

	questions := []string{"a", "b", "c", "d", "e", "f"}
	timeChan := time.NewTimer(10 * time.Second).C

	for _, v := range questions {
		fmt.Printf("type %s\n", v)
		Loop:
		for {
			fmt.Print(">")
			select {
			case <-timeChan:
				fmt.Println("Timer expired")
				return
			case s := <-ch:
				if s == v {
					fmt.Println("ええで")
					break Loop
				} else {
					fmt.Println("あかんで")
					fmt.Printf("type %s\n", v)
				}
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
