package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"./typing"
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

	words, err := typing.MakeWords(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	t := typing.NewTyping(words)
	result := t.Run(time.Duration(limit))
	switch result.Status {
	case typing.ResultClear:
		fmt.Printf("Clear! 正解=%d, 不正解=%d.\n", result.Corrects, result.Incorrects)
	case typing.ResultCancel:
		fmt.Printf("Cancel... 正解=%d, 不正解=%d.\n", result.Corrects, result.Incorrects)
	case typing.ResultTimeOver:
		fmt.Printf("Time over... 正解=%d, 不正解=%d.\n", result.Corrects, result.Incorrects)
	}
}
