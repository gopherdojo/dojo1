package encode

import (
	"image"
	"image/gif"
	"io"
)

type GifEncoder struct {
	image.Image
}

func NewGifEncoder(reader io.Reader) (Encoder, error) {
	img, err := gif.Decode(reader)
	if err != nil {
		return nil, err
	}
	return GifEncoder{img}, nil
}

func (ge GifEncoder) Encode(writer io.Writer) error {
	err := gif.Encode(writer, ge, nil)
	if err != nil {
		return err
	}
	return nil
}