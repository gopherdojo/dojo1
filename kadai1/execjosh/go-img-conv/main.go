package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/execjosh/go-img-conv/imgconv"
)

type options struct {
	From string
	To   string
}

func main() {
	opts := &options{}
	flag.StringVar(&opts.From, "from", "jpeg", "image format from which to convert")
	flag.StringVar(&opts.To, "to", "png", "image format to which to convert")
	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 {
		die(errors.New("Must specify a directory"))
	}

	inputDir := args[0] // Only use first non-flag arg
	ic := imgconv.New(inputDir)

	if err := ic.Convert(opts.From, opts.To); err != nil {
		die(err)
	}
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
