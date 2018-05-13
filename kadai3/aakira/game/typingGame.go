package game

import (
	"context"
	"time"
	"os"
	"fmt"
	"github.com/aakira/typinggame/command"
)

func Start(wordList []string, second time.Duration) int {
	bc := context.Background()
	t := second * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	count := 0

	inputChannel := command.Input(os.Stdin)

LABEL:
	for {
		word := wordList[count]
		fmt.Printf("Input [%s] > ", word)

		select {
		case input := <-inputChannel:
			if word == input {
				fmt.Println("Right!")
				count++
			} else {
				fmt.Println("Wrong.")
			}
		case <-ctx.Done():
			fmt.Println("\n========= Finish! =========")
			break LABEL
		}
	}
	return count
}
