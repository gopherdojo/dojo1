package image

import (
	"io"
	"image"
	"image/jpeg"
)

const jpgExtension = ".jpg"

type JpgImage struct {
	Path string
}

func (f *JpgImage) GetPath() string {
	return f.Path
}

func (f *JpgImage) GetExt() string {
	return jpgExtension
}

func (f *JpgImage) Encode(out io.Writer, img image.Image) error {
	err := jpeg.Encode(out, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		return err
	}
	return nil
}

func (f *JpgImage) Decode(file io.Reader) (image.Image, error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}