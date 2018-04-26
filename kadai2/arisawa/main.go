package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arisawa/go-imgconv/imgconv"
)

var (
	in   = flag.String("in", "", "Input directory")
	out  = flag.String("out", "", "Output directory")
	from = flag.String("from", "", "Image format before conversion")
	to   = flag.String("to", "", "Image format after conversion")
)

func main() {
	flag.Usage = func() {
		fmt.Printf(`Usage:
  %s -in INPUT_DIR -out OUTPUT_DIR -from FROM_FORMAT -to TO_FORMAT

  Convert image files under speicfied directory recursively.
  Supported src formats: [%s]
  Supported dest formats: [%s]

`, os.Args[0], strings.Join(imgconv.SourceFormats, ", "), strings.Join(imgconv.DestFormats, ", "))
		flag.PrintDefaults()
	}
	flag.Parse()

	if *in == "" || *out == "" || *from == "" || *to == "" {
		log.Fatalf("%s: invalid argument", os.Args[0])
	}

	rc, err := imgconv.NewRecursiveConverter(*in, *out, *from, *to)
	if err != nil {
		log.Fatal(err)
	}
	if err := rc.Convert(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: conversion finished\n", os.Args[0])
}
