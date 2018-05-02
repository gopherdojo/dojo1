// Package converter provides image convert.
// support fomat GIF, JPEG and PNG.
package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func replaceExt(filePath, newExt string) string {
	ext := filepath.Ext(filePath)
	return strings.TrimSuffix(filePath, ext) + newExt
}

func decodeImage(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func convertToJPEG(filename string, src io.Reader) error {
	img, err := decodeImage(src)
	if err != nil {
		return err
	}

	destFile, err := os.Create(replaceExt(filename, ".jpg"))
	err = jpeg.Encode(destFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func convertToPNG(filename string, src io.Reader) error {
	img, err := decodeImage(src)
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

func convertToGIF(filename string, src io.Reader) error {
	img, err := decodeImage(src)
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
	srcFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	switch destExt {
	case "png":
		err := convertToPNG(filename, srcFile)
		if err != nil {
			return err
		}
	case "jpeg", "jpg":
		err := convertToJPEG(filename, srcFile)
		if err != nil {
			return err
		}
	case "gif":
		err := convertToGIF(filename, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}
