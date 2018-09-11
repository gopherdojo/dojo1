package typing

import (
	"fmt"
	"math/rand"
	"time"
)

// タイピング問題の一覧
var q = [...]string{"golang", "java", "c++", "javascript", "lisp", "kotlin", "swift", "python", "c", "ruby"}

// 問題を表示する
// 引数で受け取るチャネルからユーザの入力を受けとる
// 正誤判定（正解：true、間違い：false）を関数の戻り値チャネルに送る
func question(ch <-chan string) <-chan bool {
	ret := make(chan bool)
	go func() {
		for {
			// 問題一覧からランダムに問題を出す
			t := time.Now().UnixNano()
			rand.Seed(t)
			s := rand.Intn(len(q))
			// 問題を出す
			fmt.Println("[", q[s], "]")
			// 入力を受け取り判定
			if q[s] == <-ch {
				ret <- true
			}
		}
	}()
	return ret
}
