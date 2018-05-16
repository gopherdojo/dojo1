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

var words = make([]string, 0, 1000)

func init() {
	f, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
}

func genWord() string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
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
				fmt.Print("OK! ")
			} else {
				fmt.Print("MISS! ")
			}

			fmt.Println("Point: ", successed)
		}
	}(ch)

	time.Sleep(LimitTimeSec * time.Second)
	close(ch)

	fmt.Println("Finish!")
	fmt.Println("Your point: ", successed)
}
