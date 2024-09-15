package helpers

import (
	"os"
	"testing"
)

var file_tests = []struct {
	in  string
	out string
}{
	{base_dir + "/testdata/jpeg/video-001.221212.jpeg", "jpeg"},
	{base_dir + "/testdata/png/video-001.221212.png", "png"},
}

var image_tests = []struct {
	in  string
	out ConvertSpec
}{
	{base_dir + "/testdata/jpeg/video-001.221212.jpeg", ConvertSpec{
		Src:    base_dir + "/testdata/jpeg/video-001.221212.jpeg",
		Dst:    test_dir + "/video-001.221212.png",
		Format: "png",
	}},
	{base_dir + "/testdata/png/video-001.221212.png", ConvertSpec{
		Src:    base_dir + "/testdata/png/video-001.221212.png",
		Dst:    test_dir + "/video-001.221212.jpeg",
		Format: "jpeg",
	}},
}

var bulk_test = []ConvertSpec{
	ConvertSpec{
		Src:    base_dir + "/testdata/jpeg/video-001.221212.jpeg",
		Dst:    test_dir + "/video-001.221212.png",
		Format: "png",
	},
	ConvertSpec{
		Src:    base_dir + "/testdata/png/video-001.221212.png",
		Dst:    test_dir + "/video-001.221212.jpeg",
		Format: "jpeg",
	},
}

func TestCheckFileType(t *testing.T) {
	for _, tt := range file_tests {
		expect := tt.out
		actual := checkFileType(tt.in)
		if expect != actual {
			t.Errorf(`expect="%s" actual="%s"`, expect, actual)
		}
	}
}

func TestCreateEmptyFile(t *testing.T) {
	test_file := "../TestCreateEmptyFile.txt"
	defer func() {
		if r := recover(); r != nil {
			t.Error("Test finished with Panic!")
		}
	}()
	file := createEmptyFile(test_file)
	file.Close()
	err := os.Remove(test_file)
	if err != nil {
		t.Error(`Test file: %s could not deleted.`, test_file)
	}
}

func TestOpenImageFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Test finished with Panic!")
		}
	}()
	for _, tt := range file_tests {
		expect := tt.out
		_, actual := openImageFile(tt.in)
		if expect != actual {
			t.Errorf(`expect="%s" actual="%s"`, expect, actual)
		}
	}
}

func TestFileOpen(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Test finished with Panic!")
		}
	}()
	for _, tt := range file_tests {
		expect := tt.out
		file, actual := fileOpen(tt.in)
		if expect != actual {
			t.Errorf(`expect="%s" actual="%s"`, expect, actual)
		}
		file.Close()
	}
}

func TestConvertImageFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Test finished with Panic!")
		}
	}()
	for _, tt := range image_tests {
		target := tt.in
		actual := ConvertImageFile(tt.out.Src, tt.out.Dst, tt.out.Format)
		if actual != nil {
			t.Errorf(`Target file: %s could not convert. %s`, target, actual)
		}
	}
	err := os.RemoveAll(test_dir)
	if err != nil {
		t.Error(`Test directory: %s could not deleted.`, test_dir)
	}
}

func TestBulkConvert(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Test finished with Panic!")
		}
	}()
	BulkConvert(bulk_test)
	err := os.RemoveAll(test_dir)
	if err != nil {
		t.Error(`Test directory: %s could not deleted.`, test_dir)
	}
}
