package questioner

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// QuizConf は
type QuizConf struct {
	quizTime       time.Duration
	quiz           map[string]string // answer: question
	collectAnswers int
}

// New は、初期化関数
func New(timeoutSecond time.Duration, quizMap map[string]string) *QuizConf {
	q := QuizConf{}
	q.quizTime = timeoutSecond
	q.quiz = quizMap
	q.collectAnswers = 0
	return &q
}

// getterInput は、標準入力から受け付ける
func getterInput(r io.Reader) <-chan string {
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

func choiceQuiz(quizMap map[string]string) (answer, question string) {
	rand.Seed(time.Now().UnixNano())
	i := 0
	index := rand.Intn(len(quizMap))
	for answer, question := range quizMap {
		if index == i {
			return answer, question
			break
		} else {
			i++
		}
	}
	return
}

func (q *QuizConf) quizInfo() {
	fmt.Println("日本語の意味が表示されるので、それを表す英単語を答えてください。")
	fmt.Println("制限時間は、", q.quizTime, "です")
	fmt.Println("Quiz Start! \n")
}

func (q *QuizConf) judge(userInput, answer string) {
	if userInput == answer {
		q.collectAnswers++
		fmt.Println("正解!!", answer, " 現在の正解数=", q.collectAnswers)
	} else {
		fmt.Println("不正解。正解は、", answer)
	}

}

// QuizStart はクイズを開始する
func (q *QuizConf) QuizStart() {
	var (
		answer   string
		question string
	)
	q.quizInfo()
	ch := getterInput(os.Stdin)
	timeout := time.After(q.quizTime)
	for {
		select {
		case <-timeout:
			fmt.Println("\n\ntimeout !!")
			fmt.Println("クイズ終了。あなたの成績=", q.collectAnswers)
			return
		default:
			answer, question = choiceQuiz(q.quiz)
			fmt.Println(question)
			if v, ok := <-ch; ok {
				q.judge(v, answer)
			} else {
				os.Exit(1)
			}
		}
	}
}
