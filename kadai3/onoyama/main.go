package main

import (
	"./utils"
	"fmt"
	"os"
)

// 問題と解答のリスト
var json_path = "./json/words.json"

// タイムアウト設定
var timeout_seconds = 60

func main() {
	var question utils.Question
	questions := utils.LoadQuestions(json_path)
	questioning := false
	results := utils.Result{0, 0}

	input_ch := utils.InputChannel(os.Stdin)
	timeout_ch := utils.TimeoutChannel(timeout_seconds)

	fmt.Printf("制限時間は、%d秒です。\n", timeout_seconds)

	for {
		select {
		// 回答があっているかの確認
		case answer := <-input_ch:
			if result := utils.CheckAnswer(question, answer); result == true {
				results.Correct++
			}
			questioning = false
		case <-timeout_ch:
			// タイムアウトで結果を表示
			utils.ShowResult(results)
			os.Exit(1)
		default:
			//質問がされていない場合は問題を出力
			if !questioning {
				results.Count++
				question = utils.PopQuestion(questions, results.Count)
				questioning = true
			}
		}
	}
}
