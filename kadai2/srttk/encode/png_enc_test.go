package encode_test

import (
	"testing"
	"os"

	"github.com/srttk/imgconv/encode"
)

func testNewPngEncoder(t *testing.T, readFrom string) encode.Encoder {
	t.Helper()
	reader, err := os.Open(readFrom)
	if err != nil {
		t.Errorf("ファイルの展開に失敗しました: %v", err)
	}
	encoder, err := encode.NewPngEncoder(reader)
	if err != nil {
		t.Errorf("エンコーダの作成に失敗しました: %v", err)
	}
	return encoder
}

func TestPngEncoder_Encode(t *testing.T) {
	encoder := testNewPngEncoder(t, "../testdata/test3.png")
	testEncode(t, encoder, "../testdata/tmp/test3.jpg")
}
