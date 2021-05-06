package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/itosho/gonverter/convert"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
)

func main() {
	var fromExt = flag.String("f", ".jpg", "from extension")
	var toExt = flag.String("t", ".png", "to extension")
	flag.Usage = usage

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify only one direcoty.")
	}

	args := flag.Args()
	directory := args[0]

	fromType := convert.GetImageType(*fromExt)
	if fromType == nil {
		log.Fatalln("Invalid from extenstion type. Please specify `.jpg(jpeg)`, `.png` or `.gif` for option.")
	}

	toType := convert.GetImageType(*toExt)
	if toType == nil {
		log.Fatalln("Invalid from extenstion type. Please specify `.jpg(jpeg)`, `.png` or `.gif` for option.")
	}

	code := convert.Convert(directory, fromType, toType)
	os.Exit(code)
}

func usage() {
	fmt.Println("usage: gonverter [-f from extension] [-t to extension] [directory]")
	flag.PrintDefaults()
	os.Exit(ExitSuccess)
}
