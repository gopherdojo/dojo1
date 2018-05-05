package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"./typing"
)

func main() {
	var seconds time.Duration
	seconds = 10
	fmt.Println("Start typingo game!")
	fmt.Println(fmt.Sprintf("Time limit is %d seconds", seconds))
	fmt.Println("Are you ready? [Y/n]")

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	if s == []byte("Y")[0] || s == []byte("y")[0] {
		fmt.Println("Ready Go!")
	} else if s == []byte("N")[0] || s == []byte("n")[0] {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, seconds*time.Second)
	defer cancel()

	typing.Start(ctx)
}
