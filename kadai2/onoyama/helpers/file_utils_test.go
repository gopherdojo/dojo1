package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

var base_dir, _ = filepath.Abs("../")
var test_dir = base_dir + "/test"

var ext_tests = []struct {
	in  string
	out bool
}{
	{".jpg", true},
	{".jpeg", true},
	{".png", true},
	{".pdf", false},
	{".zip", false},
}

var dir_tests = []struct {
	in  string
	out bool
}{
	{base_dir + "/testdata/jpeg", true},
	{base_dir + "/testdata/png", true},
	{base_dir + "/testdata/gif", false},
	{base_dir + "/testdata/pdf", false},
}

var dir_walker_tests = []struct {
	in  string
	out FileSpec
}{
	{base_dir + "/testdata/jpeg/video-001.221212.jpeg", FileSpec{
		DirPath:  base_dir + "/testdata/jpeg",
		FileName: "video-001.221212.jpeg",
		FileExt:  ".jpeg",
		BaseName: "video-001.221212",
	}},
	{base_dir + "/testdata/jpeg/video-001.q50.410.jpeg", FileSpec{
		DirPath:  base_dir + "/testdata/jpeg",
		FileName: "video-001.q50.410.jpeg",
		FileExt:  ".jpeg",
		BaseName: "video-001.q50.410",
	}},
	{base_dir + "/testdata/png/video-001.221212.png", FileSpec{
		DirPath:  base_dir + "/testdata/png",
		FileName: "video-001.221212.png",
		FileExt:  ".png",
		BaseName: "video-001.221212",
	}},
}

var make_convertspec_tests = []struct {
	in  string
	out ConvertSpec
}{
	{base_dir + "/testdata/jpeg/video-001.221212.jpeg", ConvertSpec{
		Src:    base_dir + "/testdata/jpeg/video-001.221212.jpeg",
		Dst:    test_dir + "/testdata/video-001.221212.png",
		Format: "png",
	}},
	{base_dir + "/testdata/jpeg/video-001.q50.410.jpeg", ConvertSpec{
		Src:    base_dir + "/testdata/jpeg/video-001.q50.410.jpeg",
		Dst:    test_dir + "/testdata/video-001.q50.410.png",
		Format: "png",
	}},
	{base_dir + "/testdata/png/video-001.221212.png", ConvertSpec{
		Src:    base_dir + "/testdata/png/video-001.221212.png",
		Dst:    test_dir + "/testdata/video-001.221212.png",
		Format: "png",
	}},
}

func TestDirWalker(t *testing.T) {
	actuals := DirWalker(base_dir + "/testdata")
	if len(actuals) != 31 {
		t.Errorf(`slice length : expect="%s" actual="%s"`, 31, len(actuals))
	}
	for _, tt := range dir_walker_tests {
		for _, f := range actuals {
			if f.DirPath+"/"+f.FileName == tt.in {
				if f.DirPath != tt.out.DirPath {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.DirPath, f.DirPath)
				}
				if f.FileName != tt.out.FileName {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.FileName, f.FileName)
				}
				if f.FileExt != tt.out.FileExt {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.FileExt, f.FileExt)
				}
				if f.DirPath != tt.out.DirPath {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.DirPath, f.DirPath)
				}
			}
		}
	}
}

func TestMakeConvertSpec(t *testing.T) {
	actual_filespecs := DirWalker(base_dir + "/testdata")
	actuals := MakeConvertSpec(actual_filespecs, test_dir+"/testdata", "png")
	if len(actuals) != 31 {
		t.Errorf(`slice length : expect="%s" actual="%s"`, 31, len(actuals))
	}
	for _, tt := range make_convertspec_tests {
		for _, c := range actuals {
			if c.Src == tt.in {
				if c.Src != tt.out.Src {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.Src, c.Src)
				}
				if c.Dst != tt.out.Dst {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.Dst, c.Dst)
				}
				if c.Format != tt.out.Format {
					t.Errorf(`%s : expect="%s" actual="%s"`, tt.in, tt.out.Format, c.Format)
				}
			}
		}
	}
}

func TestCheckExt(t *testing.T) {
	var actual bool
	for _, tt := range ext_tests {
		actual = checkExt(tt.in)
		if tt.out != actual {
			t.Errorf(`expect="%s" actual="%s"`, tt.out, actual)
		}
	}
}

func TestDirCheck(t *testing.T) {
	var actual bool
	for _, tt := range dir_tests {
		actual = dirCheck(tt.in)
		if tt.out != actual {
			t.Errorf(`expect="%s" actual="%s"`, tt.out, actual)
		}
	}
}

func TestdirCheckAndMkdir(t *testing.T) {
	temp_dir := test_dir + "/test_dir_check_and_mkdir"
	expect := true
	actual := dirCheckAndMkdir(test_dir)
	if expect != actual {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
	err := os.RemoveAll(temp_dir)
	if err != nil {
		t.Error(`後処理エラー: %s の削除に失敗しました。`, temp_dir)
	}
}
