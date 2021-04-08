package option

import (
	"bytes"
	"fmt"
)

type CliOption struct {
	Help        bool   `short:"h" long:"help"`
	OutputDir   string `short:"o" long:"output"`
}


func (co CliOption) usage() []byte {
	buf := bytes.Buffer{}

	fmt.Fprintf(&buf, `Usage: multidown [options] URL
  Options:
  -h,  --help                   print usage and exit
  -o,  --outputDir <filename>   output dir
`)
	return buf.Bytes()
}