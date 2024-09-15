/*
[image_converter]
JPEG画像をPNG画像に変換。またはPNG画像をJPEG画像に変換

[コマンド]
image_converter -d=出力先フォルダ -f=変換後の画像形式(jpegまたはpng) 変換元フォルダ

Example:
    image_converter -d=./jpeg -f=png ./output/png
    「./jpeg」ディレクトリ内のJPEG画像をPNG画像に変換して、「./output/png」ディレクトリに出力
*/
package main

import (
  "fmt"
	"./helpers"
)

func convert(srcDir string, dstDir string, dstFormat string) {
  files := helpers.DirWalker(srcDir)
  specs := helpers.MakeConvertSpec(files, dstDir, dstFormat)
  helpers.BulkConvert(specs)
  return
}

func main() {
  srcDir, dstDir, dstFormat := helpers.CheckParams()
  fmt.Printf("Start Image Convert... \n") 
  fmt.Printf("Source dirrectory: %s, Destination directory: %s, Image Format: %s \n", srcDir, dstDir, dstFormat)
  convert(srcDir, dstDir, dstFormat)
}
