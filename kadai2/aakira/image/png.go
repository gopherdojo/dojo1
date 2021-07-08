package image

import (
	"io"
	"image"
	"image/png"
)

const pngExtension = ".png"

type PngImage struct {
	Path string
}

func (f *PngImage) GetPath() string {
	return f.Path
}

func (f *PngImage) GetExt() string {
	return pngExtension
}

func (f *PngImage) Encode(out io.Writer, img image.Image) error {
	err := png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

func (f *PngImage) Decode(file io.Reader) (image.Image, error) {
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}