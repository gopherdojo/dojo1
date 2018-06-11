package convertor

import (
	"fmt"
)

//ImageExtention サポートする画像形式
var ImageExtention = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"gif":  true,
}

//SupportDescription サポートする画像形式出力
var SupportDescription = func() string {
	return fmt.Sprintf("<サポートする画像形式>\n%v\n", Keys(ImageExtention))
}

//Keys ImageExtentionのキーだけ取り出す
func Keys(m map[string]bool) []string {
	ks := []string{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
