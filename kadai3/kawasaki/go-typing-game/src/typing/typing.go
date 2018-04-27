package typing

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// Start タイピングゲームを開始する
// 入力先と出力先を指定してください
func Start(r io.Reader, w io.Writer) {
	fmt.Fprintln(w, "######################################")
	fmt.Fprintln(w, "###  Golang Typing Game (ver.beta) ###")
	fmt.Fprintln(w, "### * Type more words in 10 second ###")
	fmt.Fprintln(w, "### * To start, press any button.. ###")
	fmt.Fprintln(w, "######################################")

	// スタートボタンの入力をまつ
	s := bufio.NewScanner(r)
	s.Scan()
	s.Text()

	// ユーザからの入力を待つ関数
	ch := input(r)

	// 問題をだして、正誤判定する関数
	ret := question(ch)

	// 正解数をカウントする関数
	count := counter(ret)

	// 10秒待つ
	<-time.After(10 * time.Second)
	fmt.Fprintln(w, "### TIME UP!! Your score is", *count, ":) ###")
}
