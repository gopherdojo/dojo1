package image

import (
	"io"
	"image"
	"path/filepath"
	"errors"
)

type Image interface {
	GetPath() string
	GetExt() string
	Encode(out io.Writer, img image.Image) error
	Decode(file io.Reader) (image.Image, error)
}

// convert to image struct
func ToImageFile(path string) (Image, error) {
	switch filepath.Ext(path) {
	case ".jpg":
		return &JpgImage{Path: path}, nil
	case ".png":
		return &PngImage{Path: path}, nil
	default:
		return nil, errors.New("file not found")
	}
}