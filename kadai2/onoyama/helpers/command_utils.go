// [image_converter]コマンドの実行支援用パッケージ
package helpers

import (
	"flag"
	"fmt"
	"os"
)

// コマンドの使用例を表示
func PrintUsage() {
	fmt.Println("Usage: image_converter -d=出力先フォルダ -f=変換後の画像形式(jpegまたはpng) 変換元フォルダ")
}

// コメント、エラー、使用例を表示して、異常終了
func PrintMsgAndDie(comment string, message error) {
	if message != nil {
		fmt.Println(message)
	}
	fmt.Println(comment + "\n")
	PrintUsage()
	os.Exit(1)
}

// コマンドの引数の確認
func CheckParams() (string, string, string) {
	dstDir := flag.String("d", "./output", "出力先フォルダ")
	dstFormat := flag.String("f", "jpg", "変換後の画像形式(jpegまたはpng)")
	flag.Parse()

	var srcDir string
	if len(flag.Args()) != 1 {
		PrintMsgAndDie("変換元フォルダの指定が正しくありません", nil)
	} else {
		srcDir = flag.Args()[0]
	}
	return srcDir, *dstDir, *dstFormat
}
