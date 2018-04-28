package helpers

import (
	"fmt"
	"image"
	"os"
	//"image/gif"
	"image/jpeg"
	"image/png"
  "strings"
)

// 画像タイプ判別用変数
var magicTable = map[string]string {
    "\xff\xd8\xff":      "jpeg",
    "\x89PNG\r\n\x1a\n": "png",
    "GIF87a":            "gif",
    "GIF89a":            "gif",
}

// 画像タイプのチェック
func checkFileType(src string) string {
  file, err := os.Open(src)
	if err != nil {
    panic(err)
	}
  buffer := make([]byte, 512)
  _, err = file.Read(buffer)
  if err != nil {
    panic(err)
  }
  filemime := ""
  for magic, mime := range magicTable {
    if strings.HasPrefix(string(buffer), magic) {
      filemime = mime
    }
  }
  defer file.Close()
  return filemime
}

// 指定ファイルを開いて、os.Fileを返す
func fileOpen(src string) (*os.File, string) {
  mimetype := checkFileType(src)
  if mimetype == "" {
    panic("Invalid MIME Type: " + src)
  }
  file, err := os.Open(src)
	if err != nil {
    panic(err)
	}
	return file, mimetype
}

// 書き込み可能なos.Fileを返す
func createEmptyFile(dst string) *os.File {
	file, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	return file
}

// 画像ファイルを開いて、image.Imageを返す
func openImageFile(src string) (image.Image, string)  {
  srcFile, fileType := fileOpen(src)
  var img image.Image
  var err error
  if fileType == "jpeg" { 
    img, err = jpeg.Decode(srcFile)
  } else if fileType == "png" {
    img, err = png.Decode(srcFile)
  }
	if err != nil {
		panic(err)
	}
  defer srcFile.Close()
  return img, fileType
}

// 指定のフォーマットで画像を書き込み 
func writeImageFile(img image.Image, dst string, fileType string) error {
  dirCheckAndMkdir(dst)
  file := createEmptyFile(dst)
  var result error
  if fileType == "jpeg" {
    result = jpeg.Encode(file, img, nil)
  } else if fileType == "png" {
    result = png.Encode(file, img)
  } else {
    panic("Invalid File Type! ")
  }
  defer file.Close()
  return result
}

// 指定のフォーマットで画像を変換する
func ConvertImageFile(src string, dst string, toFormat string) error {
  srcFile, _ := openImageFile(src)
  result := writeImageFile(srcFile, dst, toFormat)
	return result
}

// ConvertSpecに格納された情報で、複数の画像を変換
func BulkConvert(specs []ConvertSpec) {
  for _, spec := range specs {
    srcFile := spec.Src
    dstFile := spec.Dst
    toFormat := spec.Format
    if len(dstFile) > 0 && len(toFormat) > 0 {
      result := ConvertImageFile(srcFile, dstFile, toFormat)
	    if result != nil {
        fmt.Printf("[ERROR] from: %s, to: %s\n", srcFile, dstFile)
		    fmt.Println(result)
	    } else {
        fmt.Printf("[OK] from: %s, to: %s\n", srcFile, dstFile)
      }
    } else {
      fmt.Printf("[IGNORE] Unsupported file type: %s \n", srcFile)
    }
  }
}
