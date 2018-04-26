package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli"
	"bufio"
)

// this file uses library which is urfave/cli(https://github.com/urfave/cli)
func main() {
	app := cli.NewApp()
	app.Name = "Myhead"
	app.Usage = "Display lines of a file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "number, n",
			Value: "10",
			Usage: "specify counts of output line",
		},
	}
	app.Action = func(c *cli.Context) error {
		myhead(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func myhead(c *cli.Context) {
	count := c.Int("n")
	if !checkInt(count) {
		return
	}

	for _, path := range c.Args() {
		readFile(path, count)
	}
}

func checkInt(arg int) bool {
	if arg > 0 {
		return true
	} else {
		fmt.Println("head: illegal line count.")
		return false
	}
}

func readFile(path string, length int) {
	var fp *os.File
	var err error

	fp, err = os.Open(path)

	if err != nil {
		fmt.Printf("==> File not found : %s <==\n", path)
		return
	}
	defer fp.Close()

	fmt.Printf("==> %s <==\n", path)
	scanner := bufio.NewScanner(fp)
	for i := 0; scanner.Scan() && i < length; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("File error.")
	}
}