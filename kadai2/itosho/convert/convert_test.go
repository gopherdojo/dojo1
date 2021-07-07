package convert

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	path := testTempFile(t)
	actual := RemoveFile(path)
	if actual != nil {
		t.Errorf(`actual="%s"`, actual)
	}

	_, err := os.Stat(path)
	if err == nil {
		t.Error("Test file is not removed.")
	}
}

func TestConvertFilePath(t *testing.T) {
	expect := "gopherdojo.jpg"
	actual := convertFilePath("gopherdojo.png", ".png", ".jpg")
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}

func testTempFile(t *testing.T) string {
	t.Helper()
	p, _ := os.Getwd()
	tf, err := ioutil.TempFile(p, "test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	tf.Close()
	return tf.Name()
}
