// cmd は、画像変換などのコマンド集
package cmd

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

//  処理対象のディレクトリ
var trgDir string

//  変換ルール (変換元の画像形式:変換後の画像形式)
var convCommand string

//  Example:
//  convImgDef = map[string]string
//  	"srcDir":  trgDir,
//  	"srcExt":  cmd[0],
//  	"destExt": cmd[1],
//  }
var convImgDef map[string]string

//  処理対象の画像パスの配列
type TargetImages []string

func init() {
	flag.StringVar(&trgDir, "d", "./", "画像変換したいディレクトリを指定する")
	flag.StringVar(&convCommand, "f", "jpg:png", "変換元の画像形式:変換後の画像形式")
	flag.Parse()
}

//  searchChildPaths は、指定したディレクトリ配下にある画像ファイル（処理対象の拡張子）のパス配列を取得する
func (t TargetImages) searchChildPaths(target_dir string) (TargetImages, error) {
	var paths TargetImages
	err := filepath.Walk(target_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == "."+convImgDef["srcExt"] {
			paths = append(paths, path)
		}
		return nil
	})
	return paths, err
}

// convertToDestExt は、対象の画像を指定した形式に変換する
func (t TargetImages) convertToDestExt(quality int) {
	for _, src := range t {
		// open src file
		file, err := os.Open(src)
		if err != nil {
			fmt.Println("src open err=", err)
			continue
		}
		defer file.Close()

		// open dest file
		destBaseWithoutExt := src[:len(src)-len(filepath.Ext(src))]
		dest := destBaseWithoutExt + "." + convImgDef["destExt"]
		out, err := os.Create(dest)
		if err != nil {
			fmt.Println("dest open err=", err)
			continue
		}
		defer out.Close()

		// decode
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println("decode err=", err)
			continue
		}

		// convert images
		switch convImgDef["destExt"] {
		case "jpg", "jpeg", "JPG", "JPEG":
			opts := &jpeg.Options{Quality: quality}
			err = jpeg.Encode(out, img, opts)
		case "png", "PNG":
			err = png.Encode(out, img)
		case "gif", "GIF":
			opts := &gif.Options{}
			err = gif.Encode(out, img, opts)
		}

		// delete src image
		if err != nil {
			fmt.Println("conv err=", err)
			continue
		}
		if err = os.Remove(src); err != nil {
			fmt.Println("delete err=", err)
		}
	} // end each src_image
}

// ConvImages は、指定したディレクトリ配下の画像ファイルの形式を変換する
func ConvImages() {
	cmd := strings.Split(convCommand, ":")
	convImgDef = map[string]string{
		"srcDir":  trgDir,
		"srcExt":  cmd[0],
		"destExt": cmd[1],
	}

	// 変換対象ファイルの取得
	var srcImgs TargetImages
	srcImgs, err := srcImgs.searchChildPaths(convImgDef["srcDir"])
	if err != nil {
		fmt.Println(err)
	}
	// 変換
	srcImgs.convertToDestExt(100)

	fmt.Println("opt:", convImgDef)
	fmt.Println("paths:", srcImgs)
}
