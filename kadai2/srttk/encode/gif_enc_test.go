package encode_test

import (
	"testing"
	"os"

	"github.com/srttk/imgconv/encode"
)

func testNewGifEncoder(t *testing.T, readFrom string) encode.Encoder {
	t.Helper()
	reader, err := os.Open(readFrom)
	if err != nil {
		t.Errorf("ファイルの展開に失敗しました: %v", err)
	}
	encoder, err := encode.NewGifEncoder(reader)
	if err != nil {
		t.Errorf("エンコーダの作成に失敗しました: %v", err)
	}
	return encoder
}

func TestGifEncoder_Encode(t *testing.T) {
	encoder := testNewGifEncoder(t, "../testdata/test4.gif")
	testEncode(t, encoder, "../testdata/tmp/test4.png")
}
