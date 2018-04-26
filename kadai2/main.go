package main

import (
	"flag"

	"./imgconv"
)

var outType string
var srcDir string

func init() {
	flag.StringVar(&outType, "out", "png", "set output image type")
	flag.StringVar(&outType, "o", "png", "shorthand 'out flag")
}

func main() {
	flag.Parse()
	srcDir = flag.Arg(0)

	i := &imgconv.Imagefile{}
	imgconv.ImgConv(i, srcDir, outType)
}
