package typing

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

const FilePath = "./typing/question/proverbs.txt"

func getQuestions() ([]string, error) {
	f, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	questions := make([]string, 0, 10)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		questions = append(questions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func getRandQuestion(questions []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(questions))

	return questions[i]
}
