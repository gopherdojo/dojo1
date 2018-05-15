package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const LimitTimeSec time.Duration = 10

var words [8000]string

func init() {
	f, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for i := 0; s.Scan(); i++ {
		words[i] = s.Text()
		i++
	}
}

func genWord() string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(7000)]
}

func main() {
	ch := make(chan string)
	successed := 0

	go func() {
		for {
			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				ch <- s.Text()
			}
		}
	}()

	go func(ch chan string) {
		for {
			question := genWord()
			fmt.Println("[", question, "]")

			input, ok := <-ch
			if !ok && input == "" {
				break
			}

			if question == input {
				successed++
				fmt.Println("OK!")
			} else {
				fmt.Println("MISS!")
			}

			fmt.Println("Point: ", successed)
		}
	}(ch)

	time.Sleep(LimitTimeSec * time.Second)
	close(ch)

	fmt.Println("Finish!")
	fmt.Println("Your point: ", successed)
}
