package img

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
	"strings"
)

func getOutputFile(path, ext string) string {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return path[:len(path)-len(filepath.Ext(path))] + ext
}

func isTargetFile(path, ext string) bool {
	return strings.HasSuffix(path, ext)
}

// Convert change extension of image
func Convert(r io.Reader, w io.Writer, ext string) error {
	m, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	switch ext {
	case ".jpg", ".jpeg":
		if err := png.Encode(w, m); err != nil {
			return err
		}
	case ".png":
		opts := &jpeg.Options{
			Quality: 100,
		}
		if err := jpeg.Encode(w, m, opts); err != nil {
			return err
		}
	default:
		return errors.New("specify jpeg or png for the image format.")
	}

	if err := png.Encode(w, m); err != nil {
		return err
	}

	return nil
}
