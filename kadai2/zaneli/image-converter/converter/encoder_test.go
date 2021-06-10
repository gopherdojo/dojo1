package converter_test

import (
	"fmt"
	"testing"

	"../converter"
)

func TestNewEncoder_ValidArg(t *testing.T) {
	args := [][]string{
		[]string{"gif", "&converter.gifEncoder{}"},
		[]string{"jpg", "&converter.jpgEncoder{}"},
		[]string{"jpeg", "&converter.jpgEncoder{}"},
		[]string{"png", "&converter.pngEncoder{}"},
		[]string{"GIF", "&converter.gifEncoder{}"},
		[]string{"JPG", "&converter.jpgEncoder{}"},
		[]string{"JPEG", "&converter.jpgEncoder{}"},
		[]string{"PNG", "&converter.pngEncoder{}"},
	}
	for _, arg := range args {
		d, err := converter.NewEncoder(arg[0])
		if err != nil {
			t.Error(err)
		}
		if fmt.Sprintf("%#v", d) != arg[1] {
			t.Errorf("expected=%s, but actual=%#v", arg[1], d)
		}
	}
}

func TestNewEncoder_InvalidArg(t *testing.T) {
	args := []string{"gifjpg", "xxx", "g if", ""}
	for _, arg := range args {
		e, err := converter.NewEncoder(arg)
		if err == nil {
			t.Errorf("unexpected encoder: %#v", e)
		}
	}
}
