package convert

import (
	"image"
	"image/png"
	"io"
)

type cpng struct {
	ext string
}

func (p *cpng) GetExt() string {
	return p.ext
}

func (p *cpng) Decode(file io.Reader) (image.Image, error) {
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (p *cpng) Encode(out io.Writer, img image.Image) error {
	err := png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}
