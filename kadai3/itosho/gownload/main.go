package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify url.")
	}
	args := flag.Args()
	url := args[0]

	print(url)
}

func usage() {
	fmt.Println("usage: gownload [url]")
	flag.PrintDefaults()
	os.Exit(0)
}
