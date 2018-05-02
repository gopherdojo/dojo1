package main

import (
	"flag"
	"os"
	"fmt"
	"path/filepath"

	"github.com/srttk/imgconv/encode"
)

var srcExt string
var distExt string
var srcDir  string

func init() {
	flag.StringVar(&srcExt, "src", "jpeg", "set source image ext")
	flag.StringVar(&srcExt, "s", "jpeg", "shorthand 'src flag'")
	flag.StringVar(&distExt, "out", "png", "set output image ext")
	flag.StringVar(&distExt, "o", "png", "shorthand 'out flag")
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Print("required 'filedir' params")
		os.Exit(2)
	}
	srcDir = flag.Arg(0)
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != "."+srcExt {
			return nil
		}
		srcFile, err := os.Open(path)
		defer srcFile.Close()
		if err != nil {
			return err
		}
		encoder, err := encode.NewEncoder(srcExt, srcFile)
		if err != nil {
			return err
		}
		distFile, err := os.Create(encode.GetDistPath(path, srcExt, distExt))
		if err != nil {
			return err
		}
		err = encoder.Encode(distFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
}
