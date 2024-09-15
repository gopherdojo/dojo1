package helpers

import (
	"testing"
)

func TestPermitExt(t *testing.T) {
	expect := []string{".jpeg", ".jpg", ".png"}
	actual := PermitExt
	for i := range actual {
		if expect[i] != actual[i] {
			t.Errorf(`expect="%s" actual="%s"`, expect[i], actual[i])
		}
	}
}

func TestTargetExt(t *testing.T) {
	expect := map[string]string{".jpeg": ".jpg", ".jpg": ".jpg", ".png": ".png"}
	actual := TargetExt
	for k, _ := range expect {
		if expect[k] != actual[k] {
			t.Errorf(`map key: %s expect="%s" actual="%s"`, k, expect[k], actual[k])
		}
	}
}

func TestCovertSpec(t *testing.T) {
	expect_src := "source/path"
	expect_dst := "dest/path"
	expect_format := "format"
	actual := ConvertSpec{
		Src:    "source/path",
		Dst:    "dest/path",
		Format: "format",
	}
	if expect_src != actual.Src {
		t.Errorf(`expect="%s" actual="%s"`, expect_src, actual.Src)
	}
	if expect_dst != actual.Dst {
		t.Errorf(`expect="%s" actual="%s"`, expect_dst, actual.Dst)
	}
	if expect_format != actual.Format {
		t.Errorf(`expect="%s" actual="%s"`, expect_format, actual.Src)
	}
}

func TestFileSpec(t *testing.T) {
	expect_dirpath := "dir/path"
	expect_filename := "filename"
	expect_basename := "basename"
	expect_fileext := ".ext"
	actual := FileSpec{
		DirPath:  "dir/path",
		FileName: "filename",
		BaseName: "basename",
		FileExt:  ".ext",
	}
	if expect_dirpath != actual.DirPath {
		t.Errorf(`expect="%s" actual="%s"`, expect_dirpath, actual.DirPath)
	}
	if expect_filename != actual.FileName {
		t.Errorf(`expect="%s" actual="%s"`, expect_filename, actual.FileName)
	}
	if expect_basename != actual.BaseName {
		t.Errorf(`expect="%s" actual="%s"`, expect_basename, actual.BaseName)
	}
	if expect_fileext != actual.FileExt {
		t.Errorf(`expect="%s" actual="%s"`, expect_fileext, actual.FileExt)
	}
}
