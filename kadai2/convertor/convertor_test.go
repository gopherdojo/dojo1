package convertor_test

import (
	"reflect"
	"testing"

	"github.com/yusukemisa/goConvImgExtention/convertor"
)

type TestCase struct {
	name string      //ケース名
	in   []string    //os.Args[]
	out  interface{} //合格条件
}

/*
	convertor.New
	異常系ケース定義
*/
var errCases = []TestCase{
	{
		name: "異常系_引数なし",
		in:   []string{"goConvImgExtention"},
		out:  "変換対象とするディレクトリを１つ指定してください",
	},
	{
		name: "異常系_引数あり_存在しないパス指定",
		in:   []string{"goConvImgExtention", "notExistDir"},
		out:  "stat notExistDir: no such file or directory",
	},
	{
		name: "異常系_フラグあり_f_サポート外",
		in:   []string{"goConvImgExtention", "-f", "bmp", "."},
		out:  "サポート対象外の画像形式が指定されています。",
	},
	{
		name: "異常系_フラグあり_t_サポート外",
		in:   []string{"goConvImgExtention", "-t", "bmp", "."},
		out:  "サポート対象外の画像形式が指定されています。",
	},
}

/*
	convertor.New
	正常系ケース定義
*/
var nomalCases = []TestCase{
	{
		name: "正常系_引数あり_フラグなし",
		in:   []string{"goConvImgExtention", "."},
		out: &convertor.Convertor{
			From:       "jpg",
			To:         "png",
			TargetPath: ".",
		},
	},
	{
		name: "正常系_引数あり_フラグあり_f=png",
		in:   []string{"goConvImgExtention", "-f", "png", "."},
		out: &convertor.Convertor{
			From:       "png",
			To:         "png",
			TargetPath: ".",
		},
	},
	{
		name: "正常系_引数あり_フラグあり_t=jpg",
		in:   []string{"goConvImgExtention", "-t", "jpg", "."},
		out: &convertor.Convertor{
			From:       "jpg",
			To:         "jpg",
			TargetPath: ".",
		},
	},
	{
		name: "正常系_引数あり_フラグあり_f=gif_t=jpeg",
		in:   []string{"goConvImgExtention", "-f", "gif", "-t", "jpeg", "."},
		out: &convertor.Convertor{
			From:       "gif",
			To:         "jpeg",
			TargetPath: ".",
		},
	},
}

/*
	Test Suite Run
*/
func TestAll(t *testing.T) {
	t.Run("New異常系", func(t *testing.T) {
		for _, target := range errCases {
			testNewError(t, target)
		}
	})
	t.Run("New正常系", func(t *testing.T) {
		for _, target := range nomalCases {
			testNewNormal(t, target)
		}
	})
}

//異常系テストコード
func testNewError(t *testing.T, target TestCase) {
	t.Helper()
	actual, err := convertor.New(target.in)
	if actual != nil {
		t.Error("actual expected nil")
	}
	if err.Error() != target.out {
		t.Errorf("case:%v => %q, want %q ,actual %q", target.name, target.in, target.out, err.Error())
	}
}

//正常系テストコード
func testNewNormal(t *testing.T, target TestCase) {
	t.Helper()
	actual, err := convertor.New(target.in)
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
