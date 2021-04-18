package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo1/kadai3/execjosh/go-typing-game/internal/game"
	"github.com/gopherdojo/dojo1/kadai3/execjosh/go-typing-game/internal/wordbank"
)

const (
	wordsFilePath = "words_alpha.txt"
	timeout       = 30 * time.Second
)

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	wb, err := loadWordBankFromFile(wordsFilePath)
	if err != nil {
		die(err)
	}

	fmt.Printf("You have %v!\n", timeout)

	stats := game.Run(os.Stdin, os.Stdout, wb, timeout)

	fmt.Println("\n---\nTime's Up, Grasshopper!")
	fmt.Println("\x1b[92mSuccess count: ", stats.SuccessCount(), "\x1b[m")
	fmt.Println("\x1b[31mFailure count: ", stats.FailureCount(), "\x1b[m")
}

func loadWordBankFromFile(filepath string) (wordbank.WordProvider, error) {
	f, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New(fmt.Sprint("expected ", wordsFilePath, " to exist"))
		}
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var words []string

	for s.Scan() {
		w := s.Text()
		words = append(words, w)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	wb := wordbank.NewRandomWordBank(words, time.Now().UnixNano())

	return wb, nil
}
