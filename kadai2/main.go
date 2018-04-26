package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	convimg "github.com/matsu0228/go_sandbox/02_convimg_test/convimg"
)

var (
	trgDir      string // 処理対象のディレクトリ
	convCommand string // 変換ルール (変換元の画像形式:変換後の画像形式)
	srcFormat   string
	destFormat  string
)

func init() {
	flag.StringVar(&trgDir, "d", "./", "画像変換したいディレクトリを指定する")
	flag.StringVar(&convCommand, "f", "jpg:png", "変換元の画像形式:変換後の画像形式 (png/jpg/gifのみサポート)")
	flag.Parse()
}

func errorExit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

func main() {
	// 変換形式のフォーマット数チェック
	formats := strings.Split(convCommand, ":")
	if len(formats) != 2 {
		errorExit(errors.New(
			"画像フォーマットの指定数が誤っています=> (" + strconv.Itoa(len(formats)) + ") " + convCommand,
		))
	}
	srcFormat = strings.ToLower(formats[0])
	destFormat = strings.ToLower(formats[1])

	c, err := convimg.New(trgDir, srcFormat, destFormat)
	if err != nil {
		errorExit(err)
	}
	if _, err := c.ConvImages(); err != nil {
		errorExit(err)
	}
}
