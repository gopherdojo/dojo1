package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// 質問格納用struct
type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// 回答結果の格納用struct
type Result struct {
	Count   int
	Correct int
}

// JSONファイルから問題をロード
func LoadQuestions(path string) []Question {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	if err := json.Unmarshal(bytes, &questions); err != nil {
		log.Fatal(err)
	}
	return questions
}

// 問題の配列からランダムに一問抽出
func GetQuestion(questions []Question) Question {
	rand.Seed(time.Now().UnixNano())
	rand_index := rand.Intn(len(questions))
	return questions[rand_index]
}

// 問題の表示
func PopQuestion(questions []Question, q_count int) Question {
	question := GetQuestion(questions)
	fmt.Printf("Q%d \"%s\" は日本語で？ : ", q_count, question.Question)
	return question
}

// 回答があっているか確認して結果を表示
func CheckAnswer(question Question, answer string) bool {
	correct := false
	if answer == question.Answer {
		fmt.Printf("正解です!\n")
		correct = true
	} else {
		fmt.Printf("不正解です。 正解は、\"%s\" です。\n", question.Answer)
	}
	return correct
}

// 正解の合計を表示
func ShowResult(result Result) {
	fmt.Printf("\n!!! タイムオーバーです。!!!\n")
	fmt.Printf("問題%d問中、正解は%d問でした。\n", result.Count, result.Correct)
}
