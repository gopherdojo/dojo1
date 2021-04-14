package utils

import (
	"path/filepath"
	"os"
	"image"
	"image/jpeg"
	"image/gif"
	"image/png"
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

			filename := getFileName(path)
			dirpath := getPathName(path)

			out, err := os.Create(dirpath + "/" + filename + "." + p.ToFormat)
			logError(err)
			defer out.Close()

			switch p.ToFormat {
			case "jpg", "jpeg":

				jpeg.Encode(out, img, nil)

			case "png":

				png.Encode(out, img)

			case "gif":
				gif.Encode(out, img, nil)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// get fileneme without filepath and extention
func getFileName(path string) string{
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// get dirpath
func getPathName(path string) string {
	return filepath.Dir(path)
}