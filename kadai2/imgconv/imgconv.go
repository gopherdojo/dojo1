// Package imgconv provides some simple image convert function.
// These are sample functions to practice golang.
package imgconv

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ImgConv(i Imageconverter, srcdir, outputfiletype string) {
	imagefilelist, err := i.GetImage(srcdir)
	if err != nil {
		fmt.Println(err)
	}
	err = i.ConvertImage(outputfiletype, imagefilelist)
	if err != nil {
		fmt.Println(err)
	}
}

type Imageconverter interface {
	GetImage(srcdir string) ([]Imagefile, error)
	ConvertImage(outputfiletype string, imagefile []Imagefile) error
}

type Imagefile struct {
	image         image.Image
	imagefilepath string
}

func (i *Imagefile) GetImage(srcdir string) ([]Imagefile, error) {
	imagefilelist := []Imagefile{}

	err := filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// GEt image
		img, err := getImg(path)
		if err != nil {
			return err
		}

		i.image = img
		i.imagefilepath = path
		imagefilelist = append(imagefilelist, *i)

		return nil
	})
	return imagefilelist, err
}

func (i *Imagefile) ConvertImage(outputfiletype string, imagefile []Imagefile) error {
	if _, err := os.Stat("out"); err != nil {
		if err := os.Mkdir("out", 0755); err != nil {
			fmt.Println(err)
		}
	}

	for _, imagefile := range imagefile {
		f_pos := strings.LastIndex(imagefile.imagefilepath, "/")
		p_pos := strings.LastIndex(imagefile.imagefilepath, ".")
		out, err := os.Create("out/" + imagefile.imagefilepath[f_pos+1:p_pos] + "." + outputfiletype)
		if err != nil {
			return err
		}
		switch outputfiletype {
		case "jpeg", "jpg":
			err = jpeg.Encode(out, imagefile.image, nil)
			if err != nil {
				return err
			}
		case "png":
			err = png.Encode(out, imagefile.image)
			if err != nil {
				return err
			}
		default:
			return errors.New("sorry. not support this outputfiletype extend")
		}
	}
	return nil

}

func getImg(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
