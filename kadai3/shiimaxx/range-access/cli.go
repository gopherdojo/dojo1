package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/shiimaxx/gurl/gurl"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	var (
		parallel int
		output   string
		version  bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.outStream)

	flags.StringVar(&output, "output", "", "output file")
	flags.StringVar(&output, "o", "", "output file(Short)")
	flags.IntVar(&parallel, "parallel", 10, "number of parallel")
	flags.IntVar(&parallel, "p", 10, "number of parallel(Short)")

	flags.BoolVar(&version, "version", false, "print version information")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(c.outStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if len(flags.Args()) < 1 {
		fmt.Fprintln(c.errStream, "missing arguments")
		return ExitCodeError
	}

	if output == "" {
		fmt.Fprint(c.errStream, "require output option\n")
		return ExitCodeError
	}

	if _, err := os.Stat(output); os.IsExist(err) {
		fmt.Fprintf(c.errStream, "%s: already exits\n", output)
		return ExitCodeError
	}

	url := flags.Args()[0]

	client := gurl.NewClient(parallel, output)
	if err := client.Get(url); err != nil {
		fmt.Fprintf(c.errStream, "%s\n", err)
		return ExitCodeError
	}

	return ExitCodeOK
}
