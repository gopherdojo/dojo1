package convimg_test

import (
	"testing"

	convimg "github.com/matsu0228/go_sandbox/02_convimg_test/convimg"
)

var validPramTests = []struct {
	name       string
	trgDir     string
	srcFormat  string
	destFormat string
}{
	{"jpgTopng", "./testdata", "jpg", "png"}, {"jpgTogif", "./", "jpg", "gif"},
	{"jpegTopng", "./", "jpeg", "png"}, {"jpegTogif", "./testdata", "jpeg", "gif"},
	{"pngTojpg", "./testdata", "png", "jpg"}, {"pngTogif", "./testdata", "png", "gif"},
	{"gitTojpeg", "./", "gif", "jpeg"}, {"gifTopng", "./testdata", "gif", "png"},
	{"anyTojpeg", "./", "*", "jpeg"}, {"anyTopng", "./testdata", "*", "png"},
}

var invalidPramTests = []struct {
	name       string
	trgDir     string
	srcFormat  string
	destFormat string
}{
	{"notExistDir", "./hoge", "jpg", "png"}, {"notDir", "./main.go", "jpg", "gif"},
	{"invalidSrcFormat", "./testdata", "pjg", "png"}, {"invalidDestFormat", "./testdata", "jpeg", "igf"},
	{"invalidDestFormat", "./testdata", "png", "*"},
}

func TestValidParameter(t *testing.T) {
	for _, vp := range validPramTests {
		t.Run(vp.name, func(t *testing.T) {
			c, err := convimg.New(vp.trgDir, vp.srcFormat, vp.destFormat)
			if err != nil {
				t.Errorf("convimg.New(%s, %s, %s) is valid, but err occured of %q", vp.trgDir, vp.srcFormat, vp.destFormat, err)
			}
			if c.TrgDir != vp.trgDir {
				t.Errorf("convimg.New(%s, %s, %s) wants %s but %s", vp.trgDir, vp.srcFormat, vp.destFormat, vp.trgDir, c.TrgDir)
			}
			if c.SrcFormat != vp.srcFormat {
				t.Errorf("convimg.New(%s, %s, %s) wants %s but %s", vp.trgDir, vp.srcFormat, vp.destFormat, vp.srcFormat, c.SrcFormat)
			}
			if c.DestFormat != vp.destFormat {
				t.Errorf("convimg.New(%s, %s, %s) wants %s but %s", vp.trgDir, vp.srcFormat, vp.destFormat, vp.destFormat, c.DestFormat)
			}
		})
	}
}

func TestInValidParameter(t *testing.T) {
	for _, vp := range invalidPramTests {
		t.Run(vp.name, func(t *testing.T) {
			_, err := convimg.New(vp.trgDir, vp.srcFormat, vp.destFormat)
			if err == nil {
				t.Errorf("convimg.New(%s, %s, %s) is invalid, but err didnt occur.", vp.trgDir, vp.srcFormat, vp.destFormat)
			}
		})
	}
}
