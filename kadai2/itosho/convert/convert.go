package convert

import (
	"image"
	"io"
	"os"
)

type DecodeEncoder interface {
	Decoder
	Encoder
}

type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

type Encoder interface {
	Encode(io.Writer, image.Image) error
}

var images = map[string]DecodeEncoder{}

func Register(ext string, image DecodeEncoder) {
	images[ext] = image
}

func IsConvertibleImage(ext string) bool {
	_, ok := images[ext]
	return ok
}

func CreateFile(path string, fromExt string, toExt string) error {
	inFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inFile.Close()

	fromImagePkg, _ := images[fromExt]
	fromImage, err := fromImagePkg.Decode(inFile)
	if err != nil {
		return err
	}

	convertFilePath := convertFilePath(path, fromExt, toExt)
	outFile, err := os.Create(convertFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	toImagePkg, _ := images[toExt]
	if err := toImagePkg.Encode(outFile, fromImage); err != nil {
		return err
	}

	return nil
}

func RemoveFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func convertFilePath(path string, fromExt string, toExt string) string {
	return path[:len(path)-len(fromExt)] + toExt
}
