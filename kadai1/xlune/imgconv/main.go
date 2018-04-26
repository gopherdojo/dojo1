package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/xlune/dojo1/kadai1/tenntenn/imgconv/cfile"
	"github.com/xlune/dojo1/kadai1/tenntenn/imgconv/cimage"
)

// 引数変数セット
var (
	inputType  string
	outputType string
	outputPath string
)

func init() {
	flag.StringVar(&inputType, "i", "jpg", "変換対象の画像タイプ (default: jpg)")
	flag.StringVar(&outputType, "o", "png", "変換する画像タイプ (default: png)")
	flag.StringVar(&outputPath, "d", "", "画像出力ディレクトリパス")
}

func main() {
	flag.Parse()
	// 指定拡張子チェック
	inputType = strings.ToLower(inputType)
	if !cimage.IsAllowExt(inputType) {
		errorMessage("入力拡張子が不正です")
	}
	outputType = strings.ToLower(outputType)
	if !cimage.IsAllowExt(inputType) {
		errorMessage("出力拡張子が不正です")
	}

	// ターゲットディレクトリチェック
	sourcePath := flag.Arg(0)
	d, err := cfile.NewConvDirInfo(sourcePath)
	if err != nil {
		errorMessage("入力元設定失敗")
	}
	err = d.SetOutputDir(outputPath)
	if err != nil {
		errorMessage("出力先設定失敗")
	}

	// 変換処理リスト取得
	list, err := d.GetFiles(inputType, outputType)
	if err != nil {
		errorMessage("変換リスト取得失敗")
	}

	// 変換実行
	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)
		go func(info cfile.ConvFile) {
			defer wg.Done()
			statusText := "Success"
			err := cimage.Convert(info.InputFile, info.OutputFile)
			if err != nil {
				statusText = "Error"
			}
			fmt.Printf("[%s] %s\n", statusText, info.OutputFile)
		}(v)
	}
	wg.Wait()

	fmt.Println("完了！")
}

// errorMessage エラーを表示してexit
func errorMessage(message string) {
	fmt.Printf("error: %s\n", message)
	os.Exit(1)
}
