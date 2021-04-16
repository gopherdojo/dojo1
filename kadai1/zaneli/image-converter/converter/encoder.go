package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

// Encoder has image encode function.
type Encoder interface {
	Encode(w io.Writer, i image.Image) error
	ext() string
}

// NewEncoder creates Encoder.
func NewEncoder(format string) (Encoder, error) {
	switch strings.ToLower(format) {
	case "gif":
		return &gifEncoder{}, nil
	case "jpg", "jpeg":
		return &jpgEncoder{}, nil
	case "png":
		return &pngEncoder{}, nil
	default:
		return nil, fmt.Errorf("unsupported encoder format: %s", format)
	}
}

type gifEncoder struct{}

func (e *gifEncoder) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, nil)
}

func (e *gifEncoder) ext() string {
	return "gif"
}

type jpgEncoder struct{}

func (e *jpgEncoder) Encode(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, nil)
}

func (e *jpgEncoder) ext() string {
	return "jpg"
}

type pngEncoder struct{}

func (e *pngEncoder) Encode(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}

func (e *pngEncoder) ext() string {
	return "png"
}
