// Package convimg は、画像変換などのコマンド集
package convimg

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type imgConvCommand struct {
	TrgDir        string
	SrcFormat     string
	DestFormat    string
	qualityOfJpeg int
}

// New is validate parameters and set on imgConvCommand
func New(trgDir, srcFormat, destFormat string) (*imgConvCommand, error) {
	var err error
	// サポートしている画像形式かチェック
	if err = isAvailableImgformat(srcFormat); err != nil {
		return &imgConvCommand{}, err
	}
	if err = isAvailableImgformat(destFormat); err != nil {
		return &imgConvCommand{}, err
	} else if destFormat == "*" { // 変換先フォーマットは指定が必要
		return &imgConvCommand{}, errors.New("変換先フォーマットは指定が必要です =>" + destFormat)
	}

	// ディレクトリが存在しているかチェック
	err = isExistDir(trgDir)
	if err != nil {
		return &imgConvCommand{}, err
	}

	c := imgConvCommand{}
	c.TrgDir = trgDir
	c.SrcFormat = srcFormat
	c.DestFormat = destFormat
	c.qualityOfJpeg = 100 //TODO
	return &c, nil
}

// ConvImages は、指定したディレクトリ配下の画像ファイルの形式を変換する
func (c *imgConvCommand) ConvImages() ([]string, error) {
	var paths []string
	err := filepath.Walk(c.TrgDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if (c.SrcFormat == "*" && filepath.Ext(path) != "."+c.DestFormat) || filepath.Ext(path) == "."+c.SrcFormat {
			if err = isExistDir(path); err == nil {
				return nil
			}

			dest, err := c.Convert(path) // １ファイルづつ変換
			if err != nil {
				return err
			}
			paths = append(paths, dest) // 変換成功したファイルリスト
		}
		return nil
	})

	// 処理結果の返却
	if err != nil {
		return []string{}, errors.Wrapf(err, "ConvImages() with %s", c.TrgDir)
	}
	return paths, nil
}

func isExistDir(targetDir string) error {
	file, err := os.Open(targetDir)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	return errors.New("ディレクトリではありません: " + targetDir)
}

func isAvailableImgformat(imgFormat string) error {
	switch imgFormat {
	case "jpg", "jpeg", "png", "gif", "*":
		return nil
	default:
		return errors.New("未対応のフォーマットです =>" + imgFormat)
	}
}

// convert は、１ファイルごと変換
func (c *imgConvCommand) Convert(srcImg string) (string, error) {
	// open src file
	file, err := os.Open(srcImg)
	if err != nil {
		return "", errors.Wrapf(err, "convert() with %s", srcImg)
	}
	defer file.Close()

	// open dest file
	destBaseWithoutExt := srcImg[:len(srcImg)-len(filepath.Ext(srcImg))]
	dest := destBaseWithoutExt + "." + c.DestFormat
	out, err := os.Create(dest)
	if err != nil {
		return "", errors.Wrapf(err, "convert() with %s", dest)
	}
	defer out.Close()

	// decode
	img, _, err := image.Decode(file)
	if err != nil {
		return "", errors.Wrapf(err, "convert() decode fails of %s", srcImg)
	}

	// convert images
	switch c.DestFormat {
	case "jpg", "jpeg":
		opts := &jpeg.Options{Quality: c.qualityOfJpeg}
		err = jpeg.Encode(out, img, opts)
	case "png", "PNG":
		err = png.Encode(out, img)
	case "gif", "GIF":
		opts := &gif.Options{}
		err = gif.Encode(out, img, opts)
	}
	if err != nil {
		return "", errors.Wrapf(err, "convert() convert fails of %s", srcImg)
	}

	// delete src image
	if err = os.Remove(srcImg); err != nil {
		return "", errors.Wrapf(err, "convert() delete fails of %s", srcImg)
	}
	return dest, nil
}
