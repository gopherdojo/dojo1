package encode

import (
	"image"
	"image/jpeg"
	"io"
)

type JpegEncoder struct {
	image.Image
}

func NewJpegEncoder(reader io.Reader) (Encoder, error) {
	img, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}
	return JpegEncoder{img}, nil
}

func (je JpegEncoder) Encode(writer io.Writer) error {
	err := jpeg.Encode(writer, je, nil)
	if err != nil {
		return err
	}
	return nil
}