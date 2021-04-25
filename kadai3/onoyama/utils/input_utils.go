package utils

import (
	"bufio"
	"io"
	"time"
)

// 入力用のチャネル
func InputChannel(r io.Reader) <-chan string {
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

// タイムアウト用のチャネル
func TimeoutChannel(timeout int) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		ch <- true
		close(ch)
	}()
	return ch
}
