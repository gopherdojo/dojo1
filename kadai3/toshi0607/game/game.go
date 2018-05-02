package game

import (
	"io"
	"os"
	"time"
	"fmt"
	"bufio"
	"github.com/toshi0607/dojo1/kadai3/toshi0607/word"
	"math"
)

type Game struct {
	OutStream, ErrStream io.Writer
}

func (g *Game) Run() int {
	sch := start(os.Stdin)
	<-sch

	ch := input(os.Stdin)

	words, err := word.GetWords()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	timeChan := time.NewTimer(30 * time.Second).C

	correctAnswerCount := 0
	answerCount := 0
Outer:
	for _, v := range words {
		fmt.Printf("type %s\n", v)
	Inner:
		for {
			fmt.Print(">")
			select {
			case <-timeChan:
				break Outer
			case s := <-ch:
				if s == v {
					fmt.Println("ええで")
					correctAnswerCount++
					answerCount++
					break Inner
				} else {
					fmt.Println("あかんで")
					answerCount++
					fmt.Printf("type %s\n", v)
				}
			}

		}
	}
	fmt.Println("\n終わりやで")
	fmt.Printf("正解数: %d words\n", correctAnswerCount)
	correctness := float64(0)
	if answerCount != 0 {
		correctness = roundPlus(float64(correctAnswerCount) / float64(answerCount) * 100, 2)
	}
	fmt.Printf("正確さ: %v %%\n", correctness)
	return 0
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func roundPlus(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return round(f * shift) / shift;
}

func start(r io.Reader) <-chan struct{} {
	fmt.Println("\nタイピングゲームやで")
	fmt.Println("\n制限時間は30秒")
	fmt.Println(">>> press any key to start <<<")
	ch := make(chan struct{})
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- struct{}{}
			break
		}
		close(ch)
	}()
	return ch
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
