package iria

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

//New create Downloader
func New(args []string) (*Downloader, error) {
	if len(args) != 2 {
		return nil, errors.New("取得対象とするURLを１つ指定してください")
	}
	//取得対象ファイルと同名のファイルが既にある場合を許さない
	targetFileName := filepath.Base(args[1])
	if exists(targetFileName) {
		return nil, fmt.Errorf("取得対象のファイルが既に存在しています:%v", targetFileName)
	}
	return &Downloader{
		URL:      args[1],
		SplitNum: runtime.NumCPU(), //CPUコア数だけダウンロードを分割する
	}, nil
}

//rangeヘッダに指定する値を算出する
//@return []string	rangeヘッダ指定値	{"0-N","N+1-M",..."M-contentLength"}
func getByteRange(contentLength int64, splitNum int) (rangeArr []string) {
	var from, to int64
	chunkLength := contentLength / int64(splitNum)
	for i := 0; i < splitNum; i++ {
		switch i {
		case 0:
			from = 0
			to = chunkLength
		case splitNum - 1:
			from = to + 1
			to = contentLength
		default:
			from = to + 1
			to += chunkLength
		}
		rangeArr = append(rangeArr, fmt.Sprintf("%v-%v", from, to))
	}
	return rangeArr
}

//ファイル存在チェック
func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
