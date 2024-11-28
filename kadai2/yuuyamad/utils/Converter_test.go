package utils

import (
	"testing"
	"image"
	"os"
)

func Test_createConvertFile(t *testing.T) {

	test, err := testFileDecode(t)
	if err != nil {
		t.Fatal("file open error")
	}

	var tests = []struct{
		path string
		ext  string
		img  image.Image
		want error
	}{

		{path: "./testdata/Lenna.png", ext: "gif",img: test, want: nil},
		{path: "./testdata/Lenna.gif", ext: "jpeg",img: test, want: nil},
		{path: "./testdata/Lenna.jpeg", ext: "png",img: test, want: nil},
	}
	for _, test := range tests {
		result := createConvertFile(test.path, test.ext, test.img)
		if result != nil {
			t.Errorf(`getFileName(%s, %s) = false`, test.path, test.ext)
		}
	}

}

func testFileDecode(t *testing.T) (image.Image, error){
	t.Helper()
	tf, err := os.Open("testdata/Lenna.png")
	if err != nil {
		return nil, err
	}
	defer tf.Close()
	ti , _, err := image.Decode(tf)
	if err != nil {
		return nil, err
	}

	return ti, nil

}

