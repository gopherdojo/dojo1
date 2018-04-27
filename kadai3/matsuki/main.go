package main

import (
	"time"

	"./questioner"
)

func main() {
	quiz := map[string]string{
		"sort":   "～を仕分ける。s***",
		"figure": "〜と思う。f*****",
		"delay":  "遅れる",
		"depart": "出発する。d*****",
		"eager":  "〜したがる。e***",
		"bother": "悩ます。b*****",
	}
	quizTime := 30 * time.Second
	q := questioner.New(quizTime, quiz)
	q.QuizStart()
}
