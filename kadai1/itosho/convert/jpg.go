package convert

import (
	"image"
	"image/jpeg"
	"io"
)

type cjpg struct {
	ext string
}

func (j *cjpg) GetExt() string {
	return j.ext
}

func (j *cjpg) Decode(file io.Reader) (image.Image, error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (j *cjpg) Encode(out io.Writer, img image.Image) error {
	err := jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	return nil
}
