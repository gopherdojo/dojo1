package gif

import (
	"image"
	"image/gif"
	"io"

	"github.com/itosho/gonverter/convert"
)

const (
	NumColors = 256
)

type Gif struct{}

func init() {
	convert.Register(".gif", Gif{})
}

func (g Gif) Decode(file io.Reader) (image.Image, error) {
	return gif.Decode(file)
}

func (g Gif) Encode(out io.Writer, img image.Image) error {
	return gif.Encode(out, img, &gif.Options{NumColors: NumColors})
}
