package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// テスト問題用のJSONファイル
var test_json_path = "../json/questions.json"
var questions []Question

// テスト用問題の解答
var expect_questions = []Question{
	Question{"test1", "テスト1"},
	Question{"test2", "テスト2"},
	Question{"test3", "テスト3"},
}

// 問題のロードのテスト
func TestLoadQuestions(t *testing.T) {
	questions = LoadQuestions(test_json_path)
	expect_len := 3
	if actual_len := len(questions); expect_len != actual_len {
		t.Error(`expect="%s" actual="%s"`, expect_len, actual_len)
	}
	for i, question := range questions {
		if expect_questions[i] != question {
			t.Error(`expect="%s" actual="%s"`, expect_questions[i], question)
		}
	}
}

// 問題の解答が正しいかテスト
func TestGetQuestion(t *testing.T) {
	question := GetQuestion(questions)
	is_question := false
	for _, expect := range expect_questions {
		if question.Question == expect.Question {
			if question.Answer != expect.Answer {
				t.Error(`expect="%s" actual="%s"`, expect.Answer, question.Answer)
			}
			is_question = true
		}
	}
	if !is_question {
		t.Error(`Invalid question : "%v"`, question)
	}
}

// 問題の表示が正しいかテスト
func TestPopQuestion(t *testing.T) {
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	expect_question := PopQuestion(questions, 5)
	expect := fmt.Sprintf("Q%d \"%s\" は日本語で？ : ", 5, expect_question.Question)
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = out
	actual := <-outC

	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

// 回答が正解の場合の表示のテスト
func TestCheckAnswer1(t *testing.T) {
	question := GetQuestion(questions)
	expect := fmt.Sprintf("正解です!\n")
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	CheckAnswer(question, question.Answer)
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = out
	actual := <-outC

	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

// 回答が不正解の場合のテスト
func TestCheckAnswer2(t *testing.T) {
	question := GetQuestion(questions)
	expect := fmt.Sprintf("不正解です。 正解は、\"%s\" です。\n", question.Answer)
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	CheckAnswer(question, "xxxxxxxxxxxx")
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = out
	actual := <-outC

	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

// タイムアウト後の結果表示のテスト
func TestShowResult(t *testing.T) {
	expect_result := Result{10, 5}
	expect := fmt.Sprintf("\n!!! タイムオーバーです。!!!\n問題%d問中、正解は%d問でした。\n", expect_result.Count, expect_result.Correct)
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	ShowResult(expect_result)
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = out
	actual := <-outC

	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
