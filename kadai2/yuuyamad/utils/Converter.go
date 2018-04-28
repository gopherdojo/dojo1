package utils

import (
	"path/filepath"
	"os"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"fmt"
)

type Options struct {
	Dir string
	FromFormat string
	ToFormat string
}

// package for image convert from jpen,jpg,png,gif to jpeg,jpg,png,gif
func (p *Options)ConvertImage() error {
	err := filepath.Walk(p.Dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "." + p.FromFormat {
			file, err := os.Open(path)
			logError(err)
			defer file.Close()

			img, _, err := image.Decode(file)

			logError(err)

			err = createConvertFile(path, p.ToFormat, img)
			logError(err)
		}
		return nil
	})
	if err != nil {
		logError(err)
	}
	return nil
}

// get Convert file name and filepath
func createConvertFile(path string, ext string, img image.Image) error{
	filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	filedir := filepath.Dir(path)

	fullfilaname := filedir + "/" + filename + "." + ext

	switch ext {
	case "jpg", "jpeg":
		out, err := os.Create(fullfilaname)
		if err != nil {
			return err
		}
		defer out.Close()
		jpeg.Encode(out, img, nil)

	case "png":
		out, err := os.Create(fullfilaname)
		if err != nil {
			return err
		}
		defer out.Close()
		png.Encode(out, img)

	case "gif":
		out, err := os.Create(fullfilaname)
		if err != nil {
			return err
		}
		defer out.Close()
		gif.Encode(out, img, nil)
	default:
		return fmt.Errorf("failed convert ext.")
	}
	return nil
}
