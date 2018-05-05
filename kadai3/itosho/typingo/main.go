package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"./typing"
)

const (
	ExitSuccess = iota
	ExitError
)

func main() {
	var seconds = flag.Int("s", 30, "seconds")

	flag.Usage = usage
	flag.Parse()

	timeLimit := time.Duration(*seconds)
	if timeLimit > 300 {
		log.Fatal("please set within 300 seconds")
	}

	fmt.Println("=====TYPINGO START=====")
	result := typing.Run(timeLimit)
	fmt.Println("=====TYPINGO END=====")
	if !result {
		os.Exit(ExitError)
	}

	os.Exit(ExitSuccess)
}

func usage() {
	fmt.Println("usage: typingo [-s seconds]")
	flag.PrintDefaults()
	os.Exit(ExitSuccess)
}
