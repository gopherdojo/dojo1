package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var correctNum int

//Game は正解数とゲーム終了の契機を保持します
type Game struct {
	CorrectNum     int
	GameOverReason string
}

//StartWithTimer the typing game
func StartWithTimer(timer time.Duration) *Game {
	//Game終了かタイムアップイベントを受け取る
	done := make(chan string)

	go setTimer(done, timer)
	go playGame(done)

	//Game終了かタイムアップするまでブロック
	gameOver := <-done
	close(done)

	//Game結果
	return &Game{
		CorrectNum:     correctNum,
		GameOverReason: gameOver,
	}
}

func setTimer(done chan<- string, timer time.Duration) {
	time.Sleep(time.Second * timer)
	done <- "TIME UP!!!"
}

//play the game
//done:ゲーム終了時このチャネルに通知を出す
func playGame(done chan<- string) {
	ch := input(os.Stdin)
	for i, word := range getQuestion() { //テストするときgetQuestionモックにしたい
		//出題
		fmt.Printf("Q%v %v\n>", i+1, word)
		if answer, ok := <-ch; ok {
			//答え合わせ
			if answer == word {
				fmt.Println("GREAT WORK!!!")
				correctNum++
			} else {
				fmt.Println("おまえは何をやっているんだ！？")
			}
		} else {
			break
		}
	}
	done <- "YOU HAVE ANSERED ALL QUESTION!!!"
}

func input(r io.Reader) <-chan string {
	recv := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			recv <- s.Text()
		}
		close(recv)
	}()
	return recv
}
