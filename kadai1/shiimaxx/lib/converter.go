// Package converter provides image convert.
// support fomat GIF, JPEG and PNG.
package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func replaceExt(filePath, newExt string) string {
	ext := filepath.Ext(filePath)
	return strings.TrimSuffix(filePath, ext) + newExt
}

func decodeImage(filename string) (image.Image, error) {
	srcFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func convertToJPEG(filename string) error {
	img, err := decodeImage(filename)
	if err != nil {
		return err
	}

	destFile, err := os.Create(replaceExt(filename, ".jpg"))
	if err != nil {
		return err
	}
	defer destFile.Close()

	err = jpeg.Encode(destFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func convertToPNG(filename string) error {
	img, err := decodeImage(filename)
	if err != nil {
		return err
	}

	destFile, err := os.Create(replaceExt(filename, ".png"))
	if err != nil {
		return err
	}
	defer destFile.Close()

	err = png.Encode(destFile, img)
	if err != nil {
		return err
	}
	return nil
}

func convertToGIF(filename string) error {
	img, err := decodeImage(filename)
	if err != nil {
		return err
	}

	destFile, err := os.Create(replaceExt(filename, ".gif"))
	if err != nil {
		return err
	}
	defer destFile.Close()

	err = gif.Encode(destFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}

// Convert convert image to destExt
func Convert(filename, destExt string) error {
	switch destExt {
	case "png":
		err := convertToPNG(filename)
		if err != nil {
			return err
		}
	case "jpeg", "jpg":
		err := convertToJPEG(filename)
		if err != nil {
			return err
		}
	case "gif":
		err := convertToGIF(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
