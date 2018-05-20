package typing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// Typing has word list to use typing game.
type Typing struct {
	words []string
}

// Result has game result status and correct, incorrect counts.
type Result struct {
	Status     int
	Corrects   int
	Incorrects int
}

const (
	// ResultClear express game clear.
	ResultClear = iota
	// ResultCancel express game cancel.
	ResultCancel
	// ResultTimeOver express time over.
	ResultTimeOver
)

// NewTyping create typing struct.
func NewTyping(words []string) Typing {
	return Typing{words}
}

// Run start typing game.
func (t Typing) Run(seconds time.Duration) Result {
	var corrects, incorrects int
	cancel := make(chan struct{})
	timeout := time.After(time.Second * seconds)
	ch := t.answer(cancel)

	for {
		select {
		case result, ok := <-ch:
			if !ok {
				return Result{ResultClear, corrects, incorrects}
			}
			if result {
				corrects++
			} else {
				incorrects++
			}
		case <-cancel:
			return Result{ResultCancel, corrects, incorrects}
		case <-timeout:
			return Result{ResultTimeOver, corrects, incorrects}
		}
	}
}

func (t Typing) answer(cancel chan<- struct{}) <-chan bool {
	in := input(os.Stdin)
	out := make(chan bool)

	go func() {
		for _, w := range t.words {
			fmt.Printf("\r%s[%s]%s%s ", green, w, green, reset)
			for v := range in {
				if v == ":q" {
					cancel <- struct{}{}
					return
				} else if v == w {
					out <- true
					break
				} else {
					out <- false
					fmt.Printf("%s[%s]%s%s ", red, w, red, reset)
				}
			}
		}
		close(out)
	}()
	return out
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
