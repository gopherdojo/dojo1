package converter

import (
	"testing"
	"os"

	"github.com/toshi0607/gopher-dojo/extchanger/testing/helper"
)

func TestConvertExt(t *testing.T) {
	tests := []struct {
		src         string
		from        string
		to          string
		resultCount int
		wantError   bool
	}{
		{"testdata/sample", "jpg", "png", 2, false},
		{"testdata/sample", "PNG", "jpeg", 7, false},
		{"testdata/sample", "png", "jpg", 7, false},
		{"testdata/sample", "hoge", "jpg", 0, true},
		{"testdata/sample", "jpg", "hoge", 0, true},
		{"testdata/sample", "jpg", "jpg", 0, true},
		{"testdata/sample", "", "", 0, true},
	}

	if err := os.MkdirAll("output", 0777); err != nil {
		t.Error("failed to make an output folder")
	}

	for _, te := range tests {

		count, err := ConvertExt(te.src, te.from, te.to);
		helper.TestWantError(t, err, te.wantError)
		if count != te.resultCount {
			t.Errorf("ConvertExt(%v, %v, %v) = %d, got %d",
				te.src, te.from, te.to, te.resultCount, count)
		}
	}

	if err := os.RemoveAll("output"); err != nil {
		t.Error("failed to delete an output folder")
	}

}
