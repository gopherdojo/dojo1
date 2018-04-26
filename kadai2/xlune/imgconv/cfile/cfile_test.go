package cfile_test

import (
	"testing"

	"github.com/spf13/afero"
	. "github.com/xlune/dojo1/kadai2/tenntenn/imgconv/cfile"
)

var flgNewConvDirInfo = []struct {
	path    string
	success bool
}{
	{"./", true}, {"./hoge", false},
	{"../", true}, {"../hoge", false},
}

func TestNewConvDirInfo(t *testing.T) {
	fs := afero.NewMemMapFs()
	for _, v := range flgNewConvDirInfo {
		_, err := NewConvDirInfo(fs, v.path)
		if s := err == nil; s != v.success {
			t.Fatalf("open %s result %t, (need %t) ", v.path, s, v.success)
		}
	}
}

func TestSetOutputDir(t *testing.T) {
	fs := afero.NewMemMapFs()
	obj, err := NewConvDirInfo(fs, "./")
	if err != nil {
		t.Fatal("open errror")
	}
	err = obj.SetOutputDir("./")
	if err != nil {
		t.Fatal("open errror")
	}
	_, err = fs.Create("./testfile.txt")
	if err != nil {
		t.Fatal("create errror")
	}
	defer fs.Remove("./testfile.txt")
	err = obj.SetOutputDir("./testfile.txt")
	if err == nil {
		t.Fatal("need IsDir error")
	}
}

func TestGetFiles(t *testing.T) {
	fs := afero.NewMemMapFs()
	obj, err := NewConvDirInfo(fs, "./")
	if err != nil {
		t.Fatal("open errror")
	}
	err = fs.MkdirAll("./test1/test2/", 0755)
	if err != nil {
		t.Fatal("make dir error")
	}
	defer fs.RemoveAll("./test1")
	_, err = fs.Create("./test1/t1.go")
	if err != nil {
		t.Fatal("create errror")
	}
	_, err = fs.Create("./test1/test2/t2.go")
	if err != nil {
		t.Fatal("create errror")
	}
	f, err := obj.GetFiles("go", "go2")
	if err != nil {
		t.Fatal("get errror")
	}
	if len(f) != 2 {
		t.Fatal("not match item counts")
	}
}
