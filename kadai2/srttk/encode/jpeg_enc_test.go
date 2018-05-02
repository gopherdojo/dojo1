package encode_test

import (
	"testing"
	"os"

	"github.com/srttk/imgconv/encode"
)

func testNewJpegEncoder(t *testing.T, readFrom string) encode.Encoder {
	t.Helper()
	reader, err := os.Open(readFrom)
	if err != nil {
		t.Errorf("ファイルの展開に失敗しました: %v", err)
	}
	encoder, err := encode.NewJpegEncoder(reader)
	if err != nil {
		t.Errorf("エンコーダの作成に失敗しました: %v", err)
	}
	return encoder
}

func TestJpegEncoder_Encode(t *testing.T) {
	encoder := testNewJpegEncoder(t, "../testdata/test1.jpg")
	testEncode(t, encoder, "../testdata/tmp/test1.png")
}
