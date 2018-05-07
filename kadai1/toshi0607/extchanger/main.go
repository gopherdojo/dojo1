package main

import (
	"os"
	"flag"
	"io"
	"github.com/toshi0607/gopher-dojo/extchanger/converter"
	"fmt"
)

type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run())
}

func (c *CLI) Run() int {
	flag.Usage = usage
	flag.Parse()
	if len(os.Args[1:]) != 3 {
		usage()
		return 1
	}

	if err := os.MkdirAll("output", 0777); err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}
	from := flag.Arg(0)
	to := flag.Arg(1)
	srcdir := flag.Arg(2)

	count, err := converter.ConvertExt(srcdir, from, to)
	if err != nil {
		fmt.Fprint(c.errStream, err)
		return 1
	}
	if count == 0 {
		fmt.Println("Files with extension you specified not found")
	} else {
		fmt.Printf("%d files converted! see under ./output\n", count)
	}

	return 0
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  main extension(from) extension(to) targetdirectory")
	fmt.Println("")
	fmt.Println("All of the args are requred.")
	flag.PrintDefaults()
}