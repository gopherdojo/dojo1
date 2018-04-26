// Package convimg_test is test
package convimg_test

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"testing"

	convimg "github.com/matsu0228/go_sandbox/02_convimg_test/convimg"
)

// TestMail is initial Execute in this tests.
func TestMain(m *testing.M) {
	beforeTests()
	code := m.Run()
	afterTests()

	os.Exit(code)
}

func beforeTests() {
	fmt.Println("before all..")
	initTestImages()
}
func afterTests() {
	fmt.Println("after all..")
	initTestImages()
}

func initTestImages() {
	var commands = []struct {
		trgDir     string
		srcFormat  string
		destFormat string
	}{
		{"./testdata/jpeg", "*", "jpeg"},
		{"./testdata/png", "*", "png"},
		{"./testdata/gif", "*", "gif"},
	}

	for _, c := range commands {
		c, err := convimg.New(c.trgDir, c.srcFormat, c.destFormat)
		if err != nil {
			panic("beforeTest() convimg.New(): " + err.Error())
		}
		if _, err := c.ConvImages(); err != nil {
			panic("beforeTest() convimg.ConvImages(): " + err.Error())
		}
	}
}

// 処理対象外は、処理されていないこと
func TestConvImages(t *testing.T) {
	var commands = []struct {
		trgDir     string
		srcFormat  string
		destFormat string
	}{
		{"./testdata/jpeg", "jpeg", "png"},
		{"./testdata/jpeg", "png", "gif"},
		{"./testdata/", "gif", "jpeg"},
	}

	for _, c := range commands {
		ci, err := convimg.New(c.trgDir, c.srcFormat, c.destFormat)
		if err != nil {
			panic("beforeTest() convimg.New(): " + err.Error())
		}
		images, err := ci.ConvImages()
		if err != nil {
			panic("beforeTest() convimg.ConvImages(): " + err.Error())
		}
		for _, img := range images {
			t.Run(img, func(t *testing.T) {
				testImgFormat(img, ci.DestFormat, t)
				testSrcImg(img, ci.SrcFormat, t)
			})
		}
	}
}

func testSrcImg(trgImg, srcFormat string, t *testing.T) {
	trgBaseWithoutExt := trgImg[:len(trgImg)-len(filepath.Ext(trgImg))]
	srcImg := trgBaseWithoutExt + "." + srcFormat
	_, err := os.Stat(srcImg)
	if err == nil {
		t.Errorf("srcImage exist err. src= \"%s\" want to be deleted", srcImg)
	}
}

func testImgFormat(trgImg, want string, t *testing.T) {
	file, err := os.Open(trgImg)
	if err != nil {
		panic("cant open image: " + err.Error())
	}
	defer file.Close()

	// 拡張子チェック
	ext := filepath.Ext(trgImg)
	if ext != "."+want {
		t.Errorf("extension err want %s but %s", want, ext)
	}
	// 画像形式のチェック
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		panic("cant get image format: " + err.Error())
	}
	if format != want {
		t.Errorf("format err want %s but %s", want, format)
	}
}
