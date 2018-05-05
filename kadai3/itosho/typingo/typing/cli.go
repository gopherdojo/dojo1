package typing

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func Run(timeLimit time.Duration) bool {
	isReady, err := ready(timeLimit)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ready failed. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	if !isReady {
		return true
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeLimit*time.Second)
	defer cancel()

	if err := play(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Play failed. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	return true
}

func ready(seconds time.Duration) (bool, error) {
	fmt.Println("Welcome dead simple typing game to learn software proverbs!")
	fmt.Println(fmt.Sprintf("Time limit is %d seconds", seconds))
	fmt.Println("Are you ready? [Y/n]")
	fmt.Print("> ")

	return yesOrNo()
}

func yesOrNo() (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadByte()
	if err != nil {
		return false, err
	}

	if s == []byte("Y")[0] || s == []byte("y")[0] {
		fmt.Println("Ready Go!")
		return true, nil
	} else if s == []byte("N")[0] || s == []byte("n")[0] {
		fmt.Println("Good Bye!")
		return false, nil
	}

	return false, errors.New("please enter yes or no")
}
