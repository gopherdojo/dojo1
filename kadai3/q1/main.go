package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"time"
	"math/rand"
)

func main() {
	// 問題集
	words := [...] string{"pen", "pineapple", "apple"}
	var score int = 0
	// OP表示
	fmt.Println("Welcome to dayamada's typing game.")
	// 入力処理
	ch := input(os.Stdin)
	// 乱数初期化
	rand.Seed(time.Now().UnixNano())
	for {
		// お題の取得
		index := rand.Intn(len(words))
		fmt.Println(words[index])
		fmt.Print(">")
		select {
		case answer := <-ch:
			if answer == words[index] {
				fmt.Println("great!!")
				score++
			} else {
				fmt.Println("typo!")
			}
		case <-time.After(time.Second * 10):
			fmt.Println("timeout!!")
			fmt.Println("Score:", score)
			return
		}
	}
}

func input(r io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()
	return c
}
