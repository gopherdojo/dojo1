package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/itosho/gonverter/convert"
	_ "github.com/itosho/gonverter/convert/gif"
	_ "github.com/itosho/gonverter/convert/jpeg"
	_ "github.com/itosho/gonverter/convert/png"
)

const (
	ExitSuccess = iota
	ExitError
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

	if !convert.IsConvertibleImage(*fromExt) || !convert.IsConvertibleImage(*toExt) {
		log.Fatal("Invalid Extension. Please specify jpg, png or gif extension.")
	}

	code := convertRecursive(directory, *fromExt, *toExt)
	os.Exit(code)
}

func convertRecursive(directory string, fromExt string, toExt string) int {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == fromExt {
			if err := convert.CreateFile(path, fromExt, toExt); err != nil {
				return err
			}
			if err := convert.RemoveFile(path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Convert Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}

	return ExitSuccess
}

func usage() {
	fmt.Println("usage: gonverter [-f from extension] [-t to extension] [directory]")
	flag.PrintDefaults()
	os.Exit(ExitSuccess)
}
