package main

import (
	"flag"
	"log"

	"./converter"
)

func main() {
	from := flag.String("from", "jpg", "conversion source format")
	to := flag.String("to", "png", "conversion destination format")
	flag.Parse()

	c, err := converter.NewConverter(*from, *to)
	if err != nil {
		log.Fatal(err)
	}

	if len(flag.Args()) < 1 {
		log.Fatal("ディレクトリを指定してください。")
	}

	err = c.Convert(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
}
