package main

import (
	"flag"
	"github.com/yuuyamad/gic/utils"
)

func main() {

		var (
			fromFormat = flag.String("s", "jpg", "変換前のフォーマット")
			toFormat = flag.String("d", "png", "変換後のフォーマット")
		)

		flag.Parse()

		args := flag.Args()

		option := utils.Options{args[0], *fromFormat, *toFormat}
		option.ConvertImage()

}