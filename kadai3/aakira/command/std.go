package command

import (
	"io"
	"bufio"
)

func Input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		defer close(ch)
	}()
	return ch
}
