// Package processes implements converting JPG image.
package processes

import (
	"errors"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"supermarine1377/types"
)

// convert a Myimage (see types package) passed as a argumaent.
func Convert(arg types.Myimage, extension string) error {
	log.Printf("started to convert %s to %s", arg.FileName, extension)

	if arg.Extention != "jpg" {
		err := errors.New(arg.FileName + "is not .jpg")
		return err
	}
	reader, err := myopen(arg.FileName, arg.Path)
	if err != nil {
		return err
	}

	m, err := mydecode(arg.FileName, reader)
	if err != nil {
		return err
	}

	if err := myencode(m, arg.FileName, extension); err != nil {
		return err
	}

	return nil
}

func myopen(name string, path string) (*os.File, error) {
	log.Printf("opening %s", name)
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return reader, err
}

func mydecode(name string, reader *os.File) (image.Image, error) {
	log.Printf("decoding %s", name)
	m, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func myencode(m image.Image, name string, extension string) error {
	log.Printf("encoding %s to %s", name, extension)
	new, err := newname(name, extension)
	if err != nil {
		return err
	}
	c, err := os.Create(new)
	defer c.Close()
	if err != nil {
		return err
	}
	jpeg.Encode(c, m, &jpeg.Options{Quality: 100})

	return nil
}

func newname(old string, extenstion string) (string, error) {
	log.Printf("renaming %s", old)
	var result string
	var position = strings.LastIndex(old, ".")
	if position == -1 {
		err := errors.New("This name has no dot")
		return result, err
	}
	result = old[0:position] + "." + extenstion
	return result, nil
}
