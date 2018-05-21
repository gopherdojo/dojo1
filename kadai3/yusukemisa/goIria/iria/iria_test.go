package iria_test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/yusukemisa/goIria/iria"
)

type NewTestCase struct {
	name string      //ケース名
	in   []string    //前提条件
	out  interface{} //合格条件
}

/*
	iria.New
	正常系ケース定義
*/
var newNomalCases = []NewTestCase{
	{
		name: "正常系_有効URL",
		in:   []string{"goIria", "http://localhost:0"},
		out: &iria.Downloader{
			URL:      "http://localhost:0",
			SplitNum: runtime.NumCPU(),
		},
	},
}

/*
	iria.New
	異常系ケース定義
*/
var newErrCases = []NewTestCase{
	{
		name: "異常系_引数なし",
		in:   []string{"goIria"},
		out:  "取得対象とするURLを１つ指定してください",
	},
	{
		name: "異常系_引数多すぎ",
		in:   []string{"goIria", "test", "ヘテロヘテロ", "地固めがすごい"},
		out:  "取得対象とするURLを１つ指定してください",
	},
	{
		name: "異常系_取得ファイル重複",
		in:   []string{"goIria", "."},
		out:  "取得対象のファイルが既に存在しています:.",
	},
}

/*
	Test Suite Run
	サブ実行:go test -v ./iria -run TestNew/New_正常系
*/
func TestNew(t *testing.T) {
	t.Run("New_正常系", func(t *testing.T) {
		for _, target := range newNomalCases {
			fmt.Println(target.name)
			testNewNormal(t, target)
		}
	})
	t.Run("New_異常系", func(t *testing.T) {
		for _, target := range newErrCases {
			fmt.Println(target.name)
			testNewError(t, target)
		}
	})
}

//正常系テストコード
func testNewNormal(t *testing.T, target NewTestCase) {
	t.Helper()
	actual, err := iria.New(target.in)

	if err != nil {
		t.Errorf("err expected nil: %v", err.Error())
	}
	if actual == nil {
		t.Error("New expected Nonnil")
	}
	//構造体の中身ごと一致するか比較
	if !reflect.DeepEqual(actual, target.out) {
		t.Errorf("case:%v => %q, want %v ,actual %v", target.name, target.in, target.out, actual)
	}
}

//異常系テストコード
func testNewError(t *testing.T, target NewTestCase) {
	t.Helper()
	_, err := iria.New(target.in)
	if err == nil {
		t.Error("error expected non nil")
	}
	if err.Error() != target.out {
		t.Errorf("case:%v => %q, want %q ,actual %q", target.name, target.in, target.out, err.Error())
	}
}
