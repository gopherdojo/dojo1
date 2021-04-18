package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gopherdojo/dojo1/kadai1/arisawa/imgconv"
)

var (
	in      = flag.String("in", "", "Input directory")
	out     = flag.String("out", "", "Output directory")
	from    = flag.String("from", "", "Image format before conversion")
	to      = flag.String("to", "", "Image format after conversion")
	verbose = flag.Bool("verbose", false, "Verbose output")
)

func init() {
	formats := make([]string, 0, len(imgconv.SupportedFormats))
	for k, _ := range imgconv.SupportedFormats {
		formats = append(formats, k)
	}

	flag.Usage = func() {
		fmt.Printf(`Usage:
  %s -in INPUT_DIR -out OUTPUT_DIR -from FROM_FORMAT -to TO_FORMAT

  Convert image files under speicfied directory recursively.
  Supported formats: %s

`, os.Args[0], strings.Join(formats, ", "))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *in == "" || *out == "" || *from == "" || *to == "" {
		log.Fatalf("%s: invalid argument", os.Args[0])
	}

	c, err := imgconv.NewImgconv(*in, *out, *from, *to, *verbose)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: conversion finished\n", os.Args[0])
}
