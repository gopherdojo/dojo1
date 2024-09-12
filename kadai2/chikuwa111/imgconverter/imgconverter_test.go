package imgconverter_test

import (
	"imgconverter"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	type decodeCase struct {
		name string
		path string
	}

	cases := []decodeCase{
		{name: "png", path: "testdata/1x1.png"},
		{name: "jpg", path: "testdata/1x1.jpg"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			img, err := imgconverter.Decode(c.path)
			if err != nil {
				t.Errorf("Got error (arg path: %s)", c.path)
			}
			if img.Image == nil {
				t.Errorf("Decoded image should NOT be nil (arg path: %s)", c.path)
			}
		})
	}

	errorCases := []decodeCase{
		{name: "inexistentFile", path: "testdata/1x1.jpeg"},
		{name: "cantDecode", path: "testdata/1x1.gif"},
	}
	for _, c := range errorCases {
		t.Run(c.name, func(t *testing.T) {
			img, err := imgconverter.Decode(c.path)
			if err == nil {
				t.Errorf("Should return error (arg path: %s)", c.path)
			}
			if img.Image != nil {
				t.Errorf("Decoded image should be nil (arg path: %s)", c.path)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type encodeCase struct {
		name      string
		inputPath string
		dest      string
	}

	cases := []encodeCase{
		{name: "png", inputPath: "testdata/1x1.jpg", dest: "testdata/1x1.jpg.png"},
		{name: "jpg", inputPath: "testdata/1x1.png", dest: "testdata/1x1.png.jpg"},
		{name: "jpeg", inputPath: "testdata/1x1.png", dest: "testdata/1x1.png.jpeg"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			img := testDecode(t, c.inputPath)
			err := img.Encode(c.dest)
			if err != nil {
				t.Errorf("Got error (arg dest: %s)", c.dest)
			}
			if _, err := os.Stat(c.dest); err != nil {
				t.Errorf("Should create file at %s (arg dest: %s)", c.dest, c.dest)
			} else {
				testRemove(t, c.dest)
			}
		})
	}

	errorCases := []encodeCase{
		{name: "gif", inputPath: "testdata/1x1.png", dest: "testdata/1x1.png.gif"},
	}
	for _, c := range errorCases {
		t.Run(c.name, func(t *testing.T) {
			img := testDecode(t, c.inputPath)
			err := img.Encode(c.dest)
			if err == nil {
				t.Errorf("Should return error (arg dest: %s)", c.dest)
			}
			if _, err := os.Stat(c.dest); err == nil {
				t.Errorf("Should NOT create file at %s (arg dest: %s)", c.dest, c.dest)
				testRemove(t, c.dest)
			}
		})
	}
}

func testDecode(t *testing.T, path string) imgconverter.Image {
	t.Helper()
	img, err := imgconverter.Decode(path)
	if err != nil {
		t.Fatalf("err %s", err)
	}
	return img
}

func testRemove(t *testing.T, path string) {
	t.Helper()
	err := os.Remove(path)
	if err != nil {
		t.Fatalf("err %s", err)
	}
}
