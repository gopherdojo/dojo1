package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/gopherdojo/dojo1/kadai3/srttk/multidown/option"
	"github.com/gopherdojo/dojo1/kadai3/srttk/multidown/downloader"
	"fmt"
	"github.com/gopherdojo/dojo1/kadai3/srttk/multidown/checker"
)

var options option.CliOption

func main() {
	args, err := flags.Parse(&options)
	if err != nil {
		fmt.Printf("\ninvalid args. %v", err)
		os.Exit(1)
	}
	url := args[0]
	c := checker.NewChecker(url)
	err = c.CheckResourceSupportRangeAccess()
	if err != nil {
		fmt.Printf("\nnot support range access. %v", err)
		os.Exit(1)
	}
	d, err := downloader.NewDownloader(options, c.Size, url)
	if err != nil {
		fmt.Printf("\nfailed to initialize workers %v", err)
	}
	err = d.Download()
	if err != nil {
		fmt.Printf("\ndownload err. %v", err)
		os.Exit(1)
	}
	err = d.Merge()
	if err != nil {
		fmt.Printf("\nmerge error. %v", err)
		os.Exit(1)
	}
}