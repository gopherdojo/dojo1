package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	green = "\u001b[32m"
	red   = "\u001b[31m"
	reset = "\u001b[0m"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("制限時間秒数を指定してください。")
	}
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) <= 2 {
		log.Fatal("出題単語ファイルのパスを指定してください。")
	}

	words, err := makeWords(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	loop(time.Duration(limit), words)
}

func loop(seconds time.Duration, words []string) {
	var corrects, incorrects int
	timeout := time.After(time.Second * seconds)
	ch := answer(words)

	for {
		select {
		case result, ok := <-ch:
			if !ok {
				fmt.Printf("中断しました。 正解=%d, 不正解=%d.\n", corrects, incorrects)
				return
			}
			if result {
				corrects++
			} else {
				incorrects++
			}
		case <-timeout:
			fmt.Printf("終了しました。 正解=%d, 不正解=%d.\n", corrects, incorrects)
			return
		}
	}
}

func answer(words []string) <-chan bool {
	in := input(os.Stdin)
	out := make(chan bool)

	go func() {
		for _, w := range words {
			fmt.Printf("\r%s[%s]%s%s ", green, w, green, reset)
			for v := range in {
				if v == ":q" {
					close(out)
					return
				} else if v == w {
					out <- true
					break
				} else {
					out <- false
					fmt.Printf("%s[%s]%s%s ", red, w, red, reset)
				}
			}
		}
	}()
	return out
}

func makeWords(wordListPath string) ([]string, error) {
	file, err := os.Open(wordListPath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 4096)
	var words []string
	for {
		line, _, err := reader.ReadLine()
		words = append(words, string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}
	}
	shuffle(words)
	return words, nil
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

func shuffle(xs []string) {
	// ref: http://marcelom.github.io/2013/06/07/goshuffle.html
	rand.Seed(time.Now().UnixNano())
	for i := range xs {
		j := rand.Intn(i + 1)
		xs[i], xs[j] = xs[j], xs[i]
	}
}
