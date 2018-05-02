package convert

import (
	"image"
	"image/gif"
	"io"
)

type cgif struct {
	ext string
}

func (g *cgif) GetExt() string {
	return g.ext
}

func (g *cgif) Decode(file io.Reader) (image.Image, error) {
	img, err := gif.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (g *cgif) Encode(out io.Writer, img image.Image) error {
	err := gif.Encode(out, img, &gif.Options{NumColors: 256})
	if err != nil {
		return err
	}
	return nil
}
