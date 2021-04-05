package main

import (
	"os"
	"fmt"
	"github.com/toshi0607/dojo1/kadai3/toshi0607/game"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%s\n", err)
			os.Exit(1)
		}
	}()
	cli := game.Game{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
