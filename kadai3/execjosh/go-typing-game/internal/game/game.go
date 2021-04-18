package game

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/gopherdojo/dojo1/kadai3/execjosh/go-typing-game/internal/wordbank"
)

type ctx struct {
	done     chan struct{}
	sch, fch chan bool

	wb wordbank.WordProvider

	stdin  io.Reader
	stdout io.Writer
}

// Run runs the game
func Run(stdin io.Reader, stdout io.Writer, wb wordbank.WordProvider, timeout time.Duration) *Stats {
	ctx := &ctx{
		done: make(chan struct{}),
		sch:  make(chan bool),
		fch:  make(chan bool),

		wb: wb,

		stdin:  stdin,
		stdout: stdout,
	}

	// Keep stats out of context to force use of channels
	stats := &Stats{}

	go ctx.gameloop()

	go func() {
		time.Sleep(timeout)
		close(ctx.done)
	}()

OUT:
	for {
		select {
		case _, ok := <-ctx.fch:
			if ok {
				stats.LogFailure()
			}
		case _, ok := <-ctx.sch:
			if ok {
				stats.LogSuccess()
			}
		case <-ctx.done:
			close(ctx.sch)
			close(ctx.fch)
			break OUT
		}
	}

	return stats
}

func (c *ctx) gameloop() {
	in := bufio.NewScanner(c.stdin)
	word := c.wb.NextWord()

	for {
		fmt.Fprintln(c.stdout, word)

		if ok := in.Scan(); !ok {
			break
		}

		if word == in.Text() {
			c.sch <- true
			fmt.Fprintf(c.stdout, "\u2705\n")
			word = c.wb.NextWord()
		} else {
			fmt.Fprintf(c.stdout, "\u274c\n")
			c.fch <- true
		}
	}
}
