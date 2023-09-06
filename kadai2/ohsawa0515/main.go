package main

import (
	"os"

	"github.com/ohsawa0515/goimgconverter/img"
)

func main() {
	cli := &img.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
