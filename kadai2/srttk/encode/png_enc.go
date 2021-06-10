package encode

import (
	"image"
	"image/png"
	"io"
)

type PngEncoder struct {
	image.Image
}

func NewPngEncoder(reader io.Reader) (Encoder, error) {
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}
	return PngEncoder{img}, nil
}

func (pe PngEncoder) Encode(writer io.Writer) error {
	err := png.Encode(writer, pe)
	if err != nil {
		return err
	}
	return nil
}