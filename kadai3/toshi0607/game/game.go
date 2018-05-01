package game

import (
	"io"
	"os"
	"time"
	"fmt"
	"bufio"
)

type Game struct {
	OutStream, ErrStream io.Writer
}

func (g *Game) Run() int {
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
				return 0
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
	return 0
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
