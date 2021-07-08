package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"time"

	"io/ioutil"

	"flag"

	"github.com/mpppk/go-typing-game/typing"
	"gopkg.in/yaml.v2"
)

var configFilePath string

const timeLimitSec = 60 * time.Second

type Config struct {
	Questions []string
}

func main() {
	flag.Parse()
	config, err := loadConfigFromYaml(configFilePath)
	if err != nil {
		fmt.Printf("failed to load config file from %s\n", configFilePath)
		os.Exit(1)
	}
	ch := input(os.Stdin)
	manager := typing.NewManager()
	manager.AddQuestions(config.Questions)

	timeoutChan := time.After(timeLimitSec)
	score := 0

	for manager.SetNewQuestion() {
		for { // 正しい解答が入力されるまでループ
			fmt.Println("Q: " + manager.GetCurrentQuestion())
			fmt.Print(">")
			answer, timeout := waitAnswerOrTimeout(ch, timeoutChan)
			if timeout {
				fmt.Printf("\ntime up! Your score is %d\n", score)
				return
			}

			if manager.ValidateAnswer(answer) {
				fmt.Println("Correct!")
				score++
				break
			} else {
				fmt.Println("invalid answer... try again")
			}
		}
	}
	fmt.Printf("all questions are answered! your score is %d\n", score)
}

func waitAnswerOrTimeout(answerCh <-chan string, timeoutChan <-chan time.Time) (string, bool) {
	for {
		select {
		case answer := <-answerCh:
			return answer, false
		case <-timeoutChan:
			return "", true
		}
	}
}

func loadConfigFromYaml(filePath string) (*Config, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func input(r io.Reader) <-chan string {
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

func init() {
	flag.StringVar(&configFilePath, "config", "config.yaml", "config file path")
}
