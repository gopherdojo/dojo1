package main

import (
	"flag"
	"runtime"

	"path"

	"github.com/mpppk/dpget/download"
)

var procs int
var outputFileName string
var targetDir string

func main() {
	flag.Parse()
	urlPath := flag.Arg(0)
	if outputFileName == "" {
		outputFileName = path.Base(urlPath)
	}

	outputFilePath := path.Join(targetDir, outputFileName)

	if err := download.DoParallel(urlPath, outputFilePath, procs); err != nil {
		panic(err)
	}
}

func init() {
	flag.IntVar(&procs, "procs", runtime.NumCPU(), "split ratio to download file")
	flag.StringVar(&outputFileName, "output", "", "output file to <filename>")
	flag.StringVar(&targetDir, "target-dir", ".", "path to the directory to save the downloaded file")
}
