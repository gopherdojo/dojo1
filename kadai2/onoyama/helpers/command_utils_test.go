package helpers

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintUsage(t *testing.T) {
	expect := "Usage: image_converter -d=出力先フォルダ -f=変換後の画像形式(jpegまたはpng) 変換元フォルダ"
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	PrintUsage()
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = out
	actual := <-outC
	actual = strings.TrimSpace(actual)

	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}
