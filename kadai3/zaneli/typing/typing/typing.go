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

// NewTyping create typing struct.
func NewTyping(words []string) Typing {
	return Typing{words}
}

// Run start typing game.
func (t Typing) Run(seconds time.Duration) {
	var corrects, incorrects int
	timeout := time.After(time.Second * seconds)
	ch := t.answer()

	for {
		select {
		case result, ok := <-ch:
			if !ok {
				fmt.Printf("中断しました。 正解=%d, 不正解=%d.\n", corrects, incorrects)
				return
			}
			if result {
				corrects++
			} else {
				incorrects++
			}
		case <-timeout:
			fmt.Printf("終了しました。 正解=%d, 不正解=%d.\n", corrects, incorrects)
			return
		}
	}
}

func (t Typing) answer() <-chan bool {
	in := input(os.Stdin)
	out := make(chan bool)

	go func() {
		for _, w := range t.words {
			fmt.Printf("\r%s[%s]%s%s ", green, w, green, reset)
			for v := range in {
				if v == ":q" {
					close(out)
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
