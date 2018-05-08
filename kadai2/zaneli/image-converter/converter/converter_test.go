package converter_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"../converter"
)

func TestNewConverter_ValidArg(t *testing.T) {
	args := []Arg{
		Arg{"gif", "jpg"},
		Arg{"gif", "JPG"},
		Arg{"gif", "jpeg"},
		Arg{"gif", "JPEG"},
		Arg{"gif", "png"},
		Arg{"gif", "PNG"},
		Arg{"jpg", "gif"},
		Arg{"jpg", "GIF"},
		Arg{"jpg", "png"},
		Arg{"jpg", "PNG"},
		Arg{"jpeg", "gif"},
		Arg{"jpeg", "GIF"},
		Arg{"jpeg", "png"},
		Arg{"jpeg", "PNG"},
		Arg{"png", "jpg"},
		Arg{"png", "gif"},
		Arg{"png", "GIF"},
		Arg{"png", "JPG"},
		Arg{"png", "jpeg"},
		Arg{"png", "JPEG"},
		Arg{"GIF", "jpg"},
		Arg{"GIF", "JPG"},
		Arg{"GIF", "jpeg"},
		Arg{"GIF", "JPEG"},
		Arg{"GIF", "png"},
		Arg{"GIF", "PNG"},
		Arg{"JPG", "gif"},
		Arg{"JPG", "GIF"},
		Arg{"JPG", "png"},
		Arg{"JPG", "PNG"},
		Arg{"JPEG", "gif"},
		Arg{"JPEG", "GIF"},
		Arg{"JPEG", "png"},
		Arg{"JPEG", "PNG"},
		Arg{"PNG", "jpg"},
		Arg{"PNG", "gif"},
		Arg{"PNG", "GIF"},
		Arg{"PNG", "JPG"},
		Arg{"PNG", "jpeg"},
		Arg{"PNG", "JPEG"},
	}
	for _, arg := range args {
		_, err := converter.NewConverter(arg.from, arg.to)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNewConverter_InvalidArg_Unsupported(t *testing.T) {
	args := []ArgWithExpected{
		ArgWithExpected{Arg{"gif", "XXX"}, "unsupported encoder format: XXX"},
		ArgWithExpected{Arg{"gif", "jpg "}, "unsupported encoder format: jpg "},
		ArgWithExpected{Arg{"gif", ""}, "unsupported encoder format: "},
		ArgWithExpected{Arg{"XXX", "png"}, "unsupported decoder format: XXX"},
		ArgWithExpected{Arg{"g if", "png"}, "unsupported decoder format: g if"},
		ArgWithExpected{Arg{"", "png"}, "unsupported decoder format: "},
	}
	for _, arg := range args {
		_, err := converter.NewConverter(arg.from, arg.to)
		if err == nil || err.Error() != arg.expected {
			t.Fatal(err)
		}
	}
}

func TestNewConverter_InvalidArg_SameFromTo(t *testing.T) {
	args := []ArgWithExpected{
		ArgWithExpected{Arg{"gif", "gif"}, "invalid same format: from=gif, to=gif"},
		ArgWithExpected{Arg{"jpg", "jpg"}, "invalid same format: from=jpg, to=jpg"},
		ArgWithExpected{Arg{"jpeg", "jpeg"}, "invalid same format: from=jpeg, to=jpeg"},
		ArgWithExpected{Arg{"jpg", "jpeg"}, "invalid same format: from=jpg, to=jpeg"},
		ArgWithExpected{Arg{"png", "png"}, "invalid same format: from=png, to=png"},
		ArgWithExpected{Arg{"gif", "GIF"}, "invalid same format: from=gif, to=GIF"},
		ArgWithExpected{Arg{"jpg", "JPEG"}, "invalid same format: from=jpg, to=JPEG"},
	}
	for _, arg := range args {
		_, err := converter.NewConverter(arg.from, arg.to)
		if err == nil || err.Error() != arg.expected {
			t.Fatal(err)
		}
	}
}

func TestBuildNewFilePath(t *testing.T) {
	args := []ArgWithPath{
		ArgWithPath{Arg{"gif", "png"}, "xxx", "xxx.png"},
		ArgWithPath{Arg{"gif", "png"}, "xxx.gif", "xxx.png"},
		ArgWithPath{Arg{"gif", "jpeg"}, "xxx", "xxx.jpeg"},
		ArgWithPath{Arg{"gif", "JPG"}, "xxx", "xxx.JPG"},
	}
	for _, arg := range args {
		c, err := converter.NewConverter(arg.from, arg.to)
		if err != nil {
			t.Fatal(err)
		}
		path := c.BuildNewFilePath(arg.path)
		if path != arg.expected {
			t.Fatalf("expected %s, but actual %s", arg.expected, path)
		}
	}
}

func TestConvert(t *testing.T) {
	args := []ArgWithExpected{
		ArgWithExpected{Arg{"gif", "jpg"}, "gif_gopher.jpg"},
		ArgWithExpected{Arg{"gif", "jpeg"}, "gif_gopher.jpeg"},
		ArgWithExpected{Arg{"gif", "png"}, "gif_gopher.png"},
		ArgWithExpected{Arg{"jpg", "gif"}, "jpg_gopher.gif"},
		ArgWithExpected{Arg{"jpg", "png"}, "jpg_gopher.png"},
		ArgWithExpected{Arg{"jpeg", "gif"}, "jpg_gopher.gif"},
		ArgWithExpected{Arg{"jpeg", "png"}, "jpg_gopher.png"},
		ArgWithExpected{Arg{"gif", "JPG"}, "gif_gopher.JPG"},
		ArgWithExpected{Arg{"PNG", "GIF"}, "png_gopher.GIF"},
	}
	for i, arg := range args {
		dir, srcFiles, destDirs := prepareFiles(t, fmt.Sprintf("case_%d", i))
		defer os.RemoveAll(dir)

		c, err := converter.NewConverter(arg.from, arg.to)
		if err != nil {
			t.Fatal(err)
		}

		if err = c.Convert(dir); err != nil {
			t.Fatal(err)
		}

		assertFiles(t, append(srcFiles, arg.expected), destDirs)
	}
}

type Arg struct {
	from string
	to   string
}

type ArgWithExpected struct {
	Arg
	expected string
}

type ArgWithPath struct {
	Arg
	path     string
	expected string
}

func prepareFiles(t *testing.T, ext string) (string, []string, []string) {
	t.Helper()

	dir := fmt.Sprintf("./testdata/%s", ext)

	if err := os.RemoveAll(dir); err != nil {
		t.Error(err)
	}

	if err := os.MkdirAll(fmt.Sprintf("%s/dir1/dir2", dir), 0777); err != nil {
		t.Error(err)
	}

	srcFiles := []string{
		"gif_gopher.gif",
		"jpg_gopher.jpg",
		"png_gopher.png",
		"not_image.txt",
	}
	destDirs := []string{
		dir,
		fmt.Sprintf("%s/dir1", dir),
		fmt.Sprintf("%s/dir1/dir2", dir),
	}
	for _, srcFile := range srcFiles {
		src, err := os.Open(fmt.Sprintf("./testdata/base/%s", srcFile))
		if err != nil {
			t.Error(err)
		}
		defer src.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, src)
		img := buf.Bytes()

		for _, destDir := range destDirs {
			dest, err := os.Create(filepath.Join(destDir, srcFile))
			if err != nil {
				t.Error(err)
			}
			defer dest.Close()

			if _, err = io.Copy(dest, bytes.NewBuffer(img)); err != nil {
				t.Error(err)
			}
		}
	}

	// 事前にテスト用ファイルが想定通り作成されている事を確認
	assertFiles(t, srcFiles, destDirs)

	return dir, srcFiles, destDirs
}

func assertFiles(t *testing.T, srcFiles []string, destDirs []string) {
	t.Helper()

	sort.Strings(srcFiles)
	sort.Strings(destDirs)

	for _, destDir := range destDirs {
		children, err := ioutil.ReadDir(destDir)
		if err != nil {
			t.Error(err)
		}
		var destFiles []string
		for _, child := range children {
			if !child.IsDir() {
				destFiles = append(destFiles, child.Name())
			}
		}
		if !reflect.DeepEqual(srcFiles, destFiles) {
			t.Errorf("testdata files expected=%v, but actual=%v", srcFiles, destFiles)
		}
	}
}
