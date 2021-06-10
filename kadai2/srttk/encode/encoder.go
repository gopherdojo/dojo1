/*
package will be controll convert.

Execution is very slow, so please avoid using it on actual programs
 */
package encode

import (
	"io"
	"strings"
	"errors"
)

type Encoder interface {
	Encode(io.Writer) error
}

func NewEncoder(srcExt string, reader io.Reader) (encoder Encoder, err error) {
	switch srcExt {
	case "jpg", "jpeg":
		encoder, err = NewJpegEncoder(reader)
	case "png":
		encoder, err = NewPngEncoder(reader)
	case "gif":
		encoder, err = NewGifEncoder(reader)
	default:
		encoder = nil
		err = errors.New("not support src ext")
	}
	return
}

func GetDistPath(srcPath string, srcExt string, distExt string) (distPath string) {
	withoutExtPath := strings.TrimRight(srcPath, "."+srcExt)
	return withoutExtPath+"."+distExt
}
