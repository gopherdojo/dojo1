package main

import (
	"flag"
	"github.com/yuuyamad/gic/utils"
)

func main() {

		var (
			fromFormat = flag.String("s", "jpg", "source image format")
			toFormat = flag.String("d", "png", "dist image format")
		)

		flag.Parse()

		args := flag.Args()

		option := utils.Options{args[0], *fromFormat, *toFormat}
		option.ConvertImage()

}