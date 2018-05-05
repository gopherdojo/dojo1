package typing

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	ExitSuccess = iota
	ExitError
)

func Run() {
	var seconds = flag.Int("s", 30, "seconds")
	flag.Usage = usage
	flag.Parse()

	timeLimit := time.Duration(*seconds)

	if timeLimit > 300 {
		log.Fatal("Please set within 300 seconds.")
	}

	if !ready(timeLimit) {
		log.Fatal("Not ready.")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeLimit*time.Second)
	defer cancel()

	if !game(ctx) {
		log.Fatal("Error.")
	}

	fmt.Println("=====TYPINGO END=====")
}

func usage() {
	fmt.Println("usage: typingo [-s seconds]")
	flag.PrintDefaults()
	os.Exit(ExitSuccess)
}

func ready(seconds time.Duration) bool {
	fmt.Println("=====TYPINGO START=====")
	fmt.Println("Welcome dead simple typing game to learn software proverbs!")
	fmt.Println(fmt.Sprintf("> Time limit is %d seconds", seconds))
	fmt.Println("Are you ready? [Y/n]")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadByte()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Read Byte Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	if s == []byte("Y")[0] || s == []byte("y")[0] {
		fmt.Println("Ready Go!")
		return true
	} else if s == []byte("N")[0] || s == []byte("n")[0] {
		fmt.Println("Good Bye!")
		return false
	} else {
		log.Fatal("Please enter yes or no.")
	}

	return false
}
