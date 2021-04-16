package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

// Decoder has image decode function.
type Decoder interface {
	Decode(r io.Reader) (image.Image, error)
	ext() string
}

// NewDecoder creates Decoder.
func NewDecoder(format string) (Decoder, error) {
	switch strings.ToLower(format) {
	case "gif":
		return &gifDecoder{}, nil
	case "jpg", "jpeg":
		return &jpgDecoder{}, nil
	case "png":
		return &pngDecoder{}, nil
	default:
		return nil, fmt.Errorf("unsupported decoder format: %s", format)
	}
}

type gifDecoder struct{}

func (d *gifDecoder) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (d *gifDecoder) ext() string {
	return "gif"
}

type jpgDecoder struct{}

func (d *jpgDecoder) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (d *jpgDecoder) ext() string {
	return "jpg"
}

type pngDecoder struct{}

func (d *pngDecoder) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (d *pngDecoder) ext() string {
	return "png"
}
