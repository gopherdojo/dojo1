package main

import (
	"os"

	"./typing"
)

func main() {
	typing.Start(os.Stdin, os.Stdout)
}
