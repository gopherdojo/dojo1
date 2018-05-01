package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"strconv"
)

func main() {

	fmt.Println("Start typing GAME!!!")
	ch := quiz()
	ch2 := time.After(10 * time.Second)

	j := 0
	for {
		select {
		case <-ch2:
			s := strconv.Itoa(j)
			fmt.Println("Correct Answer is " + s)
			return
		case <-ch:
			j++
		}
	}
}

func quiz() <-chan struct{}{
	ch := make(chan struct{})

	go func(){
		quizes := []string{"hoge", "moga", "fuga", "moga", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge"}

		for _, quiz := range quizes {
			fmt.Print(quiz + ":")
			s:= bufio.NewScanner(os.Stdin)
			s.Scan()
			scan := s.Text()
			if quiz == scan{
				ch <- struct{}{}
			}

		}
		close(ch)

	}()
	return ch
}

