package jpeg

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/itosho/gonverter/convert"
)

const (
	Quality = 100
)

type Jpeg struct{}

func init() {
	convert.Register(".jpeg", Jpeg{})
	convert.Register(".jpg", Jpeg{})
}

func (j Jpeg) Decode(file io.Reader) (image.Image, error) {
	return jpeg.Decode(file)
}

func (j Jpeg) Encode(out io.Writer, img image.Image) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: Quality})
}
