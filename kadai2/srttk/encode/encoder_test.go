package encode_test

import (
	"testing"
	"github.com/srttk/imgconv/encode"
	"os"
)

type encoderBase struct {
	srcExt     string
	readFrom string
}

var encoderBases = []encoderBase{
	{"jpg", "../testdata/test1.jpg"},
	{"jpeg", "../testdata/test2.jpeg"},
	{"png", "../testdata/test3.png"},
	{"gif", "../testdata/test4.gif"},
}

var failEncoderBases = []encoderBase {
	{"txt", "../testdata/test5.txt"},
	{"html", "../testdata/test6.html"},
}

type distPathBase struct {
	srcPath string
	srcExt  string
	distExt string
	result  string
}

var distPathBases = []distPathBase {
	{"../testdata/test1.jpg", "jpg", "png", "../testdata/test1.png"},
	{"../testdata/test2.jpeg", "jpeg", "gif", "../testdata/test2.gif"},
	{"../testdata/test3.png", "png", "jpg", "../testdata/test3.jpg"},
	{"../testdata/test4.gif", "gif", "jpeg", "../testdata/test4.jpeg"},
}

func TestNewEncoder(t *testing.T) {
	for _, encoderBase := range encoderBases {
		reader, err := os.Open(encoderBase.readFrom)
		if err != nil {
			t.Errorf("ファイルを展開できませんでした: %v", err)
		}
		_, err = encode.NewEncoder(encoderBase.srcExt, reader)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewEncoder2(t *testing.T) {
	for _, failEncoderBase := range failEncoderBases {
		reader, err := os.Open(failEncoderBase.readFrom)
		if err != nil {
			t.Errorf("ファイルと展開できませんでした: %v", err)
		}
		_, err = encode.NewEncoder(failEncoderBase.srcExt, reader)
		if err == nil {
			t.Error("未対応フォーマットに対してエラーが発生しませんでした")
		}
	}
}

func TestGetDistPath(t *testing.T) {
	for _, distPathBase := range distPathBases {
		distPath := encode.GetDistPath(distPathBase.srcPath, distPathBase.srcExt, distPathBase.distExt)
		if distPath != distPathBase.result {
			t.Error("期待する出力でありません")
		}
	}
}

func testEncode(t *testing.T, encoder encode.Encoder, tempFilePath string) {
	t.Helper()
	tmpFile, _ := os.Create(tempFilePath)
	err := encoder.Encode(tmpFile)
	if err != nil {
		t.Errorf("encode err!: %v", err)
	}
	os.Remove(tempFilePath)
}
