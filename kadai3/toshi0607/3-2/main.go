package main

import (
	"fmt"
	"os"
	"github.com/toshi0607/dojo1/kadai3/toshi0607/3-2/rangedownloader"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%s\n", err)
			os.Exit(1)
		}
	}()
	cli := rangedownloader.New()
	os.Exit(cli.Run())
}
