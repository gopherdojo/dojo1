package rangedownloader

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
)

type cliOptions struct {
	Name  string `short:"n" long:"name" description:"output file name with extension. if not provided, rangedownloader will guess a file name based on URL"`
	Procs int    `short:"p" long:"procs" description:"number of parallel" default:"1"`
	Args struct {
		URL string
	} `positional-args:"yes"`
}

func (d *Downloader) parseCommandLine() error {
	opts := &cliOptions{}
	p := flags.NewParser(opts, flags.HelpFlag)
	_, err := p.ParseArgs(d.Argv)
	if err != nil {
		return errors.Wrap(err, "failed to parse command line")
	}

	if 	opts.Args.URL == "" {
		p.WriteHelp(os.Stdout)
		return fmt.Errorf("\n please check usage above")
	}
	d.url = opts.Args.URL

	if opts.Name != "" {
		d.name = opts.Name
	} else {
		if name := guessFileName(d.url); name == "" {
			return errors.Wrap(err, "please provide output file name")
		} else {
			d.name = name
		}
	}

	d.procs = opts.Procs

	return nil
}

func guessFileName(URL string) string {
	s := strings.Split(URL, "/")
	return s[len(s)-1]
}
