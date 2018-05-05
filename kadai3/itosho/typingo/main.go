package main

import (
	"bufio"
	"fmt"
	"os"

	"./typing"
)

func main() {
	fmt.Println("Start typingo game!")
	fmt.Println("Time limit is one minute.")
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

	typing.Start()
}
