package main

import (
	"fmt"
	"os"

	"github.com/yusukemisa/goConvImgExtention/convertor"
)

// go install github.com/yusukemisa/goConvImgExtention
func main() {
	//フラグにサポート対象の拡張子が指定されている場合変換実行
	c, err := convertor.New(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		convertor.Usage()
		os.Exit(1)
	}
	if err = c.Convert(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		convertor.Usage()
		os.Exit(1)
	}
	os.Exit(0)
}
