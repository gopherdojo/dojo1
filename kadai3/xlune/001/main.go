package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/xlune/dojo1/kadai3/xlune/001/word"
)

var (
	limitSec int
)

func init() {
	flag.IntVar(&limitSec, "t", 30, "ゲームの制限時間")
}

func main() {
	flag.Parse()

	// 入力チャネル
	ch := input(os.Stdin)
	// タイムアウト設定
	timeout := time.After(time.Duration(limitSec) * time.Second)
	// 最初の問題生成
	makeQuestion()

	fmt.Printf("英語タイピングスタート!!\n(制限時間: %d 秒)\n\n", limitSec)
	for {
		fmt.Printf("出題 > %s\n", word.GetLatest())
		select {
		case result := <-ch:
			if word.CheckLatest(result) {
				fmt.Println("=> OK!!")
				// 次の問題生成
				makeQuestion()
			} else {
				fmt.Println("=> NG...")
			}
		case <-timeout:
			fmt.Printf("終了それまで!!\n--\n正解数は %d 問でした。\n\n", word.CountHistory()-1)
			os.Exit(0)
		}
	}
}

func makeQuestion() {
	_, err := word.Issue()
	if err != nil {
		fmt.Println("Error: 出題できませんでした")
		os.Exit(1)
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
