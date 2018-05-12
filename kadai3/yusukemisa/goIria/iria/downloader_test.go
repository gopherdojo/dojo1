package iria_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"testing"

	"github.com/yusukemisa/goIria/iria"
)

type TestDownloader struct {
	StatusCode    int
	ContentLength int64
	*iria.Downloader
}

func (d *TestDownloader) Head(ctx context.Context) (*http.Response, error) {
	return &http.Response{
		ContentLength: 9999,
	}, nil
}

/*
	downloader.GetContentLength
	正常系ケース定義
*/
var getContentLengthNomalCases = []NewTestCase{
	{
		name: "正常系_有効URL",
		in:   []string{"206", "99999"},
		out: &iria.Downloader{
			URL:      "http://localhost:0",
			SplitNum: runtime.NumCPU(),
		},
	},
}

/*
	downloader.GetContentLength
	異常系ケース定義
*/
var getContentLengthErrCases = []NewTestCase{
	{
		name: "異常系 Status Not OK",
		in:   []string{"451", "0"},
		out:  "取得対象とするURLを１つ指定してください",
	},
}

/*
	Test Suite Run
	サブ実行:go test -v ./iria -run TestNew/New_正常系
*/
func TestGetContentLength(t *testing.T) {
	t.Run("GetContentLength_正常系", func(t *testing.T) {
		for _, target := range getContentLengthNomalCases {
			fmt.Println(target.name)
			testGetContentLengthNormal(t, target)
		}
	})
	t.Run("GetContentLength_異常系", func(t *testing.T) {
		for _, target := range getContentLengthErrCases {
			fmt.Println(target.name)
			testGetContentLengthError(t, target)
		}
	})
}

//正常系テストコード
func testGetContentLengthNormal(t *testing.T, target NewTestCase) {
	t.Helper()

	status, err := strconv.Atoi(target.in[0])
	if err != nil {
		t.Fail()
	}
	length, err := strconv.Atoi(target.in[1])
	if err != nil {
		t.Fail()
	}

	d := &TestDownloader{
		StatusCode:    status,
		ContentLength: int64(length),
		Downloader: &iria.Downloader{
			URL:      "http://localhost:0",
			SplitNum: runtime.NumCPU(),
		},
	}
	actual, err := d.GetContentLength()
	if err != nil {
		t.Errorf("err expected nil: %v", err.Error())
	}
	if actual == 0 {
		t.Error("New expected Nonnil")
	}
	//構造体の中身ごと一致するか比較
	if !reflect.DeepEqual(actual, target.out) {
		t.Errorf("case:%v => %q, want %v ,actual %v", target.name, target.in, target.out, actual)
	}
}

//異常系テストコード
func testGetContentLengthError(t *testing.T, target NewTestCase) {
	t.Helper()
	_, err := iria.New(target.in)
	if err == nil {
		t.Error("error expected non nil")
	}
	if err.Error() != target.out {
		t.Errorf("case:%v => %q, want %q ,actual %q", target.name, target.in, target.out, err.Error())
	}
}
