package converter_test

import (
	"fmt"
	"testing"

	"../converter"
)

func TestNewDecoder_ValidArg(t *testing.T) {
	args := [][]string{
		[]string{"gif", "&converter.gifDecoder{}"},
		[]string{"jpg", "&converter.jpgDecoder{}"},
		[]string{"jpeg", "&converter.jpgDecoder{}"},
		[]string{"png", "&converter.pngDecoder{}"},
		[]string{"GIF", "&converter.gifDecoder{}"},
		[]string{"JPG", "&converter.jpgDecoder{}"},
		[]string{"JPEG", "&converter.jpgDecoder{}"},
		[]string{"PNG", "&converter.pngDecoder{}"},
	}
	for _, arg := range args {
		d, err := converter.NewDecoder(arg[0])
		if err != nil {
			t.Error(err)
		}
		if fmt.Sprintf("%#v", d) != arg[1] {
			t.Errorf("expected=%s, but actual=%#v", arg[1], d)
		}
	}
}

func TestNewDecoder_InvalidArg(t *testing.T) {
	args := []string{"gifjpg", "xxx", "g if", ""}
	for _, arg := range args {
		d, err := converter.NewDecoder(arg)
		if err == nil {
			t.Errorf("unexpected decoder: %#v", d)
		}
	}
}
