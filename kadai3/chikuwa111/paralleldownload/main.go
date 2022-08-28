package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"paralleldownload/pdownload"
)

var parallelCount int

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %v [option] [url]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "option:")
	flag.PrintDefaults()
}

func init() {
	flag.IntVar(&parallelCount, "p", 2, "The number of parallel download")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	runtime.GOMAXPROCS(parallelCount)
	url := args[0]
	err := pdownload.Run(url, parallelCount)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
