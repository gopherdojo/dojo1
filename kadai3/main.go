package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/ohsawa0515/gotyping/typing"
)

var timeout int

func main() {
	flag.IntVar(&timeout, "t", 30, "Timeout(sec)")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	typing.Run(ctx, os.Stdin)
}
