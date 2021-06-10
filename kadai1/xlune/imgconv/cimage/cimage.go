package cimage

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ImageExtPng pngの拡張子
	ImageExtPng = "png"
	// ImageExtJpg jpgの拡張子
	ImageExtJpg = "jpg"
	// ImageExtGif gifの拡張子
	ImageExtGif = "gif"
)

// IsAllowExt 変換に対応している拡張子かチェック
func IsAllowExt(ext string) bool {
	switch ext {
	case ImageExtPng, ImageExtJpg, ImageExtGif:
		return true
	default:
		return false
	}
}

// Convert 画層変換
func Convert(input string, output string) error {
	fi, err := os.Open(input)
	if err != nil {
		return err
	}
	defer fi.Close()
	img, _, err := image.Decode(fi)
	if err != nil {
		return err
	}

	err = mkdirAll(output)
	if err != nil {
		return err
	}
	fo, err := os.Create(output)
	if err != nil {
		return err
	}
	defer fo.Close()

	toExt := getExt(output)
	switch toExt {
	case ImageExtPng:
		return png.Encode(fo, img)
	case ImageExtJpg:
		opt := &jpeg.Options{Quality: 80}
		return jpeg.Encode(fo, img, opt)
	case ImageExtGif:
		opt := &gif.Options{}
		return gif.Encode(fo, img, opt)
	}
	return nil
}

// getExt 拡張子文字列取得
func getExt(path string) string {
	return strings.TrimLeft(filepath.Ext(path), ".")
}

// mkdirAll 指定パスのディレクトリ作成
func mkdirAll(path string) error {
	dirPath := filepath.Dir(path)
	return os.MkdirAll(dirPath, 0755)
}
