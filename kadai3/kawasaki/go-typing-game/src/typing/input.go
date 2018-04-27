package typing

import (
	"bufio"
	"io"
)

// ユーザからの入力を待ち受ける
// 入力された内容は、戻り値のチャネルに送る
func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			ch <- sc.Text()
		}
		close(ch)
	}()
	return ch
}
